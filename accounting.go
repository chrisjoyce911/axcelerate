package axcelerate

import (
	"encoding/json"
	"fmt"
)

// CoursesService handles all interactions with Contact
type AccountingService struct {
	client *Client
}

type Items struct {
	TotalTax float32 `json:"TOTALTAX"`
	Children struct {
	} `json:"CHILDREN"`
	UnitPriceTax   float32     `json:"UNITPRICETAX"`
	DomainID       int         `json:"DOMAINID"`
	HasChildren    int         `json:"HASCHILDREN"`
	PartID         int         `json:"PARTID"`
	UnitPriceNet   float32     `json:"UNITPRICENETT"`
	Qty            int         `json:"QTY"`
	TaxPercent     int         `json:"TAXPERCENT"`
	TotalGross     float32     `json:"TOTALGROSS"`
	ItemCode       string      `json:"ITEMCODE"`
	FinanceCode    string      `json:"FINANCECODE"`
	TotalNet       float32     `json:"TOTALNETT"`
	SourceID       int         `json:"SOURCEID"`
	ServiceDate    string      `json:"SERVICEDATE" time_format:"axc_date"`
	UnitPriceGross float32     `json:"UNITPRICEGROSS"`
	ItemID         int         `json:"ITEMID"`
	Description    string      `json:"DESCRIPTION"`
	CostCentreCode interface{} `json:"COSTCENTRECODE"`
}

// Invoice object with the full Invoice information
type Invoice struct {
	InvoiceNumber      string           `json:"INVOICENR"`
	PriceGross         float32          `json:"PRICEGROSS"`
	Address2           string           `json:"ADDRESS2"`
	OwnerContactID     int              `json:"OWNERCONTACTID"`
	Organisation       string           `json:"ORGANISATION"`
	ShipLastName       string           `json:"SHIPLASTNAME"`
	ShipHousenr        string           `json:"SHIPHOUSENR"`
	ShipOrgID          int              `json:"SHIPORGID"`
	ShipPostCode       string           `json:"SHIPPOSTCODE"`
	Comment            string           `json:"COMMENT"`
	PhoneNumber        string           `json:"PHONENR"`
	OrgID              int              `json:"ORGID"`
	Payments           []PaymentDetails `json:"PAYMENTS"`
	ShipCountry        string           `json:"SHIPCOUNTRY"`
	DueDate            string           `json:"DUEDATE" time_format:"axc_date"`
	ShipOrganisation   string           `json:"SHIPORGANISATION"`
	InvoiceID          int              `json:"INVOICEID"`
	DueDateOffset      int              `json:"DUEDATEOFFSET"`
	AreItemsLocked     bool             `json:"AREITEMSLOCKED"`
	LastName           string           `json:"LASTNAME"`
	Street             string           `json:"STREET"`
	Items              []Items          `json:"ITEMS"`
	InvoiceDate        string           `json:"INVOICEDATE" time_format:"axc_date"`
	Currency           string           `json:"CURRENCY"`
	ShipPriceNett      float32          `json:"SHIPPRICENETT"`
	ShipState          string           `json:"SHIPSTATE"`
	ContactName        string           `json:"CONTACTNAME"`
	ShipCountryISO3166 string           `json:"SHIPCOUNTRYISO3166"`
	Shipstreet         string           `json:"SHIPSTREET"`
	Housenr            string           `json:"HOUSENR"`
	Isarchived         bool             `json:"ISARCHIVED"`
	OrderDate          string           `json:"ORDERDATE" time_format:"axc_date"`
	ContactID          int              `json:"CONTACTID"`
	ShipCity           string           `json:"SHIPCITY"`
	Shippricegross     float32          `json:"SHIPPRICEGROSS"`
	CountryISO3166     string           `json:"COUNTRYISO3166"`
	Isinvoicenrlocked  bool             `json:"ISINVOICENRLOCKED"`
	Shiptaxpercent     int              `json:"SHIPTAXPERCENT"`
	Shippricetax       int              `json:"SHIPPRICETAX"`
	State              string           `json:"STATE"`
	FirstName          string           `json:"FIRSTNAME"`
	Shipaddress2       string           `json:"SHIPADDRESS2"`
	ShipFirstName      string           `json:"SHIPFIRSTNAME"`
	City               string           `json:"CITY"`
	OrderNumber        string           `json:"ORDERNR"`
	Email              string           `json:"EMAIL"`
	PriceNett          float32          `json:"PRICENETT"`
	Country            string           `json:"COUNTRY"`
	InvGUID            string           `json:"INVGUID"`
	Balance            float32          `json:"BALANCE"`
	Postcode           string           `json:"POSTCODE"`
	IsPaid             bool             `json:"ISPAID"`
}

