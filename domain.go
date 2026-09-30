package axcelerate

import "encoding/json"

// Domain is an aXcelerate domain: the reporting unit an invoice line, a
// course instance or a contact belongs to — a venue or region ("Geelong"),
// "ONLINE", "Head Office" and so on.
type Domain struct {
	DomainID   int    `json:"DOMAINID"`
	DomainName string `json:"DOMAINNAME"`
}

/*
Domains lists every domain on the account (GET /domains).

The endpoint is NOT in aXcelerate's published API docs, but it is live and
answers the same 135 ids on staging and production (checked 29 Sep 2026).

It matters because aXcelerate does NOT validate a DOMAINID: an invoice line
created with a made-up id (104157) is stored without complaint and shows a
blank Domain column in the UI. This list is the only way to check an id is
real before writing it, or to turn an id into the name the UI shows.
*/
func (s *AccountingService) Domains() ([]Domain, *Response, error) {
	var obj []Domain

	resp, err := do(s.client, "GET", Params{parms: map[string]string{}, u: "/domains"}, obj)
	if err != nil {
		return obj, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
