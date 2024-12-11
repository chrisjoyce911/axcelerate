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
