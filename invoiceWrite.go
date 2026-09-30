package axcelerate

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// Invoice writes: create, add lines, issue, apply a payment, and read the
// contact's transactions.
//
// Every rule enforced below was learned against the live API — most of
// them in production, with a customer's money involved — and the API docs
// get several of them wrong. They are validated HERE, before a request is
// sent, because aXcelerate's refusals are either vague ("The invoice items
// are invalid.") or, worse, partial: a refused payment can still record
// the money (see ApplyPayment).

// InvoiceLine is one invoice line to write (aItem on POST and PUT
// /accounting/invoice). Prices and rates are decimal STRINGS, not floats:
// this is money, and 17.10 must arrive as 17.10.
type InvoiceLine struct {
	// ItemID identifies an EXISTING line on a PUT; nil creates a new line.
	ItemID      *int   `json:"ITEMID,omitempty"`
	Description string `json:"DESCRIPTION"`
	Qty         int    `json:"QTY"`
	// ItemCode is REQUIRED on every line — present even when empty. A line
	// without the key makes aXcelerate refuse the WHOLE request with "The
	// invoice items are invalid." (HTTP 400, verified on staging 30 Sep
	// 2026). No omitempty: an empty code is sent as "".
	ItemCode       string `json:"ITEMCODE"`
	UnitPriceGross string `json:"UNITPRICEGROSS"` // gross price for one unit, dollars, e.g. "119.00"
	TaxPercent     string `json:"TAXPERCENT"`     // GST percent, e.g. "0.00" or "10.00"
	FinanceCode    string `json:"FINANCECODE,omitempty"`
	CostCentreCode string `json:"COSTCENTRECODE,omitempty"`
	ServiceDate    string `json:"SERVICEDATE,omitempty"` // YYYY-MM-DD, the date the revenue is recognised against
	// DomainID is the domain the line reports under. aXcelerate does NOT
	// check that it exists — a made-up id is stored and shows a blank
	// Domain in the UI — so validate ids against Domains().
	DomainID *int `json:"DOMAINID,omitempty"`
	PartID   *int `json:"PARTID,omitempty"`
}

// Validate checks a line before it is sent.
func (l InvoiceLine) Validate() error {
	if strings.TrimSpace(l.Description) == "" {
		return fmt.Errorf("DESCRIPTION is required")
	}
	if l.Qty <= 0 {
		return fmt.Errorf("QTY must be greater than zero (got %d)", l.Qty)
	}
	if _, err := decimal(l.UnitPriceGross, "UNITPRICEGROSS"); err != nil {
		return err
	}
	tax, err := decimal(l.TaxPercent, "TAXPERCENT")
	if err != nil {
		return err
	}
	if tax > 100 {
		return fmt.Errorf("TAXPERCENT %s is a percentage and cannot exceed 100", l.TaxPercent)
	}
	if l.ServiceDate != "" {
		if _, err := time.Parse("2006-01-02", l.ServiceDate); err != nil {
			return fmt.Errorf("SERVICEDATE %q must be YYYY-MM-DD", l.ServiceDate)
		}
	}
	if l.DomainID != nil && *l.DomainID <= 0 {
		return fmt.Errorf("DOMAINID must be a positive id (got %d)", *l.DomainID)
	}
	return nil
}

// decimal parses a required non-negative decimal string.
func decimal(s, field string) (float64, error) {
	if strings.TrimSpace(s) == "" {
		return 0, fmt.Errorf("%s is required", field)
	}
	v, err := strconv.ParseFloat(s, 64)
	if err != nil || v < 0 {
		return 0, fmt.Errorf("%s %q must be a non-negative decimal", field, s)
	}
	return v, nil
}

func validateLines(lines []InvoiceLine) error {
	if len(lines) == 0 {
		return fmt.Errorf("at least one line is required")
	}
	for i, l := range lines {
		if err := l.Validate(); err != nil {
			return fmt.Errorf("line %d: %w", i+1, err)
		}
	}
	return nil
}

// InvoiceRequest is the input to CreateInvoice.
type InvoiceRequest struct {
	ContactID int
	// FirstName and LastName are the Bill To name. Both are required:
	// without them the API answers "The Bill To Given Name and Surname
	// fields cannot be empty."
	FirstName string
	LastName  string
	// InvoiceDate and OrderDate are YYYY-MM-DD. The API refuses a call
	// without them ("The invoice order date is out of bounds."); left
	// empty, this client sends today's date for both.
	InvoiceDate       string
	OrderDate         string
	ExternalReference string // optional, at most 60 characters
	Lines             []InvoiceLine
}

