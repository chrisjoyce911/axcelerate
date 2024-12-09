package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/chrisjoyce911/axcelerate"
)

// EmailResponse struct
type EmailResponse struct {
	FailedCount    int      `json:"FAILEDCOUNT"`
	Message        string   `json:"MESSAGE"`
	Errors         []string `json:"ERRORS,omitempty"`
	AttemptedCount int      `json:"ATTEMPTEDCOUNT"`
	SuccessCount   int      `json:"SUCCESSCOUNT"`
}

var client *axcelerate.Client

func main() {

	// // JSON response string
	// responseJSON := `{"FAILEDCOUNT":0,"MESSAGE":"1 Email(s) sent successfully. ","ERRORS":[],"ATTEMPTEDCOUNT":1,"SUCCESSCOUNT":1}`

	// // Unmarshal the JSON into the EmailResponse struct
	// var emailResponse EmailResponse
	// err := json.Unmarshal([]byte(responseJSON), &emailResponse)
	// if err != nil {
	// 	log.Fatalf("Error unmarshalling JSON: %v", err)
	// }

	// // Print the result
	// log.Printf("emailResponse.FailedCount: %+v\n", emailResponse.FailedCount)
	// log.Printf("emailResponse.Message: %+v\n", emailResponse.Message)
	// log.Printf("emailResponse.AttemptedCount: %+v\n", emailResponse.AttemptedCount)
	// log.Printf("emailResponse.SuccessCount: %+v\n", emailResponse.SuccessCount)
	// log.Printf("emailResponse.Errors: %+v\n", emailResponse.Errors)

	// log.Fatalln("done")
	var apitoken string = os.Getenv("AXCELERATE_APITOKEN")
	var wstoken string = os.Getenv("AXCELERATE_WSTOKEN")

	client, _ = axcelerate.NewClient(apitoken, wstoken, axcelerate.RateLimit(10), axcelerate.BaseURL("https://awfa.stg.axcelerate.com"))

	// savedReportList(client)
	// savedReport(client)
	// contactCertificate(client)
	// contactSearch(client)
	// contactEnrolments(14446094)
	// contactEnrolments(14365825)
	// contactCertificate(client)

	// courseEnrolments(10148651)

	// SavedReport()

	//getCoursesInstanceDetail()
	// getCoursesInstanceSearch()
	// courseEnrolmentStatus()
	// templateEmail()

	getVenueDetail()

}

func getVenueDetail() {

	contactID := 12228659

	i, reps, _ := client.Venue.Venue(contactID)

	je, _ := json.MarshalIndent(i, "", "\t")
	fmt.Printf("e: \n%s", je)

	fmt.Printf("%+v\n", reps.Body)

}

func templateEmail() {

	p := axcelerate.TemplateEmailParams{
		PlanID:                  95745,
		ContactID:               11300044,
		InstanceID:              1977505,
		InvoiceID:               3378756,
		Subject:                 "Booking Confirmation - Australia Wide First Aid",
		Type:                    "w",
		InvoiceAttachmentPlanID: 3440,
	}

	eUpdate, reps, err := client.Template.TemplateEmail(p)

	if err != nil {
		fmt.Printf("Body: %s", reps.Body)
		fmt.Print(err.Error())
		return
	}

	fmt.Printf("eUpdate%+v", eUpdate)

}

func courseEnrolmentStatus() {

	contactID := 11300044
	instanceID := 1977505

	// i, _, err := client.Courses.GetCoursesInstanceDetail(instanceID, "w")

	// //

	parms := map[string]string{}

	// currentTime := time.Now()
	// formattedDate := currentTime.Format("02/01/2006")

	// $quizKey        = "ELA:" .  $courseDataArr['instanceID'] . ":" . $pd['contactID'];
	// "https://assessment.australiawidefirstaid.com.au/?k="

	// parms["customField_PFAquiz"] = "Complete"
	// parms["customField_PFAquizlink"] = "https://assessment.australiawidefirstaid.com.au/?k=ELA:1997276:11300044"
	// parms["customField_PFAquizdate"] = formattedDate
	// parms["customField_terms"] = "Yes"

	parms["logType"] = "Booked"
	parms["theMethod"] = "Online"

	eUpdate, reps, err := client.Courses.CourseEnrolmentUpdate(contactID, instanceID, "w", parms)

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("Body%s\n", reps.Body)

	fmt.Printf("eUpdate%+v", eUpdate)

}

func getInvoices() {

	contactID := 11300044

	i, reps, _ := client.Accounting.Invoices(contactID, nil)

	je, _ := json.MarshalIndent(i, "", "\t")
	fmt.Printf("e: \n%s", je)

	fmt.Printf("%+v\n", reps.Body)

}

func getCoursesInstanceDetail() {

	instanceID := 1977505

	i, reps, _ := client.Courses.GetCoursesInstanceDetail(instanceID, "w")

	je, _ := json.MarshalIndent(i, "", "\t")
	fmt.Printf("e: \n%s", je)

	fmt.Printf("%+v\n", reps.Body)

}

func getCoursesInstanceSearch() {

	instanceID := 1977505

	args := map[string]string{
		"instanceID": fmt.Sprintf("%d", instanceID), // Convert contactID to string
		"type":       "w",                           // Convert workshopID to string

	}

	// i, reps, err := client.Courses.GetCoursesInstanceSearch(args)

	i, _, err := client.Courses.GetCoursesInstanceSearch(args)

	// je, _ := json.MarshalIndent(i, "", "\t")
	// fmt.Printf("e: \n%s", je)

	fmt.Printf("%+v\n", i)

	fmt.Printf("err %+v\n", err)

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

	cert, reps, err := client.Courses.CourseEnrolmentUpdate(contactID, int(i.LinkedClassID), "p", parms)

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

	cert, _, err := client.Report.SavedReportRun(85950, map[string]string{})

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
		log.Printf("%d\t %s\n", contacts[c].ContactID, contacts[c].EmailAddress)

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

	parms := map[string]string{}

	parms["offsetRows"] = fmt.Sprintf("%d", offsetRows)
	//	parms["filterOverride"]

	parms["filterOverride"] = url.QueryEscape(`[{"VALUE2":"","OPERATOR":"IS","DISPLAY":"Cancelled","NAME":"workshops.deleted","VALUE":"0"},{"VALUE2":"","OPERATOR":"IS","DISPLAY":"USI Verified","NAME":"contacts.usi_verified","VALUE":"0"},{"VALUE2":"","OPERATOR":"IS","DISPLAY":"Coordination Type","NAME":"workshops.ptype","VALUE":"Public Workshop"},{"VALUE2":"","OPERATOR":"IN","DISPLAY":"Key Student Status","NAME":"workshopbookings.logentrytypeid","VALUE":"1,14"}]`)

	savedReport, _, err := client.Report.SavedReportRun(85951, parms)

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
