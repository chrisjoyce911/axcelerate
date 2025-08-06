package main

import (
	"fmt"
	"log"

	"github.com/chrisjoyce911/axcelerate"
)

// contactNoteAddExample demonstrates how to add a note to a contact
func contactNoteAddExample() {
	contactID := 11300044

	params := map[string]string{
		"contactNote": "test note by example",
		"noteTypeID":  "27938",
	}

	log.Printf("Adding note to contact %d with params: %+v", contactID, params)

	response, httpResp, err := client.Contact.NoteAdd(contactID, params)

	log.Println("-----")
	log.Printf("Note Add Response\n%+v", response)
	log.Println("-----")
	log.Printf("HTTP Response Body\n%s", httpResp.Body)
	log.Printf("HTTP Status Code: %d", httpResp.StatusCode)
	log.Println("-----")

	if err != nil {
		log.Printf("Error occurred: %v", err)
		return
	}

	// Check if the API returned an error response
	if response.ERROR != nil && *response.ERROR {
		log.Printf("API Error: %s", *response.MESSAGES)
		if response.CODE != nil {
			log.Printf("Error Code: %s", *response.CODE)
		}
		if response.DETAILS != nil {
			log.Printf("Error Details: %s", *response.DETAILS)
		}
	} else {
		// Success case
		log.Printf("Note added successfully!")
		if response.NOTEID != nil {
			log.Printf("Note ID: %d", *response.NOTEID)
		}
		if response.MESSAGE != nil {
			log.Printf("Message: %s", *response.MESSAGE)
		}
		if response.STATUS != nil {
			log.Printf("Status: %s", *response.STATUS)
		}
	}
	log.Println("-----")
}

// contactSearch demonstrates how to search for contacts
func contactSearch() {
	parms := map[string]string{"emailAddress": "chris@joyce.au"}
	contacts, _, err := client.Contact.ContactSearch(parms)

	if err != nil {
		fmt.Print(err)
		return
	}

	for c := range contacts {
		log.Printf("%d\t %s\n", contacts[c].ContactID, contacts[c].EmailAddress)
	}
}

// contactEnrolments demonstrates how to get contact enrollments
func contactEnrolments(contactID int) {
	parms := map[string]string{}
	enrolments, resp, err := client.Contact.ContactEnrolments(contactID, parms)

	if err != nil {
		fmt.Println(err.Error())
		fmt.Println(resp.Body)
		return
	}

	for e := range enrolments {
		log.Printf("%d\t %s\n", enrolments[e].EnrolID, enrolments[e].Code)
	}
}

// contactCertificate demonstrates how to verify a certificate
func contactCertificate() {
	cert, _, err := client.Contact.ContactVerifyCertificate("8058765-9441274")

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%+v", cert)
}

// findME demonstrates basic contact search functionality
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

// findMEandVerifyUSI demonstrates contact search and USI verification
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