// CreateInvoice creates a DRAFT invoice (POST /accounting/invoice/).
//
// The result has INVOICENR "AUTO": it is not issued, has no number, and
// cannot take a payment until ApproveInvoice issues it.
//
// The Bill To surname is sent as `lastname`. The API docs name the
// parameter `surname`, and aXcelerate ignores it: an invoice created with
// surname=Joyce came back with LASTNAME null, and one with lastname=Joyce
// came back correct (staging, 29–30 Sep 2026). `surname` is not needed
// when `lastname` is sent.
func (s *AccountingService) CreateInvoice(req InvoiceRequest) (*Invoice, *Response, error) {
	if req.ContactID <= 0 {
		return nil, nil, fmt.Errorf("invoice: contactID is required")
	}
	if strings.TrimSpace(req.FirstName) == "" || strings.TrimSpace(req.LastName) == "" {
		return nil, nil, fmt.Errorf("invoice: the Bill To FirstName and LastName are both required")
	}
	if len(req.ExternalReference) > 60 {
		return nil, nil, fmt.Errorf("invoice: externalReference is at most 60 characters (got %d)", len(req.ExternalReference))
	}
	if err := validateLines(req.Lines); err != nil {
		return nil, nil, fmt.Errorf("invoice: %w", err)
	}
	today := time.Now().Format("2006-01-02")
	invDate, orderDate := req.InvoiceDate, req.OrderDate
	if invDate == "" {
		invDate = today
	}
	if orderDate == "" {
		orderDate = today
	}
	for _, d := range []string{invDate, orderDate} {
		if _, err := time.Parse("2006-01-02", d); err != nil {
			return nil, nil, fmt.Errorf("invoice: date %q must be YYYY-MM-DD", d)
		}
	}
	items, err := json.Marshal(req.Lines)
	if err != nil {
		return nil, nil, fmt.Errorf("invoice: encoding lines: %w", err)
	}
	parms := map[string]string{
		"contactID":   strconv.Itoa(req.ContactID),
		"firstname":   req.FirstName,
		"lastname":    req.LastName,
		"invoiceDate": invDate,
		"orderDate":   orderDate,
		"aItem":       string(items),
	}
	if req.ExternalReference != "" {
		parms["externalReference"] = req.ExternalReference
	}

	var obj Invoice
	resp, err := do(s.client, "POST", Params{parms: parms, u: "/accounting/invoice/"}, obj)
	if err != nil {
		return nil, resp, err
	}
	if err := refusedIn(resp); err != nil {
		return nil, resp, fmt.Errorf("invoice: %w", err)
	}
	if err := json.Unmarshal([]byte(resp.Body), &obj); err != nil {
		return nil, resp, fmt.Errorf("invoice: decoding response: %w", err)
	}
	return &obj, resp, nil
}

// invoiceLineFields are the line fields the invoice PUT accepts. An
// existing line is re-sent with exactly these; anything aXcelerate derives
// (totals, tax amounts, HASH) is not echoed back.
var invoiceLineFields = []string{
	"ITEMID", "DESCRIPTION", "QTY", "ITEMCODE", "FINANCECODE", "TAXPERCENT",
	"UNITPRICEGROSS", "DOMAINID", "COSTCENTRECODE", "SERVICEDATE", "PARTID",
}

// AddInvoiceLines adds lines to an existing invoice
// (PUT /accounting/invoice/{invoiceID}).
//
// aItem on the PUT is the invoice's COMPLETE line list, not an append: a
// PUT carrying one line REPLACES every line already there. A real booking
// (26 Aug 2026) paid $169 and ended with a $10.80 invoice after a
// one-line PUT overwrote the workshop line. So this reads the invoice
// first and re-sends every existing line with its ITEMID — which is what
// makes it an update rather than a duplicate — plus the new lines, which
// carry no ITEMID. Existing lines are re-sent from the raw response so
// nothing is lost to typing (a DOMAINID, a fractional tax rate), and a
// null ITEMCODE is sent as "" because a line without one fails the call.
//
// Lines cannot be added once an invoice is issued and its items locked.
func (s *AccountingService) AddInvoiceLines(invoiceID int, lines ...InvoiceLine) (*Response, error) {
	if invoiceID <= 0 {
		return nil, fmt.Errorf("invoice lines: invoiceID is required")
	}
	if err := validateLines(lines); err != nil {
		return nil, fmt.Errorf("invoice lines: %w", err)
	}
	for i, l := range lines {
		if l.ItemID != nil {
			return nil, fmt.Errorf("invoice lines: line %d has an ITEMID — new lines must not carry one", i+1)
		}
	}

	existing, locked, resp, err := s.rawInvoiceLines(invoiceID)
	if err != nil {
		return resp, fmt.Errorf("invoice lines: reading invoice %d first: %w", invoiceID, err)
	}
	if locked {
		return resp, fmt.Errorf("invoice lines: invoice %d has its items locked", invoiceID)
	}
	all := make([]any, 0, len(existing)+len(lines))
	for _, e := range existing {
		all = append(all, e)
	}
	for _, l := range lines {
		all = append(all, l)
	}
	items, err := json.Marshal(all)
	if err != nil {
		return nil, fmt.Errorf("invoice lines: encoding: %w", err)
	}

	resp, err = do(s.client, "PUT", Params{
		parms: map[string]string{"aItem": string(items)},
		u:     fmt.Sprintf("/accounting/invoice/%d", invoiceID),
		body:  true, // aItem JSON must not go through the query string
	}, nil)
	if err != nil {
		return resp, err
	}
	if err := refusedIn(resp); err != nil {
		return resp, fmt.Errorf("invoice lines: %w", err)
	}
	return resp, nil
}

