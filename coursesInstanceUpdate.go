package axcelerate

import (
	"encoding/json"
	"fmt"
)

// UpdateInstanceDetail details of an activity instance.
type UpdateInstanceDetail struct {
	Messages   string `json:"MESSAGES"`
	Message    string `json:"MESSAGE"`
	Cost       int    `json:"COST"`
	InstanceID int    `json:"INSTANCEID"`
	Status     string `json:"STATUS"`
	Data       string `json:"DATA"`
	Error      bool   `json:"ERROR"`
	Code       string `json:"CODE"`
	Details    string `json:"DETAILS"`
}

/*
UpdateInstanceCost Updates existing instance cost per student

instanceID
	The instanceID of the activity you want details from.
activityType
	The type of the activity. w = workshop, p = accredited program, el = e-learning.
cost
	The Cost Per Participant / Student.
*/
func (s *CoursesService) UpdateInstanceCost(instanceID int, activityType string, cost int) (UpdateInstanceDetail, *Response, error) {

	var obj UpdateInstanceDetail

	parms := map[string]string{}

	parms["ID"] = fmt.Sprintf("%d", instanceID)
	parms["type"] = activityType
	parms["cost"] = fmt.Sprintf("%d", cost)

	resp, err := do(s.client, "PUT", Params{parms: parms, u: "/course/instance/"}, obj)

	if err != nil {
		return obj, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
