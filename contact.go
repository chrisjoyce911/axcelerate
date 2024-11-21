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
	AnzscoCode                          *string       `json:"ANZSCOCODE"`
	AnzsicCode                          *string       `json:"ANZSICCODE"`
	AtSchoolFlag                        *bool         `json:"ATSCHOOLFLAG"`
	AtSchoolName                        *string       `json:"ATSCHOOLNAME"`
	BuildingName                        *string       `json:"BUILDINGNAME"`
	CategoryIDs                         []interface{} `json:"CATEGORYIDS"`
	CitizenStatusID                     *int          `json:"CITIZENSTATUSID"`
	CitizenStatusName                   *string       `json:"CITIZENSTATUSNAME"`
	City                                string        `json:"CITY"`
	CityOfBirth                         *string       `json:"CITYOFBIRTH"`
	CoachContactID                      *int64        `json:"COACHCONTACTID"`
	Comment                             *string       `json:"COMMENT"`
	ContactActive                       bool          `json:"CONTACTACTIVE"`
	ContactEntryDate                    time.Time     `json:"CONTACTENTRYDATE" time_format:"axc_date_hours"`
	ContactID                           int           `json:"CONTACTID"`
	ContactRoleID                       *int          `json:"CONTACTROLEID"`
	Country                             string        `json:"COUNTRY"`
	CountryID                           int           `json:"COUNTRYID"`
	CountryOfBirthID                    *int          `json:"COUNTRYOFBIRTHID"`
	CountryOfBirthName                  *string       `json:"COUNTRYOFBIRTHNAME"`
	CountryOfCitizenID                  *int          `json:"COUNTRYOFCITIZENID"`
	CountryOfCitizenName                *string       `json:"COUNTRYOFCITIZENNAME"`
	CurrentSchoolLevel                  *interface{}  `json:"CURRENTSCHOOLLEVEL"`
	CustomFieldConcernsAboutTheCourse   *string       `json:"CUSTOMFIELD_CONCERNSABOUTTHECOURSE"`
	CustomFieldQRCode                   *string       `json:"CUSTOMFIELD_QRCODE"`
	CustomFieldThirdPartyDec            *string       `json:"CUSTOMFIELD_THIRDPARTYDEC"`
	CustomFieldThirdPartyEmployerEmail2 *string       `json:"CUSTOMFIELD_THIRDPARTYEMPLOYEREMAIL2"`
	CustomFieldThirdPartyEmployerName   *string       `json:"CUSTOMFIELD_THIRDPARTYEMPLOYERNAME"`
	DisabilityFlag                      *bool         `json:"DISABILITYFLAG"`
	DisabilityTypeIDs                   []interface{} `json:"DISABILITYTYPEIDS"`
	DisabilityTypeNames                 []interface{} `json:"DISABILITYTYPENAMES"`
	Division                            *string       `json:"DIVISION"`
	DOB                                 *time.Time    `json:"DOB" time_format:"axc_date"`
	DomainIDs                           []interface{} `json:"DOMAINIDS"`
	EmailAddress                        string        `json:"EMAILADDRESS"`
	EmailAddressAlternative             *string       `json:"EMAILADDRESSALTERNATIVE"`
	EmergencyContact                    *string       `json:"EMERGENCYCONTACT"`
	EmergencyContactPhone               *string       `json:"EMERGENCYCONTACTPHONE"`
	EmergencyContactRelation            *string       `json:"EMERGENCYCONTACTRELATION"`
	EmployerContactID                   *int          `json:"EMPLOYERCONTACTID"`
	EnglishAssistanceFlag               *bool         `json:"ENGLISHASSISTANCEFLAG"`
	EnglishProficiencyID                *int          `json:"ENGLISHPROFICIENCYID"`
	Fax                                 *string       `json:"FAX"`
	GivenName                           string        `json:"GIVENNAME"`
	HighestSchoolLevelID                *int          `json:"HIGHESTSCHOOLLEVELID"`
	HighestSchoolLevelYear              *string       `json:"HIGHESTSCHOOLLEVELYEAR"`
	HistoricClientID                    *string       `json:"HISTORICCLIENTID"`
	Ielts                               *string       `json:"IELTS"`
	IndigenousStatusID                  *int          `json:"INDIGENOUSSTATUSID"`
	IndigenousStatusName                *string       `json:"INDIGENOUSSTATUSNAME"`
	LabourForceID                       *int          `json:"LABOURFORCEID"`
	LabourForceName                     *string       `json:"LABOURFORCENAME"`
	LastUpdated                         time.Time     `json:"LASTUPDATED" time_format:"axc_date_hours"`
	LUI                                 *string       `json:"LUI"`
	MainLanguageID                      *int          `json:"MAINLANGUAGEID"`
	MainLanguageName                    *string       `json:"MAINLANGUAGENAME"`
	MiddleName                          *string       `json:"MIDDLENAME"`
	MobilePhone                         *string       `json:"MOBILEPHONE"`
	OptionalID                          *string       `json:"OPTIONALID"`
	Organisation                        string        `json:"ORGANISATION"`
	OrgID                               int           `json:"ORGID"`
	OtherPhone                          *string       `json:"OTHERPHONE"`
	PayerContactID                      *int64        `json:"PAYERCONTACTID"`
	Phone                               *string       `json:"PHONE"`
	Photo                               *string       `json:"PHOTO"`
	POBox                               *string       `json:"POBOX"`
	Position                            *string       `json:"POSITION"`
	Postcode                            string        `json:"POSTCODE"`
	PreferredName                       *string       `json:"PREFERREDNAME"`
	PriorEducationIDs                   []interface{} `json:"PRIOREDUCATIONIDS"`
	PriorEducationNames                 []interface{} `json:"PRIOREDUCATIONNAMES"`
	PriorEducationStatus                *bool         `json:"PRIOREDUCATIONSTATUS"`
	SACEStudentID                       *string       `json:"SACESTUDENTID"`
	SAddress1                           *string       `json:"SADDRESS1"`
	SAddress2                           *string       `json:"SADDRESS2"`
	SBuildingName                       *string       `json:"SBUILDINGNAME"`
	SCity                               *string       `json:"SCITY"`
	SCountry                            *string       `json:"SCOUNTRY"`
	SCountryID                          *int          `json:"SCOUNTRYID"`
	Section                             *string       `json:"SECTION"`
	Sex                                 *string       `json:"SEX"`
	Source                              *string       `json:"SOURCE"`
	SourceCodeID                        *int          `json:"SOURCECODEID"`
	SPOBox                              *string       `json:"SPOBOX"`
	SPostcode                           *string       `json:"SPOSTCODE"`
	SState                              *string       `json:"SSTATE"`
	SStreetName                         *string       `json:"SSTREETNAME"`
	SStreetNo                           *string       `json:"SSTREETNO"`
	State                               string        `json:"STATE"`
	StreetName                          *string       `json:"STREETNAME"`
	StreetNo                            *string       `json:"STREETNO"`
	SUnitNo                             *string       `json:"SUNITNO"`
	SupervisorContactID                 interface{}   `json:"SUPERVISORCONTACTID"`
	Surname                             string        `json:"SURNAME"`
	SurveyContactStatusCode             interface{}   `json:"SURVEYCONTACTSTATUSCODE"`
	TFNRecorded                         bool          `json:"TFN_RECORDED"`
	Title                               string        `json:"TITLE"`
	UnitNo                              string        `json:"UNITNO"`
	USI                                 string        `json:"USI"`
	USIExemption                        bool          `json:"USI_EXEMPTION"`
	USIVerified                         bool          `json:"USI_VERIFIED"`
	VSN                                 interface{}   `json:"VSN"`
	Website                             interface{}   `json:"WEBSITE"`
	WorkPhone                           string        `json:"WORKPHONE"`
	WorkReadyParticipantNumber          interface{}   `json:"WORKREADYPARTICIPANTNUMBER"`
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
	CustomfieldEmployerName              interface{} `json:"CUSTOMFIELD_EMPLOYERNAME,omitempty"`
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