// rawInvoiceLines reads an invoice's lines as the PUT accepts them.
func (s *AccountingService) rawInvoiceLines(invoiceID int) ([]map[string]any, bool, *Response, error) {
	resp, err := do(s.client, "GET", Params{parms: map[string]string{}, u: fmt.Sprintf("/accounting/invoice/%d", invoiceID)}, nil)
	if err != nil {
		return nil, false, resp, err
	}
	type raw struct {
		Locked bool             `json:"AREITEMSLOCKED"`
		Items  []map[string]any `json:"ITEMS"`
	}
	var one raw
	if err := json.Unmarshal([]byte(resp.Body), &one); err != nil {
		// Some reads answer with a single-element array.
		var many []raw
		if err2 := json.Unmarshal([]byte(resp.Body), &many); err2 != nil || len(many) == 0 {
			return nil, false, resp, fmt.Errorf("decoding invoice: %w", err)
		}
		one = many[0]
	}
	out := make([]map[string]any, 0, len(one.Items))
	for _, it := range one.Items {
		line := map[string]any{}
		for _, f := range invoiceLineFields {
			if v, ok := it[f]; ok && v != nil {
				line[f] = v
			}
		}
		if _, ok := line["ITEMCODE"]; !ok && len(line) > 0 {
			line["ITEMCODE"] = ""
		}
		if len(line) > 0 {
			out = append(out, line)
		}
	}
	return out, one.Locked, resp, nil
}

// ApproveInvoice ISSUES a draft invoice
// (PUT /accounting/invoice/{invoiceGUID}/approve) — it takes the GUID
// (INVGUID), not the invoice id. A draft has INVOICENR "AUTO" and no
// number, and a payment cannot be applied to it; see ApplyPayment.
// Issuing locks the invoice's items.
func (s *AccountingService) ApproveInvoice(invoiceGUID string) (*Response, error) {
	if strings.TrimSpace(invoiceGUID) == "" {
		return nil, fmt.Errorf("approve invoice: invoiceGUID is required (INVGUID, not the invoice id)")
	}
	resp, err := do(s.client, "PUT", Params{
		parms: map[string]string{},
		u:     fmt.Sprintf("/accounting/invoice/%s/approve", invoiceGUID),
		body:  true,
	}, nil)
	if err != nil {
		return resp, err
	}
	if err := refusedIn(resp); err != nil {
		return resp, fmt.Errorf("approve invoice: %w", err)
	}
	return resp, nil
}

// InvoicePayment is the input to ApplyPayment.
type InvoicePayment struct {
	ContactID int // the payer
	InvoiceID int
	// Amount is in DOLLARS as a decimal string ("119.00"). A payment
	// gateway usually reports cents; sending cents here records a payment
	// a hundred times too large against a customer's invoice.
	Amount      string
	Reference   string // optional; your payment/receipt reference
	Description string // optional
	// PaymentMethodID: 1=Cash, 2=Credit Card (the API default), 4=Direct
	// Deposit, 5=Cheque, 6=EFTPOS. 0 leaves the API default.
	PaymentMethodID int
}

