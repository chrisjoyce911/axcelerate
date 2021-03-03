package axcelerate

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client for the axcelerate SDK
type Client struct {
	apitoken   string
	wstoken    string
	baseURL    *url.URL
	httpClient *http.Client

	common service // Reuse a single struct instead of allocating one for each service on the heap.

	Contact *ContactService
}

type service struct {
	client *Client
}

// NewAuthClient for all operations
func NewAuthClient(baseURL *url.URL, apitoken, wstoken string, httpClient *http.Client) *Client {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	c := &Client{
		apitoken:   apitoken,
		wstoken:    wstoken,
		baseURL:    baseURL,
		httpClient: httpClient,
	}

	c.common.client = c

	c.Contact = (*ContactService)(&c.common)

	return c
}

func (c *Client) newRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)
	var buf io.ReadWriter
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, u.String(), buf)
	if err != nil {
		return nil, err
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*http.Response, error) {

	req.Header.Set("Accept", "application/json")
	req.Header.Set("apitoken", c.apitoken)
	req.Header.Set("wstoken", c.wstoken)

	fmt.Println(req.URL)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(v)
	return resp, err
}
