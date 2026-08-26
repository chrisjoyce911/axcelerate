package main

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/chrisjoyce911/axcelerate"
	"github.com/joho/godotenv"
)

// workshopEvent is the payload shape expected by the webhook receiver.
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

func main() {
	// Find .env in the current directory or any parent (it lives at the repo root).
	for _, p := range []string{".env", "../.env", "../../.env", "../../../.env"} {
		if godotenv.Load(p) == nil {
			break
		}
	}

	location := flag.String("location", "", "Location to search (matches %name%)")
	instanceID := flag.Int("instance", 0, "Print full details for a single InstanceID")
	postURL := flag.String("post", "", "Webhook URL to POST course.workshop_updated events to")
	clientID := flag.Int("clientId", 90743551, "clientId to send in webhook events")
	locations := flag.Bool("locations", false, "List all public locations with future activity")
	all := flag.Bool("all", false, "Full run: POST events for every upcoming workshop at every location")
	delay := flag.Duration("delay", 60*time.Second, "Pause between webhook POSTs during a full run")
	start := flag.String("start", "", "Full run: skip locations before this one (alphabetical, case-insensitive)")
	flag.Parse()

	if *location == "" && *instanceID == 0 && !*locations && !*all {
		log.Fatal("usage: go run ./example/dataCatchup -location <name> | -instance <id> | -locations | -all -post <url>")
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

	if *all {
		if err := fullRun(client, *postURL, *clientID, *delay, *start); err != nil {
			log.Fatal(err)
		}
		return
	}

	if *locations {
		locs, resp, err := client.Courses.GetCoursesLocations(true, true)
		if err != nil {
			if resp != nil {
				log.Fatalf("%v: %s", err, resp.Body)
			}
			log.Fatal(err)
		}
		fmt.Printf("Locations (public, future activity): %d\n\n", len(locs))
		for _, l := range locs {
			fmt.Println(l)
		}
		return
	}

	if *instanceID != 0 {
		if *postURL != "" {
			if err := postWorkshopEvent(*postURL, *clientID, *instanceID); err != nil {
				log.Fatal(err)
			}
			return
		}
		if err := printInstance(client, *instanceID); err != nil {
			log.Fatal(err)
		}
		return
	}

	instances, err := upcomingWorkshops(client, *location)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Upcoming workshops for location %q: %d\n\n", *location, len(instances))
	for _, i := range instances {
		fmt.Printf("%d\t%s\t%s\t%d/%d\t%s\n", i.InstanceID, i.StartDate.Format("2006-01-02 15:04"), i.Code, i.Participants, i.MaxParticipants, i.Name)
	}

	if *postURL != "" {
		fmt.Println()
		for _, i := range instances {
			if err := postWorkshopEvent(*postURL, *clientID, i.InstanceID); err != nil {
				log.Printf("POST %d failed: %v", i.InstanceID, err)
			}
		}
	}
}

// fullRun posts course.workshop_updated events for every upcoming workshop at every
// public location, pausing between locations. With no post URL it only reports counts.
func fullRun(client *axcelerate.Client, postURL string, clientID int, delay time.Duration, start string) error {
	locs, resp, err := client.Courses.GetCoursesLocations(true, true)
	if err != nil {
		if resp != nil {
			return fmt.Errorf("%w: %s", err, resp.Body)
		}
		return err
	}

	if start != "" {
		from := -1
		for n, loc := range locs {
			if strings.EqualFold(strings.TrimSpace(string(loc)), strings.TrimSpace(start)) {
				from = n
				break
			}
		}
		if from == -1 {
			return fmt.Errorf("start location %q not found in the %d locations", start, len(locs))
		}
		log.Printf("Starting at %s, skipping %d locations", locs[from], from)
		locs = locs[from:]
	}

	log.Printf("Full run: %d locations, %s between POSTs", len(locs), delay)
	totalWorkshops, totalPosted, totalFailed := 0, 0, 0

	for n, loc := range locs {
		instances, err := upcomingWorkshops(client, string(loc))
		if err != nil {
			log.Printf("[%d/%d] %s: search failed: %v", n+1, len(locs), loc, err)
			totalFailed++
			continue
		}

		log.Printf("[%d/%d] %s: %d upcoming workshops", n+1, len(locs), loc, len(instances))
		totalWorkshops += len(instances)

		if postURL == "" {
			continue
		}

		for _, i := range instances {
			if totalPosted+totalFailed > 0 {
				time.Sleep(delay)
			}
			if err := postWorkshopEvent(postURL, clientID, i.InstanceID); err != nil {
				log.Printf("POST %d failed: %v", i.InstanceID, err)
				totalFailed++
				continue
			}
			totalPosted++
		}
	}

	log.Printf("Full run complete: %d workshops, %d posted, %d failures", totalWorkshops, totalPosted, totalFailed)
	return nil
}

// postWorkshopEvent sends a course.workshop_updated event for one workshop instance.
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
	fmt.Printf("POST workshop %d -> %s %s\n", instanceID, resp.Status, bytes.TrimSpace(respBody))

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return fmt.Errorf("webhook returned %s", resp.Status)
	}
	return nil
}

// uuidV4 generates a random UUID v4 string without external dependencies.
func uuidV4() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		log.Fatal(err)
	}
	b[6] = (b[6] & 0x0f) | 0x40
	b[8] = (b[8] & 0x3f) | 0x80
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:16])
}

// printInstance fetches a single workshop instance by InstanceID and prints it as JSON.
func printInstance(client *axcelerate.Client, instanceID int) error {
	parms := map[string]string{
		"type":       "w",
		"InstanceID": fmt.Sprintf("%d", instanceID),
	}

	instances, resp, err := client.Courses.GetCoursesInstanceSearch(parms)
	if err != nil {
		if resp != nil {
			return fmt.Errorf("%w: %s", err, resp.Body)
		}
		return err
	}
	if len(instances) == 0 {
		return fmt.Errorf("no instance found for InstanceID %d", instanceID)
	}

	out, err := json.MarshalIndent(instances[0], "", "  ")
	if err != nil {
		return err
	}
	fmt.Println(string(out))
	return nil
}

// upcomingWorkshops returns all workshop instances at a location starting from today, paging until exhausted.
func upcomingWorkshops(client *axcelerate.Client, location string) ([]axcelerate.Instance, error) {
	const pageSize = 100
	var all []axcelerate.Instance

	for offset := 0; ; offset += pageSize {
		parms := map[string]string{
			"type":          "w",
			"location":      location,
			"startDate_min": time.Now().Format("2006-01-02"),
			"isActive":      "true",
			"offset":        fmt.Sprintf("%d", offset),
			"displayLength": fmt.Sprintf("%d", pageSize),
		}

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
