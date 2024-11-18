package axcelerate

import (
	"fmt"
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// InstanceDetail details of an activity instance.
type InstanceDetail struct {
	ComplexDates        []ComplexDate `json:"COMPLEXDATES"`
	Cost                int64         `json:"COST"`
	CustomfieldWeekends interface{}   `json:"CUSTOMFIELD_WEEKENDS"`
	DateDescriptor      string        `json:"DATEDESCRIPTOR"`
	EnrolmentOpen       bool          `json:"ENROLMENTOPEN"`
	FinishDate          time.Time     `json:"FINISHDATE"`
	CourseID            int           `json:"ID"`
	InstanceID          int           `json:"INSTANCEID"`
	Items               []interface{} `json:"ITEMS"`
	LinkedClassID       int           `json:"LINKEDCLASSID"`
	LinkedeLearning     []struct {
		Code       string      `json:"CODE"`
		Enddate    interface{} `json:"ENDDATE"`
		Instanceid int64       `json:"INSTANCEID"`
		Name       string      `json:"NAME"`
		Startdate  interface{} `json:"STARTDATE"`
	} `json:"LINKEDELEARNING"`
	Location           string      `json:"LOCATION"`
	MaxParticipants    int         `json:"MAXPARTICIPANTS"`
	MinParticipants    int         `json:"MINPARTICIPANTS"`
	Name               string      `json:"NAME"`
	Notices            interface{} `json:"NOTICES"`
	OwnerContactID     int         `json:"OWNERCONTACTID"`
	Participants       int         `json:"PARTICIPANTS"`
	ParticipantVacancy int         `json:"PARTICIPANTVACANCY"`
	Public             bool        `json:"PUBLIC"`
	StartDate          time.Time   `json:"STARTDATE"`
	TrainerContactID   int         `json:"TRAINERCONTACTID"`
	VenueContactID     int         `json:"VENUECONTACTID"`
	Status             string      `json:"STATUS"`
	SyncDateDescriptor bool        `json:"SYNCDATEDESCRIPTOR"`
	LastUpdatedUTC     time.Time   `json:"LASTUPDATEDUTC"`
	Address            string      `json:"ADDRESS"`
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

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
