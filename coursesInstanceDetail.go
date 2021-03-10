package axcelerate

import (
	"encoding/json"
	"fmt"
)

// InstanceDetail details of an activity instance.
type InstanceDetail struct {
	Complexdates        []interface{} `json:"COMPLEXDATES"`
	Cost                int64         `json:"COST"`
	CustomfieldWeekends interface{}   `json:"CUSTOMFIELD_WEEKENDS"`
	Datedescriptor      string        `json:"DATEDESCRIPTOR"`
	Enrolmentopen       bool          `json:"ENROLMENTOPEN"`
	Finishdate          string        `json:"FINISHDATE"`
	ID                  int64         `json:"ID"`
	Instanceid          int64         `json:"INSTANCEID"`
	Items               []interface{} `json:"ITEMS"`
	Linkedclassid       int64         `json:"LINKEDCLASSID"`
	Linkedelearning     []struct {
		Code       string      `json:"CODE"`
		Enddate    interface{} `json:"ENDDATE"`
		Instanceid int64       `json:"INSTANCEID"`
		Name       string      `json:"NAME"`
		Startdate  interface{} `json:"STARTDATE"`
	} `json:"LINKEDELEARNING"`
	Location           string      `json:"LOCATION"`
	Maxparticipants    int64       `json:"MAXPARTICIPANTS"`
	Minparticipants    int64       `json:"MINPARTICIPANTS"`
	Name               string      `json:"NAME"`
	Notices            interface{} `json:"NOTICES"`
	Ownercontactid     int64       `json:"OWNERCONTACTID"`
	Participants       int64       `json:"PARTICIPANTS"`
	Participantvacancy int64       `json:"PARTICIPANTVACANCY"`
	Public             bool        `json:"PUBLIC"`
	Startdate          string      `json:"STARTDATE"`
	Trainercontactid   int64       `json:"TRAINERCONTACTID"`
	Venuecontactid     int64       `json:"VENUECONTACTID"`
}

/*
GetCoursesInstanceDetail Returns details of an activity instance.

instanceID
	The instanceID of the activity you want details from.
activityType
	The type of the activity. w = workshop, p = accredited program, el = e-learning.
*/
func (s *CoursesService) GetCoursesInstanceDetail(instanceID int, activityType string) (*InstanceDetail, *Response, error) {
	a := new(InstanceDetail)

	parms := map[string]string{}

	parms["instanceID"] = fmt.Sprintf("%d", instanceID)
	parms["type"] = activityType

	resp, err := do(s.client, "GET", Params{parms: parms, u: "/course/instance/detail"}, a)

	if err != nil {
		return nil, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &a)
	return a, resp, err
}
