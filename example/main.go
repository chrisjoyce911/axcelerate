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

	parms := map[string]string{"emailAddress": "xxx@xxx"}
	e, _, _ := client.Contact.SearchContacts(parms)

	fmt.Print(e)
	fmt.Println(len(e))

}
