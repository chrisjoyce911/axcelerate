package axcelerate

import (
	"fmt"
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// TrainerService handles all interactions with Trainer
type TrainerService struct {
	client *Client
}

type Trainer struct {
	CONSULTANTID                 int         `json:"CONSULTANTID"`
	CONTACTHRSTIMEFRAME          string      `json:"CONTACTHRSTIMEFRAME"`
	EMPTYPE                      string      `json:"EMPTYPE"`
	DOMAINID                     interface{} `json:"DOMAINID"`
	CONENTRYDATE                 *time.Time  `json:"CONENTRYDATE" time_format:"axc_date_hours"`
	CONSULTANTUPDATEDDATETIMEUTC *time.Time  `json:"CONSULTANTUPDATEDDATETIMEUTC" time_format:"axc_date_hours"`
	CONTACTHRS                   int         `json:"CONTACTHRS"`
	CONTACTID                    int         `json:"CONTACTID"`
	IMAGE                        string      `json:"IMAGE"`
	RATEAMOUNT                   string      `json:"RATEAMOUNT"`
	NAME                         string      `json:"NAME"`
	ACTIVE                       bool        `json:"ACTIVE"`
	EXPERIENCE                   string      `json:"EXPERIENCE"`
	PAYTIMEFRAME                 string      `json:"PAYTIMEFRAME"`
	FILLINONLY                   bool        `json:"FILLINONLY"`
	PAGE                         string      `json:"PAGE"`
	EMAIL                        string      `json:"EMAIL"`
}

// GetTrainer Interacts with a specific trainer.
func (s *TrainerService) GetTrainer(trainerID int) (*Trainer, *Response, error) {
	var a Trainer

	resp, err := do(s.client, "GET", Params{u: fmt.Sprintf("/trainer/%d", trainerID)}, a)
	if err != nil {
		return nil, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")
	jsontime.AddTimeFormatAlias("axc_date", "2006-01-02")

	err = json.Unmarshal([]byte(resp.Body), &a)

	return &a, resp, err
}
