package axcelerate

import (
	"fmt"
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// CoursesService handles all interactions with Contact
type ReportService struct {
	client *Client
}

type SavedReport struct {
	ReportType      string          `json:"REPORTTYPE"`
	Data            []ReportData    `json:"DATA"`
	ErrorMsg        string          `json:"ERRORMSG"`
	Count           int             `json:"COUNT"`
	ReportName      string          `json:"REPORTNAME"`
	RowLimit        int             `json:"ROWLIMIT"`
	ReportID        int             `json:"REPORTID"`
	Path            string          `json:"PATH"`
	Success         bool            `json:"SUCCESS"`
	ContentType     string          `json:"CONTENTTYPE"`
	ReportReference string          `json:"REPORTREFERENCE"`
	ReportVersion   string          `json:"REPORTVERSION"`
	Filters         []ReportFilters `json:"FILTERS"`
	Description     string          `json:"DESCRIPTION"`
}

type ReportFilters struct {
	Value2   string `json:"VALUE2"`
	Operator string `json:"OPERATOR"`
	Display  string `json:"DISPLAY"`
	Name     string `json:"NAME"`
	Value    string `json:"VALUE"`
}

type ReportData struct {
	ContactID       int       `json:"CONTACTID"`
	EmailAddress    string    `json:"EMAILADDRESS"`
	AwardCourseCode string    `json:"AWARD_COURSECODE"`
	CertificateNo   string    `json:"CERTIFICATENO"`
	AwardQualName   string    `json:"AWARD_QUALNAME"`
	IssuedDate      time.Time `json:"AWARDISSUEDDATE" time_format:"axc_date"`
}

// GetInvoice will get the details of an invoice / Update an invoice (NOTE: Currently you can only Lock Items and Finalise)
// Header			Type		Required	Default	Description
// filterOverride	json		false		Override the standard filter value(s):
//											e.g: To override the Enrolment Status Value you can pass:
//											[{NAME:'enrolments.enrolstatus', VALUE:'canc', VALUE2:''}]
//											NOTE: You cannot test/use this in our API relax console simulator
// offsetRows		numeric		true		Warehoused reports only: The number of report rows to skip before returning row objects in the data[] response element. The maximum supported value is 100,000. Reports larger than this are not supported

func (s *ReportService) SavedReportRun(reportID int, displayLength int, parms map[string]string) (SavedReport, *Response, error) {
	var obj SavedReport

	parms["reportID"] = fmt.Sprintf("%d", reportID)
	parms["displayLength"] = fmt.Sprintf("%d", displayLength)

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	url := "/report/saved/run"
	resp, err := do(s.client, "POST", Params{parms: parms, u: url}, obj)

	if err != nil {
		return obj, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat

	jsontime.AddTimeFormatAlias("axc_date", "2006-01-02")

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
