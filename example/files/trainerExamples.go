package files

import (
	"log"

	"github.com/chrisjoyce911/axcelerate"
)

// ActiveTrainer gets an active trainer
func ActiveTrainer(client *axcelerate.Client) {

	trainer, resp, err := client.Trainer.GetTrainer(13509784)

	log.Printf("Response Body: %v\n", resp.Body)
	log.Printf("Response Status Code: %v\n", resp.StatusCode)

	if err != nil {
		log.Printf("API error: %v", err)
		return
	}

	if trainer != nil {
		log.Printf("Trainer found: %+v", trainer)
	} else {
		log.Printf("Trainer not found.")
	}
}
