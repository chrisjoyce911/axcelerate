package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	client, _ = axcelerate.NewClient(apitoken, wstoken, axcelerate.RateLimit(10), axcelerate.BaseURL("https://awfa.app.axcelerate.com/api"))

	// savedReportList(client)
	// savedReport(client)
	// contactCertificate(client)
	// contactSearch(client)
	// contactEnrolments(14446094)
	// contactEnrolments(14365825)
	// contactCertificate(client)

	// courseEnrolments(10148651)
	// savedReport()
	// getCoursesInstanceDetail()
	// getCoursesInstanceSearch()
	// courseEnrolmentStatus()
	// templateEmail()

	// getVenueDetail()

	updateFinCode(client)

}

func updateFinCode(client *axcelerate.Client) {
	ids := []string{
		"2003904", "2006097", "2010906", "2010906", "2006095", "2006097", "2006095", "2003904", "2010908", "2010906",
		"2006095", "2006097", "2003904", "2010906", "2003907", "2006097", "2010907", "2010907", "2006097", "2010908",
		"2006095", "2006095", "2003907", "2010906", "2003907", "2010907", "2006095", "2006101", "2010906", "2006103",
		"2006103", "2006097", "2003907", "2006103", "2006097", "2006097", "2010907", "2010907", "2010907", "2006101",
		"2003904", "2006101", "2010907", "2006097", "2003904", "2006097", "2010906", "2003907", "2010907", "2006097",
		"2006099", "2010907", "2006099", "2010908", "2010907", "2003904", "2006095", "2006095", "2010906", "2010906",
		"2003907", "2006101", "2010907", "2006103", "2003907", "2010907", "2010908", "2010907", "2006099", "2006095",
		"2006101", "2006099", "2006103", "2006103", "2010908", "2006099", "2006099", "2006097", "2006103", "2010908",
		"2010907", "2006099", "2010907", "2010907", "2010908",
	}

	for _, id := range ids {
		params := map[string]string{
			"finCodeID": "10076",
			"type":      "w",
			"ID":        id,
		}

		_, resp, err := client.Courses.UpdateInstanceDetails(params)
		if err != nil {
			log.Printf("Error updating finCodeID for ID %s: %v", id, err)
			continue
		}

		log.Printf("Updated ID %s with Response Status Code: %v", id, resp.StatusCode)
	}
}

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

func paymentVerify() {

	payment, res, err := client.Accounting.PaymentVerify("82A45263-0C31-49F3-B3C7196331B5AFCAcc")

	// Log payment details on success
	if payment != nil && payment.ErrorResponse != nil {
		log.Printf("Payment Details: %+v", payment.ErrorResponse)
	} else {
		log.Printf("Payment Details: <nil>")
	}

	if err != nil {
		log.Printf("Error: %v", err)
		if payment != nil && payment.ErrorResponse != nil {
			log.Printf("Error Details: %+v", payment.ErrorResponse)
		}
		// log.Printf("Response: %+v", res)
		return
	}

	// Log payment details on success
	if payment != nil {
		log.Printf("Payment Details: %+v", payment)
	} else {
		log.Printf("Payment Details: <nil>")
	}
	log.Printf("Response: %+v", res)
}

func savedReport() {
	offsetRows := 0

	parms := map[string]string{}

	parms["offsetRows"] = fmt.Sprintf("%d", offsetRows)

	parms["filterOverride"] = ` [
 {
     "VALUE2": "0",
     "OPERATOR": "BETWEEN N Days",
     "DISPLAY": "Workshop Start Date",
     "NAME": "workshops.pstartdate",
     "VALUE": "0"
 }]`

	savedReport, _, err := client.Report.SavedReportRun(85957, parms)

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
	instanceID := 1997276

	i, _, err := client.Courses.GetCoursesInstanceDetail(instanceID, "w")

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
