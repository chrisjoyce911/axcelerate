package axcelerate

import (
	"encoding/json"
	"fmt"
)

// CoursesService handles all interactions with Contact
type AccountingService struct {
	client *Client
}

type Payments struct {
	TransactionDate       string `json:"TRANSACTIONDATE" time_format:"axc_date_hours"`
	TransactionProviderID int    `json:"TRANSACTIONPROVIDERID"`
	GUID                  string `json:"GUID"`
	TransactionProvider   string `json:"TRANSACTIONPROVIDER"`
	FragmentAmount        int    `json:"FRAGMENT_AMOUNT"`
}

type Items struct {
	TotalTax int `json:"TOTALTAX"`
	Children struct {
	} `json:"CHILDREN"`
	UnitPriceTax   float32     `json:"UNITPRICETAX"`
	DomainID       int         `json:"DOMAINID"`
	HasChildren    int         `json:"HASCHILDREN"`
	PartID         int         `json:"PARTID"`
	UnitPriceNet   float32     `json:"UNITPRICENETT"`
	Qty            int         `json:"QTY"`
	TaxPercent     int         `json:"TAXPERCENT"`
	TotalGross     int         `json:"TOTALGROSS"`
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
	Invoicenr          string     `json:"INVOICENR"`
	PriceGross         float32    `json:"PRICEGROSS"`
	Address2           string     `json:"ADDRESS2"`
	OwnerContactID     int        `json:"OWNERCONTACTID"`
	Organisation       string     `json:"ORGANISATION"`
	ShipLastName       string     `json:"SHIPLASTNAME"`
	ShipHousenr        string     `json:"SHIPHOUSENR"`
	ShipOrgID          int        `json:"SHIPORGID"`
	ShipPostCode       string     `json:"SHIPPOSTCODE"`
	Comment            string     `json:"COMMENT"`
	PhoneNumber        string     `json:"PHONENR"`
	OrgID              int        `json:"ORGID"`
	Payments           []Payments `json:"PAYMENTS"`
	ShipCountry        string     `json:"SHIPCOUNTRY"`
	DueDate            string     `json:"DUEDATE" time_format:"axc_date"`
	ShipOrganisation   string     `json:"SHIPORGANISATION"`
	InvoiceID          int        `json:"INVOICEID"`
	DueDateOffset      int        `json:"DUEDATEOFFSET"`
	AreItemsLocked     bool       `json:"AREITEMSLOCKED"`
	LastName           string     `json:"LASTNAME"`
	Street             string     `json:"STREET"`
	Items              []Items    `json:"ITEMS"`
	InvoiceDate        string     `json:"INVOICEDATE" time_format:"axc_date"`
	Currency           string     `json:"CURRENCY"`
	ShipPriceNett      int        `json:"SHIPPRICENETT"`
	ShipState          string     `json:"SHIPSTATE"`
	ContactName        string     `json:"CONTACTNAME"`
	ShipCountryISO3166 string     `json:"SHIPCOUNTRYISO3166"`
	Shipstreet         string     `json:"SHIPSTREET"`
	Housenr            string     `json:"HOUSENR"`
	Isarchived         bool       `json:"ISARCHIVED"`
	OrderDate          string     `json:"ORDERDATE" time_format:"axc_date"`
	ContactID          int        `json:"CONTACTID"`
	ShipCity           string     `json:"SHIPCITY"`
	Shippricegross     int        `json:"SHIPPRICEGROSS"`
	CountryISO3166     string     `json:"COUNTRYISO3166"`
	Isinvoicenrlocked  bool       `json:"ISINVOICENRLOCKED"`
	Shiptaxpercent     int        `json:"SHIPTAXPERCENT"`
	Shippricetax       int        `json:"SHIPPRICETAX"`
	State              string     `json:"STATE"`
	FirstName          string     `json:"FIRSTNAME"`
	Shipaddress2       string     `json:"SHIPADDRESS2"`
	ShipFirstName      string     `json:"SHIPFIRSTNAME"`
	City               string     `json:"CITY"`
	OrderNumber        string     `json:"ORDERNR"`
	Email              string     `json:"EMAIL"`
	PriceNett          float32    `json:"PRICENETT"`
	Country            string     `json:"COUNTRY"`
	Invguid            string     `json:"INVGUID"`
	Balance            float32    `json:"BALANCE"`
	Postcode           string     `json:"POSTCODE"`
	IsPaid             bool       `json:"ISPAID"`
}

// GetInvoice will get the details of an invoice / Update an invoice (NOTE: Currently you can only Lock Items and Finalise)
// Header			Type		Required	Default	Description
// invoiceID		numeric		true				The invoiceID to retrieve. Note that this is NOT the same as the invoice number

func (s *AccountingService) GetInvoice(invoiceID int) (Invoice, *Response, error) {
	var obj Invoice

	parms := map[string]string{"invoiceID": fmt.Sprintf("/%d", invoiceID)}

	url := fmt.Sprintf("/accounting/invoice/%d", invoiceID)
	resp, err := do(s.client, "GET", Params{parms: parms, u: url}, obj)

	if err != nil {
		return obj, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
