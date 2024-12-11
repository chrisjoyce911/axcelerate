package axcelerate

type MockTemplateService struct {
	MockTemplateEmail func(params TemplateEmailParams) (*EmailResponse, *Response, error)
}

func (m *MockTemplateService) TemplateEmail(params TemplateEmailParams) (*EmailResponse, *Response, error) {
	if m.MockTemplateEmail != nil {
		return m.MockTemplateEmail(params)
	}
	return nil, nil, nil
}
