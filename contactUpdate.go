package axcelerate

import (
	"fmt"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

/*
ContactUpdate used to update a contact.

Interacts with a specfic contact. You can update (PUT) contact details. The including the parameters below.

Request Parameters
givenName

	Given (first) name

surname

	Surname (last name)

emailAddress

	Email address. Must be a valid email address

DOB

	yyyy-mm-dd formatted date

//
*/
func (s *ContactService) ContactUpdate(contactID int, parms map[string]string) (Contact, *Response, error) {
	var obj Contact

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	resp, err := do(s.client, "PUT", Params{parms: parms, u: fmt.Sprintf("/contact/%d", contactID)}, obj)

	if err != nil {
		return obj, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_time", "15:04")
	jsontime.AddTimeFormatAlias("axc_time_long", "15:04:05")
	jsontime.AddTimeFormatAlias("axc_date", "2006-01-02")
	jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")

	err = json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}
