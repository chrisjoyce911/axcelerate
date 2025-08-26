package files

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/chrisjoyce911/axcelerate"
)

// transact demonstrates creating a transaction
func Transact(client *axcelerate.Client) {
	params := map[string]string{
		"amount":      "59",
		"ContactID":   "14518907",
		"invoiceID":   "3643466",
		"description": "Stripe Payment pi_3RkeLoHiVYttPAwh1dmGSw6f",
	}

	i, reps, err := client.Accounting.CreateTransaction(params)

	log.Println("-----")
	log.Printf("Transaction\n%+v", i)
	log.Println("-----")
	log.Printf("Body\n%s", reps.Body)
	log.Println("-----")
	if err != nil {
		log.Printf("%+v", err.Error())
	} else {
		log.Printf("No error!")
	}
	log.Println("-----")
}

// InvoiceVoid demonstrates voiding an invoice
func InvoiceVoid(client *axcelerate.Client) {
	guid := "DEF95391-7FDF-4A92-8D3F7717123F0881"

	i, reps, err := client.Accounting.InvoiceVoid(guid)

	fmt.Printf("%t %+v %s", i, reps.Body, err.Error())
}

// PaymentVerify demonstrates payment verification
func PaymentVerify(client *axcelerate.Client) {
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

// GetInvoices demonstrates getting invoices for a contact
func GetInvoices(client *axcelerate.Client) {
	contactID := 11300044

	i, reps, _ := client.Accounting.Invoices(contactID, nil)

	je, _ := json.MarshalIndent(i, "", "\t")
	fmt.Printf("Invoices: \n%s", je)

	fmt.Printf("Response Body: %+v\n", reps.Body)
}
