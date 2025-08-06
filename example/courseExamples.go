package main

import (
	"fmt"
	"log"
	"time"
)

// courseEnrolmentStatus demonstrates how to update course enrollment status
func courseEnrolmentStatus() {
	contactID := 11300044
	instanceID := 1977505

	parms := map[string]string{}

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

// courseEnrolments demonstrates how to get course enrollments
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

// courseEnrolment demonstrates course enrollment update with custom fields
func courseEnrolment() {
	contactID := 11300044
	instanceID := 1997276

	i, _, err := client.Courses.GetCoursesInstanceDetail(instanceID, "w")

	parms := map[string]string{}

	currentTime := time.Now()
	formattedDate := currentTime.Format("02/01/2006")

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

// getCoursesInstanceDetail demonstrates how to get course instance details
func getCoursesInstanceDetail() {
	instanceID := 2014519

	i, reps, _ := client.Courses.GetCoursesInstanceDetail(instanceID, "w")

	fmt.Printf("Course Instance Detail: %+v\n", i)
	fmt.Printf("Response Body: %+v\n", reps.Body)
}

// getCoursesInstanceSearch demonstrates how to search course instances
func getCoursesInstanceSearch() {
	instanceID := 2014519

	args := map[string]string{
		"instanceID": fmt.Sprintf("%d", instanceID),
		"type":       "w",
	}

	i, _, err := client.Courses.GetCoursesInstanceSearch(args)

	fmt.Printf("Search Results: %+v\n", i)
	fmt.Printf("Error: %+v\n", err)
}

// updateInstanceMaxParticipants demonstrates updating max participants for workshops
func updateInstanceMaxParticipants() {
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

// updateFinCode demonstrates bulk updating of financial codes
func updateFinCode() {
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
