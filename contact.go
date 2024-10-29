package axcelerate

import (
	"fmt"
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// ContactService handles all interactions with Contact
type ContactService struct {
	client *Client
}

// Contact object with the full contact information
type Contact struct {
	Address1                            string        `json:"ADDRESS1"`
	Address2                            *string       `json:"ADDRESS2"`
	AgentContactID                      *int          `json:"AGENTCONTACTID"`
	Anzscocode                          *string       `json:"ANZSCOCODE"`
	Anzsiccode                          *string       `json:"ANZSICCODE"`
	Atschoolflag                        *bool         `json:"ATSCHOOLFLAG"`
	Atschoolname                        *string       `json:"ATSCHOOLNAME"`
	Buildingname                        *string       `json:"BUILDINGNAME"`
	CategoryIDs                         []interface{} `json:"CATEGORYIDS"`
	Citizenstatusid                     *int          `json:"CITIZENSTATUSID"`
	Citizenstatusname                   *string       `json:"CITIZENSTATUSNAME"`
	City                                string        `json:"CITY"`
	CityOfBirth                         *string       `json:"CITYOFBIRTH"`
	Coachcontactid                      *int64        `json:"COACHCONTACTID"`
	Comment                             *string       `json:"COMMENT"`
	Contactactive                       bool          `json:"CONTACTACTIVE"`
	Contactentrydate                    time.Time     `json:"CONTACTENTRYDATE" time_format:"axc_date_hours"`
	ContactID                           int           `json:"CONTACTID"`
	ContactRoleID                       *int          `json:"CONTACTROLEID"`
	Country                             string        `json:"COUNTRY"`
	CountryID                           int           `json:"COUNTRYID"`
	CountryofBirthID                    *int          `json:"COUNTRYOFBIRTHID"`
	Countryofbirthname                  *string       `json:"COUNTRYOFBIRTHNAME"`
	Countryofcitizenid                  *int          `json:"COUNTRYOFCITIZENID"`
	Countryofcitizenname                *string       `json:"COUNTRYOFCITIZENNAME"`
	Currentschoollevel                  *string       `json:"CURRENTSCHOOLLEVEL"`
	CustomfieldConcernsaboutthecourse   *string       `json:"CUSTOMFIELD_CONCERNSABOUTTHECOURSE"`
	CustomfieldQrcode                   *string       `json:"CUSTOMFIELD_QRCODE"`
	CustomfieldThirdpartydec            *string       `json:"CUSTOMFIELD_THIRDPARTYDEC"`
	CUSTOMFIELDTHIRDPARTYEMPLOYEREMAIL2 *string       `json:"CUSTOMFIELD_THIRDPARTYEMPLOYEREMAIL2"`
	CustomfieldThirdpartyemployername   *string       `json:"CUSTOMFIELD_THIRDPARTYEMPLOYERNAME"`
	Disabilityflag                      *bool         `json:"DISABILITYFLAG"`
	Disabilitytypeids                   []interface{} `json:"DISABILITYTYPEIDS"`
	Disabilitytypenames                 []interface{} `json:"DISABILITYTYPENAMES"`
	Division                            *string       `json:"DIVISION"`
	DOB                                 *time.Time    `json:"DOB" time_format:"axc_date"`
	Domainids                           []interface{} `json:"DOMAINIDS"`
	Emailaddress                        string        `json:"EMAILADDRESS"`
	Emailaddressalternative             *string       `json:"EMAILADDRESSALTERNATIVE"`
	Emergencycontact                    *string       `json:"EMERGENCYCONTACT"`
	Emergencycontactphone               *string       `json:"EMERGENCYCONTACTPHONE"`
	Emergencycontactrelation            *string       `json:"EMERGENCYCONTACTRELATION"`
	Employercontactid                   *int          `json:"EMPLOYERCONTACTID"`
	Englishassistanceflag               *bool         `json:"ENGLISHASSISTANCEFLAG"`
	Englishproficiencyid                *int          `json:"ENGLISHPROFICIENCYID"`
	Fax                                 *string       `json:"FAX"`
	Givenname                           string        `json:"GIVENNAME"`
	HighestschoollevelID                *int          `json:"HIGHESTSCHOOLLEVELID"`
	Highestschoollevelyear              *string       `json:"HIGHESTSCHOOLLEVELYEAR"`
	Historicclientid                    *string       `json:"HISTORICCLIENTID"`
	Ielts                               *string       `json:"IELTS"`
	Indigenousstatusid                  *int          `json:"INDIGENOUSSTATUSID"`
	Indigenousstatusname                *string       `json:"INDIGENOUSSTATUSNAME"`
	LabourforceID                       *int          `json:"LABOURFORCEID"`
	Labourforcename                     *string       `json:"LABOURFORCENAME"`
	LastUpdated                         time.Time     `json:"LASTUPDATED" time_format:"axc_date_hours"`
	Lui                                 *string       `json:"LUI"`
	Mainlanguageid                      *int          `json:"MAINLANGUAGEID"`
	Mainlanguagename                    *string       `json:"MAINLANGUAGENAME"`
	Middlename                          *string       `json:"MIDDLENAME"`
	Mobilephone                         *string       `json:"MOBILEPHONE"`
	OptionalID                          *string       `json:"OPTIONALID"`
	Organisation                        string        `json:"ORGANISATION"`
	OrgID                               int           `json:"ORGID"`
	OtherPhone                          *string       `json:"OTHERPHONE"`
	PayerContactID                      *int64        `json:"PAYERCONTACTID"`
	Phone                               *string       `json:"PHONE"`
	Photo                               *string       `json:"PHOTO"`
	Pobox                               *string       `json:"POBOX"`
	Position                            *string       `json:"POSITION"`
	Postcode                            string        `json:"POSTCODE"`
	PreferredName                       *string       `json:"PREFERREDNAME"`
	Prioreducationids                   []interface{} `json:"PRIOREDUCATIONIDS"`
	Prioreducationnames                 []interface{} `json:"PRIOREDUCATIONNAMES"`
	Prioreducationstatus                *bool         `json:"PRIOREDUCATIONSTATUS"`
	Sacestudentid                       *string       `json:"SACESTUDENTID"`
	SADDRESS1                           *string       `json:"SADDRESS1"`
	SADDRESS2                           *string       `json:"SADDRESS2"`
	Sbuildingname                       *string       `json:"SBUILDINGNAME"`
	Scity                               *string       `json:"SCITY"`
	Scountry                            *string       `json:"SCOUNTRY"`
	Scountryid                          *int          `json:"SCOUNTRYID"`
	Section                             *string       `json:"SECTION"`
	Sex                                 *string       `json:"SEX"`
	Source                              *string       `json:"SOURCE"`
	Sourcecodeid                        *int          `json:"SOURCECODEID"`
	Spobox                              *string       `json:"SPOBOX"`
	Spostcode                           *string       `json:"SPOSTCODE"`
	Sstate                              *string       `json:"SSTATE"`
	Sstreetname                         *string       `json:"SSTREETNAME"`
	Sstreetno                           *string       `json:"SSTREETNO"`
	State                               string        `json:"STATE"`
	Streetname                          *string       `json:"STREETNAME"`
	Streetno                            *string       `json:"STREETNO"`
	Sunitno                             *string       `json:"SUNITNO"`
	Supervisorcontactid                 interface{}   `json:"SUPERVISORCONTACTID"`
	Surname                             string        `json:"SURNAME"`
	Surveycontactstatuscode             interface{}   `json:"SURVEYCONTACTSTATUSCODE"`
	TfnRecorded                         bool          `json:"TFN_RECORDED"`
	Title                               string        `json:"TITLE"`
	Unitno                              string        `json:"UNITNO"`
	USI                                 string        `json:"USI"`
	USIExemption                        bool          `json:"USI_EXEMPTION"`
	USIVerified                         bool          `json:"USI_VERIFIED"`
	Vsn                                 interface{}   `json:"VSN"`
	Website                             interface{}   `json:"WEBSITE"`
	Workphone                           string        `json:"WORKPHONE"`
	Workreadyparticipantnumber          interface{}   `json:"WORKREADYPARTICIPANTNUMBER"`
}

type ContactEnrolment struct {
	RowID                                int         `json:"ROWID"`
	Type                                 string      `json:"TYPE"`
	ID                                   int         `json:"ID"`
	InstanceID                           int         `json:"INSTANCEID"`
	EnrolID                              int         `json:"ENROLID"`
	VicenrolmentID                       string      `json:"VICENROLMENTID,omitempty"`
	InvoiceID                            int         `json:"INVOICEID"`
	InvoicePaid                          bool        `json:"INVOICEPAID"`
	LearnerID                            int         `json:"LEARNERID"`
	Code                                 string      `json:"CODE"`
	Location                             string      `json:"LOCATION"`
	Delivery                             string      `json:"DELIVERY"`
	DeliveryMode                         string      `json:"DELIVERYMODE"`
	Activitytype                         string      `json:"ACTIVITYTYPE"`
	Name                                 string      `json:"NAME"`
	CommencedDate                        time.Time   `json:"COMMENCEDDATE" time_format:"axc_date"`
	StartDate                            time.Time   `json:"STARTDATE" time_format:"axc_date"`
	FinishDate                           time.Time   `json:"FINISHDATE" time_format:"axc_date"`
	CompletionDate                       time.Time   `json:"COMPLETIONDATE" time_format:"axc_date"`
	Mandatory                            bool        `json:"MANDATORY"`
	Status                               string      `json:"STATUS"`
	ProgramStatusidvic                   interface{} `json:"PROGRAMSTATUSIDVIC"`
	SchoolTypeID                         interface{} `json:"SCHOOLTYPEID"`
	SchoolOrgID                          interface{} `json:"SCHOOLORGID"`
	Count                                int         `json:"COUNT"`
	OutcomeCode                          string      `json:"OUTCOMECODE,omitempty"`
	LaunchURL                            string      `json:"LAUNCHURL,omitempty"`
	Pstacdatevic                         interface{} `json:"PSTACDATEVIC,omitempty"`
	VicprogramenrolmentID                string      `json:"VICPROGRAMENROLMENTID,omitempty"`
	CustomfieldAgreetoelearning          string      `json:"CUSTOMFIELD_AGREETOELEARNING,omitempty"`
	CustomfieldAgreetosoa                interface{} `json:"CUSTOMFIELD_AGREETOSOA,omitempty"`
	CustomfieldTerms                     string      `json:"CUSTOMFIELD_TERMS,omitempty"`
	CustomfieldDeclaration               interface{} `json:"CUSTOMFIELD_DECLARATION,omitempty"`
	CustomfieldAvetmissconsent           interface{} `json:"CUSTOMFIELD_AVETMISSCONSENT,omitempty"`
	CustomfieldPfaquizdate               string      `json:"CUSTOMFIELD_PFAQUIZDATE,omitempty"`
	CustomfieldPfaquizlink               string      `json:"CUSTOMFIELD_PFAQUIZLINK,omitempty"`
	CustomfieldPfaquiz                   string      `json:"CUSTOMFIELD_PFAQUIZ,omitempty"`
	CustomfieldEmployername              interface{} `json:"CUSTOMFIELD_EMPLOYERNAME,omitempty"`
	CustomfieldSoaverificationstatus     interface{} `json:"CUSTOMFIELD_SOAVERIFICATIONSTATUS,omitempty"`
	CustomfieldSoaverify                 interface{} `json:"CUSTOMFIELD_SOAVERIFY,omitempty"`
	CustomfieldSoaverifieddate           interface{} `json:"CUSTOMFIELD_SOAVERIFIEDDATE,omitempty"`
	CustomfieldStatementofattainmentlink interface{} `json:"CUSTOMFIELD_STATEMENTOFATTAINMENTLINK,omitempty"`
	CustomfieldThirdpartyconsent         interface{} `json:"CUSTOMFIELD_THIRDPARTYCONSENT,omitempty"`
	CustomfieldThirdpartyemployer        interface{} `json:"CUSTOMFIELD_THIRDPARTYEMPLOYER,omitempty"`
	CustomfieldThirdpartyemployeremail   interface{} `json:"CUSTOMFIELD_THIRDPARTYEMPLOYEREMAIL,omitempty"`
	CustomfieldTrueandcorrectdec         interface{} `json:"CUSTOMFIELD_TRUEANDCORRECTDEC,omitempty"`
	CustomfieldCprlink                   interface{} `json:"CUSTOMFIELD_CPRLINK,omitempty"`
	// Complexdates                         []EnrolmentComplexdates `json:"COMPLEXDATES,omitempty"`
	CustomfieldPriceBeat   interface{} `json:"CUSTOMFIELD_PRICEBEAT,omitempty"`
	CustomfieldCompeteName interface{} `json:"CUSTOMFIELD_COMPETENAME,omitempty"`
	CustomfieldWeekends    interface{} `json:"CUSTOMFIELD_WEEKENDS,omitempty"`
}

type EnrolmentComplexdates struct {
	Date      time.Time `json:"DATE" time_format:"axc_date"`
	StartTime time.Time `json:"STARTTIME" time_format:"axc_time_long"`
	EndTime   time.Time `json:"ENDTIME" time_format:"axc_time_long"`

	TrainerContactID int    `json:"TRAINERCONTACTID"`
	Location         string `json:"LOCATION"`
	RoomID           int    `json:"ROOMID"`
}

// ContactOptions for Updateing
type ContactOptions struct {
	GivenName    int `url:"givenName"`
	Surname      int `url:"surname"`
	Title        int `url:"title"`
	EmailAddress int `url:"emailAddress"`
}

type USIstatus struct {
	Data        USIdata `json:"DATA"`
	Msg         string  `json:"MSG"`
	UsiVerified bool    `json:"USI_VERIFIED"`
}

type USIdata struct {
	DateOfBirth string `json:"dateOfBirth"`
	FamilyName  string `json:"familyName"`
	FirstName   string `json:"firstName"`
	UsiStatus   string `json:"usiStatus"`
}

// GetContact Interacts with a specific contact.
func (s *ContactService) GetContact(contactID int) (Contact, *Response, error) {
	var a Contact

	resp, err := do(s.client, "GET", Params{u: fmt.Sprintf("/contact/%d", contactID)}, a)
	if err != nil {
		return a, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")
	jsontime.AddTimeFormatAlias("axc_date", "2006-01-02")

	err = json.Unmarshal([]byte(resp.Body), &a)

	return a, resp, err
}

// GetContact Interacts with a specific contact.
func (s *ContactService) SearchContacts(parms map[string]string) ([]Contact, *Response, error) {
	var obj []Contact

	resp, err := do(s.client, "GET", Params{parms: parms, u: "/contacts/search"}, obj)

	if err != nil {
		return obj, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")
	jsontime.AddTimeFormatAlias("axc_date", "2006-01-02")

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
