package axcelerate

import (
	"fmt"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// ContactEnrolments Returns enrolments for a specific contact
func (s *ContactService) ContactEnrolments(contactID int, parms map[string]string) ([]ContactEnrolment, *Response, error) {
	var obj []ContactEnrolment

	resp, err := do(s.client, "GET", Params{parms: parms, u: fmt.Sprintf("/contact/enrolments/%d", contactID)}, obj)

	if err != nil {
		return obj, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_time", "15:04")
	jsontime.AddTimeFormatAlias("axc_date", "2006-01-02")

	json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}
