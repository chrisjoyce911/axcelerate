package axcelerate

import (
	"errors"
	"net/url"
	"testing"
	"time"
)

// MockClient is a mock implementation of the axcelerate Client.
type MockClient struct {
	MockNewRequest        func(dat map[string]string, action string) (*AxRequest, error)
	MockDo                func(req *AxRequest) (*Response, error)
	MockContactCreate     func(params map[string]string) (*Contact, *Response, error)
	MockCoursesEnrol      func(params map[string]string) (*Enrolment, *Response, error)
	MockCheckResponse     func(r *Response) error
	MockSanitizeURL       func(uri *url.URL) *url.URL
	MockResponse          func() *Response
	MockErrorResponse     func(r *Response) error
	MockFormatBool        func(b bool) string
	MockParmDate          func(d time.Time) string
	MockInt               func(v int) *int
	MockBool              func(v bool) *bool
	MockString            func(v string) *string
	MockAccountingService *MockAccountingService
	MockContactService    *MockContactService
	MockCoursesService    *MockCoursesService
	MockReportService     *MockReportService
	MockTemplateService   *MockTemplateService
	MockVenueService      *MockVenueService
}

// NewRequest mocks the NewRequest function.
func (m *MockClient) NewRequest(dat map[string]string, action string) (*AxRequest, error) {
	if m.MockNewRequest != nil {
		return m.MockNewRequest(dat, action)
	}
	return nil, errors.New("MockNewRequest not implemented")
}

// Do mocks the execution of an API request.
func (m *MockClient) Do(req *AxRequest) (*Response, error) {
	if m.MockDo != nil {
		return m.MockDo(req)
	}
	return nil, errors.New("MockDo not implemented")
}

// ContactCreate mocks creating a contact.
func (m *MockClient) ContactCreate(params map[string]string) (*Contact, *Response, error) {
	if m.MockContactCreate != nil {
		return m.MockContactCreate(params)
	}
	return nil, nil, errors.New("MockContactCreate not implemented")
}

// CheckResponse mocks checking the API response for errors.
func (m *MockClient) CheckResponse(r *Response) error {
	if m.MockCheckResponse != nil {
		return m.MockCheckResponse(r)
	}
	return nil
}

// SanitizeURL mocks sanitizing a URL.
func (m *MockClient) SanitizeURL(uri *url.URL) *url.URL {
	if m.MockSanitizeURL != nil {
		return m.MockSanitizeURL(uri)
	}
	return nil
}

// MockContactService provides mock methods for the ContactService.
type MockContactService struct {
	MockCreate func(params map[string]string) (*Contact, *Response, error)
}

// Create mocks creating a contact.
func (m *MockContactService) Create(params map[string]string) (*Contact, *Response, error) {
	if m.MockCreate != nil {
		return m.MockCreate(params)
	}
	return nil, nil, errors.New("MockCreate not implemented")
}

// Mock other services similarly...
// MockCoursesService, MockAccountingService, MockReportService, MockTemplateService, MockVenueService

// Example usage of the mock in a test.
func TestContactCreate(t *testing.T) {
	mockClient := &MockClient{
		MockContactCreate: func(params map[string]string) (*Contact, *Response, error) {
			if params["givenName"] == "John" && params["surname"] == "Doe" {
				return &Contact{GivenName: "John", Surname: "Doe"}, &Response{StatusCode: 201}, nil
			}
			return nil, nil, errors.New("Invalid parameters")
		},
	}

	params := map[string]string{
		"givenName": "John",
		"surname":   "Doe",
	}

	contact, resp, err := mockClient.ContactCreate(params)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
	if contact.GivenName != "John" {
		t.Errorf("Expected GivenName to be John, got %v", contact.GivenName)
	}
	if resp.StatusCode != 201 {
		t.Errorf("Expected status code 201, got %v", resp.StatusCode)
	}
}
