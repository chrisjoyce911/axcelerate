package main

import (
	"encoding/json"
	"fmt"
)

// getVenueDetail demonstrates getting venue details
func getVenueDetail() {
	contactID := 12228659

	i, reps, _ := client.Venue.Venue(contactID)

	je, _ := json.MarshalIndent(i, "", "\t")
	fmt.Printf("Venue Details: \n%s", je)

	fmt.Printf("Response Body: %+v\n", reps.Body)
}
