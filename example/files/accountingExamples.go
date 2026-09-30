package files

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/chrisjoyce911/axcelerate"
)

// Transact demonstrates creating a transaction with the raw parameter map.
// The amount is in DOLLARS. Against an invoice, prefer ApplyPayment: a
// payment against a DRAFT invoice answers 400 yet still records the money
// as an unassigned receipt, so a retry here would pay twice.
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

// Domains lists every domain id and name (GET /domains — undocumented by
// aXcelerate). aXcelerate stores any DOMAINID on an invoice line without
// checking it exists, so this is how an id is validated.
func Domains(client *axcelerate.Client) {
	domains, _, err := client.Accounting.Domains()
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	for _, d := range domains {
		fmt.Printf("%6d  %s\n", d.DomainID, d.DomainName)
	}
}

// InvoiceFlow creates an invoice, adds a line, issues it and applies the
// payment — in the order aXcelerate requires. Run it against STAGING: it
// writes a real invoice and a real payment.
func InvoiceFlow(client *axcelerate.Client) {
	a := client.Accounting
	contactID := 14518907

	// 1. A DRAFT invoice (INVOICENR "AUTO"). LastName is sent as `lastname`
	//    — the documented `surname` is ignored. Dates default to today.
	inv, _, err := a.CreateInvoice(axcelerate.InvoiceRequest{
		ContactID: contactID, FirstName: "Chris", LastName: "Joyce",
		ExternalReference: "SDK-EXAMPLE",
		Lines: []axcelerate.InvoiceLine{{
			Description: "Pocket Mask Training", Qty: 1,
			ItemCode:       "", // always sent, even empty
			UnitPriceGross: "19.00", TaxPercent: "0.00",
			FinanceCode: "43420", DomainID: axcelerate.Int(5566), // ONLINE
		}},
	})
	if err != nil {
		log.Printf("create: %v", err)
		return
	}
	log.Printf("created invoice %d (number %s)", inv.InvoiceID, inv.InvoiceNumber)

	// 2. Add a line. The PUT replaces every line, so AddInvoiceLines
	//    re-sends the existing ones for you.
	if _, err := a.AddInvoiceLines(inv.InvoiceID, axcelerate.InvoiceLine{
		Description: "Afterpay transaction fee", Qty: 1,
		UnitPriceGross: "2.00", TaxPercent: "0.00",
		FinanceCode: "14180", DomainID: axcelerate.Int(5494), // Head Office
	}); err != nil {
		log.Printf("add line: %v", err)
		return
	}

	// 3. Issue it — a draft cannot take a payment. Approve takes the GUID.
	if _, err := a.ApproveInvoice(inv.InvGUID); err != nil {
		log.Printf("approve: %v", err)
		return
	}

	// 4. Apply the payment, in dollars. ApplyPayment refuses a draft
	//    before sending anything.
	txn, _, err := a.ApplyPayment(axcelerate.InvoicePayment{
		ContactID: contactID, InvoiceID: inv.InvoiceID,
		Amount: "21.00", Reference: "SDK-EXAMPLE",
	})
	if err != nil {
		// The payment may have been recorded even so: check Transactions
		// before retrying.
		log.Printf("apply payment: %v", err)
		return
	}
	log.Printf("paid: transaction %d", txn.TransactionID)
}

// Transactions lists a contact's transactions; UnassignedAmount is money
// no invoice has claimed.
func Transactions(client *axcelerate.Client) {
	txns, _, err := client.Accounting.Transactions(14518907)
	if err != nil {
		log.Printf("Error: %v", err)
		return
	}
	for _, t := range txns {
		fmt.Printf("%d  %s  $%.2f  unassigned $%.2f  %s\n",
			t.TransactionID, t.TransDate, float64(t.Amount), float64(t.UnassignedAmount), t.Reference)
	}
}
