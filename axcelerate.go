// Package axcelerate provides a simple interface to theRESTFul API for interfacing with aXcelerate.
// All responses can be access as the returned as JSON structures as well as an object
// All endpoint parameters can be set for all endpoints using a simple map
package axcelerate

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/ratelimit"
)

// Client for the axcelerate SDK
type Client struct {
	apitoken string
	wstoken  string
	BaseURL  *url.URL
	client   *http.Client
	rl       ratelimit.Limiter

	APIEndSux string

	Contact    *ContactService
	Courses    *CoursesService
	Accounting *AccountingService
	Report     *ReportService
}

// APIerr may happen along with a status code
// aXcelerate has a standard error struct that is returned whenever something goes wrong.
// This includes validation errors, as well as unexpected application errors.
type APIerr struct {
	Code     string `json:"code"`
	Data     string `json:"data"`
	Details  string `json:"details"`
	Error    bool   `json:"error"`
	Messages string `json:"messages"`
}

// NewClient for all operations
func NewClient(apitoken, wstoken string, baseURL *url.URL, httpClient *http.Client) *Client {

	if baseURL == nil {
		baseURL, _ = url.Parse("https://awfa.app.axcelerate.com")
	}

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		apitoken:  apitoken,
		wstoken:   wstoken,
		BaseURL:   baseURL,
		client:    httpClient,
		APIEndSux: "api",
		rl:        ratelimit.New(150, ratelimit.Per(time.Minute)), // per Minute]
	}

	c.Contact = &ContactService{client: c}
	c.Courses = &CoursesService{client: c}
	c.Accounting = &AccountingService{client: c}
	c.Report = &ReportService{client: c}
	return c
}

// Params specifies the optional parameters to various List methods that
// support pagination.
type Params struct {
	parms map[string]string
	u     string
}

// A AxRequest manages communication with the axe API.
type AxRequest struct {
	data   *url.Values
	method string
	url    *url.URL
}

// addFormValues adds the parameters in opt as URL values parameters.
func addFormValues(opt map[string]string) *url.Values {
	uv := url.Values{}
	for k, v := range opt {
		uv.Set(k, v)
	}
	return &uv
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(dat map[string]string, action string) (*AxRequest, error) {
	rel, err := url.Parse(fmt.Sprintf("%s%s", c.APIEndSux, action))

	if err != nil {
		return nil, err
	}
	u := c.BaseURL.ResolveReference(rel)

	return &AxRequest{url: u, data: addFormValues(dat)}, nil
}

// Response is a aXcelerate API response.  This wraps the standard http.Response
// returned from aXcelerate
type Response struct {
	Status        string // e.g. "200 OK"
	StatusCode    int    // e.g. 200
	Body          string
	ContentLength int64
}

// newResponse creates a new Response for the provided http.Response.
func newResponse(r *http.Response) *Response {
	body, _ := io.ReadAll(r.Body)
	response := &Response{
		Status:        r.Status,
		StatusCode:    r.StatusCode,
		Body:          string(body),
		ContentLength: r.ContentLength,
	}
	return response
}

// Do sends an API request and returns the API response.  The API response is
// JSON decoded and stored in the value pointed to by v, or returned as an
// error if an API error has occurred.  If v implements the io.Writer
// interface, the raw response body will be written to v, without attempting to
// first decode it.
func (c *Client) Do(req AxRequest, v interface{}) (*Response, error) {

	c.rl.Take()
	req.url.RawQuery, _ = url.QueryUnescape(req.data.Encode())

	thisReq, err := http.NewRequest(req.method, req.url.String(), nil)
	if err != nil {
		return nil, err
	}

	thisReq.Header.Set("Accept", "application/json")
	thisReq.Header.Set("apitoken", c.apitoken)
	thisReq.Header.Set("wstoken", c.wstoken)

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

/*
ErrorResponse reports one or more errors caused by an API request.
*/
type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:"message"` // error message
	Errors   []Error        `json:"errors"`  // more detail on individual errors
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, sanitizeURL(r.Response.Request.URL),
		r.Response.StatusCode, r.Message, r.Errors)
}

// sanitizeURL redacts the client_secret parameter from the URL which may be
// exposed to the user, specifically in the ErrorResponse error message.
func sanitizeURL(uri *url.URL) *url.URL {
	if uri == nil {
		return nil
	}
	params := uri.Query()
	if len(params.Get("accesskey")) > 0 {
		params.Set("accesskey", "REDACTED")
		uri.RawQuery = params.Encode()
	}
	return uri
}

/*
Error reports more details on an individual error in an ErrorResponse.

	 These are the possible validation error codes:

		 missing:
			 resource does not exist
		 missing_field:
			 a required field on a resource has not been set
		 invalid:
			 the formatting of a field is invalid
		 already_exists:
			 another resource has the same valid as this field
*/
type Error struct {
	Resource string `json:"resource"` // resource on which the error occurred
	Field    string `json:"field"`    // field on which the error occurred
	Code     string `json:"code"`     // validation error code
}

func (e *Error) Error() string {
	return fmt.Sprintf("%v error caused by %v field on %v resource",
		e.Code, e.Field, e.Resource)
}

// CheckResponse checks the API response for errors, and returns them if
// present.  A response is considered an error if it has a status code outside
// the 200 range.  API error responses are expected to have either no response
// body, or a JSON response body that maps to ErrorResponse.  Any other
// response body will be silently ignored.
func CheckResponse(r *Response) error {
	if c := r.StatusCode; 200 <= c && c <= 299 {
		return nil
	}
	return errors.New(r.Body)
}

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool {
	p := new(bool)
	*p = v
	return p
}

// Int is a helper routine that allocates a new int32 value
// to store v and returns a pointer to it, but unlike Int32
// its argument value is an int.
func Int(v int) *int {
	p := new(int)
	*p = v
	return p
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string {
	p := new(string)
	*p = v
	return p
}

// FormatBool returns "1" or "0" according to the value of b
func FormatBool(b bool) string {
	if b {
		return "1"
	}
	return "0"
}

// ParmDate returns time as string in YYYY-MM-DD hh:mm format
func ParmDate(d time.Time) string {
	return strings.Replace(d.Format("2006-01-02 15:04"), " ", "%20", -1)
}
