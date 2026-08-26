package axcelerate

import (
	"bytes"
	"io"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const creditNoteBody = `{"CREDITNOTEID":"88123","CREDITNOTENR":"CN-00042","CONTACTID":"14754152",` +
	`"GUID":"a1b2c3","FIRSTNAME":"Jinky","LASTNAME":"Abayan","BALANCE":"59.00",` +
	`"PRICEGROSS":"59.00","ISPAID":false,"ISVOID":false}`

func TestAccountingService_CreateCreditNote(t *testing.T) {
	var got url.Values
	tclient := NewTestClient(func(req *http.Request) *http.Response {
		body, _ := io.ReadAll(req.Body)
		got, _ = url.ParseQuery(string(body))
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(creditNoteBody)),
			Header:     make(http.Header),
		}
	})
	client, _ := NewClient("", "", HttpClient(tclient))
	s := &AccountingService{client: client}

	note, resp, err := s.CreateCreditNote(CreditNoteRequest{
		ContactID: 14754152,
		FirstName: "Jinky",
		Surname:   "Abayan",
		Items: []CreditNoteItem{{
			Description:    "Payment held — enrolment could not be completed",
			Qty:            1,
			ItemCode:       "CREDIT",
			TaxPercent:     "0.00",
			UnitPriceGross: "59.00",
			FinanceCode:    "14180",
		}},
		Date: "2026-08-26",
	})
	require.NoError(t, err)
	require.NotNil(t, resp)

	// The wire shape the API documents.
	assert.Equal(t, "14754152", got.Get("contactID"))
	assert.Equal(t, "Jinky", got.Get("firstname"))
	assert.Equal(t, "Abayan", got.Get("surname"))
	assert.Contains(t, got.Get("aItem"), `"UNITPRICEGROSS":"59.00"`)
	assert.Contains(t, got.Get("aItem"), `"TAXPERCENT":"0.00"`)
	assert.Contains(t, got.Get("aItem"), `"FINANCECODE":"14180"`)
	// The API documents creditnoteDate as optional and then refuses the
	// call without it (production, 26 Aug 2026).
	assert.Equal(t, "2026-08-26", got.Get("creditnoteDate"))

	// The response the caller acts on.
	assert.Equal(t, 88123, int(note.CreditNoteID))
	assert.Equal(t, 14754152, int(note.ContactID))
	require.NotNil(t, note.Balance)
	assert.Equal(t, "59.00", *note.Balance)
	assert.False(t, note.IsPaid)
}

// The two inputs the API requires and we must not send without.
func TestAccountingService_CreateCreditNoteValidates(t *testing.T) {
	s := &AccountingService{}

	_, _, err := s.CreateCreditNote(CreditNoteRequest{FirstName: "A", Surname: "B",
		Items: []CreditNoteItem{{Description: "x"}}})
	assert.Error(t, err, "contactID is required")

	_, _, err = s.CreateCreditNote(CreditNoteRequest{ContactID: 1, FirstName: "A", Surname: "B"})
	assert.Error(t, err, "at least one item is required")
}

func TestAccountingService_CreditNotes(t *testing.T) {
	tclient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, "14754152", req.URL.Query().Get("contactID"))
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString("[" + creditNoteBody + "]")),
			Header:     make(http.Header),
		}
	})
	client, _ := NewClient("", "", HttpClient(tclient))
	s := &AccountingService{client: client}

	notes, _, err := s.CreditNotes(14754152)
	require.NoError(t, err)
	require.Len(t, notes, 1)
	assert.Equal(t, 88123, int(notes[0].CreditNoteID))
}

// Left empty, the client supplies today's date rather than reproducing
// the API's UNDEFINED_DATE refusal.
func TestAccountingService_CreateCreditNoteDefaultsDate(t *testing.T) {
	var got url.Values
	tclient := NewTestClient(func(req *http.Request) *http.Response {
		body, _ := io.ReadAll(req.Body)
		got, _ = url.ParseQuery(string(body))
		return &http.Response{StatusCode: 200,
			Body: io.NopCloser(bytes.NewBufferString(creditNoteBody)), Header: make(http.Header)}
	})
	client, _ := NewClient("", "", HttpClient(tclient))
	s := &AccountingService{client: client}

	_, _, err := s.CreateCreditNote(CreditNoteRequest{
		ContactID: 1, FirstName: "A", Surname: "B",
		Items: []CreditNoteItem{{Description: "x", Qty: 1, ItemCode: "CREDIT",
			TaxPercent: "0.00", UnitPriceGross: "1.00"}},
	})
	require.NoError(t, err)
	assert.NotEmpty(t, got.Get("creditnoteDate"), "always sent")
}
