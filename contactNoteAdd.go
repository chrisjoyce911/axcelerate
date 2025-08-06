package axcelerate

import (
	"encoding/json"
	"fmt"
)

// NoteResponse represents the response from adding a note to a contact.
// This type handles both success and error response formats from the Axcelerate API.
//
// Success Response Example:
//
//	{
//	  "NOTEID": 128830815,
//	  "MESSAGE": "Note was added to Contact: 11300044.",
//	  "STATUS": "success"
//	}
//
// Error Response Example:
//
//	{
//	  "DATA": "",
//	  "ERROR": true,
//	  "MESSAGES": "Invalid NoteTypeID!",
//	  "CODE": "412",
//	  "DETAILS": "NoteTypeID: 27938x was not found in this account."
//	}
//
// All fields are optional pointers to accommodate the different response structures.
// Check the ERROR field to determine if the response indicates an error condition.
type NoteResponse struct {
	// Success response fields
	NOTEID  *int    `json:"NOTEID,omitempty"`  // The unique identifier of the created note
	MESSAGE *string `json:"MESSAGE,omitempty"` // Success message describing the operation result
	STATUS  *string `json:"STATUS,omitempty"`  // Status indicator (typically "success" for successful operations)

	// Error response fields
	DATA     *string `json:"DATA,omitempty"`     // Additional data field (often empty in error responses)
	ERROR    *bool   `json:"ERROR,omitempty"`    // Boolean flag indicating if an error occurred
	MESSAGES *string `json:"MESSAGES,omitempty"` // Error message describing what went wrong
	CODE     *string `json:"CODE,omitempty"`     // HTTP or application error code
	DETAILS  *string `json:"DETAILS,omitempty"`  // Detailed error information for debugging
}

// NoteAdd adds a note against a specific Contact.
//
// This function creates a new note for the specified contact in the Axcelerate system.
//
// Parameters:
//   - contactID: The ContactID to add the note to (required)
//   - params: A map of additional parameters including:
//   - "contactNote": The note content to add to the Contact (required)
//   - "noteTypeID": The type of note to add to the Contact (optional, defaults to 88 - System Note)
//   - "emailNote": A comma-separated list of ContactIDs to email the note to (optional)
//
// Returns:
//   - NoteResponse: Contains either success data (NOTEID, MESSAGE, STATUS) or error information
//   - *Response: HTTP response details
//   - error: Any error that occurred during the request
//
// Example Usage:
//
//	params := map[string]string{
//	    "contactNote": "This is a test note",
//	    "noteTypeID": "88",
//	    "emailNote": "12345,67890",
//	}
//	response, httpResp, err := contactService.NoteAdd(12345, params)
//	if err != nil {
//	    // Handle error
//	}
//	if response.ERROR != nil && *response.ERROR {
//	    // Handle API error response
//	    fmt.Printf("Error: %s", *response.MESSAGES)
//	} else {
//	    // Success
//	    fmt.Printf("Note ID: %d", *response.NOTEID)
//	}
func (s *ContactService) NoteAdd(contactID int, params map[string]string) (NoteResponse, *Response, error) {
	var obj NoteResponse

	parms := map[string]string{
		"contactID": fmt.Sprintf("%d", contactID),
	}

	// Add all parameters from the provided map
	for key, value := range params {
		parms[key] = value
	}

	// Set default noteTypeID if not provided
	if _, exists := parms["noteTypeID"]; !exists {
		parms["noteTypeID"] = "88"
	}

	resp, err := do(s.client, "POST", Params{parms: parms, u: "/contact/note/"}, obj)

	if err != nil {
		return obj, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}
