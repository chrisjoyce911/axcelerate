package axcelerate

import (
	"fmt"
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// InstanceDetail details of an activity instance.
type InstanceDetail struct {
	ComplexDates        []interface{} `json:"COMPLEXDATES"`
	Cost                int64         `json:"COST"`
	CustomfieldWeekends interface{}   `json:"CUSTOMFIELD_WEEKENDS"`
	DateDescriptor      string        `json:"DATEDESCRIPTOR"`
	EnrolmentOpen       bool          `json:"ENROLMENTOPEN"`
	FinishDate          time.Time     `json:"FINISHDATE"`
	ID                  int64         `json:"ID"`
	InstanceID          int64         `json:"INSTANCEID"`
	Items               []interface{} `json:"ITEMS"`
	LinkedClassID       int64         `json:"LINKEDCLASSID"`
	LinkedeLearning     []struct {
		Code       string      `json:"CODE"`
		Enddate    interface{} `json:"ENDDATE"`
		Instanceid int64       `json:"INSTANCEID"`
		Name       string      `json:"NAME"`
		Startdate  interface{} `json:"STARTDATE"`
	} `json:"LINKEDELEARNING"`
	Location           string      `json:"LOCATION"`
	MaxParticipants    int64       `json:"MAXPARTICIPANTS"`
	MinParticipants    int64       `json:"MINPARTICIPANTS"`
	Name               string      `json:"NAME"`
	Notices            interface{} `json:"NOTICES"`
	OwnerContactID     int64       `json:"OWNERCONTACTID"`
	Participants       int64       `json:"PARTICIPANTS"`
	ParticipantVacancy int64       `json:"PARTICIPANTVACANCY"`
	Public             bool        `json:"PUBLIC"`
	StartDate          time.Time   `json:"STARTDATE"`
	TrainerContactID   int64       `json:"TRAINERCONTACTID"`
	VenueContactID     int64       `json:"VENUECONTACTID"`
}

/*
GetCoursesInstanceDetail Returns details of an activity instance.

instanceID
	The instanceID of the activity you want details from.
activityType
	The type of the activity. w = workshop, p = accredited program, el = e-learning.
*/
func (s *CoursesService) GetCoursesInstanceDetail(instanceID int, activityType string) (InstanceDetail, *Response, error) {
	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.SetDefaultTimeFormat("2006-01-02 15:04", time.Local)

	// jsontime.AddTimeFormatAlias("axcelerate_datetime", "2006-01-02 15:04")
	time.LoadLocation("Asia/Shanghai")
	var obj InstanceDetail

	parms := map[string]string{}

	parms["instanceID"] = fmt.Sprintf("%d", instanceID)
	parms["type"] = activityType

	resp, err := do(s.client, "GET", Params{parms: parms, u: "/course/instance/detail"}, obj)

	if err != nil {
		return obj, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
