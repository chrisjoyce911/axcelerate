package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/chrisjoyce911/axcelerate"
)

var client *axcelerate.Client

func main() {

	var apitoken string = os.Getenv("AXCELERATE_APITOKEN")
	var wstoken string = os.Getenv("AXCELERATE_WSTOKEN")

	client = axcelerate.NewClient(apitoken, wstoken, nil, nil)

	// savedReportList(client)
	// savedReport(client)
	// contactCertificate(client)
	// contactSearch(client)
	contactEnrolments(14446094)
	// contactCertificate(client)

}

func savedReportList() {

	cert, _, err := client.Report.SavedReportList()

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%+v", cert)

}

func savedReport() {

	cert, _, err := client.Report.SavedReportRun(85950, 10, map[string]string{})

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%+v", cert)
}

func contactCertificate() {

	cert, _, err := client.Contact.ContactVerifyCertificate("8058765-9441274")

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%+v", cert)

}

func ContactCertificate() {

	cert, _, err := client.Contact.ContactEnrolmentsCertificate(12787538)

	if err != nil {
		fmt.Print(err)
		return
	}

	saveMediaToDisk(cert, "./example/")

}

func saveMediaToDisk(media axcelerate.Media, directory string) error {
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		return err
	}

	filePath := filepath.Join(directory, media.FileName)
	err = os.WriteFile(filePath, media.Data, 0644)
	return err
}

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

func contactSearch() {

	parms := map[string]string{"emailAddress": "chris@joyce.au"}
	contacts, _, err := client.Contact.ContactSearch(parms)

	if err != nil {
		fmt.Print(err)
		return
	}

	for c := range contacts {
		log.Printf("%d\t %s\n", contacts[c].ContactID, contacts[c].Emailaddress)

	}

}

func UpdateInstanceMaxParticipants() {

	max := 10
	workshops := []int{
		1904663,
		1904664,
		1913826,
	}

	for w := range workshops {
		c, _, err := client.Courses.UpdateInstanceMaxParticipants(workshops[w], "w", max)

		log.Printf("%d\t %s\n", c.InstanceID, c.Message)
		if err != nil {
			fmt.Print(err)
		}

	}

}
