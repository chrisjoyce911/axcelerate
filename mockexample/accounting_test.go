package main

import (
	"testing"

	"github.com/chrisjoyce911/axcelerate"
)

func TestGetInvoice(t *testing.T) {
	// Initialize the mock client
	mockClient := axcelerate.NewMockClient()

	// Configure the mock GetInvoice method
	mockClient.Accounting.MockGetInvoice = func(invoiceID int) (*axcelerate.Invoice, *axcelerate.Response, error) {
		return &axcelerate.Invoice{InvoiceID: invoiceID, PriceGross: 100.50}, nil, nil
	}

	// Use the mock Accounting service
	invoiceID := 12345
	invoice, _, err := mockClient.Accounting.GetInvoice(invoiceID)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if invoice.InvoiceID != invoiceID {
		t.Errorf("expected invoice ID %d, got %d", invoiceID, invoice.InvoiceID)
	}

	if invoice.PriceGross != 100.50 {
		t.Errorf("expected invoice total 100.50, got %.2f", invoice.PriceGross)
	}
}

func TestPaymentURL(t *testing.T) {
	mockService := &axcelerate.MockAccountingService{
		MockPaymentURL: func(reference, invoiceGUID, redirectURL, cancelURL string) (*axcelerate.PaymentRequest, *axcelerate.Response, error) {
			return &axcelerate.PaymentRequest{
				Meta:       map[string]interface{}{"exampleKey": "exampleValue"},
				FormMethod: "POST",
				HTML:       "<form>Mock Form</form>",
				Action:     "https://mock-payment-action.com",
			}, &axcelerate.Response{Status: "200 OK"}, nil
		},
	}

	reference := "testReference"
	invoiceGUID := "testGUID"
	redirectURL := "https://redirect.com"
	cancelURL := "https://cancel.com"

	paymentRequest, resp, err := mockService.PaymentURL(reference, invoiceGUID, redirectURL, cancelURL)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Status != "200 OK" {
		t.Errorf("expected status '200 OK', got '%s'", resp.Status)
	}

	if paymentRequest.Action != "https://mock-payment-action.com" {
		t.Errorf("expected action URL 'https://mock-payment-action.com', got '%s'", paymentRequest.Action)
	}
}

func TestPaymentVerify(t *testing.T) {
	mockService := &axcelerate.MockAccountingService{
		MockPaymentVerify: func(reference string) (*axcelerate.FullPaymentResponse, *axcelerate.Response, error) {
			return &axcelerate.FullPaymentResponse{
				CurrentState: "Completed",
				ResultDetails: axcelerate.PaymentResultDetails{
					IsSuccessful: true,
				},
			}, &axcelerate.Response{Status: "200 OK"}, nil
		},
	}

	reference := "testReference"

	paymentResponse, resp, err := mockService.PaymentVerify(reference)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Status != "200 OK" {
		t.Errorf("expected status '200 OK', got '%s'", resp.Status)
	}

	if paymentResponse.CurrentState != "Completed" {
		t.Errorf("expected state 'Completed', got '%s'", paymentResponse.CurrentState)
	}

	if !paymentResponse.ResultDetails.IsSuccessful {
		t.Errorf("expected IsSuccessful to be true, got false")
	}
}
