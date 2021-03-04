package axcelerate

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// ContactService handles all interactions with Contact
type ContactService service

// Contact object with the full contact information
type Contact struct {
	ContactID                           int           `json:"CONTACTID"`
	GivenName                           string        `json:"GIVENNAME"`
	Surname                             string        `json:"SURNAME"`
	EmailAddress                        string        `json:"EMAILADDRESS"`
	EmailAddressAlternative             interface{}   `json:"EMAILADDRESSALTERNATIVE"`
	Sex                                 interface{}   `json:"SEX"`
	DOB                                 string        `json:"DOB"`
	HISTORICCLIENTID                    interface{}   `json:"HISTORICCLIENTID"`
	OPTIONALID                          interface{}   `json:"OPTIONALID"`
	USI                                 string        `json:"USI"`
	USIVERIFIED                         bool          `json:"USI_VERIFIED"`
	USIEXEMPTION                        bool          `json:"USI_EXEMPTION"`
	VSN                                 interface{}   `json:"VSN"`
	LUI                                 interface{}   `json:"LUI"`
	TFNRECORDED                         bool          `json:"TFN_RECORDED"`
	WORKREADYPARTICIPANTNUMBER          interface{}   `json:"WORKREADYPARTICIPANTNUMBER"`
	SACESTUDENTID                       interface{}   `json:"SACESTUDENTID"`
	Title                               interface{}   `json:"TITLE"`
	MIDDLENAME                          interface{}   `json:"MIDDLENAME"`
	PREFERREDNAME                       interface{}   `json:"PREFERREDNAME"`
	POSITION                            string        `json:"POSITION"`
	SECTION                             interface{}   `json:"SECTION"`
	DIVISION                            interface{}   `json:"DIVISION"`
	ORGANISATION                        string        `json:"ORGANISATION"`
	ORGID                               int           `json:"ORGID"`
	BUILDINGNAME                        interface{}   `json:"BUILDINGNAME"`
	UNITNO                              string        `json:"UNITNO"`
	STREETNO                            string        `json:"STREETNO"`
	STREETNAME                          string        `json:"STREETNAME"`
	POBOX                               interface{}   `json:"POBOX"`
	ADDRESS1                            string        `json:"ADDRESS1"`
	ADDRESS2                            string        `json:"ADDRESS2"`
	CITY                                string        `json:"CITY"`
	STATE                               string        `json:"STATE"`
	POSTCODE                            string        `json:"POSTCODE"`
	COUNTRYID                           int           `json:"COUNTRYID"`
	COUNTRY                             string        `json:"COUNTRY"`
	SBUILDINGNAME                       interface{}   `json:"SBUILDINGNAME"`
	SUNITNO                             string        `json:"SUNITNO"`
	SSTREETNO                           string        `json:"SSTREETNO"`
	SSTREETNAME                         string        `json:"SSTREETNAME"`
	SPOBOX                              interface{}   `json:"SPOBOX"`
	SADDRESS1                           string        `json:"SADDRESS1"`
	SADDRESS2                           string        `json:"SADDRESS2"`
	SCITY                               string        `json:"SCITY"`
	SSTATE                              string        `json:"SSTATE"`
	SPOSTCODE                           string        `json:"SPOSTCODE"`
	SCOUNTRYID                          int           `json:"SCOUNTRYID"`
	SCOUNTRY                            string        `json:"SCOUNTRY"`
	PHONE                               interface{}   `json:"PHONE"`
	MOBILEPHONE                         string        `json:"MOBILEPHONE"`
	WORKPHONE                           interface{}   `json:"WORKPHONE"`
	FAX                                 interface{}   `json:"FAX"`
	OTHERPHONE                          interface{}   `json:"OTHERPHONE"`
	SOURCECODEID                        int           `json:"SOURCECODEID"`
	SOURCE                              interface{}   `json:"SOURCE"`
	COMMENT                             interface{}   `json:"COMMENT"`
	WEBSITE                             interface{}   `json:"WEBSITE"`
	CITIZENSTATUSID                     interface{}   `json:"CITIZENSTATUSID"`
	CITIZENSTATUSNAME                   interface{}   `json:"CITIZENSTATUSNAME"`
	COUNTRYOFBIRTHID                    interface{}   `json:"COUNTRYOFBIRTHID"`
	COUNTRYOFBIRTHNAME                  interface{}   `json:"COUNTRYOFBIRTHNAME"`
	CITYOFBIRTH                         interface{}   `json:"CITYOFBIRTH"`
	COUNTRYOFCITIZENID                  interface{}   `json:"COUNTRYOFCITIZENID"`
	COUNTRYOFCITIZENNAME                interface{}   `json:"COUNTRYOFCITIZENNAME"`
	INDIGENOUSSTATUSID                  interface{}   `json:"INDIGENOUSSTATUSID"`
	INDIGENOUSSTATUSNAME                interface{}   `json:"INDIGENOUSSTATUSNAME"`
	MAINLANGUAGEID                      interface{}   `json:"MAINLANGUAGEID"`
	MAINLANGUAGENAME                    interface{}   `json:"MAINLANGUAGENAME"`
	ENGLISHPROFICIENCYID                interface{}   `json:"ENGLISHPROFICIENCYID"`
	ENGLISHASSISTANCEFLAG               interface{}   `json:"ENGLISHASSISTANCEFLAG"`
	HIGHESTSCHOOLLEVELID                interface{}   `json:"HIGHESTSCHOOLLEVELID"`
	HIGHESTSCHOOLLEVELYEAR              interface{}   `json:"HIGHESTSCHOOLLEVELYEAR"`
	CURRENTSCHOOLLEVEL                  interface{}   `json:"CURRENTSCHOOLLEVEL"`
	ATSCHOOLFLAG                        bool          `json:"ATSCHOOLFLAG"`
	ATSCHOOLNAME                        interface{}   `json:"ATSCHOOLNAME"`
	PRIOREDUCATIONIDS                   []interface{} `json:"PRIOREDUCATIONIDS"`
	PRIOREDUCATIONNAMES                 []interface{} `json:"PRIOREDUCATIONNAMES"`
	PRIOREDUCATIONSTATUS                bool          `json:"PRIOREDUCATIONSTATUS"`
	DISABILITYFLAG                      bool          `json:"DISABILITYFLAG"`
	DISABILITYTYPEIDS                   []interface{} `json:"DISABILITYTYPEIDS"`
	DISABILITYTYPENAMES                 []interface{} `json:"DISABILITYTYPENAMES"`
	LABOURFORCEID                       interface{}   `json:"LABOURFORCEID"`
	LABOURFORCENAME                     interface{}   `json:"LABOURFORCENAME"`
	EMERGENCYCONTACT                    interface{}   `json:"EMERGENCYCONTACT"`
	EMERGENCYCONTACTRELATION            interface{}   `json:"EMERGENCYCONTACTRELATION"`
	EMERGENCYCONTACTPHONE               interface{}   `json:"EMERGENCYCONTACTPHONE"`
	ANZSCOCODE                          interface{}   `json:"ANZSCOCODE"`
	ANZSICCODE                          interface{}   `json:"ANZSICCODE"`
	IELTS                               interface{}   `json:"IELTS"`
	SURVEYCONTACTSTATUSCODE             interface{}   `json:"SURVEYCONTACTSTATUSCODE"`
	EMPLOYERCONTACTID                   interface{}   `json:"EMPLOYERCONTACTID"`
	PAYERCONTACTID                      interface{}   `json:"PAYERCONTACTID"`
	SUPERVISORCONTACTID                 interface{}   `json:"SUPERVISORCONTACTID"`
	COACHCONTACTID                      interface{}   `json:"COACHCONTACTID"`
	AGENTCONTACTID                      interface{}   `json:"AGENTCONTACTID"`
	CONTACTROLEID                       interface{}   `json:"CONTACTROLEID"`
	CUSTOMFIELDCONCERNSABOUTTHECOURSE   interface{}   `json:"CUSTOMFIELD_CONCERNSABOUTTHECOURSE"`
	CUSTOMFIELDTHIRDPARTYEMPLOYEREMAIL2 interface{}   `json:"CUSTOMFIELD_THIRDPARTYEMPLOYEREMAIL2"`
	CUSTOMFIELDTHIRDPARTYEMPLOYERNAME   interface{}   `json:"CUSTOMFIELD_THIRDPARTYEMPLOYERNAME"`
	CUSTOMFIELDQRCODE                   interface{}   `json:"CUSTOMFIELD_QRCODE"`
	CUSTOMFIELDTHIRDPARTYDEC            interface{}   `json:"CUSTOMFIELD_THIRDPARTYDEC"`
	CATEGORYIDS                         []interface{} `json:"CATEGORYIDS"`
	LASTUPDATED                         string        `json:"LASTUPDATED"`
	CONTACTENTRYDATE                    string        `json:"CONTACTENTRYDATE"`
	CONTACTACTIVE                       bool          `json:"CONTACTACTIVE"`
	PHOTO                               interface{}   `json:"PHOTO"`
	DOMAINIDS                           []interface{} `json:"DOMAINIDS"`
}

// ContactOptions for Updateing
type ContactOptions struct {
	GivenName    int `url:"givenName"`
	Surname      int `url:"surname"`
	Title        int `url:"title"`
	EmailAddress int `url:"emailAddress"`
}

// GetContact Interacts with a specfic contact.
func (c *ContactService) GetContact(contactID int) (Contact, error) {
	var contact Contact
	path := fmt.Sprintf("/api/contact/%d", contactID)
	rel := &url.URL{Path: path}
	u := c.client.baseURL.ResolveReference(rel)
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return contact, err
	}

	resp, err := c.client.do(req, &contact)
	if err != nil {
		return contact, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&contact)
	return contact, err
}
