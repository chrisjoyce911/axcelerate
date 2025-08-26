package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/chrisjoyce911/axcelerate"
	"github.com/chrisjoyce911/axcelerate/example/files"
	"github.com/stretchr/testify/assert"
)

// TestFindME tests the findME function, which interacts with the axcelerate API client.
// It uses a table-driven approach to verify different scenarios by simulating API responses
// with an httptest server.
func TestFindME(t *testing.T) {
	// Define test cases using a table-driven structure
	tests := []struct {
		name         string  // Description of the test case
		mockResponse string  // The mock response returned by the API
		mockStatus   int     // HTTP status code returned by the API
		expectedName *string // Expected name from the findME function
		expectedErr  string  // Expected error message (if any)
	}{
		{
			name:         "Success with two contacts",
			mockResponse: `[{"CONTACTID":1,"GIVENNAME":"John"},{"CONTACTID":2,"GIVENNAME":"Doe"}]`,
			mockStatus:   http.StatusOK,
			expectedName: strPtr("Doe"), // Expecting the second contact's GivenName
			expectedErr:  "",
		},
		{
			name:         "No contacts found",
			mockResponse: `[]`, // No contacts in the response
			mockStatus:   http.StatusOK,
			expectedName: nil,
			expectedErr:  "second contact not found",
		},
		{
			name:         "API returns error",
			mockResponse: `Internal Server Error`,
			mockStatus:   http.StatusInternalServerError, // Simulating a server error
			expectedName: nil,
			expectedErr:  "API error", // Expecting the error message to start with "API error"
		},
	}

	// Iterate through each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up an httptest server to mock the API
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				// Check if the request is for the expected endpoint
				if strings.HasPrefix(r.URL.Path, "/api/contacts/search") {
					// Respond with the mock response and status code
					w.WriteHeader(tc.mockStatus)
					_, _ = io.WriteString(w, tc.mockResponse)
				} else {
					// Return a 404 if the endpoint is incorrect
					http.Error(w, "not found", http.StatusNotFound)
				}
			}))
			defer testServer.Close()

			// Create an axcelerate client pointing to the mock server
			client, err := axcelerate.NewClient(
				"test-apitoken", // Mock API token
				"test-wstoken",  // Mock WS token
				axcelerate.BaseURL(testServer.URL+"/api"), // Use the mock server's URL
			)
			assert.NoError(t, err, "Client creation should not return an error")

			// Call the findME function
			result, respBody, err := files.FindME(client)

			// Validate the result
			if tc.expectedErr != "" {
				// If an error is expected, check if it contains the expected message
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				// If no error is expected, ensure the result matches the expected name
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedName, result)
			}

			// Log the response body for debugging
			if respBody != nil {
				t.Logf("Response Body: %s", *respBody)
			}
		})
	}
}

func TestMockServerFindME(t *testing.T) {
	// Define test cases using a table-driven structure
	tests := []struct {
		name          string                    // Description of the test case
		mockResponses []axcelerate.MockResponse // Mock responses for the API
		expectedName  *string                   // Expected name from the findME function
		expectedErr   string                    // Expected error message (if any)
	}{
		{
			name: "Success with two contacts",
			mockResponses: []axcelerate.MockResponse{
				{
					Endpoint:     "/contacts/search",
					StatusCode:   http.StatusOK,
					ResponseBody: `[{"CONTACTID":1,"GIVENNAME":"John"},{"CONTACTID":2,"GIVENNAME":"Doe"}]`,
					Method:       "GET",
				},
			},
			expectedName: strPtr("Doe"),
			expectedErr:  "",
		},
		{
			name: "No contacts found",
			mockResponses: []axcelerate.MockResponse{
				{
					Endpoint:     "/contacts/search",
					StatusCode:   http.StatusOK,
					ResponseBody: `[]`,
					Method:       "GET",
				},
			},
			expectedName: nil,
			expectedErr:  "second contact not found",
		},
		{
			name: "API returns error",
			mockResponses: []axcelerate.MockResponse{
				{
					Endpoint:     "/contacts/search",
					StatusCode:   http.StatusInternalServerError,
					ResponseBody: `Internal Server Error`,
					Method:       "GET",
				},
			},
			expectedName: nil,
			expectedErr:  "API error",
		},
	}

	// Iterate through each test case
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Set up the mock server and client
			mockServer, mockClient := axcelerate.SetupMockServer(t, tc.mockResponses)
			defer mockServer.Close()

			// Call the findME function
			result, respBody, err := files.FindME(mockClient)

			// Validate the result
			if tc.expectedErr != "" {
				// If an error is expected, check if it contains the expected message
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				// If no error is expected, ensure the result matches the expected name
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedName, result)
			}

			// Log the response body for debugging
			if respBody != nil {
				t.Logf("Response Body: %s", *respBody)
			}
		})
	}
}

// strPtr is a helper function to create a pointer to a string value.
// This is useful for comparing pointer values in test cases.
func strPtr(s string) *string {
	return &s
}

