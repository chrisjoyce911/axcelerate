// mockexample/mocktest.go

package main

import (
	"testing"

	"github.com/chrisjoyce911/axcelerate"
)

func TestCreateContact(t *testing.T) {
	// Initialize the mock client
	mockClient := axcelerate.NewMockClient()

	// Configure the mock Contact service
	mockClient.Contact.MockContactCreate = func(parms map[string]string) (*axcelerate.Contact, *axcelerate.Response, error) {
		return &axcelerate.Contact{ContactID: 123, GivenName: "John", Surname: "Doe"}, nil, nil
	}

	// Use the mock Contact service
	contact, _, err := mockClient.Contact.ContactCreate(map[string]string{
		"givenName": "John",
		"surname":   "Doe",
	})

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if contact.ContactID != 123 {
		t.Errorf("expected contact ID 123, got %d", contact.ContactID)
	}
}

func TestContactSearch(t *testing.T) {
	// Initialize the mock client
	mockClient := axcelerate.NewMockClient()

	// Configure the mock ContactSearch method
	mockClient.Contact.MockContactSearch = func(parms map[string]string) ([]axcelerate.Contact, *axcelerate.Response, error) {
		return []axcelerate.Contact{
			{ContactID: 1, GivenName: "John", Surname: "Doe"},
			{ContactID: 2, GivenName: "Jane", Surname: "Smith"},
		}, nil, nil
	}

	// Use the mock Contact service
	params := map[string]string{"q": "John Doe"}
	contacts, _, err := mockClient.Contact.ContactSearch(params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(contacts) != 2 {
		t.Errorf("expected 2 contacts, got %d", len(contacts))
	}

	if contacts[0].GivenName != "John" {
		t.Errorf("expected first contact given name to be 'John', got %s", contacts[0].GivenName)
	}

	if contacts[1].Surname != "Smith" {
		t.Errorf("expected second contact surname to be 'Smith', got %s", contacts[1].Surname)
	}
}
