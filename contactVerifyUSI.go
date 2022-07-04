package axcelerate

import (
	"fmt"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

func (s *ContactService) VerifyUSI(contactID int) (USIstatus, *Response, error) {
	var obj USIstatus

	parms := map[string]string{
		"contactID": fmt.Sprintf("%d", contactID),
	}

	resp, err := do(s.client, "POST", Params{parms: parms, u: "/contact/verifyUSI"}, obj)

	if err != nil {
		return obj, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_datetime", "2006-01-02 15:04:05")

	err = json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}
