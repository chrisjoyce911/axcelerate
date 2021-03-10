package axcelerate

import (
	"encoding/json"
	"fmt"
)

// Attendeds for a Course Instance.
type Attendeds struct {
	Complexid  interface{} `json:"COMPLEXID"`
	Date       string      `json:"DATE"`
	Duration   interface{} `json:"DURATION"`
	Enrollees  []Attended  `json:"ENROLLEES"`
	Finishtime string      `json:"FINISHTIME"`
	Instanceid int64       `json:"INSTANCEID"`
	Starttime  string      `json:"STARTTIME"`
	Type       string      `json:"TYPE"`
}

// Attended for a Course Instance.
type Attended struct {
	Attendanceid     interface{} `json:"ATTENDANCEID"`
	Attendedduration interface{} `json:"ATTENDEDDURATION"`
	Attendedflag     interface{} `json:"ATTENDEDFLAG"`
	Comment          string      `json:"COMMENT"`
	Completedflag    bool        `json:"COMPLETEDFLAG"`
	Contactid        int64       `json:"CONTACTID"`
	Enrolid          int64       `json:"ENROLID"`
	Finishtime       string      `json:"FINISHTIME"`
	Givenname        string      `json:"GIVENNAME"`
	Starttime        string      `json:"STARTTIME"`
	Surname          string      `json:"SURNAME"`
}

/*
GetCoursesAttendeds returns a list of instances

instanceID
	The Instance ID of the activity you want to retrieve the attendance for (currently only works for workshops [type=w])

Request Parameters

type
	The activity type of the activity. w = workshop [ not currently availabile: p = accredited program & el = e-learning].
contactID
	The contactID of the Enrollee. (required when updating)
attended
	The attended flag, 1 = attended, 0 = did not attend. (required when updating)
complexID
	The complexID of the Session. (required when updating)
duration
	This field is no longer supported for PUT or POST requests. aXcelerate will calculate the session duration from the Start and Finish times.
arrival
	The date & time the student arrived (yyyy-mm-dd hh:mm)
departure
	The date & time the student departed (yyyy-mm-dd hh:mm)
comment
	Associate a comment with this attendance record.
*/
func (s *CoursesService) GetCoursesAttendeds(instanceID int, parms map[string]string) (*[]Attendeds, *Response, error) {
	a := new([]Attendeds)

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	parms["instanceID"] = fmt.Sprintf("%d", instanceID)
	parms["type"] = "w"

	resp, err := do(s.client, "GET", Params{parms: parms, u: "/course/instance/attendance"}, a)

	if err != nil {
		return nil, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &a)
	return a, resp, err
}
