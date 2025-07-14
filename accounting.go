package axcelerate

import (
	"encoding/json"
	"fmt"
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
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

// TransactionFragment represents a fragment of a transaction,
// such as a payment applied to a particular invoice or credit note.
type TransactionFragment struct {
	InvoiceID    *StringInt  `json:"INVOICEID"`                                // InvoiceID is the ID of the invoice this fragment is applied to (nullable).
	CreditNoteID *StringInt  `json:"CREDITNOTEID"`                             // CreditNoteID is the ID of the credit note if this fragment is related to one (nullable).
	LockedDate   *time.Time  `json:"LOCKEDDATE" time_format:"axc_date_hours"`  // LockedDate is the date when the fragment was locked, if applicable (nullable).
	AppliedDate  *time.Time  `json:"APPLIEDDATE" time_format:"axc_date_hours"` // AppliedDate is the date when this fragment was applied (nullable).
	IsLocked     *bool       `json:"ISLOCKED"`                                 // IsLocked is a flag or identifier for whether this fragment is locked (nullable).
	Amount       StringFloat `json:"AMOUNT"`                                   // Amount is the amount applied in this fragment.
	FragmentID   StringInt   `json:"FRAGMENTID"`                               // FragmentID is the unique identifier for this fragment.
}

// Transaction represents a single transaction object as returned by the accounting API.
type Transaction struct {
	TransactionID         StringInt             `json:"TRANSACTIONID"`                          // TransactionID is the unique identifier for the transaction.
	ChequeNr              *string               `json:"CHEQUENR"`                               // ChequeNr is the cheque number, if applicable (nullable).
	ContactID             StringInt             `json:"CONTACTID"`                              // ContactID is the ID of the contact the transaction is associated with.
	GUID                  string                `json:"GUID"`                                   // GUID is the globally unique identifier for this transaction.
	ChequeDrawer          *string               `json:"CHEQUEDRAWER"`                           // ChequeDrawer is the name of the cheque drawer (nullable, for cheque payments).
	Start                 *time.Time            `json:"START" time_format:"axc_date_hours"`     // Start is the transaction start time (nullable).
	UnassignedAmount      StringInt             `json:"UNASSIGNEDAMOUNT"`                       // UnassignedAmount is any remaining amount not applied to invoices.
	Reference             *string               `json:"REFERENCE"`                              // Reference is the reference or receipt number for the transaction (nullable).
	IsCompleted           bool                  `json:"ISCOMPLETED"`                            // IsCompleted indicates whether the transaction is completed.
	TransDate             time.Time             `json:"TRANSDATE" time_format:"axc_date_hours"` // TransDate is the date and time the transaction was made.
	Finish                *time.Time            `json:"FINISH" time_format:"axc_date_hours"`    // Finish is the transaction finish time (nullable).
	BankName              *string               `json:"BANKNAME"`                               // BankName is the name of the bank (nullable, for cheque payments).
	PaymentMethodID       StringInt             `json:"PAYMENTMETHODID"`                        // PaymentMethodID is the ID for the payment method used (e.g., 1=Cash, 2=Credit Card, etc.).
	TransactionProviderID StringInt             `json:"TRANSACTIONPROVIDERID"`                  // TransactionProviderID is the provider ID for this transaction.
	TransactionTypeID     StringInt             `json:"TRANSACTIONTYPEID"`                      // TransactionTypeID is the type of transaction.
	Organisation          *StringInt            `json:"ORGANISATION"`                           // Organisation is the name of the associated organisation (nullable).
	OrgID                 *int                  `json:"ORGID"`                                  // OrgID is the unique identifier of the organisation (nullable).
	BankBSB               *string               `json:"BANKBSB"`                                // BankBSB is the BSB number for the bank (nullable, 6-digit string, cheque payments).
	Currency              string                `json:"CURRENCY"`                               // Currency is the currency code for the transaction (e.g., "AUD").
	Fragments             []TransactionFragment `json:"FRAGMENTS"`                              // Fragments contains one or more fragments (e.g., how payment is applied to invoices).
	Description           *string               `json:"DESCRIPTION"`                            // Description is an optional description for the transaction (nullable).
	Amount                StringFloat           `json:"AMOUNT"`                                 // Amount is the dollar amount of the transaction.
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

// CreateTransaction sends a POST request to create a new accounting transaction.
//
// Allowed params (all should be provided in the 'parms' map):
//
//	contactID        (required, numeric)   - The contact ID this transaction is for (payer).
//	amount           (required, numeric)   - Transaction amount in dollars.
//	invoiceID        (optional, numeric)   - If provided, applies this transaction to the given invoice ID.
//	paymentMethodID  (optional, numeric)   - Payment method: 1=Cash, 2=Credit Card (default), 4=Direct Deposit, 5=Cheque, 6=EFTPOS.
//	transDate        (optional, datetime)  - Date of the transaction. Format: "2006-01-02 15:04". Defaults to now.
//	reference        (optional, string)    - An optional reference or receipt number.
//	description      (optional, string)    - An optional description for the transaction.
//	ChequeNr         (optional, string)    - Cheque number (only for paymentMethodID=5).
//	ChequeDrawer     (optional, string)    - Cheque drawer (only for paymentMethodID=5).
//	BankName         (optional, string)    - Bank name (only for paymentMethodID=5).
//	BankBSB          (optional, string)    - 6-digit BSB number (only for paymentMethodID=5).
//
// Returns:
//
//	The created Transaction object, the API Response, and any error encountered.
func (s *AccountingService) CreateTransaction(parms map[string]string) (*Transaction, *Response, error) {

	var obj Transaction

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	resp, err := do(s.client, "POST", Params{parms: parms, u: "/accounting/transaction/"}, obj)

	if err != nil {
		return nil, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_datetime", "2006-01-02 15:04:05")
	jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")

	// Unmarshal the response body into the result Transaction
	err = json.Unmarshal([]byte(resp.Body), &obj)

	return &obj, resp, err
}
