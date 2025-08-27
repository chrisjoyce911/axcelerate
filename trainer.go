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
	ConsultantID                 int         `json:"CONSULTANTID"`
	ContactHrsTimeFrame          string      `json:"CONTACTHRSTIMEFRAME"`
	EmpType                      string      `json:"EMPTYPE"`
	DomainID                     interface{} `json:"DOMAINID"`
	ConEntryDate                 *time.Time  `json:"CONENTRYDATE" time_format:"axc_date_hours"`
	ConsultantUpdatedDateTimeUTC *time.Time  `json:"CONSULTANTUPDATEDDATETIMEUTC" time_format:"axc_date_hours"`
	ContactHrs                   int         `json:"CONTACTHRS"`
	ContactID                    int         `json:"CONTACTID"`
	Image                        string      `json:"IMAGE"`
	RateAmount                   string      `json:"RATEAMOUNT"`
	Name                         string      `json:"NAME"`
	Active                       bool        `json:"ACTIVE"`
	Experience                   string      `json:"EXPERIENCE"`
	PayTimeFrame                 string      `json:"PAYTIMEFRAME"`
	FillInOnly                   bool        `json:"FILLINONLY"`
	Page                         string      `json:"PAGE"`
	Email                        string      `json:"EMAIL"`
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
