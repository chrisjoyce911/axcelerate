package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

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
	// contactEnrolments(14446094)
	// contactEnrolments(14365825)
	// contactCertificate(client)

	// courseEnrolments(10148651)

	SavedReport()

}

func courseEnrolment() {

	contactID := 11300044
	instranceID := 1997276

	i, _, err := client.Courses.GetCoursesInstanceDetail(instranceID, "w")

	//

	parms := map[string]string{}

	currentTime := time.Now()
	formattedDate := currentTime.Format("02/01/2006")

	// $quizKey        = "ELA:" .  $courseDataArr['instanceID'] . ":" . $pd['contactID'];
	// "https://assessment.australiawidefirstaid.com.au/?k="

	parms["customField_PFAquiz"] = "Complete"
	parms["customField_PFAquizlink"] = "https://assessment.australiawidefirstaid.com.au/?k=ELA:1997276:11300044"
	parms["customField_PFAquizdate"] = formattedDate
	parms["customField_terms"] = "Yes"

	cert, reps, err := client.Courses.UpdateCourseEnrolment(contactID, int(i.LinkedClassID), "p", parms)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%+v\n", reps)

	fmt.Printf("%+v", cert)

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

func courseEnrolments(contactID int) {

	parms := map[string]string{}

	parms["contactID"] = fmt.Sprintf("%d", contactID)

	enrolments, _, err := client.Courses.GetEnrolments(parms)

	if err != nil {
		fmt.Print(err)
		return
	}

	for e := range enrolments {
		log.Printf("%d\t%s\t%d\n", enrolments[e].EnrolmentID, enrolments[e].Code, enrolments[e].ContactID)

	}

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

func SavedReport() {
	offsetRows := 0
	displayLength := 10

	parms := map[string]string{}

	parms["offsetRows"] = fmt.Sprintf("%d", offsetRows)
	//	parms["filterOverride"]

	parms["filterOverride"] = url.QueryEscape(`[{"VALUE2":"","OPERATOR":"IS","DISPLAY":"Cancelled","NAME":"workshops.deleted","VALUE":"0"},{"VALUE2":"","OPERATOR":"IS","DISPLAY":"USI Verified","NAME":"contacts.usi_verified","VALUE":"0"},{"VALUE2":"","OPERATOR":"IS","DISPLAY":"Coordination Type","NAME":"workshops.ptype","VALUE":"Public Workshop"},{"VALUE2":"","OPERATOR":"IN","DISPLAY":"Key Student Status","NAME":"workshopbookings.logentrytypeid","VALUE":"1,14"}]`)

	savedReport, _, err := client.Report.SavedReportRun(85951, displayLength, parms)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(savedReport.Data)
	// for w := range workshops {
	// 	c, _, err := client.Courses.UpdateInstanceMaxParticipants(workshops[w], "w", max)

	// 	log.Printf("%d\t %s\n", c.InstanceID, c.Message)
	// 	if err != nil {
	// 		fmt.Print(err)
	// 	}

	// }

}
