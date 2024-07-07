package axcelerate

import (
	"fmt"
)

/*
UpdateCourseEnrolment Update an enrolment

# Request Parameters

contactID

	The ID of the Contact.

activityInstanceID

	The ID of the Activity Instance, this is not the workshop instance (Not required when using subjectCode)

type

	The type of the activity. w = workshop, p = accredited program, el = e-learning. Only p & w work at this time
*/
func (s *CoursesService) UpdateCourseEnrolment(contactID, activityInstanceID int, activityType string, parms map[string]string) ([]Enrolment, *Response, error) {
	var obj []Enrolment

	parms["type"] = activityType
	parms["contactID"] = fmt.Sprintf("%d", contactID)
	parms["instanceID"] = fmt.Sprintf("%d", activityInstanceID)

	resp, err := do(s.client, "PUT", Params{parms: parms, u: "/course/enrolment"}, obj)

	if err != nil {
		return obj, resp, err
	}

	return obj, resp, err
}
