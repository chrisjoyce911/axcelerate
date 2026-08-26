// groupedAudit reports every upcoming workshop that belongs to a grouped course,
// and which of those groups are at or over their shared cap.
//
// It only ever reads. Nothing here writes to Axcelerate or posts an event — the
// point is to see what a close-group sweep would touch, and to capture each
// session's current MaxParticipants first, because closing a group overwrites
// those and nothing else records them.
//
// Given -post, it then asks for those groups to be closed, by sending one
// course.workshop_updated event per group to the webhook receiver. That goes
// through the deployed chain — to-lmsdb decides, webhook-events searches the
// group, writes each session's cap down to its headcount, and reads the write
// back to prove it landed. One event closes a whole group, so this is tens of
// posts rather than thousands. Without -post nothing is sent.
//
// It walks every public location with future activity, pacing one Axcelerate
// call every 2 seconds. That takes a while, so progress is saved to
// grouped-audit-progress.json after each location: stop it whenever, run it
// again, and it picks up at the first location it has not done. Delete that file
// to start fresh. grouped-audit.json and grouped-audit.csv are rewritten as it
// goes, so a partial walk still leaves a usable report.
//
//	go run ./example/groupedAudit                  # report only
//	go run ./example/groupedAudit -post <url>      # report, then close the full groups
package main

import (
	"bytes"
	"crypto/rand"
	"encoding/csv"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"time"

	"github.com/chrisjoyce911/axcelerate"
	"github.com/joho/godotenv"
)

// maxActions is a runaway guard on the closing pass. A close is one-way, so if
// the audit turns up more workshops needing adjustment than this, that is a
// result to look at rather than a batch to fire off unattended.
const maxActions = 200

// callDelay paces every Axcelerate request. The client also rate limits, but a
// steady 2 seconds keeps a whole-fleet walk gentle and predictable.
const callDelay = 2 * time.Second

const (
	jsonPath     = "grouped-audit.json"
	csvPath      = "grouped-audit.csv"
	actionsPath  = "grouped-audit-actions.csv"
	progressPath = "grouped-audit-progress.json"
)

// Session is one workshop instance inside a grouped course, captured as it was
// at audit time. MaxParticipants is the field a close overwrites, so this record
// is the only way back.
type Session struct {
	InstanceID      int       `json:"instanceID"`
	Name            string    `json:"name"`
	CourseName      string    `json:"courseName"`
	Code            string    `json:"code"`
	Location        string    `json:"location"`
	StartDate       time.Time `json:"startDate"`
	Participants    int       `json:"participants"`
	MaxParticipants int       `json:"maxParticipants"`
	Vacancy         int       `json:"participantVacancy"`
	EnrolmentOpen   bool      `json:"enrolmentOpen"`
	IsActive        bool      `json:"isActive"`

	GroupedCourseID   int    `json:"groupedCourseID"`
	GroupedCourseName string `json:"groupedCourseName"`
	GroupParticipants int    `json:"groupParticipants"`
	GroupMax          int    `json:"groupMaxParticipants"`
	Simultaneous      bool   `json:"simultaneous"`
}

// Group is one grouped course: the sessions that share a participant pool.
type Group struct {
	GroupedCourseID   int       `json:"groupedCourseID"`
	GroupedCourseName string    `json:"groupedCourseName"`
	Location          string    `json:"location"`
	State             string    `json:"state"`
	Simultaneous      bool      `json:"simultaneous"`
	GroupParticipants int       `json:"groupParticipants"`
	GroupMax          int       `json:"groupMaxParticipants"`
	SeenParticipants  int       `json:"seenParticipants"`
	Complete          bool      `json:"complete"`
	StillSelling      bool      `json:"stillSelling"`
	PhantomSeats      int       `json:"phantomSeats"`
	Sessions          []Session `json:"sessions"`
}

