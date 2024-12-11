package axcelerate

type MockAccountingService struct {
	MockGetInvoice           func(invoiceID int) (*Invoice, *Response, error)
	MockInvoices             func(contactID int, extra *map[string]string) ([]InvoiceSummary, *Response, error)
	MockPaymentAxcelerateURL func(invoiceID int) (*PaymentURL, *Response, error)
	MockPaymentURL           func(reference, invoiceGUID, redirectURL, cancelURL string) (*PaymentRequest, *Response, error)
	MockPaymentForm          func(reference, invoiceGUID, redirectURL, cancelURL string) (*PaymentResponse, *Response, error)
	MockPaymentVerify        func(reference string) (*FullPaymentResponse, *Response, error)
}

func defaultMockResponse() *Response {
	return &Response{Status: "200 OK"}
}

func defaultInvoice() *Invoice {
	return &Invoice{
		InvoiceID:   1,
		InvoiceDate: "2024-01-01",
		PriceGross:  100.00,
		IsPaid:      true,
	}
}

func (m *MockAccountingService) GetInvoice(invoiceID int) (*Invoice, *Response, error) {
	if m.MockGetInvoice != nil {
		return m.MockGetInvoice(invoiceID)
	}
	return defaultInvoice(), defaultMockResponse(), nil
}

func (m *MockAccountingService) Invoices(contactID int, extra *map[string]string) ([]InvoiceSummary, *Response, error) {
	if m.MockInvoices != nil {
		return m.MockInvoices(contactID, extra)
	}
	return nil, nil, nil
}

func (m *MockAccountingService) PaymentAxcelerateURL(invoiceID int) (*PaymentURL, *Response, error) {
	if m.MockPaymentAxcelerateURL != nil {
		return m.MockPaymentAxcelerateURL(invoiceID)
	}
	return nil, nil, nil
}

func (m *MockAccountingService) PaymentURL(reference, invoiceGUID, redirectURL, cancelURL string) (*PaymentRequest, *Response, error) {
	if m.MockPaymentURL != nil {
		return m.MockPaymentURL(reference, invoiceGUID, redirectURL, cancelURL)
	}
	return nil, nil, nil
}

func (m *MockAccountingService) PaymentForm(reference, invoiceGUID, redirectURL, cancelURL string) (*PaymentResponse, *Response, error) {
	if m.MockPaymentForm != nil {
		return m.MockPaymentForm(reference, invoiceGUID, redirectURL, cancelURL)
	}
	return nil, nil, nil
}

func (m *MockAccountingService) PaymentVerify(reference string) (*FullPaymentResponse, *Response, error) {
	if m.MockPaymentVerify != nil {
		return m.MockPaymentVerify(reference)
	}
	return nil, nil, nil
}
