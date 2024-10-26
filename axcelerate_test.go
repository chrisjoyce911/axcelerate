package axcelerate

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"reflect"
	"strings"
	"testing"
)

type wonkyReader struct{}

func (wr wonkyReader) Read(p []byte) (n int, err error) {
	return 0, nil
}

type testDoer struct {
	response     string
	responseCode int
	http.Header
}

func (nd testDoer) Do(*http.Request) (*http.Response, error) {
	return &http.Response{
		Body:       io.NopCloser(bytes.NewReader([]byte(nd.response))),
		StatusCode: nd.responseCode,
		Header:     nd.Header,
	}, nil
}

type values map[string]string

// RoundTripFunc .
type RoundTripFunc func(req *http.Request) *http.Response

// RoundTrip .
func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

// NewTestClient returns *http.Client with Transport replaced to avoid making real calls
func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

// LoadTestData get test data for body response
func LoadTestData(file string) string {

	b, err := os.ReadFile(fmt.Sprintf("testdata/%s", file))
	if err != nil {
		fmt.Print(err)
	}
	return string(b)
}

func TestAddFormValues(t *testing.T) {
	tests := []struct {
		name string
		opt  map[string]string
		want url.Values
	}{
		{
			name: "Single parameter",
			opt:  map[string]string{"param1": "value1"},
			want: url.Values{"param1": {"value1"}},
		},
		{
			name: "Multiple parameters",
			opt:  map[string]string{"param1": "value1", "param2": "value2"},
			want: url.Values{"param1": {"value1"}, "param2": {"value2"}},
		},
		{
			name: "Empty map",
			opt:  map[string]string{},
			want: url.Values{},
		},
		{
			name: "Parameter with empty value",
			opt:  map[string]string{"param1": ""},
			want: url.Values{"param1": {""}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addFormValues(tt.opt)
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("addFormValues() = %v, want %v", *got, tt.want)
			}
		})
	}
}

func Test_newResponse(t *testing.T) {
	tests := []struct {
		name           string
		httpResponse   *http.Response
		expectedStatus string
		expectedCode   int
		expectedBody   string
		expectedLength int64
	}{
		{
			name: "Basic response with status 200",
			httpResponse: &http.Response{
				Status:        "200 OK",
				StatusCode:    200,
				Body:          io.NopCloser(strings.NewReader("Hello, world!")),
				ContentLength: 13,
			},
			expectedStatus: "200 OK",
			expectedCode:   200,
			expectedBody:   "Hello, world!",
			expectedLength: 13,
		},
		{
			name: "Empty body response with status 404",
			httpResponse: &http.Response{
				Status:        "404 Not Found",
				StatusCode:    404,
				Body:          io.NopCloser(strings.NewReader("")),
				ContentLength: 0,
			},
			expectedStatus: "404 Not Found",
			expectedCode:   404,
			expectedBody:   "",
			expectedLength: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resp := newResponse(tt.httpResponse)

			// Validate Status
			if resp.Status != tt.expectedStatus {
				t.Errorf("Expected Status %s, got %s", tt.expectedStatus, resp.Status)
			}

			// Validate StatusCode
			if resp.StatusCode != tt.expectedCode {
				t.Errorf("Expected StatusCode %d, got %d", tt.expectedCode, resp.StatusCode)
			}

			// Validate Body
			if resp.Body != tt.expectedBody {
				t.Errorf("Expected Body %s, got %s", tt.expectedBody, resp.Body)
			}

			// Validate ContentLength
			if resp.ContentLength != tt.expectedLength {
				t.Errorf("Expected ContentLength %d, got %d", tt.expectedLength, resp.ContentLength)
			}
		})
	}
}

func TestSanitizeURL(t *testing.T) {
	tests := []struct {
		name        string
		inputURL    *url.URL
		expectedURL string
	}{
		{
			name: "URL with accesskey",
			inputURL: func() *url.URL {
				u, _ := url.Parse("https://example.com/api?accesskey=12345&otherparam=value")
				return u
			}(),
			expectedURL: "https://example.com/api?accesskey=REDACTED&otherparam=value",
		},
		{
			name: "URL without accesskey",
			inputURL: func() *url.URL {
				u, _ := url.Parse("https://example.com/api?otherparam=value")
				return u
			}(),
			expectedURL: "https://example.com/api?otherparam=value",
		},
		{
			name: "URL with empty accesskey",
			inputURL: func() *url.URL {
				u, _ := url.Parse("https://example.com/api?accesskey=&otherparam=value")
				return u
			}(),
			expectedURL: "https://example.com/api?accesskey=&otherparam=value",
		},
		{
			name:        "Nil URL",
			inputURL:    nil,
			expectedURL: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := sanitizeURL(tt.inputURL)

			// Handle nil URL case
			if tt.inputURL == nil && result != nil {
				t.Errorf("Expected nil result, got %v", result)
				return
			} else if tt.inputURL == nil && result == nil {
				return
			}

			// Compare the sanitized URL with the expected result
			if result.String() != tt.expectedURL {
				t.Errorf("sanitizeURL() = %v, want %v", result.String(), tt.expectedURL)
			}
		})
	}
}

func TestCheckResponse(t *testing.T) {
	tests := []struct {
		name       string
		response   *Response
		wantErr    bool
		errMessage string
	}{
		{
			name: "Successful response with 200 status",
			response: &Response{
				StatusCode: 200,
				Body:       "OK",
			},
			wantErr:    false,
			errMessage: "",
		},
		{
			name: "Successful response with 299 status",
			response: &Response{
				StatusCode: 299,
				Body:       "OK",
			},
			wantErr:    false,
			errMessage: "",
		},
		{
			name: "Error response with 400 status",
			response: &Response{
				StatusCode: 400,
				Body:       "Bad Request",
			},
			wantErr:    true,
			errMessage: "Bad Request",
		},

		{
			name: "Error response with 429 Too Many Requests Retry-After: 30",
			response: &Response{
				StatusCode: 429,
				Body:       "Bad Request",
			},
			wantErr:    true,
			errMessage: "Bad Request",
		},

		{
			name: "Error response with 500 status",
			response: &Response{
				StatusCode: 500,
				Body:       "Internal Server Error",
			},
			wantErr:    true,
			errMessage: "Internal Server Error",
		},
		{
			name: "Error response with empty body",
			response: &Response{
				StatusCode: 404,
				Body:       "",
			},
			wantErr:    true,
			errMessage: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CheckResponse(tt.response)

			// Check if an error was expected
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckResponse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// If an error was expected, check the error message
			if tt.wantErr && err.Error() != tt.errMessage {
				t.Errorf("CheckResponse() error message = %v, want %v", err.Error(), tt.errMessage)
			}
		})
	}
}
