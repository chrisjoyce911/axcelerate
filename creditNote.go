package axcelerate

import (
	"encoding/json"
	"fmt"
	"time"
)

// CreditNoteItem is one line of a credit note. The field names mirror
// the API's aItem entries (DESCRIPTION, QTY, ITEMCODE, TAXPERCENT,
// UNITPRICEGROSS are required by POST /accounting/creditnote/).
type CreditNoteItem struct {
	Description    string  `json:"DESCRIPTION"`              // What the credit is for
	Qty            int     `json:"QTY"`                      // Quantity, normally 1
	ItemCode       string  `json:"ITEMCODE"`                 // Catalogue item code
	TaxPercent     string  `json:"TAXPERCENT"`               // e.g. "0.00" or "10.00"
	UnitPriceGross string  `json:"UNITPRICEGROSS"`           // Gross unit price as a decimal string
	FinanceCode    string  `json:"FINANCECODE,omitempty"`    // Optional finance (account) code
	CostCentreCode string  `json:"COSTCENTRECODE,omitempty"` // Optional cost centre
	ServiceDate    string  `json:"SERVICEDATE,omitempty"`    // REQUIRED — see CreateCreditNote
	DomainID       *int    `json:"DOMAINID,omitempty"`       // Optional domain
	PartID         *int    `json:"PARTID,omitempty"`         // Optional part id
	Data           *string `json:"DATA,omitempty"`           // Optional free-form data
}

// CreditNote is a credit note as the accounting API returns it: money
// held FOR a contact (it carries its own balance and paid flag), as
// distinct from an invoice, which is money the contact owes.
type CreditNote struct {
	CreditNoteID StringInt `json:"CREDITNOTEID"` // Unique identifier
	CreditNoteNr *string   `json:"CREDITNOTENR"` // Human-facing number (nullable)
	ContactID    StringInt `json:"CONTACTID"`    // Contact the credit is held for
	GUID         string    `json:"GUID"`         // Globally unique identifier
	FirstName    *string   `json:"FIRSTNAME"`    // Contact's given name (nullable)
	LastName     *string   `json:"LASTNAME"`     // Contact's surname (nullable)
	Balance      *string   `json:"BALANCE"`      // Remaining credit (nullable)
	PriceGross   *string   `json:"PRICEGROSS"`   // Gross total of the note (nullable)
	IsPaid       bool      `json:"ISPAID"`       // Whether the note has been settled
	IsVoid       bool      `json:"ISVOID"`       // Whether the note has been voided
	Date         *string   `json:"CREDITNOTEDATE"`
}

// CreditNoteRequest is the input to CreateCreditNote. ContactID,
// FirstName, Surname and at least one Item are required by the API.
type CreditNoteRequest struct {
	ContactID int
	FirstName string
	Surname   string
	Items     []CreditNoteItem
	// Date is the credit note's date, YYYY-MM-DD. The API documents it
	// as OPTIONAL and then refuses the call without it —
	// `key [UNDEFINED_DATE] doesn't exist in the request scope`
	// (observed in production, 26 Aug 2026). Left empty, this client
	// sends today's date rather than reproducing that error.
	Date string
}

