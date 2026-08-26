package main

import (
	"fmt"
	"log"
	"os"

	"github.com/chrisjoyce911/axcelerate"
	"github.com/joho/godotenv"
	// Import your example packages
)

var client *axcelerate.Client

func main() {
	_ = godotenv.Load()

	var apitoken string = os.Getenv("AXCELERATE_APITOKEN")
	var wstoken string = os.Getenv("AXCELERATE_WSTOKEN")
	var baseURL string = os.Getenv("AXCELERATE_BASEURL")

	fmt.Println("API Token:", apitoken)
	fmt.Println("WS Token:", wstoken)
	fmt.Println("Base URL:", baseURL)

	client, _ = axcelerate.NewClient(apitoken, wstoken, axcelerate.RateLimit(10), axcelerate.BaseURL(baseURL))

	ContactEnrolments(10148651, client)

}

// ContactEnrolments demonstrates how to get contact enrollments
func ContactEnrolments(contactID int, client *axcelerate.Client) {
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