// ApplyPayment records a payment AGAINST an invoice
// (POST /accounting/transaction/ with invoiceID), refusing first if the
// invoice cannot take it.
//
// Why it checks: a payment against an invoice that is not issued is
// refused with HTTP 400 — "Transaction recorded in aXcelerate but unable
// to apply against invoice. An invoice must be issued with an invoice
// number before transactions can be applied against it." The money IS
// recorded anyway, as an unassigned receipt on the contact (verified on
// staging, 30 Sep 2026). A caller that treats the 400 as "nothing
// happened" and retries records the payment twice. So this reads the
// invoice first and refuses a draft (INVOICENR "AUTO") or a void invoice
// before any money is written. If ApplyPayment returns an error AFTER the
// transaction call, check Transactions() before retrying.
//
// To record money with no invoice (a receipt on account), use
// CreateTransaction without an invoiceID.
func (s *AccountingService) ApplyPayment(req InvoicePayment) (*Transaction, *Response, error) {
	if req.ContactID <= 0 {
		return nil, nil, fmt.Errorf("apply payment: contactID is required")
	}
	if req.InvoiceID <= 0 {
		return nil, nil, fmt.Errorf("apply payment: invoiceID is required (use CreateTransaction for money on account)")
	}
	amt, err := decimal(req.Amount, "amount")
	if err != nil {
		return nil, nil, fmt.Errorf("apply payment: %w", err)
	}
	if amt == 0 {
		return nil, nil, fmt.Errorf("apply payment: amount must be greater than zero")
	}

	inv, resp, err := s.GetInvoice(req.InvoiceID)
	if err != nil {
		return nil, resp, fmt.Errorf("apply payment: reading invoice %d first: %w", req.InvoiceID, err)
	}
	if inv.InvoiceNumber == "" || strings.EqualFold(inv.InvoiceNumber, "AUTO") {
		return nil, resp, fmt.Errorf("apply payment: invoice %d is a draft (no invoice number) — ApproveInvoice it first; "+
			"aXcelerate would record the money and refuse to apply it", req.InvoiceID)
	}

	parms := map[string]string{
		"contactID": strconv.Itoa(req.ContactID),
		"invoiceID": strconv.Itoa(req.InvoiceID),
		"amount":    req.Amount,
	}
	if req.Reference != "" {
		parms["reference"] = req.Reference
	}
	if req.Description != "" {
		parms["description"] = req.Description
	}
	if req.PaymentMethodID != 0 {
		parms["paymentMethodID"] = strconv.Itoa(req.PaymentMethodID)
	}
	return s.CreateTransaction(parms)
}

// AccountTransaction is one money movement on a contact's account, as
// GET /accounting/transaction/ lists it. Amounts arrive as strings
// ("59", "61.36") and are decoded as floats — an integer type cannot hold
// a surcharged amount.
type AccountTransaction struct {
	TransactionID     StringInt   `json:"TRANSACTIONID"`
	ContactID         StringInt   `json:"CONTACTID"`
	GUID              string      `json:"GUID"`
	TransDate         string      `json:"TRANSDATE"` // "2006-01-02 15:04"
	TransactionType   string      `json:"TRANSACTIONTYPE"`
	TransactionTypeID StringInt   `json:"TRANSACTIONTYPEID"`
	PaymentMethod     string      `json:"PAYMENTMETHOD"`
	PaymentMethodID   StringInt   `json:"PAYMENTMETHODID"`
	Amount            StringFloat `json:"AMOUNT"`
	// UnassignedAmount is money that arrived and was never applied to an
	// invoice — a payment refused against a draft invoice ends up here.
	UnassignedAmount StringFloat `json:"UNASSIGNEDAMOUNT"`
	Reference        string      `json:"REFERENCE"`
	Description      string      `json:"DESCRIPTION"`
	Currency         string      `json:"CURRENCY"`
}

// Transactions lists a contact's transactions
// (GET /accounting/transaction/?contactID=N) — the way to see money on
// the account that no invoice has claimed.
func (s *AccountingService) Transactions(contactID int) ([]AccountTransaction, *Response, error) {
	if contactID <= 0 {
		return nil, nil, fmt.Errorf("transactions: contactID is required")
	}
	var out []AccountTransaction
	resp, err := do(s.client, "GET", Params{
		parms: map[string]string{"contactID": strconv.Itoa(contactID)},
		u:     "/accounting/transaction/",
	}, out)
	if err != nil {
		return nil, resp, err
	}
	if err := json.Unmarshal([]byte(resp.Body), &out); err != nil {
		return nil, resp, fmt.Errorf("transactions: decoding response: %w", err)
	}
	return out, resp, nil
}

// refusedIn reports an ERROR carried in a 2xx body. aXcelerate mostly
// refuses with a 4xx, which the client already turns into an error, but
// an ERROR field alongside a 200 is still a refusal.
func refusedIn(resp *Response) error {
	if resp == nil || resp.Body == "" {
		return nil
	}
	var out struct {
		Error    json.RawMessage `json:"ERROR"`
		Messages any             `json:"MESSAGES"`
	}
	if err := json.Unmarshal([]byte(resp.Body), &out); err != nil {
		return nil // not an object (e.g. an array) — nothing to report
	}
	if e := string(out.Error); e != "" && e != "null" && e != "false" {
		return fmt.Errorf("aXcelerate refused: %v", out.Messages)
	}
	return nil
}
