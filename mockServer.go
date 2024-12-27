package axcelerate

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// MockResponse defines the structure for a mock server response
type MockResponse struct {
	Endpoint     string // Partial or full endpoint path (e.g., "/contact/")
	StatusCode   int    // HTTP status code to return
	ResponseBody string // Response body to return
	Method       string // Optional: Defaults to "GET" if not specified
}

// SetupMockServer creates a mock server and client for testing.
// It matches requests to mock responses based on Endpoint patterns.
func SetupMockServer(t *testing.T, responses []MockResponse) (*httptest.Server, *Client) {
	t.Helper()

	mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t.Logf("Received request: %s %s", r.Method, r.URL.Path) // Log request for debugging
		for _, response := range responses {
			// Adjusted to check the full path including the "/api" prefix
			expectedPath := "/api" + response.Endpoint
			if strings.HasPrefix(r.URL.Path, expectedPath) && (response.Method == "" || response.Method == r.Method) {
				t.Logf("Matched response for: %s %s", response.Method, response.Endpoint) // Log match
				w.WriteHeader(response.StatusCode)
				_, _ = io.WriteString(w, response.ResponseBody)
				return
			}
		}

		// Default to 404 if no match is found
		t.Logf("No match for request: %s %s", r.Method, r.URL.Path)
		http.Error(w, "not found", http.StatusNotFound)
	}))

	mockClient, err := NewClient(
		"test-apitoken",
		"test-wstoken",
		BaseURL(mockServer.URL+"/api"),
	)
	if err != nil {
		t.Fatalf("Failed to create mock client: %v", err)
	}

	return mockServer, mockClient
}