type PaymentRequestDetails struct {
	CancelURL   string `json:"CANCELURL"`   // URL to cancel the payment process
	WebhookURL  string `json:"WEBHOOKURL"`  // URL for webhook notifications
	RedirectURL string `json:"REDIRECTURL"` // URL to redirect after payment
	InvoiceGUID string `json:"INVOICEGUID"` // GUID of the invoice
}

type PaymentErrorDetails struct {
	ErrorCode    int    `json:"ERROR_CODE"` // Error code
	ErrorMessage string `json:"ERROR_MSG"`  // Error message
	Code         int    `json:"CODE"`       // Error code
	Msg          string `json:"MSG"`        // Error message
}

type PaymentResultDetails struct {
	Error           PaymentErrorDetails `json:"ERROR"`           // Details of any error encountered
	TransactionGUID string              `json:"TRANSACTIONGUID"` // GUID of the transaction
	TransactionID   string              `json:"TRANSACTIONID"`   // ID of the transaction
	IsSuccessful    bool                `json:"OK"`              // Indicates if the transaction was successful
	PlatformName    string              `json:"PLATFORM"`        // Platform used for the transaction
}

type FullPaymentResponse struct {
	RequestDetails  PaymentRequestDetails `json:"REQUEST"`           // Details of the payment request
	CurrentState    string                `json:"STATE"`             // Current state of the payment
	PaymentPlatform string                `json:"PLATFORM"`          // Platform used for payment
	ResultDetails   PaymentResultDetails  `json:"RESULT"`            // Result details of the payment process
	PlatformRefGUID string                `json:"PLATFORMREFERENCE"` // Reference for the platform transaction
	ErrorResponse   *PaymentErrorResponse `json:"-"`                 // Populated if an error response is detected
}

// PaymentErrorResponse represents an alternative error response structure
type PaymentErrorResponse struct {
	Data     string `json:"DATA"`     // Encoded metadata and variables
	Error    bool   `json:"ERROR"`    // Indicates if the response is an error
	Messages string `json:"MESSAGES"` // Error messages
	Code     string `json:"CODE"`     // Error code
	Details  string `json:"DETAILS"`  // Error details
}

type PaymentRequest struct {
	Meta       map[string]interface{} `json:"META"`       // Metadata (deprecated; avoid use in new implementations)
	FormMethod string                 `json:"FORMMETHOD"` // Form method, e.g., POST
	Script     string                 `json:"SCRIPT"`     // Script field (deprecated; avoid use in new implementations)
	HTML       string                 `json:"HTML"`       // The HTML that should be rendered inside of an HTML form tag
	Action     string                 `json:"ACTION"`     // The action attribute for the form into which the returned HTML is inserted
}

type PaymentResponse struct {
	Data    PaymentRequest `json:"DATA"`    // Contains the payment request details
	Success bool           `json:"SUCCESS"` // Returns true if a checkout form could be generated
}

type PaymentDetails struct {
	TransactionDate       string  `json:"TRANSACTIONDATE" time_format:"axc_date_hours"` // Date and time of the transaction
	TransactionProviderID int     `json:"TRANSACTIONPROVIDERID"`                        // Provider ID for the transaction
	GUID                  string  `json:"GUID"`                                         // Unique transaction GUID
	ProviderName          string  `json:"TRANSACTIONPROVIDER"`                          // Name of the transaction provider
	FragmentAmount        float32 `json:"FRAGMENT_AMOUNT"`                              // Partial amount of the transaction
}

type PaymentURL struct {
	URL string `json:"PAYMENTURL"` // URL to proceed with the payment
}