// CreateCreditNote records a credit note for a contact — money held for
// them rather than owed by them.
//
// The usual reason is money taken for something not delivered: a
// payment captured while an enrolment could not be created leaves the
// customer with cash on account and no invoice to apply it to. A credit
// note documents that obligation, carries its own balance, and can be
// settled against an invoice once one exists.
//
// POST /accounting/creditnote/
//
// Required params (per the API docs):
//
//	contactID   numeric   the contact the credit is held for
//	firstname   string    contact's given name
//	surname     string    contact's surname
//	aItem       json      array of line items; each needs DESCRIPTION,
//	                      QTY, ITEMCODE, TAXPERCENT, UNITPRICEGROSS
//
// SERVICEDATE (YYYY-MM-DD) is documented as OPTIONAL on a line and is
// not: without it every call fails with
//
//	http 500: key [UNDEFINED_DATE] doesn't exist in the request scope
//
// which names no field and reads like a server fault. It is not the
// header date — creditnoteDate, creditNoteDate, date and every other
// spelling change nothing; the date the endpoint wants is on each LINE.
// This client fills a missing SERVICEDATE with the note's date so the
// call cannot be made without one. (Established against the live API on
// 26 Aug 2026, after six failed credits for real customers whose money
// went unrecorded.)
//
// Optional: FINANCECODE, COSTCENTRECODE, DOMAINID, DATA.
//
// Usage:
//
//	note, resp, err := client.Accounting.CreateCreditNote(axcelerate.CreditNoteRequest{
//	    ContactID: 14754152, FirstName: "Jinky", Surname: "Abayan",
//	    Items: []axcelerate.CreditNoteItem{{
//	        Description: "Payment held — enrolment could not be completed",
//	        Qty: 1, ItemCode: "CREDIT", TaxPercent: "0.00", UnitPriceGross: "59.00",
//	    }},
//	})
func (s *AccountingService) CreateCreditNote(req CreditNoteRequest) (*CreditNote, *Response, error) {
	if req.ContactID <= 0 {
		return nil, nil, fmt.Errorf("credit note: contactID is required")
	}
	if len(req.Items) == 0 {
		return nil, nil, fmt.Errorf("credit note: at least one item is required")
	}
	date := req.Date
	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
	// Every line needs a service date or the whole call fails; see the
	// doc comment. Filling it here means a caller cannot get this wrong.
	lines := make([]CreditNoteItem, len(req.Items))
	copy(lines, req.Items)
	for i := range lines {
		if lines[i].ServiceDate == "" {
			lines[i].ServiceDate = date
		}
	}

	items, err := json.Marshal(lines)
	if err != nil {
		return nil, nil, fmt.Errorf("credit note: encoding items: %w", err)
	}
	// Header-level fields are contactID, firstname, surname,
	// creditnoteDate and aItem — everything else (finance code, cost
	// centre, service date, data) belongs on the LINE, not here.
	parms := map[string]string{
		"contactID":      fmt.Sprintf("%d", req.ContactID),
		"firstname":      req.FirstName,
		"surname":        req.Surname,
		"creditnoteDate": date,
		"aItem":          string(items),
	}

	var obj CreditNote
	resp, err := do(s.client, "POST", Params{parms: parms, u: "/accounting/creditnote/"}, obj)
	if err != nil {
		return nil, resp, err
	}
	if err := json.Unmarshal([]byte(resp.Body), &obj); err != nil {
		return nil, resp, fmt.Errorf("credit note: decoding response: %w", err)
	}
	return &obj, resp, nil
}

// CreditNotes returns the credit notes held for a contact.
//
// GET /accounting/creditnote/?contactID=N
func (s *AccountingService) CreditNotes(contactID int) ([]CreditNote, *Response, error) {
	if contactID <= 0 {
		return nil, nil, fmt.Errorf("credit notes: contactID is required")
	}
	var out []CreditNote
	resp, err := do(s.client, "GET", Params{
		parms: map[string]string{"contactID": fmt.Sprintf("%d", contactID)},
		u:     "/accounting/creditnote/",
	}, out)
	if err != nil {
		return nil, resp, err
	}
	if err := json.Unmarshal([]byte(resp.Body), &out); err != nil {
		return nil, resp, fmt.Errorf("credit notes: decoding response: %w", err)
	}
	return out, resp, nil
}

// GetCreditNote returns one credit note by id.
//
// GET /accounting/creditnote/{creditnoteID}
func (s *AccountingService) GetCreditNote(creditNoteID int) (*CreditNote, *Response, error) {
	if creditNoteID <= 0 {
		return nil, nil, fmt.Errorf("credit note: creditNoteID is required")
	}
	var obj CreditNote
	resp, err := do(s.client, "GET", Params{
		parms: map[string]string{},
		u:     fmt.Sprintf("/accounting/creditnote/%d", creditNoteID),
	}, obj)
	if err != nil {
		return nil, resp, err
	}
	if err := json.Unmarshal([]byte(resp.Body), &obj); err != nil {
		return nil, resp, fmt.Errorf("credit note: decoding response: %w", err)
	}
	return &obj, resp, nil
}
