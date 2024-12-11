package axcelerate

import "errors"

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
