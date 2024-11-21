package axcelerate

import "errors"

// MockVenueService provides mock methods for the VenueService.
type MockVenueService struct {
	MockCreate func(params map[string]string) (*Venue, *Response, error)
	MockGet    func(params map[string]string) (*Venue, *Response, error)
}

// Create mocks creating a venue.
func (m *MockVenueService) Create(params map[string]string) (*Venue, *Response, error) {
	if m.MockCreate != nil {
		return m.MockCreate(params)
	}
	return nil, nil, errors.New("MockCreate not implemented")
}

// Get mocks retrieving a venue.
func (m *MockVenueService) Get(params map[string]string) (*Venue, *Response, error) {
	if m.MockGet != nil {
		return m.MockGet(params)
	}
	return nil, nil, errors.New("MockGet not implemented")
}

// MockAccountingService provides mock methods for the AccountingService.
type MockAccountingService struct {
	MockCreateInvoice func(params map[string]string) (*Invoice, *Response, error)
}

// CreateInvoice mocks creating an invoice.
func (m *MockAccountingService) CreateInvoice(params map[string]string) (*Invoice, *Response, error) {
	if m.MockCreateInvoice != nil {
		return m.MockCreateInvoice(params)
	}
	return nil, nil, errors.New("MockCreateInvoice not implemented")
}

// MockTemplateService provides mock methods for the TemplateService.
type MockTemplateService struct {
	MockRender func(params map[string]string) (string, *Response, error)
}

// Render mocks rendering a template.
func (m *MockTemplateService) Render(params map[string]string) (string, *Response, error) {
	if m.MockRender != nil {
		return m.MockRender(params)
	}
	return "", nil, errors.New("MockRender not implemented")
}

// MockCoursesService provides mock methods for the CoursesService.
type MockCoursesService struct {
	MockEnrol func(params map[string]string) (*Enrolment, *Response, error)
}

// Enrol mocks enrolling a contact in a course.
func (m *MockCoursesService) Enrol(params map[string]string) (*Enrolment, *Response, error) {
	if m.MockEnrol != nil {
		return m.MockEnrol(params)
	}
	return nil, nil, errors.New("MockEnrol not implemented")
}

// MockReportService provides mock methods for the ReportService.
type MockReportService struct {
	MockGenerate func(params map[string]string) (*SavedReport, *Response, error)
}

// Generate mocks generating a report.
func (m *MockReportService) Generate(params map[string]string) (*SavedReport, *Response, error) {
	if m.MockGenerate != nil {
		return m.MockGenerate(params)
	}
	return nil, nil, errors.New("MockGenerate not implemented")
}
