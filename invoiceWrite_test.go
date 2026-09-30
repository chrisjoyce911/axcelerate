package axcelerate

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// route answers each request from a method+path table and records what
// was sent, so a test can assert the exact wire shape.
type sent struct {
	method, path string
	form         url.Values // body for POST / opted-in PUT
	query        url.Values
}

func routedClient(t *testing.T, routes map[string]string, log *[]sent) *AccountingService {
	t.Helper()
	tclient := NewTestClient(func(req *http.Request) *http.Response {
		var form url.Values
		if req.Body != nil {
			b, _ := io.ReadAll(req.Body)
			form, _ = url.ParseQuery(string(b))
		}
		*log = append(*log, sent{req.Method, req.URL.Path, form, req.URL.Query()})
		body, ok := routes[req.Method+" "+req.URL.Path]
		code := 200
		if !ok {
			code, body = 404, `{"ERROR":true,"MESSAGES":"no route in test"}`
		}
		return &http.Response{StatusCode: code, Body: io.NopCloser(bytes.NewBufferString(body)), Header: make(http.Header)}
	})
	client, _ := NewClient("", "", HttpClient(tclient))
	return &AccountingService{client: client}
}

func line() InvoiceLine {
	return InvoiceLine{Description: "Pocket Mask Training", Qty: 1, ItemCode: "",
		UnitPriceGross: "19.00", TaxPercent: "0.00", FinanceCode: "43420", DomainID: Int(5566)}
}

// --- validation -----------------------------------------------------------

func TestInvoiceLineValidate(t *testing.T) {
	assert.NoError(t, line().Validate())
	for name, mut := range map[string]func(*InvoiceLine){
		"no description":      func(l *InvoiceLine) { l.Description = " " },
		"zero qty":            func(l *InvoiceLine) { l.Qty = 0 },
		"no price":            func(l *InvoiceLine) { l.UnitPriceGross = "" },
		"price not decimal":   func(l *InvoiceLine) { l.UnitPriceGross = "$19" },
		"negative price":      func(l *InvoiceLine) { l.UnitPriceGross = "-1" },
		"no tax":              func(l *InvoiceLine) { l.TaxPercent = "" },
		"tax over 100":        func(l *InvoiceLine) { l.TaxPercent = "110" },
		"bad service date":    func(l *InvoiceLine) { l.ServiceDate = "30/09/2026" },
		"non-positive domain": func(l *InvoiceLine) { l.DomainID = Int(0) },
	} {
		l := line()
		mut(&l)
		assert.Error(t, l.Validate(), name)
	}
}

// ITEMCODE must be on the wire even when empty — a line without the key
// fails the whole request with "The invoice items are invalid."
func TestInvoiceLineAlwaysSendsItemCode(t *testing.T) {
	b, err := json.Marshal(line())
	require.NoError(t, err)
	assert.Contains(t, string(b), `"ITEMCODE":""`)
	assert.Contains(t, string(b), `"UNITPRICEGROSS":"19.00"`, "money is sent as the decimal string given")
	assert.NotContains(t, string(b), `"ITEMID"`, "a new line carries no ITEMID")
}

// --- CreateInvoice ----------------------------------------------------------

func TestCreateInvoiceSendsLastnameAndDates(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{
		"POST /api/accounting/invoice/": `{"INVOICEID":3868280,"INVOICENR":"AUTO","INVGUID":"E088521D","FIRSTNAME":"Chris","LASTNAME":"Joyce","ERROR":null}`,
	}, &log)

	inv, _, err := s.CreateInvoice(InvoiceRequest{ContactID: 14518907, FirstName: "Chris", LastName: "Joyce",
		ExternalReference: "REF-1", Lines: []InvoiceLine{line()}})
	require.NoError(t, err)
	assert.Equal(t, 3868280, inv.InvoiceID)
	assert.Equal(t, "AUTO", inv.InvoiceNumber, "a created invoice is a draft")

	require.Len(t, log, 1)
	f := log[0].form
	assert.Equal(t, "Joyce", f.Get("lastname"), "the Bill To surname is `lastname` — `surname` is ignored")
	assert.Empty(t, f.Get("surname"))
	assert.NotEmpty(t, f.Get("invoiceDate"), "missing dates fail with 'order date is out of bounds'")
	assert.NotEmpty(t, f.Get("orderDate"))
	var items []map[string]any
	require.NoError(t, json.Unmarshal([]byte(f.Get("aItem")), &items))
	assert.Equal(t, float64(5566), items[0]["DOMAINID"])
}

