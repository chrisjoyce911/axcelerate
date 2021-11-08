package main

import (
	"fmt"
	"os"
	"time"

	"github.com/chrisjoyce911/axcelerate"
)

func main() {

	var apitoken string = os.Getenv("AXCELERATE_APITOKEN")
	var wstoken string = os.Getenv("AXCELERATE_WSTOKEN")

	client := axcelerate.NewClient(apitoken, wstoken, nil, nil)

	// Search for somone via their email
	parms := map[string]string{"emailAddress": "xxxx@xxxx"}
	c, _, _ := client.Contact.SearchContacts(parms)

	for _, i := range c {
		fmt.Printf("%s\t%d\t%s\t%s\t%s\t%s\n", i.USI, i.ContactID, i.Emailaddress, i.Givenname, i.Surname, i.Scity)

		// For everyone we find get their enrolments
		eparms := map[string]string{"type": "s"}
		// eparms := map[string]string{}
		e, _, _ := client.Contact.ContactEnrolments(int(i.ContactID), eparms)

		for _, i := range e {

			if i.Code == "HLTAID001" || i.Code == "HLTAID009" {

				days := time.Now().Sub(i.StartDate).Hours() / 24

				fmt.Printf("%s\t%s\t%s\t%s\t%f\t%s\n", i.Type, i.Status, i.StartDate.Format("02-01-2006"), i.Code, days, i.Location)
			}
		}
	}

	parms = map[string]string{"instanceID": "1702030", "type": "w"}
	obj, _, _ := client.Courses.GetEnrolments(parms)

	for _, i := range obj {
		fmt.Printf("%d\t%f\t%s\t%v\n", i.ContactID, i.AmountPaid, i.CompeteName, i.PriceBeat)
	}

}
