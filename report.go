package axcelerate

import (
	"fmt"
	"log"

	jsontime "github.com/liamylian/jsontime/v2/v2"

	"encoding/json"
)

// CoursesService handles all interactions with Contact
type ReportService struct {
	client *Client
}

type SavedReport struct {
	ReportType      string          `json:"REPORTTYPE"`
	Data            interface{}     `json:"DATA"`
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

type ReportList []struct {
	ReportName         string `json:"REPORTNAME"`
	CreatedByContactID int    `json:"CREATEDBYCONTACTID"`
	CreatedBy          string `json:"CREATEDBY"`
	ReportID           int    `json:"REPORTID"`
	Reference          string `json:"REPORTREFERENCE"`
	Active             bool   `json:"ACTIVE"`
	ReportVersion      string `json:"REPORTVERSION"`
	Description        string `json:"DESCRIPTION"`
}

// SavedReportRun Interacts with the aXcelerate Report Builder to run a specified saved report. It will return the results in JSON format.
// Header			Type		Required	Default	Description
// filterOverride	json		false		Override the standard filter value(s):
//											e.g: To override the Enrolment Status Value you can pass:
//											[{NAME:'enrolments.enrolstatus', VALUE:'canc', VALUE2:''}]
//											NOTE: You cannot test/use this in our API relax console simulator
// offsetRows		numeric		true		Warehoused reports only: The number of report rows to skip before returning row objects in the data[] response element. The maximum supported value is 100,000. Reports larger than this are not supported

func (s *ReportService) SavedReportRun(reportID int, parms map[string]string) (SavedReport, *Response, error) {
	var obj SavedReport

	fmt.Printf("SavedReportRun : %+v", parms)

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	parms["reportID"] = fmt.Sprintf("%d", reportID)

	fmt.Printf("SavedReportRun : %+v", parms)

	for key, value := range parms {
		log.Printf("Key: %s, Value: %s", key, value)
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

// SavedReportList Returns an array of available reports, their names and descriptions.
func (s *ReportService) SavedReportList() (ReportList, *Response, error) {
	var obj ReportList

	parms := map[string]string{}

	url := "/report/saved/list"
	resp, err := do(s.client, "GET", Params{parms: parms, u: url}, obj)

	if err != nil {
		return obj, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}

// SavedReportWarehousedRun Interacts with the aXcelerate Report Builder to run a specified saved report. It will return the results in JSON format.
// Header			Type		Required	Default	Description
// filterOverride	json		false		Override the standard filter value(s):
//											e.g: To override the Enrolment Status Value you can pass:
//											[{NAME:'enrolments.enrolstatus', VALUE:'canc', VALUE2:''}]
//											NOTE: You cannot test/use this in our API relax console simulator
// offsetRows		numeric		true		Warehoused reports only: The number of report rows to skip before returning row objects in the data[] response element. The maximum supported value is 100,000. Reports larger than this are not supported

func (s *ReportService) SavedReportWarehousedRun(reportID int, displayLength int, parms map[string]string) (SavedReport, *Response, error) {
	var obj SavedReport

	parms["reportID"] = fmt.Sprintf("%d", reportID)
	parms["displayLength"] = fmt.Sprintf("%d", displayLength)

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	for key, value := range parms {
		log.Printf("Key: %s, Value: %s", key, value)
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