type InvoiceSummary struct {
	InvoiceNr         string  `json:"INVOICENR"`         // Invoice number
	PriceGross        string  `json:"PRICEGROSS"`        // Gross price of the invoice
	DueDate           string  `json:"DUEDATE"`           // Due date of the invoice
	InvoiceID         string  `json:"INVOICEID"`         // Unique identifier for the invoice
	AreItemsLocked    bool    `json:"AREITEMSLOCKED"`    // Indicates if items in the invoice are locked
	LastName          string  `json:"LASTNAME"`          // Last name associated with the invoice
	IsCancelled       bool    `json:"ISCANCELLED"`       // Indicates if the invoice is cancelled
	ExternalReference *string `json:"EXTERNALREFERENCE"` // External reference, nullable
	Balance           string  `json:"BALANCE"`           // Remaining balance on the invoice
	FirstName         string  `json:"FIRSTNAME"`         // First name associated with the invoice
	InvoiceDate       string  `json:"INVOICEDATE"`       // Invoice creation date
	IsVoid            bool    `json:"ISVOID"`            // Indicates if the invoice is void
	IsPaid            bool    `json:"ISPAID"`            // Indicates if the invoice is paid
}

type InvoiceCollection []InvoiceSummary

// GetInvoice will get the details of an invoice / Update an invoice (NOTE: Currently you can only Lock Items and Finalise)
// Header			Type		Required	Default	Description
// invoiceID		numeric		true				The invoiceID to retrieve. Note that this is NOT the same as the invoice number

func (s *AccountingService) GetInvoice(invoiceID int) (*Invoice, *Response, error) {
	var obj Invoice

	parms := map[string]string{"invoiceID": fmt.Sprintf("/%d", invoiceID)}

	url := fmt.Sprintf("/accounting/invoice/%d", invoiceID)
	resp, err := do(s.client, "GET", Params{parms: parms, u: url}, obj)

	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return &obj, resp, err
}

// Invoices will get an array of invoices for an account
// Header			Type		Required	Default	Description
// contactID		numeric		true				The contactID to return invoices for that contact

