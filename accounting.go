package axcelerate

import (
	"encoding/json"
	"fmt"
	"log"
)

// CoursesService handles all interactions with Contact
type AccountingService struct {
	client *Client
}

type Payments struct {
	Transactiondate       string `json:"TRANSACTIONDATE" time_format:"axc_date_hours"`
	Transactionproviderid int    `json:"TRANSACTIONPROVIDERID"`
	GUID                  string `json:"GUID"`
	Transactionprovider   string `json:"TRANSACTIONPROVIDER"`
	FragmentAmount        int    `json:"FRAGMENT_AMOUNT"`
}

type Items struct {
	Totaltax int `json:"TOTALTAX"`
	Children struct {
	} `json:"CHILDREN"`
	Unitpricetax   int         `json:"UNITPRICETAX"`
	DomainID      int         `json:"DOMAINID"`
	HasChildren    int         `json:"HASCHILDREN"`
	PartID         int         `json:"PARTID"`
	Unitpricenett  int         `json:"UNITPRICENETT"`
	Qty            int         `json:"QTY"`
	TaxPercent     int         `json:"TAXPERCENT"`
	TotalGross     int         `json:"TOTALGROSS"`
	ItemCode       string      `json:"ITEMCODE"`
	Financecode    string      `json:"FINANCECODE"`
	Totalnett      int         `json:"TOTALNETT"`
	SourceID       int         `json:"SOURCEID"`
	ServiceDate    string      `json:"SERVICEDATE" time_format:"axc_date"`
	Unitpricegross int         `json:"UNITPRICEGROSS"`
	Itemid         int         `json:"ITEMID"`
	Description    string      `json:"DESCRIPTION"`
	Costcentrecode interface{} `json:"COSTCENTRECODE"`
}

// Invoice object with the full Invoice information
type Invoice struct {
	Invoicenr          string      `json:"INVOICENR"`
	PriceGross         int         `json:"PRICEGROSS"`
	Address2           interface{} `json:"ADDRESS2"`
	OwnerContactID     int         `json:"OWNERCONTACTID"`
	Organisation       interface{} `json:"ORGANISATION"`
	ShipLastName       string      `json:"SHIPLASTNAME"`
	Shiphousenr        string      `json:"SHIPHOUSENR"`
	ShipOrgID          interface{} `json:"SHIPORGID"`
	ShipPostCode       interface{} `json:"SHIPPOSTCODE"`
	Comment            interface{} `json:"COMMENT"`
	Phonenr            string      `json:"PHONENR"`
	OrgID              interface{} `json:"ORGID"`
	Payments           []Payments  `json:"PAYMENTS"`
	Shipcountry        interface{} `json:"SHIPCOUNTRY"`
	DueDate            string      `json:"DUEDATE" time_format:"axc_date"`
	Shiporganisation   interface{} `json:"SHIPORGANISATION"`
	Invoiceid          int         `json:"INVOICEID"`
	Duedateoffset      int         `json:"DUEDATEOFFSET"`
	Areitemslocked     bool        `json:"AREITEMSLOCKED"`
	Lastname           string      `json:"LASTNAME"`
	Street             string      `json:"STREET"`
	Items              []Items     `json:"ITEMS"`
	InvoiceDate        string      `json:"INVOICEDATE" time_format:"axc_date"`
	Currency           string      `json:"CURRENCY"`
	Shippricenett      int         `json:"SHIPPRICENETT"`
	Shipstate          interface{} `json:"SHIPSTATE"`
	Contactname        string      `json:"CONTACTNAME"`
	Shipcountryiso3166 interface{} `json:"SHIPCOUNTRYISO3166"`
	Shipstreet         string      `json:"SHIPSTREET"`
	Housenr            string      `json:"HOUSENR"`
	Isarchived         bool        `json:"ISARCHIVED"`
	OrderDate          string      `json:"ORDERDATE" time_format:"axc_date"`
	ContactID         int         `json:"CONTACTID"`
	Shipcity           interface{} `json:"SHIPCITY"`
	Shippricegross     int         `json:"SHIPPRICEGROSS"`
	Countryiso3166     interface{} `json:"COUNTRYISO3166"`
	Isinvoicenrlocked  bool        `json:"ISINVOICENRLOCKED"`
	Shiptaxpercent     int         `json:"SHIPTAXPERCENT"`
	Shippricetax       int         `json:"SHIPPRICETAX"`
	State              interface{} `json:"STATE"`
	Firstname          string      `json:"FIRSTNAME"`
	Shipaddress2       interface{} `json:"SHIPADDRESS2"`
	Shipfirstname      string      `json:"SHIPFIRSTNAME"`
	City               interface{} `json:"CITY"`
	Ordernr            interface{} `json:"ORDERNR"`
	Email              string      `json:"EMAIL"`
	Pricenett          int         `json:"PRICENETT"`
	Country            interface{} `json:"COUNTRY"`
	Invguid            string      `json:"INVGUID"`
	Balance            int         `json:"BALANCE"`
	Postcode           interface{} `json:"POSTCODE"`
	IsPaid             bool        `json:"ISPAID"`
}

// GetInvoice will get the details of an invoice / Update an invoice (NOTE: Currently you can only Lock Items and Finalise)
// Header			Type		Required	Default	Description
// invoiceID		numeric		true				The invoiceID to retrieve. Note that this is NOT the same as the invoice number

func (s *AccountingService) GetInvoice(invoiceID int) (Invoice, *Response, error) {
	var obj Invoice

	parms := map[string]string{"invoiceID": fmt.Sprintf("/%d",invoiceID), }

	url :=fmt.Sprintf("/accounting/invoice/%d",invoiceID)
	log.Print(url)

	resp, err := do(s.client, "GET", Params{parms: parms, u: url}, obj)

	if err != nil {
		return obj, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