/*
ContactCreate creates a new contact in the system. This method does not search for existing contacts and will directly create a new one.
If duplicates are created, they can be merged manually through the aXcelerate interface.

Request Parameters (passed as a map[string]string):
--------------------------------------------
Header                      Type        Required    Default     Description
givenName                   string      true                    Given (first) name. Maximum 40 characters.
surname                     string      true                    Surname (last name). Maximum 40 characters.
title                       string      false                   Title or salutation.
emailAddress                string      false                   A valid email address.
ContactActive               boolean     false       true        Retrieve active contact records. Passing false returns inactive records.
dob                         date        false                   Date of Birth in the format YYYY-MM-DD. Cannot be a future date.
sex                         character   false                   Must be one letter: M, F, or X (for Other).
middleName                  string      false                   Middle name(s). Maximum 40 characters.
phone                       string      false                   Home phone number.
mobilephone                 string      false                   Mobile phone number.
workphone                   string      false                   Work phone number.
fax                         string      false                   Fax number.
organisation                string      false                   Organisation name.
position                    string      false                   Position within the organisation.
section                     string      false                   Section within the organisation.
division                    string      false                   Division within the organisation.
SourceCodeID                numeric     false                   Client-specific source ID from the list set up in aXcelerate.
HistoricClientID            string      false                   Historical ID used with the student record.
USI                         string      false                   Unique Student Identifier (10 characters, excludes I, 1, 0, O).
LUI                         string      false                   Learner Unique Identifier (10-digit numeric code for Queensland).
TFN                         string      false                   Tax File Number, validated using the ATO's algorithm.
VSN                         string      false                   Victorian Student Number for students in VET institutions.
WorkReadyParticipantNumber  string      false                   South Australian student identifier.
SACEStudentID               string      false                   South Australian Certificate of Education ID (6 digits + 1 letter).
EmergencyContact            string      false                   Name of an emergency contact.
EmergencyContactRelation    string      false                   Relationship of the emergency contact (e.g., sister).
EmergencyContactPhone       string      false                   Phone number of the emergency contact.
buildingName                string      false                   AVETMISS 7.0 field; prioritized over address1/address2 if passed.
address1                    string      false                   First line of postal address (used if AVETMISS fields not passed).
address2                    string      false                   Second line of postal address.
city                        string      false                   Postal suburb, locality, or town.
state                       string      false                   Postal state/territory (NSW, VIC, QLD, SA, WA, TAS, NT, ACT, OTH, OVS).
postcode                    string      false                   Postal postcode.
country                     string      false                   Postal country.

Returns:
--------------------------------------------
- *Contact: The created contact object.
- *Response: API response metadata.
- error: Error object if any issue occurs during the request.

Example Usage:
--------------------------------------------
params := map[string]string{
    "givenName":   "John",
    "surname":     "Doe",
    "emailAddress":"john.doe@example.com",
    "dob":         "1990-01-01",
}

contact, resp, err := contactService.ContactCreate(params)
if err != nil {
    log.Fatalf("Error creating contact: %v", err)
}
fmt.Printf("Contact created with ID: %d\n", contact.ContactID)
*/

func (s *ContactService) ContactCreate(parms map[string]string) (*Contact, *Response, error) {
	var obj Contact

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	resp, err := do(s.client, "POST", Params{parms: parms, u: "/contact/"}, obj)

	if err != nil {
		return nil, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_time", "15:04")
	jsontime.AddTimeFormatAlias("axc_time_long", "15:04:05")
	jsontime.AddTimeFormatAlias("axc_date", "2006-01-02")
	jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")

	err = json.Unmarshal([]byte(resp.Body), &obj)

	return &obj, resp, err
}
