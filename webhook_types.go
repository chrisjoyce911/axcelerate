package axcelerate

import "time"

type WebhookPayload struct {
	Type    string `json:"type,omitempty"`
	Message struct {
		Enrolment struct {
			Student struct {
				ContactID int `json:"contactId,omitempty"`
			} `json:"student,omitempty"`
			Class struct {
				ID int `json:"id,omitempty"`
			} `json:"class,omitempty"`
		} `json:"enrolment,omitempty"`
	} `json:"message,omitempty"`
	ClientID  int       `json:"clientId,omitempty"`
	MessageID string    `json:"messageId,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
}
