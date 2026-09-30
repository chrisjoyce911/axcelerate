package axcelerate

import (
	"bytes"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The shape GET /domains returned on 29 Sep 2026 (a slice of the 135).
const domainsBody = `[{"DOMAINID":514,"DOMAINNAME":"ZZZ-Training Compliance"},` +
	`{"DOMAINID":5494,"DOMAINNAME":"Head Office"},` +
	`{"DOMAINID":5566,"DOMAINNAME":"ONLINE"},` +
	`{"DOMAINID":4527,"DOMAINNAME":"Geelong"}]`

func TestAccountingService_Domains(t *testing.T) {
	tclient := NewTestClient(func(req *http.Request) *http.Response {
		assert.Equal(t, http.MethodGet, req.Method)
		assert.Contains(t, req.URL.Path, "/domains")
		return &http.Response{
			StatusCode: 200,
			Body:       io.NopCloser(bytes.NewBufferString(domainsBody)),
			Header:     make(http.Header),
		}
	})
	client, _ := NewClient("", "", HttpClient(tclient))
	s := &AccountingService{client: client}

	got, resp, err := s.Domains()
	require.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	require.Len(t, got, 4)
	assert.Equal(t, Domain{DomainID: 5566, DomainName: "ONLINE"}, got[2])
	assert.Equal(t, Domain{DomainID: 5494, DomainName: "Head Office"}, got[1])
}

func TestAccountingService_DomainsError(t *testing.T) {
	tclient := NewTestClient(func(req *http.Request) *http.Response {
		return &http.Response{
			StatusCode: 401,
			Body:       io.NopCloser(bytes.NewBufferString(`{"ERROR":true,"MESSAGES":"Invalid token"}`)),
			Header:     make(http.Header),
		}
	})
	client, _ := NewClient("", "", HttpClient(tclient))
	s := &AccountingService{client: client}

	_, _, err := s.Domains()
	assert.Error(t, err)
}
