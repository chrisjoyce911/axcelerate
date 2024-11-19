package axcelerate

import (
	"io"
	"net/http"
	"net/url"
	"strings"
)

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req AxRequest, v interface{}) (*Response, error) {

	c.rl.Take()

	var body io.Reader

	if req.method == "POST" {
		body = strings.NewReader(req.data.Encode()) // Encode form data for POST
	} else {
		req.url.RawQuery, _ = url.QueryUnescape(req.data.Encode()) // For non-POST, add query parameters
	}

	// Create the HTTP request
	thisReq, err := http.NewRequest(req.method, req.url.String(), body)
	if err != nil {
		return nil, err
	}

	thisReq.Header.Set("Accept", "application/json")
	thisReq.Header.Set("apitoken", c.apitoken)
	thisReq.Header.Set("wstoken", c.wstoken)

	// Set Content-Type for POST requests
	if req.method == "POST" {
		thisReq.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}

	resp, err := c.client.Do(thisReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := newResponse(resp)
	err = CheckResponse(response)
	if err != nil {
		// even though there was an error, we still return the response
		// in case the caller wants to inspect it further
		return response, err
	}

	return response, err
}

func do(c *Client, m string, p Params, a interface{}) (*Response, error) {

	req, err := c.NewRequest(p.parms, p.u)
	if err != nil {
		return nil, err
	}
	req.method = m

	resp, err := c.Do(*req, a)
	if err != nil {
		return resp, err
	}

	return resp, err
}
