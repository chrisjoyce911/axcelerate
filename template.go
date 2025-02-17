package axcelerate

import (
	"fmt"

	"encoding/json"
)

// CoursesService handles all interactions with Contact
type TemplateService struct {
	client *Client
}

type EmailResponse struct {
	FailedCount int    `json:"FAILEDCOUNT"`
	Message     string `json:"MESSAGE"`
	// Errors         []string `json:"ERRORS,omitempty"`
	AttemptedCount int `json:"ATTEMPTEDCOUNT"`
	SuccessCount   int `json:"SUCCESSCOUNT"`
}

type TemplateEmailParams struct {
	PlanID                  int    `json:"planID,omitempty"`
	ContactID               int    `json:"contactID" validate:"required"`
	InstanceID              int    `json:"instanceID,omitempty"`
	InvoiceID               int    `json:"invoiceID,omitempty"`
	Subject                 string `json:"subject,omitempty"`
	Type                    string `json:"type,omitempty"`
	InvoiceAttachmentPlanID int    `json:"invoiceAttachmentPlanID,omitempty"`
	HasIcalAttachment       bool
	// Content            *string `json:"content,omitempty"`            // Optional
	// From               *string `json:"from,omitempty"`               // Optional
	// To                 *string `json:"to,omitempty"`                 // Optional
	// ID                 *int    `json:"id,omitempty"`                 // Optional
	// IncludeStatus      *string `json:"includeStatus,omitempty"`      // Optional
	// HonourUnsubscribed *bool   `json:"honourUnsubscribed,omitempty"` // Optional
	// AttachmentPlanID   *int    `json:"attachmentPlanID,omitempty"`   // Optional

	// AttachAccreditedCert *bool              `json:"attachAccreditedCertificate,omitempty"`
	// AttachWorkshopCert   *bool              `json:"attachWorkshopCertificate,omitempty"`
	// UserRoleID           *[]int             `json:"userRoleID,omitempty"`
	// ContactRoleID        *[]int             `json:"contactRoleID,omitempty"`
	// SkillGroupID         *[]int             `json:"skillGroupID,omitempty"`
	// ReplaceContent       *map[string]string `json:"replaceContent,omitempty"`
	// HasIcalAttachment    *bool              `json:"hasIcalAttachment,omitempty"`
	// CopyToAlternateEmail *bool              `json:"copyToAlternateEmailAddress,omitempty"`
}

// TemplateEmail sends a template email to the specified contact(s) using the provided parameters.
// HTTP Parameters are passed to the resource via the URL or via FORM posts depending on the HTTP verb used.
//
// # Header Parameters
//
// verbose                 boolean    false    If passed, an additional key REPORT will be included, which will contain detailed information on each contact emailed.
// planID                  numeric    false    The ID of a template to use. One of planID or content is required.
// content                 string     false    The template content to use. Text such as [Trainee Full Name] will be replaced with the contact's name. One of planID or content is required.
// subject                 string     false    The email subject. Required if content is used. Defaults to the template name if planID is used and subject is not passed.
// from                    string     false    Must be a valid email address or the contactID of a contact with a valid email address. Defaults to the API user if not passed.
// to                      string     false    Defaults to contactID. Supports a contactID, a list of contactIDs, or the special values: [student,creator,client,owner,trainer] when used with type=w or p and an instanceID.
// contactID               numeric    false    The contact to email the template to. Either contactID or search parameters (e.g., instanceID/type, userRoleID) must be provided.
// type                    string     false    The type of course. Used with instanceID or ID parameters.
// instanceID              numeric    false    The course instance. Used with type to email templates to students on the course.
// ID                      numeric    false    The course ID. Prefer instanceID as courses at the ID level lack student context. Requires type and contactID if used.
// includeStatus           string     false    A list of statuses determining which students to email. Options: [All,Enrolled,Completed,Tentative,Cancelled,Suspended,Deferred]. Works with type and instanceID.
// honourUnsubscribed      boolean    false    If passed, students who unsubscribed from merge documents will NOT receive an email.
// attachmentPlanID        numeric    false    The ID of a template to attach as a PDF. The filename will match the template name.
// invoiceID               numeric    false    The invoiceID associated with this email. Helps parse invoice-related fields. Use with invoiceAttachmentPlanID to include an invoice as a PDF.
// invoiceAttachmentPlanID numeric    false    The ID of a template to attach an invoice as a PDF. Names the attachment Invoice_[N].pdf, where [N] is the invoice number.
// attachAccreditedCertificate boolean false   Attaches the latest accredited certificate issued to the student if true. Works with accredited instanceID (type=p) or workshops linked to accredited enrolments (type=w). Only attaches when emailing the student (to=student).
// attachWorkshopCertificate boolean  false    Attaches the latest non-accredited certificate for the student if true. Works with workshop instanceID (type=w) and emails sent to the student (to=student).
// userRoleID              numeric    false    Email users with this role. Can be a list.
// contactRoleID           numeric    false    Email contacts belonging to this role. Enterprise only. Can be a list.
// skillGroupID            numeric    false    Email contacts whose role contains this skill group. Enterprise only. Can be a list.
// replaceContent          string     false    JSON string containing key-value pairs to replace template fields with custom content.
// hasIcalAttachment       boolean    false    Attaches events for the workshop in iCalendar format if true.
// copyToAlternateEmailAddress boolean false   Sends the template to the recipient's alternate email addresses if valid.
//
// # Parameters
// contactID   int    The ID of the contact to send the email to.
// planID      int    The ID of the email template to use.
// parms       map[string]string Additional parameters for the email.
//
// Returns:
// - EmailResponse: The response containing email details and results.
// - Response: The raw HTTP response.
// - error: Any error that occurred during execution.

