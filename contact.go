package axcelerate

import (
	"encoding/json"
	"fmt"
)

// ContactService handles all interactions with Contact
type ContactService struct {
	client *Client
}

// Contact object with the full contact information
type Contact struct {
	Address1                            string        `json:"ADDRESS1"`
	Address2                            string        `json:"ADDRESS2"`
	Agentcontactid                      interface{}   `json:"AGENTCONTACTID"`
	Anzscocode                          interface{}   `json:"ANZSCOCODE"`
	Anzsiccode                          interface{}   `json:"ANZSICCODE"`
	Atschoolflag                        bool          `json:"ATSCHOOLFLAG"`
	Atschoolname                        interface{}   `json:"ATSCHOOLNAME"`
	Buildingname                        interface{}   `json:"BUILDINGNAME"`
	Categoryids                         []interface{} `json:"CATEGORYIDS"`
	Citizenstatusid                     interface{}   `json:"CITIZENSTATUSID"`
	Citizenstatusname                   interface{}   `json:"CITIZENSTATUSNAME"`
	City                                string        `json:"CITY"`
	Cityofbirth                         interface{}   `json:"CITYOFBIRTH"`
	Coachcontactid                      interface{}   `json:"COACHCONTACTID"`
	Comment                             interface{}   `json:"COMMENT"`
	Contactactive                       bool          `json:"CONTACTACTIVE"`
	Contactentrydate                    string        `json:"CONTACTENTRYDATE"`
	Contactid                           int64         `json:"CONTACTID"`
	Contactroleid                       interface{}   `json:"CONTACTROLEID"`
	Country                             string        `json:"COUNTRY"`
	Countryid                           int64         `json:"COUNTRYID"`
	Countryofbirthid                    interface{}   `json:"COUNTRYOFBIRTHID"`
	Countryofbirthname                  interface{}   `json:"COUNTRYOFBIRTHNAME"`
	Countryofcitizenid                  interface{}   `json:"COUNTRYOFCITIZENID"`
	Countryofcitizenname                interface{}   `json:"COUNTRYOFCITIZENNAME"`
	Currentschoollevel                  interface{}   `json:"CURRENTSCHOOLLEVEL"`
	CustomfieldConcernsaboutthecourse   interface{}   `json:"CUSTOMFIELD_CONCERNSABOUTTHECOURSE"`
	CustomfieldQrcode                   interface{}   `json:"CUSTOMFIELD_QRCODE"`
	CustomfieldThirdpartydec            interface{}   `json:"CUSTOMFIELD_THIRDPARTYDEC"`
	CUSTOMFIELDTHIRDPARTYEMPLOYEREMAIL2 interface{}   `json:"CUSTOMFIELD_THIRDPARTYEMPLOYEREMAIL2"`
	CustomfieldThirdpartyemployername   interface{}   `json:"CUSTOMFIELD_THIRDPARTYEMPLOYERNAME"`
	Disabilityflag                      bool          `json:"DISABILITYFLAG"`
	Disabilitytypeids                   []interface{} `json:"DISABILITYTYPEIDS"`
	Disabilitytypenames                 []interface{} `json:"DISABILITYTYPENAMES"`
	Division                            interface{}   `json:"DIVISION"`
	Dob                                 string        `json:"DOB"`
	Domainids                           []interface{} `json:"DOMAINIDS"`
	Emailaddress                        string        `json:"EMAILADDRESS"`
	Emailaddressalternative             interface{}   `json:"EMAILADDRESSALTERNATIVE"`
	Emergencycontact                    interface{}   `json:"EMERGENCYCONTACT"`
	Emergencycontactphone               interface{}   `json:"EMERGENCYCONTACTPHONE"`
	Emergencycontactrelation            interface{}   `json:"EMERGENCYCONTACTRELATION"`
	Employercontactid                   interface{}   `json:"EMPLOYERCONTACTID"`
	Englishassistanceflag               interface{}   `json:"ENGLISHASSISTANCEFLAG"`
	Englishproficiencyid                interface{}   `json:"ENGLISHPROFICIENCYID"`
	Fax                                 interface{}   `json:"FAX"`
	Givenname                           string        `json:"GIVENNAME"`
	Highestschoollevelid                interface{}   `json:"HIGHESTSCHOOLLEVELID"`
	Highestschoollevelyear              interface{}   `json:"HIGHESTSCHOOLLEVELYEAR"`
	Historicclientid                    interface{}   `json:"HISTORICCLIENTID"`
	Ielts                               interface{}   `json:"IELTS"`
	Indigenousstatusid                  interface{}   `json:"INDIGENOUSSTATUSID"`
	Indigenousstatusname                interface{}   `json:"INDIGENOUSSTATUSNAME"`
	Labourforceid                       interface{}   `json:"LABOURFORCEID"`
	Labourforcename                     interface{}   `json:"LABOURFORCENAME"`
	Lastupdated                         string        `json:"LASTUPDATED"`
	Lui                                 interface{}   `json:"LUI"`
	Mainlanguageid                      interface{}   `json:"MAINLANGUAGEID"`
	Mainlanguagename                    interface{}   `json:"MAINLANGUAGENAME"`
	Middlename                          interface{}   `json:"MIDDLENAME"`
	Mobilephone                         string        `json:"MOBILEPHONE"`
	Optionalid                          interface{}   `json:"OPTIONALID"`
	Organisation                        string        `json:"ORGANISATION"`
	Orgid                               int64         `json:"ORGID"`
	Otherphone                          interface{}   `json:"OTHERPHONE"`
	Payercontactid                      interface{}   `json:"PAYERCONTACTID"`
	Phone                               interface{}   `json:"PHONE"`
	Photo                               interface{}   `json:"PHOTO"`
	Pobox                               interface{}   `json:"POBOX"`
	Position                            string        `json:"POSITION"`
	Postcode                            string        `json:"POSTCODE"`
	Preferredname                       interface{}   `json:"PREFERREDNAME"`
	Prioreducationids                   []interface{} `json:"PRIOREDUCATIONIDS"`
	Prioreducationnames                 []interface{} `json:"PRIOREDUCATIONNAMES"`
	Prioreducationstatus                bool          `json:"PRIOREDUCATIONSTATUS"`
	Sacestudentid                       interface{}   `json:"SACESTUDENTID"`
	SADDRESS1                           string        `json:"SADDRESS1"`
	SADDRESS2                           string        `json:"SADDRESS2"`
	Sbuildingname                       interface{}   `json:"SBUILDINGNAME"`
	Scity                               string        `json:"SCITY"`
	Scountry                            string        `json:"SCOUNTRY"`
	Scountryid                          int64         `json:"SCOUNTRYID"`
	Section                             interface{}   `json:"SECTION"`
	Sex                                 interface{}   `json:"SEX"`
	Source                              interface{}   `json:"SOURCE"`
	Sourcecodeid                        int64         `json:"SOURCECODEID"`
	Spobox                              interface{}   `json:"SPOBOX"`
	Spostcode                           string        `json:"SPOSTCODE"`
	Sstate                              string        `json:"SSTATE"`
	Sstreetname                         string        `json:"SSTREETNAME"`
	Sstreetno                           string        `json:"SSTREETNO"`
	State                               string        `json:"STATE"`
	Streetname                          string        `json:"STREETNAME"`
	Streetno                            string        `json:"STREETNO"`
	Sunitno                             string        `json:"SUNITNO"`
	Supervisorcontactid                 interface{}   `json:"SUPERVISORCONTACTID"`
	Surname                             string        `json:"SURNAME"`
	Surveycontactstatuscode             interface{}   `json:"SURVEYCONTACTSTATUSCODE"`
	TfnRecorded                         bool          `json:"TFN_RECORDED"`
	Title                               interface{}   `json:"TITLE"`
	Unitno                              string        `json:"UNITNO"`
	Usi                                 string        `json:"USI"`
	UsiExemption                        bool          `json:"USI_EXEMPTION"`
	UsiVerified                         bool          `json:"USI_VERIFIED"`
	Vsn                                 interface{}   `json:"VSN"`
	Website                             interface{}   `json:"WEBSITE"`
	Workphone                           interface{}   `json:"WORKPHONE"`
	Workreadyparticipantnumber          interface{}   `json:"WORKREADYPARTICIPANTNUMBER"`
}

// ContactOptions for Updateing
type ContactOptions struct {
	GivenName    int `url:"givenName"`
	Surname      int `url:"surname"`
	Title        int `url:"title"`
	EmailAddress int `url:"emailAddress"`
}

// GetContact Interacts with a specfic contact.
func (s *ContactService) GetContact(contactID int) (*Contact, *Response, error) {
	a := new(Contact)

	resp, err := do(s.client, "GET", Params{u: fmt.Sprintf("/contact/%d", contactID)}, a)
	if err != nil {
		return nil, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &a)

	return a, resp, err
}
