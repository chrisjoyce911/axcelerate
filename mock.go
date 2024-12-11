// mock.go

package axcelerate

import (
	"net/http"
	"net/url"
)

type MockClient struct {
	apitoken string
	wstoken  string
	BaseURL  *url.URL
	client   *http.Client

	Contact    *MockContactService
	Courses    *MockCoursesService
	Accounting *MockAccountingService
	Report     *MockReportService
	Template   *MockTemplateService
	Venue      *MockVenueService
}

func NewMockClient() *MockClient {
	return &MockClient{
		Contact:    &MockContactService{},
		Courses:    &MockCoursesService{},
		Accounting: &MockAccountingService{},
		Report:     &MockReportService{},
		Template:   &MockTemplateService{},
		Venue:      &MockVenueService{},
	}
}

func NewMockClientWithServices(
	contact *MockContactService,
	courses *MockCoursesService,
	accounting *MockAccountingService,
	report *MockReportService,
	template *MockTemplateService,
	venue *MockVenueService,
) *MockClient {
	return &MockClient{
		Contact:    contact,
		Courses:    courses,
		Accounting: accounting,
		Report:     report,
		Template:   template,
		Venue:      venue,
	}
}

func NewMockAccountingServiceWithDefaults() *MockAccountingService {
	return &MockAccountingService{
		MockPaymentURL: func(reference, invoiceGUID, redirectURL, cancelURL string) (*PaymentRequest, *Response, error) {
			return &PaymentRequest{
				HTML:   "<form>Mock Payment Form</form>",
				Action: "https://mock-payment-url.com",
			}, defaultMockResponse(), nil
		},
		MockGetInvoice: func(invoiceID int) (*Invoice, *Response, error) {
			return defaultInvoice(), defaultMockResponse(), nil
		},
	}
}

func NewFullyMockedClient() *MockClient {
	return NewMockClientWithServices(
		&MockContactService{}, // Add meaningful defaults if needed
		&MockCoursesService{}, // Add meaningful defaults if needed
		NewMockAccountingServiceWithDefaults(),
		&MockReportService{},   // Add meaningful defaults if needed
		&MockTemplateService{}, // Add meaningful defaults if needed
		&MockVenueService{},    // Add meaningful defaults if needed
	)
}
