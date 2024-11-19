package axcelerate

import (
	"fmt"

	"encoding/json"
)

// CoursesService handles all interactions with Contact
type VenueService struct {
	client *Client
}

type Venue struct {
	RowID       int    `json:"ROWID"`
	Name        string `json:"NAME"`
	Address1    string `json:"SADDRESS1"`
	Address2    string `json:"SADDRESS2"`
	City        string `json:"SCITY"`
	State       string `json:"SSTATE"`
	Postcode    string `json:"SPOSTCODE"`
	MobilePhone string `json:"MOBILEPHONE"`
	ContactID   int    `json:"CONTACTID"`
	Count       int    `json:"COUNT"`
}

// Venue A venue search.
//
// # Parameters
// contactID   int    The ID of the related Contact

//
// Returns:
// - Venue: The response containing Venue details and results.
// - Response: The raw HTTP response.
// - error: Any error that occurred during execution.

func (s *VenueService) Venue(contactID int) (*Venue, *Response, error) {

	// Initialize the map with required parameters
	parms := map[string]string{
		"contactID": fmt.Sprintf("%d", contactID),
	}

	var obj []Venue

	// API call
	resp, err := do(s.client, "POST", Params{parms: parms, u: "/venues"}, obj)
	if err != nil {
		return nil, resp, err
	}

	if len(obj) > 0 {
		err = json.Unmarshal([]byte(resp.Body), &obj)
		return &obj[0], resp, err

	}

	return nil, resp, err
}

// Venues searches for venues.
//
// # Parameters
//
// - contactID: The ID of the related Contact.
//
// # HTTP Parameters
//
// HTTP Parameters are passed to the resource via the URL or via FORM posts depending on the HTTP verb used.
//
// Header         Type     Required  Default  Description
// -----------------------------------------------------------------------
// offset         numeric  false     0        Used for paging - start at record.
// displayLength  numeric  false     10       Maximum number of records to return (up to a system maximum of 100).
// contactID      numeric  false              The ID of the related Contact.
// name           string   false              The Name of the Venue.
// sAddress1      string   false              Venue address line 1.
// sAddress2      string   false              Venue address line 2.
// sCity          string   false              Venue City.
// sState         string   false              Venue State.
// sPostcode      string   false              Venue Post Code.
// mobilePhone    string   false              Venue Contact mobile phone.
//
// # Returns
//
// - []Venue: The response containing Venues details and results.
// - Response: The raw HTTP response.
// - error: Any error that occurred during execution.
func (s *VenueService) Venues(parms map[string]string) ([]Venue, *Response, error) {

	var obj []Venue

	// API call
	resp, err := do(s.client, "POST", Params{parms: parms, u: "/venues"}, obj)
	if err != nil {
		return nil, resp, err
	}

	if len(obj) > 0 {
		err = json.Unmarshal([]byte(resp.Body), &obj)
		return obj, resp, err
	}

	return nil, resp, err
}
