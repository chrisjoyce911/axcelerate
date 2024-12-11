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
