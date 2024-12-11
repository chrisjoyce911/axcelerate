package axcelerate

type MockContactService struct {
	MockContactCreate     func(parms map[string]string) (*Contact, *Response, error)
	MockContactEnrolments func(contactID int, parms map[string]string) ([]ContactEnrolment, *Response, error)
	MockContactUpdate     func(contactID int, parms map[string]string) (Contact, *Response, error)
	MockContactSearch     func(parms map[string]string) ([]Contact, *Response, error)
}

func (m *MockContactService) ContactCreate(parms map[string]string) (*Contact, *Response, error) {
	if m.MockContactCreate != nil {
		return m.MockContactCreate(parms)
	}
	return nil, nil, nil
}

func (m *MockContactService) ContactEnrolments(contactID int, parms map[string]string) ([]ContactEnrolment, *Response, error) {
	if m.MockContactEnrolments != nil {
		return m.MockContactEnrolments(contactID, parms)
	}
	return nil, nil, nil
}

func (m *MockContactService) ContactUpdate(contactID int, parms map[string]string) (Contact, *Response, error) {
	if m.MockContactUpdate != nil {
		return m.MockContactUpdate(contactID, parms)
	}
	return Contact{}, nil, nil
}

func (m *MockContactService) ContactSearch(parms map[string]string) ([]Contact, *Response, error) {
	if m.MockContactSearch != nil {
		return m.MockContactSearch(parms)
	}
	return nil, nil, nil
}
