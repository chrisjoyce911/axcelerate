package exampleTestServer

import (
	"fmt"
	"log"

	"github.com/chrisjoyce911/axcelerate"
)

func findME(client *axcelerate.Client) (*string, *string, error) {
	params := map[string]string{"name": "John"}

	contacts, resp, err := client.Contact.SearchContacts(params)

	log.Printf("Response Body: %v\n", resp.Body)
	log.Printf("Response Status Code: %v\n", resp.StatusCode)

	if err != nil {
		return nil, &resp.Body, fmt.Errorf("API error: %v", err)
	}

	if len(contacts) > 1 {
		return &contacts[1].GivenName, &resp.Body, nil
	}

	return nil, &resp.Body, fmt.Errorf("second contact not found")
}

func findMEandVerifyUSI(client *axcelerate.Client) (bool, error) {
	params := map[string]string{"name": "John"}

	contacts, resp, err := client.Contact.SearchContacts(params)

	log.Printf("Response Body: %v\n", resp.Body)
	log.Printf("Response Status Code: %v\n", resp.StatusCode)

	if err != nil {
		return false, fmt.Errorf("API error: %v", err)
	}

	if len(contacts) > 1 {
		status, _, err := client.Contact.VerifyUSI(contacts[1].ContactID)
		if err != nil {
			return false, fmt.Errorf("API error: %v", err)
		}

		return status.UsiVerified, nil
	}

	return false, fmt.Errorf("second contact not found")
}