func (s *TemplateService) TemplateEmail(params TemplateEmailParams) (*EmailResponse, *Response, error) {
	// Validate the required parameter
	if params.ContactID == 0 {
		return nil, nil, fmt.Errorf("contactID is required")
	}

	// Initialize the map with required parameters
	parms := map[string]string{
		"planID":                  fmt.Sprintf("%d", params.PlanID),
		"contactID":               fmt.Sprintf("%d", params.ContactID),
		"instanceID":              fmt.Sprintf("%d", params.InstanceID),
		"invoiceID":               fmt.Sprintf("%d", params.InvoiceID),
		"subject":                 params.Subject,
		"type":                    params.Type,
		"invoiceAttachmentPlanID": fmt.Sprintf("%d", params.InvoiceAttachmentPlanID),
	}

	if params.HasIcalAttachment {
		parms["hasIcalAttachment"] = "true"
	}

	var obj EmailResponse

	// API call
	resp, err := do(s.client, "POST", Params{parms: parms, u: "/template/email"}, obj)
	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return &obj, resp, err
}

// TemplateEmail sends a template email to the specified contact(s) using the provided parameters.
// HTTP Parameters are passed to the resource via the URL or via FORM posts depending on the HTTP verb used.
//
// # Header Parameters
//
// verbose                 boolean    false    If passed, an additional key REPORT will be included, which will contain detailed information on each contact emailed.
// planID                  numeric    false    The ID of a template to use. One of planID or content is required.
// content                 string     false    The template content to use. Text such as [Trainee Full Name] will be replaced with the contact's name. One of planID or content is required.
// subject                 string     false    The email subject. Required if content is used. Defaults to the template name if planID is used and subject is not passed.
// from                    string     false    Must be a valid email address or the contactID of a contact with a valid email address. Defaults to the API user if not passed.
// to                      string     false    Defaults to contactID. Supports a contactID, a list of contactIDs, or the special values: [student,creator,client,owner,trainer] when used with type=w or p and an instanceID.
// contactID               numeric    false    The contact to email the template to. Either contactID or search parameters (e.g., instanceID/type, userRoleID) must be provided.
// type                    string     false    The type of course. Used with instanceID or ID parameters.
// instanceID              numeric    false    The course instance. Used with type to email templates to students on the course.
// ID                      numeric    false    The course ID. Prefer instanceID as courses at the ID level lack student context. Requires type and contactID if used.
// includeStatus           string     false    A list of statuses determining which students to email. Options: [All,Enrolled,Completed,Tentative,Cancelled,Suspended,Deferred]. Works with type and instanceID.
// honourUnsubscribed      boolean    false    If passed, students who unsubscribed from merge documents will NOT receive an email.
// attachmentPlanID        numeric    false    The ID of a template to attach as a PDF. The filename will match the template name.
// invoiceID               numeric    false    The invoiceID associated with this email. Helps parse invoice-related fields. Use with invoiceAttachmentPlanID to include an invoice as a PDF.
// invoiceAttachmentPlanID numeric    false    The ID of a template to attach an invoice as a PDF. Names the attachment Invoice_[N].pdf, where [N] is the invoice number.
// attachAccreditedCertificate boolean false   Attaches the latest accredited certificate issued to the student if true. Works with accredited instanceID (type=p) or workshops linked to accredited enrolments (type=w). Only attaches when emailing the student (to=student).
// attachWorkshopCertificate boolean  false    Attaches the latest non-accredited certificate for the student if true. Works with workshop instanceID (type=w) and emails sent to the student (to=student).
// userRoleID              numeric    false    Email users with this role. Can be a list.
// contactRoleID           numeric    false    Email contacts belonging to this role. Enterprise only. Can be a list.
// skillGroupID            numeric    false    Email contacts whose role contains this skill group. Enterprise only. Can be a list.
// replaceContent          string     false    JSON string containing key-value pairs to replace template fields with custom content.
// hasIcalAttachment       boolean    false    Attaches events for the workshop in iCalendar format if true.
// copyToAlternateEmailAddress boolean false   Sends the template to the recipient's alternate email addresses if valid.
//
// # Parameters
// contactID   int    The ID of the contact to send the email to.
// planID      int    The ID of the email template to use.
// parms       map[string]string Additional parameters for the email.
//
// Returns:
// - EmailResponse: The response containing email details and results.
// - Response: The raw HTTP response.
// - error: Any error that occurred during execution.

func (s *TemplateService) EmailTemplate(parms map[string]string) (*EmailResponse, *Response, error) {

	var obj EmailResponse

	// API call
	resp, err := do(s.client, "POST", Params{parms: parms, u: "/template/email"}, obj)
	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return &obj, resp, err
}
