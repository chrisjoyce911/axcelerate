package axcelerate

import (
	"encoding/json"
	"fmt"
)

type VerifyCertificate struct {
	Document struct {
		Havefile bool `json:"HAVEFILE"`
		Detail   struct {
			Awardtype    string `json:"AWARDTYPE"`
			Contactid    int    `json:"CONTACTID"`
			Surname      string `json:"SURNAME"`
			Enrolid      int    `json:"ENROLID"`
			Issuedby     string `json:"ISSUEDBY"`
			Activityname string `json:"ACTIVITYNAME"`
			Givenname    string `json:"GIVENNAME"`
		} `json:"DETAIL"`
	} `json:"DOCUMENT"`
	Result bool   `json:"RESULT"`
	Msg    string `json:"MSG"`
}

// ContactVerifyCertificate Checks the system for a certificate issued by the Client.
func (s *ContactService) ContactVerifyCertificate(certificateID string) (VerifyCertificate, *Response, error) {
	var obj VerifyCertificate

	resp, err := do(s.client, "GET", Params{parms: map[string]string{}, u: fmt.Sprintf("/contact/enrolment/verifyCertificate/%s", certificateID)}, obj)

	if err != nil {
		return obj, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}