func TestCreateInvoiceRefusesBeforeSending(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{}, &log)
	ok := InvoiceRequest{ContactID: 1, FirstName: "A", LastName: "B", Lines: []InvoiceLine{line()}}
	for name, mut := range map[string]func(*InvoiceRequest){
		"no contact":    func(r *InvoiceRequest) { r.ContactID = 0 },
		"no first name": func(r *InvoiceRequest) { r.FirstName = "" },
		"no last name":  func(r *InvoiceRequest) { r.LastName = "" },
		"no lines":      func(r *InvoiceRequest) { r.Lines = nil },
		"bad line":      func(r *InvoiceRequest) { r.Lines = []InvoiceLine{{Description: "x"}} },
		"bad date":      func(r *InvoiceRequest) { r.OrderDate = "2026-13-01" },
		"reference too long": func(r *InvoiceRequest) {
			r.ExternalReference = "0123456789012345678901234567890123456789012345678901234567890"
		},
	} {
		r := ok
		mut(&r)
		_, _, err := s.CreateInvoice(r)
		assert.Error(t, err, name)
	}
	assert.Empty(t, log, "nothing is sent for a request that would be refused")
}

// --- AddInvoiceLines ----------------------------------------------------------

// The PUT's aItem REPLACES every line, so existing lines are read back and
// re-sent with their ITEMID (and DOMAINID, and a "" for a null ITEMCODE),
// and the JSON travels in the form body, not the query string.
func TestAddInvoiceLinesKeepsExistingLines(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{
		"GET /api/accounting/invoice/3906271": `{"INVOICEID":3906271,"AREITEMSLOCKED":false,"ITEMS":[
			{"ITEMID":6469380,"DESCRIPTION":"Workshop Booking for A & B #1","QTY":1,"UNITPRICEGROSS":59,
			 "TAXPERCENT":0,"ITEMCODE":"EFA-012","FINANCECODE":"41230","DOMAINID":4677,"TOTALGROSS":59,"HASH":"x"},
			{"ITEMID":6469381,"DESCRIPTION":"Safe Manual Handling","QTY":1,"UNITPRICEGROSS":15.29,
			 "TAXPERCENT":0,"ITEMCODE":null,"FINANCECODE":"43430","DOMAINID":5566}]}`,
		"PUT /api/accounting/invoice/3906271": `{"INVOICEID":3906271,"ERROR":null}`,
	}, &log)

	_, err := s.AddInvoiceLines(3906271, line())
	require.NoError(t, err)

	require.Len(t, log, 2)
	put := log[1]
	assert.Equal(t, "PUT", put.method)
	assert.Empty(t, put.query.Get("aItem"), "aItem must not travel in the query string")
	var items []map[string]any
	require.NoError(t, json.Unmarshal([]byte(put.form.Get("aItem")), &items))
	require.Len(t, items, 3, "both existing lines are re-sent, plus the new one")
	assert.Equal(t, float64(6469380), items[0]["ITEMID"])
	assert.Equal(t, "Workshop Booking for A & B #1", items[0]["DESCRIPTION"], "survives the form body intact")
	assert.Equal(t, float64(4677), items[0]["DOMAINID"])
	assert.NotContains(t, items[0], "HASH", "derived fields are not echoed")
	assert.Equal(t, "", items[1]["ITEMCODE"], "a null ITEMCODE is sent as empty, never omitted")
	assert.NotContains(t, items[2], "ITEMID", "the new line has no ITEMID")
	assert.Equal(t, float64(5566), items[2]["DOMAINID"])
}

func TestAddInvoiceLinesRefusesALockedInvoice(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{
		"GET /api/accounting/invoice/1": `{"INVOICEID":1,"AREITEMSLOCKED":true,"ITEMS":[]}`,
	}, &log)
	_, err := s.AddInvoiceLines(1, line())
	assert.ErrorContains(t, err, "locked")
	assert.Len(t, log, 1, "read only; no PUT")
}

func TestAddInvoiceLinesRefusesAnItemIDOnANewLine(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{}, &log)
	l := line()
	l.ItemID = Int(5)
	_, err := s.AddInvoiceLines(1, l)
	assert.Error(t, err)
	assert.Empty(t, log)
}

// --- ApproveInvoice -----------------------------------------------------------

func TestApproveInvoiceTakesTheGUID(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{
		"PUT /api/accounting/invoice/C0DD48E0-222C/approve": `{"PAYMENTURL":"https://example/payment"}`,
	}, &log)
	_, err := s.ApproveInvoice("C0DD48E0-222C")
	require.NoError(t, err)
	_, err = s.ApproveInvoice("")
	assert.Error(t, err)
	assert.Len(t, log, 1)
}

// --- ApplyPayment -------------------------------------------------------------

// A payment against a DRAFT invoice is refused with a 400 — and the money
// is recorded anyway. ApplyPayment must refuse before any money is sent.
func TestApplyPaymentRefusesADraftInvoice(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{
		"GET /api/accounting/invoice/3868280": `{"INVOICEID":3868280,"INVOICENR":"AUTO","INVGUID":"E088"}`,
	}, &log)
	_, _, err := s.ApplyPayment(InvoicePayment{ContactID: 14518907, InvoiceID: 3868280, Amount: "21.00"})
	assert.ErrorContains(t, err, "draft")
	for _, r := range log {
		assert.NotEqual(t, "/api/accounting/transaction/", r.path, "no transaction may be sent")
	}
}