// Progress is what makes a stopped run resumable: the locations already walked
// and every grouped session found so far.
type Progress struct {
	StartedAt      time.Time `json:"startedAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
	LocationsDone  []string  `json:"locationsDone"`
	WorkshopsSeen  int       `json:"upcomingWorkshopsSeen"`
	GroupedSession []Session `json:"groupedSessions"`
}

// Audit is the report built from whatever has been captured so far.
type Audit struct {
	CapturedAt     time.Time `json:"capturedAt"`
	Complete       bool      `json:"walkComplete"`
	LocationsDone  int       `json:"locationsWalked"`
	LocationsTotal int       `json:"locationsTotal"`
	Workshops      int       `json:"upcomingWorkshopsSeen"`
	Groups         int       `json:"groupsFound"`
	AtCap          int       `json:"groupsAtCap"`
	OverCap        int       `json:"groupsOverCap"`
	HasRoom        int       `json:"groupsWithRoom"`
	Incomplete     int       `json:"groupsIncomplete"`
	PhantomSeats   int       `json:"phantomSeatsAtOrOverCap"`
	GroupList      []Group   `json:"groups"`
}

// Group states.
const (
	stateAtCap   = "at_cap"   // participants == cap: a close-group sweep would act on this
	stateOverCap = "over_cap" // participants > cap: already closed, or genuinely overbooked
	stateHasRoom = "has_room" // participants < cap: nothing to do
	stateNoCap   = "no_cap"   // grouped but no cap recorded
)

func main() {
	postURL := flag.String("post", "", "Webhook URL to send one course.workshop_updated per full group to. Omitted, nothing is sent.")
	clientID := flag.Int("clientId", 90743551, "clientId to send in webhook events")
	flag.Parse()

	for _, p := range []string{".env", "../.env", "../../.env", "../../../.env"} {
		if godotenv.Load(p) == nil {
			break
		}
	}

	apitoken := os.Getenv("AXCELERATE_APITOKEN")
	wstoken := os.Getenv("AXCELERATE_WSTOKEN")
	baseURL := os.Getenv("AXCELERATE_BASEURL")
	if apitoken == "" || wstoken == "" || baseURL == "" {
		log.Fatal("missing AXCELERATE_APITOKEN / AXCELERATE_WSTOKEN / AXCELERATE_BASEURL — no .env found in current or parent directories?")
	}

	client, err := axcelerate.NewClient(apitoken, wstoken, axcelerate.RateLimit(10), axcelerate.BaseURL(baseURL))
	if err != nil {
		log.Fatal(err)
	}

	if err := run(client, *postURL, *clientID); err != nil {
		log.Fatal(err)
	}
}

func run(client *axcelerate.Client, postURL string, clientID int) error {
	progress, err := loadProgress()
	if err != nil {
		return err
	}

	locs, resp, err := client.Courses.GetCoursesLocations(true, true)
	if err != nil {
		if resp != nil {
			return fmt.Errorf("%w: %s", err, resp.Body)
		}
		return err
	}

	var locations []string
	for _, l := range locs {
		locations = append(locations, string(l))
	}
	sort.Strings(locations)

	done := map[string]bool{}
	for _, l := range progress.LocationsDone {
		done[l] = true
	}

	if len(done) > 0 {
		log.Printf("resuming: %d of %d locations already walked, %d grouped sessions captured",
			len(done), len(locations), len(progress.GroupedSession))
	}
	log.Printf("%d locations, one call every %s", len(locations), callDelay)

	for n, loc := range locations {
		if done[loc] {
			continue
		}

		instances, err := upcomingWorkshops(client, loc)
		if err != nil {
			// Leave the location unmarked so the next run retries it, and keep
			// going: one bad location should not cost the whole walk.
			log.Printf("[%d/%d] %s: search failed, will retry on the next run: %v", n+1, len(locations), loc, err)
			continue
		}

		grouped := 0
		for _, i := range instances {
			if i.GroupedCourseID == nil || *i.GroupedCourseID <= 0 {
				continue
			}
			progress.GroupedSession = append(progress.GroupedSession, session(i))
			grouped++
		}

		progress.WorkshopsSeen += len(instances)
		progress.LocationsDone = append(progress.LocationsDone, loc)
		progress.UpdatedAt = time.Now()

		log.Printf("[%d/%d] %s: %d upcoming, %d grouped", n+1, len(locations), loc, len(instances), grouped)

		audit := buildAudit(progress, len(locations))
		if err := save(progress, audit); err != nil {
			return err
		}
	}

	audit := buildAudit(progress, len(locations))
	if err := save(progress, audit); err != nil {
		return err
	}

	report(audit)

	if postURL == "" {
		return nil
	}
	return closeFullGroups(audit, postURL, clientID)
}

// closeFullGroups sends one course.workshop_updated per full group.
//
// One event per group, not per workshop: the consumer searches the group and
// closes every sibling from a single message. It also re-reads Axcelerate
// before writing anything, so a stale audit cannot cause a wrong close — a
// group that has freed a place since the walk is simply left alone.
func closeFullGroups(audit Audit, postURL string, clientID int) error {
	list := actions(audit)

	groups := map[int]bool{}
	for _, a := range list {
		groups[a.Group.GroupedCourseID] = true
	}

	if len(list) >= maxActions {
		return fmt.Errorf("%d workshops need adjusting across %d groups, at or above the %d guard - nothing sent; review %s first",
			len(list), len(groups), maxActions, actionsPath)
	}

	if !audit.Complete {
		log.Printf("note: walk is partial (%d/%d locations), so this covers only what has been seen so far",
			audit.LocationsDone, audit.LocationsTotal)
	}

	log.Printf("closing %d group(s) covering %d workshop(s), one post every %s", len(groups), len(list), callDelay)

	// One representative session per group.
	posted := map[int]bool{}
	var sent, failed, seats int

	for _, a := range list {
		if posted[a.Group.GroupedCourseID] {
			seats += a.Session.MaxParticipants - a.Session.Participants
			continue
		}

		if sent+failed > 0 {
			time.Sleep(callDelay)
		}

		if err := postWorkshopEvent(postURL, clientID, a.Session.InstanceID); err != nil {
			log.Printf("group %d (%s %s): POST %d failed: %v", a.Group.GroupedCourseID, a.Group.Location, a.Group.GroupedCourseName, a.Session.InstanceID, err)
			failed++
			continue
		}

		posted[a.Group.GroupedCourseID] = true
		sent++
		seats += a.Session.MaxParticipants - a.Session.Participants
		log.Printf("group %d (%s %s %d/%d): sent via instance %d", a.Group.GroupedCourseID, a.Group.Location, a.Group.GroupedCourseName,
			a.Group.GroupParticipants, a.Group.GroupMax, a.Session.InstanceID)
	}

	fmt.Printf("\n%d group(s) sent for closing, %d failed, covering %d workshop(s) and %d unfulfillable seat(s)\n", sent, failed, len(list), seats)
	fmt.Printf("%s holds the caps as they were — it is the only record of them\n", actionsPath)
	return nil
}

// workshopEvent is the payload shape the webhook receiver expects.
type workshopEvent struct {
	Message struct {
		Workshop struct {
			ID int `json:"id"`
		} `json:"workshop"`
	} `json:"message"`
	Timestamp string `json:"timestamp"`
	ClientID  int    `json:"clientId"`
	MessageID string `json:"messageId"`
	Type      string `json:"type"`
}

func postWorkshopEvent(url string, clientID, instanceID int) error {
	var ev workshopEvent
	ev.Message.Workshop.ID = instanceID
	ev.Timestamp = time.Now().UTC().Format("2006-01-02T15:04:05.000Z")
	ev.ClientID = clientID
	ev.MessageID = uuidV4()
	ev.Type = "course.workshop_updated"

	body, err := json.Marshal(ev)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewReader(body))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("webhook returned %s: %s", resp.Status, bytes.TrimSpace(respBody))
	}
	return nil
}

func uuidV4() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		log.Fatal(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

func session(i axcelerate.Instance) Session {
	s := Session{
		InstanceID:        i.InstanceID,
		Name:              i.Name,
		CourseName:        i.CourseName,
		Code:              i.Code,
		Location:          i.Location,
		StartDate:         i.StartDate,
		Participants:      i.Participants,
		MaxParticipants:   i.MaxParticipants,
		Vacancy:           i.ParticipantVacancy,
		EnrolmentOpen:     i.EnrolmentOpen,
		IsActive:          i.IsActive,
		GroupedCourseID:   *i.GroupedCourseID,
		GroupParticipants: i.GroupedParticipants,
		GroupMax:          i.GroupedMaxParticipants,
		Simultaneous:      i.GroupedCourseSimultaneous,
	}
	if i.GroupedCourseName != nil {
		s.GroupedCourseName = *i.GroupedCourseName
	}
	return s
}

func buildAudit(progress Progress, locationsTotal int) Audit {
	audit := Audit{
		CapturedAt:     time.Now(),
		LocationsDone:  len(progress.LocationsDone),
		LocationsTotal: locationsTotal,
		Workshops:      progress.WorkshopsSeen,
		Complete:       len(progress.LocationsDone) == locationsTotal,
	}

	// Location searches match on %name%, so overlapping names can capture the
	// same instance more than once — and a resumed run may re-walk a location
	// whose save did not land.
	unique := map[int]Session{}
	for _, s := range progress.GroupedSession {
		unique[s.InstanceID] = s
	}

	byGroup := map[int][]Session{}
	for _, s := range unique {
		byGroup[s.GroupedCourseID] = append(byGroup[s.GroupedCourseID], s)
	}

	for gid, members := range byGroup {
		audit.GroupList = append(audit.GroupList, buildGroup(gid, members))
	}

	sort.Slice(audit.GroupList, func(a, b int) bool {
		if audit.GroupList[a].Location != audit.GroupList[b].Location {
			return audit.GroupList[a].Location < audit.GroupList[b].Location
		}
		return audit.GroupList[a].GroupedCourseID < audit.GroupList[b].GroupedCourseID
	})

	for _, g := range audit.GroupList {
		audit.Groups++
		switch g.State {
		case stateAtCap:
			audit.AtCap++
		case stateOverCap:
			audit.OverCap++
		case stateHasRoom:
			audit.HasRoom++
		}
		if !g.Complete {
			audit.Incomplete++
		}
		if g.State == stateAtCap || g.State == stateOverCap {
			audit.PhantomSeats += g.PhantomSeats
		}
	}

	return audit
}

func buildGroup(gid int, members []Session) Group {
	sort.Slice(members, func(a, b int) bool { return members[a].InstanceID < members[b].InstanceID })

	group := Group{GroupedCourseID: gid, Sessions: members}

	for _, m := range members {
		if group.Location == "" {
			group.Location = m.Location
		}
		if group.GroupedCourseName == "" {
			group.GroupedCourseName = m.GroupedCourseName
		}
		if m.Simultaneous {
			group.Simultaneous = true
		}
		// Every sibling carries the same group totals; the first that reports a
		// cap answers for all of them.
		if group.GroupMax == 0 && m.GroupMax > 0 {
			group.GroupMax = m.GroupMax
			group.GroupParticipants = m.GroupParticipants
		}

		group.SeenParticipants += m.Participants
		if vacancy := m.MaxParticipants - m.Participants; vacancy > 0 {
			group.PhantomSeats += vacancy
			if m.IsActive && m.StartDate.After(time.Now()) {
				group.StillSelling = true
			}
		}
	}

	switch {
	case group.GroupMax <= 0:
		group.State = stateNoCap
	case group.GroupParticipants == group.GroupMax:
		group.State = stateAtCap
	case group.GroupParticipants > group.GroupMax:
		group.State = stateOverCap
	default:
		group.State = stateHasRoom
	}

	// The sessions found should account for every enrolment the group reports.
	// When they do not, a sibling is missing from this capture — it sits at a
	// location not yet walked, or has already started — and this group's phantom
	// seat count is understated.
	group.Complete = group.SeenParticipants == group.GroupParticipants

	return group
}

func report(audit Audit) {
	fmt.Printf("\nGrouped course audit — %s\n", audit.CapturedAt.Format("2006-01-02 15:04"))
	fmt.Printf("%d/%d locations walked, %d upcoming workshops, %d grouped courses\n\n",
		audit.LocationsDone, audit.LocationsTotal, audit.Workshops, audit.Groups)

	for _, g := range audit.GroupList {
		if g.State == stateHasRoom || g.State == stateNoCap {
			continue
		}
		flag := ""
		if !g.Complete {
			flag = fmt.Sprintf("  ** INCOMPLETE: sessions found account for %d of %d enrolments **", g.SeenParticipants, g.GroupParticipants)
		}
		fmt.Printf("%-9s %-22s %-7d %s (%d/%d)%s\n", g.State, g.Location, g.GroupedCourseID, g.GroupedCourseName, g.GroupParticipants, g.GroupMax, flag)
		for _, s := range g.Sessions {
			fmt.Printf("    %d  %s  %-34s %2d/%-2d  vacancy %d\n", s.InstanceID, s.StartDate.Format("2006-01-02 15:04"), s.Name, s.Participants, s.MaxParticipants, s.Vacancy)
		}
		fmt.Println()
	}

	// An over-cap group with every session at its headcount has already been
	// closed. One still advertising seats has not - it is genuinely overbooked
	// and still selling, which no automated pass will touch: deciding who gets
	// turned away is a person's job.
	var problems []Group
	for _, g := range audit.GroupList {
		if g.State == stateOverCap && g.StillSelling {
			problems = append(problems, g)
		}
	}

	if len(problems) > 0 {
		fmt.Printf("\n--- NEEDS A HUMAN: over cap and still selling ---\n\n")
		for _, g := range problems {
			over := g.GroupParticipants - g.GroupMax
			fmt.Printf("%-22s %-7d %s  %d over cap (%d/%d), %d seat(s) still on sale\n",
				g.Location, g.GroupedCourseID, g.GroupedCourseName, over, g.GroupParticipants, g.GroupMax, g.PhantomSeats)
			for _, s := range g.Sessions {
				fmt.Printf("    %d  %s  %-34s %2d/%-2d  vacancy %d\n", s.InstanceID, s.StartDate.Format("2006-01-02 15:04"), s.Name, s.Participants, s.MaxParticipants, s.Vacancy)
			}
			fmt.Println()
		}
	}

	fmt.Printf("at cap %d   over cap %d (%d still selling)   with room %d   incomplete %d\n",
		audit.AtCap, audit.OverCap, len(problems), audit.HasRoom, audit.Incomplete)
	fmt.Printf("seats on sale that the group cannot honour: %d\n", audit.PhantomSeats)
	list := actions(audit)
	seats := 0
	for _, a := range list {
		seats += a.Session.MaxParticipants - a.Session.Participants
	}
	fmt.Printf("\n%d workshop(s) need adjusting, removing %d unfulfillable seat(s)\n", len(list), seats)
	fmt.Printf("written to %s, %s and %s\n", jsonPath, csvPath, actionsPath)

	if !audit.Complete {
		fmt.Printf("\nWalk is partial (%d/%d locations). Run again to continue, or delete %s to start over.\n",
			audit.LocationsDone, audit.LocationsTotal, progressPath)
	}
	if audit.Incomplete > 0 {
		fmt.Printf("\n%d group(s) are missing a sibling from this capture — treat their numbers as a floor.\n", audit.Incomplete)
	}
}

func loadProgress() (Progress, error) {
	b, err := os.ReadFile(progressPath)
	if errors.Is(err, fs.ErrNotExist) {
		return Progress{StartedAt: time.Now()}, nil
	}
	if err != nil {
		return Progress{}, err
	}

	var p Progress
	if err := json.Unmarshal(b, &p); err != nil {
		return Progress{}, fmt.Errorf("%s is unreadable (%w) — delete it to start a fresh walk", progressPath, err)
	}
	return p, nil
}

func save(progress Progress, audit Audit) error {
	if err := writeJSON(progressPath, progress); err != nil {
		return err
	}
	if err := writeJSON(jsonPath, audit); err != nil {
		return err
	}
	if err := writeCSV(csvPath, audit); err != nil {
		return err
	}
	return writeActions(actionsPath, audit)
}

// Action is one workshop that needs its cap adjusted, and to what.
type Action struct {
	Group   Group
	Session Session
}

// actions lists the sessions a close-group sweep would actually change.
//
// Only groups sitting exactly at their cap: a group reading over its cap has
// already been closed, and closing it again is what ratchets capacity away.
// Only sessions that are active, still to run, and advertising seats the group
// cannot honour — one already at its headcount needs no change.
func actions(audit Audit) []Action {
	now := time.Now()
	var out []Action

	for _, g := range audit.GroupList {
		if g.State != stateAtCap || !g.Complete {
			continue
		}
		for _, s := range g.Sessions {
			if !s.IsActive || !s.StartDate.After(now) {
				continue
			}
			if s.MaxParticipants <= s.Participants {
				continue
			}
			out = append(out, Action{Group: g, Session: s})
		}
	}
	return out
}

// writeJSON writes via a temp file and renames, so a run stopped mid-write
// leaves the previous good file rather than a truncated one.
func writeJSON(path string, v interface{}) error {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return err
	}
	tmp := path + ".tmp"
	if err := os.WriteFile(tmp, append(b, '\n'), 0o644); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

func writeCSV(path string, audit Audit) error {
	tmp := path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	header := []string{
		"state", "location", "groupedCourseID", "groupedCourseName",
		"groupParticipants", "groupMax", "complete",
		"instanceID", "startDate", "name", "code",
		"participants", "maxParticipants", "vacancy", "enrolmentOpen",
	}
	if err := w.Write(header); err != nil {
		f.Close()
		return err
	}

	for _, g := range audit.GroupList {
		for _, s := range g.Sessions {
			row := []string{
				g.State, g.Location, strconv.Itoa(g.GroupedCourseID), g.GroupedCourseName,
				strconv.Itoa(g.GroupParticipants), strconv.Itoa(g.GroupMax), strconv.FormatBool(g.Complete),
				strconv.Itoa(s.InstanceID), s.StartDate.Format(time.RFC3339), s.Name, s.Code,
				strconv.Itoa(s.Participants), strconv.Itoa(s.MaxParticipants), strconv.Itoa(s.Vacancy),
				strconv.FormatBool(s.EnrolmentOpen),
			}
			if err := w.Write(row); err != nil {
				f.Close()
				return err
			}
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}

// upcomingWorkshops returns all workshop instances at a location starting from
// today, paging until exhausted.
func upcomingWorkshops(client *axcelerate.Client, location string) ([]axcelerate.Instance, error) {
	const pageSize = 100
	var all []axcelerate.Instance

	for offset := 0; ; offset += pageSize {
		parms := map[string]string{
			"type":          "w",
			"location":      location,
			"startDate_min": time.Now().Format("2006-01-02"),
			"isActive":      "true",
			"offset":        strconv.Itoa(offset),
			"displayLength": strconv.Itoa(pageSize),
		}

		time.Sleep(callDelay)

		page, resp, err := client.Courses.GetCoursesInstanceSearch(parms)
		if err != nil {
			if resp != nil {
				return nil, fmt.Errorf("%w: %s", err, resp.Body)
			}
			return nil, err
		}

		all = append(all, page...)
		if len(page) < pageSize {
			return all, nil
		}
	}
}

// writeActions writes the to-do list: one row per workshop needing its cap
// adjusted, with the value it has now and the value it should have. The "from"
// column is the record to restore from — closing a group overwrites it in
// Axcelerate and nothing else keeps it.
func writeActions(path string, audit Audit) error {
	tmp := path + ".tmp"
	f, err := os.Create(tmp)
	if err != nil {
		return err
	}

	w := csv.NewWriter(f)
	header := []string{
		"location", "groupedCourseID", "groupedCourseName", "groupParticipants", "groupMax",
		"instanceID", "startDate", "name", "code",
		"participants", "maxFrom", "maxTo", "seatsRemoved",
	}
	if err := w.Write(header); err != nil {
		f.Close()
		return err
	}

	for _, a := range actions(audit) {
		g, s := a.Group, a.Session
		row := []string{
			g.Location, strconv.Itoa(g.GroupedCourseID), g.GroupedCourseName,
			strconv.Itoa(g.GroupParticipants), strconv.Itoa(g.GroupMax),
			strconv.Itoa(s.InstanceID), s.StartDate.Format(time.RFC3339), s.Name, s.Code,
			strconv.Itoa(s.Participants), strconv.Itoa(s.MaxParticipants), strconv.Itoa(s.Participants),
			strconv.Itoa(s.MaxParticipants - s.Participants),
		}
		if err := w.Write(row); err != nil {
			f.Close()
			return err
		}
	}

	w.Flush()
	if err := w.Error(); err != nil {
		f.Close()
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return os.Rename(tmp, path)
}