func (s *AccountingService) Invoices(contactID int, extra *map[string]string) ([]InvoiceSummary, *Response, error) {
	var obj []InvoiceSummary

	// Initialize parms as an empty map
	parms := map[string]string{}

	// If extra is not nil, merge its contents into parms
	if extra != nil {
		for key, value := range *extra {
			parms[key] = value
		}
	}

	// Add contactID to parms
	parms["contactID"] = fmt.Sprintf("%d", contactID)

	resp, err := do(s.client, "GET", Params{parms: parms, u: "/accounting/invoice"}, obj)

	if err != nil {
		return obj, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}

// PaymentAxcelerateURL Get a payable URL for an invoice.
// Header			Type		Required	Default	Description
// invoiceID		numeric		true				The invoiceID to pay

func (s *AccountingService) PaymentAxcelerateURL(invoiceID int) (*PaymentURL, *Response, error) {
	var obj PaymentURL

	parms := map[string]string{}
	url := fmt.Sprintf("/accounting/invoice/%d/paymenturl", invoiceID)

	resp, err := do(s.client, "GET", Params{parms: parms, u: url}, obj)
	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return &obj, resp, err
}

// PaymentURL will get an array of invoices for an account
// Header			Type		Required	Default	Description
// reference		string		true		The external identifier for the payment flow process.
// invoiceGUID		string		true		The GUID of the aXcelerate invoice for which payment should be collected.
// redirectURL		string		true		The URL to which the client will be redirected after payment processing.
// cancelURL		string		true		The URL to which the client is redirected if the client decides to cancel payment processing.

func (s *AccountingService) PaymentURL(reference, invoiceGUID, redirectURL, cancelURL string) (*PaymentRequest, *Response, error) {
	var obj PaymentRequest

	parms := map[string]string{}

	parms["reference"] = reference
	parms["invoiceGUID"] = invoiceGUID
	parms["redirectURL"] = redirectURL
	parms["cancelURL"] = cancelURL

	url := "/accounting/ecommerce/payment/url"
	resp, err := do(s.client, "GET", Params{parms: parms, u: url}, obj)

	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return &obj, resp, err
}

// PaymentForm will get an array of invoices for an account
// Header			Type		Required	Default	Description
// reference		string		true		The external identifier for the payment flow process.
// invoiceGUID		string		true		The GUID of the aXcelerate invoice for which payment should be collected.
// redirectURL		string		true		The URL to which the client will be redirected after payment processing.
// cancelURL		string		true		The URL to which the client is redirected if the client decides to cancel payment processing.

func (s *AccountingService) PaymentForm(reference, invoiceGUID, redirectURL, cancelURL string) (*PaymentResponse, *Response, error) {
	var obj PaymentResponse

	parms := map[string]string{}

	parms["reference"] = reference
	parms["invoiceGUID"] = invoiceGUID
	parms["redirectURL"] = redirectURL
	parms["cancelURL"] = cancelURL

	url := "/accounting/ecommerce/payment/url"
	resp, err := do(s.client, "GET", Params{parms: parms, u: url}, obj)

	if err != nil {
		return nil, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return &obj, resp, err
}

// UnmarshalJSON custom unmarshals PaymentResultDetails to handle string or object cases
func (r *PaymentResultDetails) UnmarshalJSON(data []byte) error {
	if string(data) == `""` { // Handle the case where RESULT is an empty string
		return nil
	}

	// Otherwise, unmarshal into the struct as usual
	type Alias PaymentResultDetails
	aux := (*Alias)(r)
	return json.Unmarshal(data, aux)
}

// PaymentVerify Returns the current state of a payment flow process.
// Header			Type		Required	Default	Description
// reference		string		true		The external identifier for the payment flow process.

// PaymentVerify returns the current state of a payment flow process, handling dynamic JSON structures
func (s *AccountingService) PaymentVerify(reference string) (*FullPaymentResponse, *Response, error) {
	url := fmt.Sprintf("/accounting/ecommerce/payment/ref/%s", reference)

	// Make the request
	resp, err := do(s.client, "GET", Params{parms: nil, u: url}, nil)
	if err != nil {
		return nil, resp, err
	}

	// Initialize the FullPaymentResponse struct
	var fullPaymentResp FullPaymentResponse

	// Try to unmarshal into FullPaymentResponse
	if err := json.Unmarshal([]byte(resp.Body), &fullPaymentResp); err == nil {
		// Handle cases where RESULT is an empty string
		if fullPaymentResp.ResultDetails == (PaymentResultDetails{}) && resp.Body != "" {
			var errorResp PaymentErrorResponse
			if json.Unmarshal([]byte(resp.Body), &errorResp) == nil && errorResp.Error {
				fullPaymentResp.ErrorResponse = &errorResp
				return &fullPaymentResp, resp, fmt.Errorf("error: %s (code: %s, details: %s)", errorResp.Messages, errorResp.Code, errorResp.Details)
			}
		}
		return &fullPaymentResp, resp, nil
	}

	// Handle direct error response format
	var errorResp PaymentErrorResponse
	if err := json.Unmarshal([]byte(resp.Body), &errorResp); err == nil {
		fullPaymentResp.ErrorResponse = &errorResp
		return &fullPaymentResp, resp, fmt.Errorf("error: %s (code: %s, details: %s)", errorResp.Messages, errorResp.Code, errorResp.Details)
	}

	// Fallback for unknown formats
	return nil, resp, fmt.Errorf("unknown response format: %s", resp.Body)
}

// InvoiceVoid Void an invoice. Note that invoices that have had payments applied cannot be voided.
// Header			Type		Required	Default	Description
// invoiceGUID		numeric		true				The invoiceGUID to void

func (s *AccountingService) InvoiceVoid(invoiceGUID string) (bool, *Response, error) {
	var obj interface{}

	parms := map[string]string{}
	url := fmt.Sprintf("/accounting/invoice/%s/void", invoiceGUID)

	resp, err := do(s.client, "POST", Params{parms: parms, u: url}, obj)
	if err != nil {
		return false, resp, err
	}

	if resp.StatusCode == 204 {
		return true, resp, err
	}

	return false, resp, err

}
