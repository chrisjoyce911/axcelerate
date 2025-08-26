package files

import (
	"encoding/json"
	"fmt"

	"github.com/chrisjoyce911/axcelerate"
)

// GetVenueDetail demonstrates getting venue details
func GetVenueDetail(client *axcelerate.Client) {
	contactID := 12228659

	i, reps, _ := client.Venue.Venue(contactID)

	je, _ := json.MarshalIndent(i, "", "\t")
	fmt.Printf("Venue Details: \n%s", je)

	fmt.Printf("Response Body: %+v\n", reps.Body)
}