func TestFindMEandVerifyUSI(t *testing.T) {
	// Define test cases
	tests := []struct {
		name             string
		mockContacts     string // Mock response for SearchContacts
		mockContactsErr  string // Error for SearchContacts
		mockVerifyUSI    string // Mock response for VerifyUSI
		mockVerifyUSIErr string // Error for VerifyUSI
		mockStatusCode   int    // HTTP status code for both APIs
		expectedResult   bool   // Expected result from findMEandVerifyUSI
		expectedErr      string // Expected error message (if any)
	}{
		{
			name:           "Success with verified USI",
			mockContacts:   `[{"CONTACTID":1,"GIVENNAME":"John"},{"CONTACTID":2,"GIVENNAME":"Doe"}]`,
			mockVerifyUSI:  `{"DATA":{},"MSG":"Verification successful","USI_VERIFIED":true}`,
			mockStatusCode: http.StatusOK,
			expectedResult: true,
			expectedErr:    "",
		},
		{
			name:           "Second contact not found",
			mockContacts:   `[{"CONTACTID":1,"GIVENNAME":"John"}]`,
			mockVerifyUSI:  "",
			mockStatusCode: http.StatusOK,
			expectedResult: false,
			expectedErr:    "second contact not found",
		},
		{
			name:            "Error in SearchContacts",
			mockContacts:    ``,
			mockContactsErr: "SearchContacts error",
			mockStatusCode:  http.StatusInternalServerError,
			expectedResult:  false,
			expectedErr:     "API error",
		},
		{
			name:             "Error in VerifyUSI",
			mockContacts:     `[{"CONTACTID":1,"GIVENNAME":"John"},{"CONTACTID":2,"GIVENNAME":"Doe"}]`,
			mockVerifyUSI:    "",
			mockVerifyUSIErr: "VerifyUSI error",
			mockStatusCode:   http.StatusOK,
			expectedResult:   false,
			expectedErr:      "API error",
		},
	}

	// Iterate through test cases
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Create a mock server
			testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if strings.HasPrefix(r.URL.Path, "/api/contacts/search") {
					if tc.mockContactsErr != "" {
						http.Error(w, tc.mockContactsErr, tc.mockStatusCode)
					} else {
						w.WriteHeader(tc.mockStatusCode)
						_, _ = io.WriteString(w, tc.mockContacts)
					}
				} else if strings.HasPrefix(r.URL.Path, "/api/contact/verifyUSI") {
					if tc.mockVerifyUSIErr != "" {
						http.Error(w, tc.mockVerifyUSIErr, tc.mockStatusCode)
					} else {
						w.WriteHeader(tc.mockStatusCode)
						_, _ = io.WriteString(w, tc.mockVerifyUSI)
					}
				} else {
					http.Error(w, "not found", http.StatusNotFound)
				}
			}))
			defer testServer.Close()

			// Create an axcelerate client
			client, err := axcelerate.NewClient(
				"test-apitoken",
				"test-wstoken",
				axcelerate.BaseURL(testServer.URL+"/api"),
			)
			assert.NoError(t, err)

			// Call the findMEandVerifyUSI function
			result, err := files.FindMEandVerifyUSI(client)

			// Validate results
			if tc.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}

func TestMockServerFindMEandVerifyUSI(t *testing.T) {
	// Define test cases
	tests := []struct {
		name           string
		mockResponses  []axcelerate.MockResponse
		expectedResult bool
		expectedErr    string
	}{
		{
			name: "Success with verified USI",
			mockResponses: []axcelerate.MockResponse{
				{
					Endpoint:     "/contacts/search",
					StatusCode:   http.StatusOK,
					ResponseBody: `[{"CONTACTID":1,"GIVENNAME":"John"},{"CONTACTID":2,"GIVENNAME":"Doe"}]`,
					Method:       "GET",
				},
				{
					Endpoint:     "/contact/verifyUSI",
					StatusCode:   http.StatusOK,
					ResponseBody: `{"DATA":{},"MSG":"Verification successful","USI_VERIFIED":true}`,
					Method:       "POST",
				},
			},
			expectedResult: true,
			expectedErr:    "",
		},
		{
			name: "Second contact not found",
			mockResponses: []axcelerate.MockResponse{
				{
					Endpoint:     "/contacts/search",
					StatusCode:   http.StatusOK,
					ResponseBody: `[{"CONTACTID":1,"GIVENNAME":"John"}]`,
					Method:       "GET",
				},
			},
			expectedResult: false,
			expectedErr:    "second contact not found",
		},
		{
			name: "Error in SearchContacts",
			mockResponses: []axcelerate.MockResponse{
				{
					Endpoint:     "/contacts/search",
					StatusCode:   http.StatusInternalServerError,
					ResponseBody: "SearchContacts error",
					Method:       "GET",
				},
			},
			expectedResult: false,
			expectedErr:    "API error",
		},
		{
			name: "Error in VerifyUSI",
			mockResponses: []axcelerate.MockResponse{
				{
					Endpoint:     "/contacts/search",
					StatusCode:   http.StatusOK,
					ResponseBody: `[{"CONTACTID":1,"GIVENNAME":"John"},{"CONTACTID":2,"GIVENNAME":"Doe"}]`,
					Method:       "GET",
				},
				{
					Endpoint:     "/contact/verifyUSI",
					StatusCode:   http.StatusInternalServerError,
					ResponseBody: "VerifyUSI error",
					Method:       "POST",
				},
			},
			expectedResult: false,
			expectedErr:    "API error",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			mockServer, mockClient := axcelerate.SetupMockServer(t, tc.mockResponses)
			defer mockServer.Close()

			result, err := files.FindMEandVerifyUSI(mockClient)

			if tc.expectedErr != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tc.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tc.expectedResult, result)
			}
		})
	}
}