func TestApplyPaymentOnAnIssuedInvoice(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{
		"GET /api/accounting/invoice/3868278": `{"INVOICEID":3868278,"INVOICENR":"3867608","INVGUID":"C0DD"}`,
		"POST /api/accounting/transaction/":   `{"TRANSACTIONID":"4656265","CONTACTID":"14518907","UNASSIGNEDAMOUNT":"0"}`,
	}, &log)
	txn, _, err := s.ApplyPayment(InvoicePayment{ContactID: 14518907, InvoiceID: 3868278, Amount: "21.00",
		Reference: "CLAUDE-DOMAIN-TEST-2"})
	require.NoError(t, err)
	assert.Equal(t, 4656265, int(txn.TransactionID))
	f := log[len(log)-1].form
	assert.Equal(t, "21.00", f.Get("amount"), "dollars, exactly as given")
	assert.Equal(t, "3868278", f.Get("invoiceID"))
}

func TestApplyPaymentValidates(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{}, &log)
	for name, p := range map[string]InvoicePayment{
		"no contact": {InvoiceID: 1, Amount: "1.00"},
		"no invoice": {ContactID: 1, Amount: "1.00"},
		"no amount":  {ContactID: 1, InvoiceID: 1},
		"zero":       {ContactID: 1, InvoiceID: 1, Amount: "0"},
		"not money":  {ContactID: 1, InvoiceID: 1, Amount: "$21"},
	} {
		_, _, err := s.ApplyPayment(p)
		assert.Error(t, err, name)
	}
	assert.Empty(t, log)
}

// --- Transactions -------------------------------------------------------------

// Amounts arrive as strings and can be fractional; the shape is staging's
// (30 Sep 2026) with a surcharged amount.
func TestTransactionsDecodesStringAmounts(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{
		"GET /api/accounting/transaction/": `[{"PAYMENTMETHOD":"Credit Card","CONTACTID":"14518907",
			"UNASSIGNEDAMOUNT":"61.36","TRANSACTIONID":"4329020","REFERENCE":"6581225",
			"TRANSACTIONTYPE":"Money Received","TRANSDATE":"2024-10-21 12:05","PAYMENTMETHODID":"2",
			"CURRENCY":"AUD","DESCRIPTION":"TXN: 00000000","AMOUNT":"61.36","TRANSACTIONTYPEID":"1"}]`,
	}, &log)
	txns, _, err := s.Transactions(14518907)
	require.NoError(t, err)
	require.Len(t, txns, 1)
	assert.Equal(t, 4329020, int(txns[0].TransactionID))
	assert.InDelta(t, 61.36, float64(txns[0].Amount), 0.001)
	assert.InDelta(t, 61.36, float64(txns[0].UnassignedAmount), 0.001)
	assert.Equal(t, "14518907", log[0].query.Get("contactID"))
}

func TestRefusedIn(t *testing.T) {
	assert.NoError(t, refusedIn(&Response{Body: `{"INVOICEID":1,"ERROR":null}`}))
	assert.NoError(t, refusedIn(&Response{Body: `[{"a":1}]`}))
	assert.Error(t, refusedIn(&Response{Body: `{"ERROR":true,"MESSAGES":"The invoice items are invalid."}`}))
}

// The body POST /accounting/transaction/ returned on staging, 30 Sep 2026.
// ORGANISATION is a NAME, ORGID a quoted id and UNASSIGNEDAMOUNT a decimal
// string; the Transaction type had them as numbers, so a payment that
// SUCCEEDED came back as a decode error — and a caller that retried on
// that error would record the money twice.
func TestCreateTransactionDecodesTheRealResponse(t *testing.T) {
	var log []sent
	s := routedClient(t, map[string]string{
		"POST /api/accounting/transaction/": LoadTestData("transaction_created.json"),
	}, &log)
	txn, _, err := s.CreateTransaction(map[string]string{"contactID": "14518907", "amount": "0.01", "invoiceID": "3868282"})
	require.NoError(t, err)
	assert.Equal(t, 4656268, int(txn.TransactionID))
	require.NotNil(t, txn.Organisation)
	assert.Equal(t, "Chris Joyce", *txn.Organisation)
	require.NotNil(t, txn.OrgID)
	assert.Equal(t, 1143205, int(*txn.OrgID))
	assert.InDelta(t, 0.01, float64(txn.Amount), 0.0001)
	assert.Equal(t, 0.0, float64(txn.UnassignedAmount))
	require.Len(t, txn.Fragments, 1)
	assert.Equal(t, 3868282, int(*txn.Fragments[0].InvoiceID))
}
