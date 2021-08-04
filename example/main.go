package main

import (
	"fmt"
	"os"

	"github.com/chrisjoyce911/axcelerate"
)

func main() {

	var apitoken string = os.Getenv("AXCELERATE_APITOKEN")
	var wstoken string = os.Getenv("AXCELERATE_WSTOKEN")

	client := axcelerate.NewClient(apitoken, wstoken, nil, nil)

	// Search for somone via their email
	parms := map[string]string{"emailAddress": "xxx@xxx"}
	c, _, _ := client.Contact.SearchContacts(parms)

	for _, i := range c {
		fmt.Printf("%s\t%d\t%s\t%s\t%s\t%s\n", i.USI, i.ContactID, i.Emailaddress, i.Givenname, i.Surname, i.Source)

		// For everyone we find get their enrolments
		eparms := map[string]string{"type": "w"}
		e, _, _ := client.Contact.ContactEnrolments(int(i.ContactID), eparms)

		for _, i := range e {
			fmt.Printf("%s\t%s\t%s\t%s\n", i.Type, i.Name, i.StartDate, i.Location)
		}
	}

}
