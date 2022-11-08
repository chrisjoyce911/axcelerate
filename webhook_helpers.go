package axcelerate

import (
	"net/url"
	"strconv"
)

type WebhookBooking struct {
	ContactID     int
	StatusID      int
	NewStatus     string
	CurrentStatus string
	EventType     string
	WorkshopID    int
}

type WebhookContact struct {
	ContactID   int
	EventType   string
	ContactName string
}

type WebhookWorkshop struct {
	WorkshopTriggerID int
	ProcessGUID       string
	PDataID           int
}

func GetWebHookBooking(payload []byte) (WebhookBooking, error) {
	var b WebhookBooking

	params, err := url.ParseQuery(string(payload))
	if err != nil {
		return b, err
	}

	contactID, err := strconv.Atoi(params.Get("contactID"))
	if err != nil {
		return b, err
	}
	b.ContactID = contactID

	statusID, err := strconv.Atoi(params.Get("statusID"))
	if err != nil {
		return b, err
	}
	b.StatusID = statusID

	PDataID, err := strconv.Atoi(params.Get("PDataID"))
	if err != nil {
		return b, err
	}
	b.WorkshopID = PDataID

	b.NewStatus = params.Get("new_status")
	b.CurrentStatus = params.Get("current_status")
	b.EventType = params.Get("eventType")

	return b, nil
}

func GetWebHookContact(payload []byte) (WebhookContact, error) {
	var c WebhookContact

	params, err := url.ParseQuery(string(payload))
	if err != nil {
		return c, err
	}

	i, err := strconv.Atoi(params.Get("contactID"))
	if err != nil {
		return c, err
	}

	c.ContactID = i
	c.EventType = params.Get("eventType")
	c.ContactName = params.Get("contactName")

	return c, nil
}

func GetWebHookWorkshop(payload []byte) (WebhookWorkshop, error) {
	var w WebhookWorkshop

	params, err := url.ParseQuery(string(payload))
	if err != nil {
		return w, err
	}

	w.WorkshopTriggerID, _ = strconv.Atoi(params.Get("workshopTriggerID"))
	w.ProcessGUID = params.Get("processGUID")
	w.PDataID, _ = strconv.Atoi(params.Get("PDataID"))

	return w, nil
}
