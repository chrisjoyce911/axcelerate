package axcelerate

import (
	"fmt"
)

type EnrolmentUpdate struct {
	Data     string `json:"DATA"`
	Error    bool   `json:"ERROR"`
	Messages string `json:"MESSAGES"`
	Code     string `json:"CODE"`
	Details  string `json:"DETAILS"`
}

/*
CourseEnrolmentUpdate Update an enrolment

# Request Parameters

contactID

	The ID of the Contact.

activityInstanceID

	The ID of the Activity Instance, this is not the workshop instance (Not required when using subjectCode)

type

	The type of the activity. w = workshop, p = accredited program, el = e-learning. Only p & w work at this time
*/
func (s *CoursesService) CourseEnrolmentUpdate(contactID, instanceID int, activityType string, parms map[string]string) (*EnrolmentUpdate, *Response, error) {
	var obj EnrolmentUpdate

	parms["type"] = activityType
	parms["contactID"] = fmt.Sprintf("%d", contactID)
	parms["instanceID"] = fmt.Sprintf("%d", instanceID)

	resp, err := do(s.client, "PUT", Params{parms: parms, u: "/course/enrolment"}, obj)

	if err != nil {
		return nil, resp, err
	}

	return &EnrolmentUpdate{Data: "ok", Error: false, Messages: "ok", Code: resp.Status, Details: "ok"}, resp, err
}
