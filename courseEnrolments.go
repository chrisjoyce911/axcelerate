package axcelerate

import (
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// Enrolment for a course
type Enrolment struct {
	Activities []struct {
		Activities []interface{} `json:"ACTIVITIES"`
		Amountpaid interface{}   `json:"AMOUNTPAID"`
		Code       string        `json:"CODE"`
		ContactID  int           `json:"CONTACTID"`
		Delivery   struct {
			Code        int    `json:"CODE"`
			Description string `json:"DESCRIPTION"`
		} `json:"DELIVERY"`
		Email              string      `json:"EMAIL"`
		EnrolmentID        int         `json:"ENROLID"`
		EnrolmentDate      time.Time   `json:"ENROLMENTDATE" time_format:"axc_datetime"`
		FinishDate         time.Time   `json:"FINISHDATE" time_format:"axc_datetime"`
		Givenname          string      `json:"GIVENNAME"`
		ID                 int         `json:"ID"`
		InstanceID         int         `json:"INSTANCEID"`
		LearnerID          int         `json:"LEARNERID"`
		Mobilephone        string      `json:"MOBILEPHONE"`
		Name               string      `json:"NAME"`
		Outcomecode        string      `json:"OUTCOMECODE"`
		StartDate          time.Time   `json:"STARTDATE" time_format:"axc_datetime"`
		Status             string      `json:"STATUS"`
		Surname            string      `json:"SURNAME"`
		Type               string      `json:"TYPE"`
		VicenrolmentID     string      `json:"VICENROLMENTID"`
		VirtualClassroomID interface{} `json:"VIRTUALCLASSROOMID"`
	} `json:"ACTIVITIES"`
	AmountPaid                           int         `json:"AMOUNTPAID"`
	Code                                 string      `json:"CODE"`
	CustomfieldAgreetoelearning          string      `json:"CUSTOMFIELD_AGREETOELEARNING"`
	CustomfieldAgreetosoa                interface{} `json:"CUSTOMFIELD_AGREETOSOA"`
	CustomfieldAvetmissconsent           interface{} `json:"CUSTOMFIELD_AVETMISSCONSENT"`
	CustomfieldCprlink                   interface{} `json:"CUSTOMFIELD_CPRLINK"`
	CustomfieldDeclaration               interface{} `json:"CUSTOMFIELD_DECLARATION"`
	CustomfieldEmployername              interface{} `json:"CUSTOMFIELD_EMPLOYERNAME"`
	CustomfieldPfaquiz                   interface{} `json:"CUSTOMFIELD_PFAQUIZ"`
	CustomfieldPfaquizdate               interface{} `json:"CUSTOMFIELD_PFAQUIZDATE"`
	CustomfieldPfaquizlink               interface{} `json:"CUSTOMFIELD_PFAQUIZLINK"`
	CustomfieldSoaverificationstatus     interface{} `json:"CUSTOMFIELD_SOAVERIFICATIONSTATUS"`
	CustomfieldSoaverifieddate           interface{} `json:"CUSTOMFIELD_SOAVERIFIEDDATE"`
	CustomfieldSoaverify                 interface{} `json:"CUSTOMFIELD_SOAVERIFY"`
	CustomfieldStatementofattainmentlink interface{} `json:"CUSTOMFIELD_STATEMENTOFATTAINMENTLINK"`
	CustomfieldTerms                     interface{} `json:"CUSTOMFIELD_TERMS"`
	CustomfieldThirdpartyconsent         interface{} `json:"CUSTOMFIELD_THIRDPARTYCONSENT"`
	CustomfieldThirdpartyemployer        interface{} `json:"CUSTOMFIELD_THIRDPARTYEMPLOYER"`
	CustomfieldThirdpartyemployeremail   interface{} `json:"CUSTOMFIELD_THIRDPARTYEMPLOYEREMAIL"`
	CustomfieldTrueandcorrectdec         interface{} `json:"CUSTOMFIELD_TRUEANDCORRECTDEC"`
	PriceBeat                            interface{} `json:"CUSTOMFIELD_PRICEBEAT"`
	Delivery                             string      `json:"DELIVERY"`
	ContactID                            int         `json:"CONTACTID"`
	CourseID                             int         `json:"ID"`
	InstanceID                           int         `json:"INSTANCEID"`
	LearnerID                            int         `json:"LEARNERID"`
	EnrolmentID                          int         `json:"ENROLID"`
	EnrolmentDate                        time.Time   `json:"ENROLMENTDATE" time_format:"axc_datetime"`
	Startdate                            time.Time   `json:"STARTDATE" time_format:"axc_datetime"`
	Finishdate                           time.Time   `json:"FINISHDATE" time_format:"axc_datetime"`
	Email                                string      `json:"EMAIL"`
	MobilePhone                          string      `json:"MOBILEPHONE"`
	Givenname                            string      `json:"GIVENNAME"`
	Surname                              string      `json:"SURNAME"`
	Name                                 string      `json:"NAME"`
	OwnerID                              int         `json:"OWNERID"`
	Pstacdatevic                         interface{} `json:"PSTACDATEVIC"`
	Schooldeliverylocationid             interface{} `json:"SCHOOLDELIVERYLOCATIONID"`
	SchoolorgID                          interface{} `json:"SCHOOLORGID"`
	SchooltypeID                         string      `json:"SCHOOLTYPEID"`
	Status                               string      `json:"STATUS"`
	Type                                 string      `json:"TYPE"`
	VicprogramenrolmentID                string      `json:"VICPROGRAMENROLMENTID"`
	VicprogramstatusID                   string      `json:"VICPROGRAMSTATUSID"`
}

/*
GetEnrolments returns an array of struct containing the unique learnerID and contactID for each student's enrolment

Request Parameters

contactID
	The ID of the Contact.
orgID
	The ID of the Organisation.
instanceID
	The ID of the Activity Instance.
ID
	The ID Activity Type - Use this in combination with enrolmentDateOlderThan for speed.
lastUpdated_min
	In YYYY-MM-DD hh:mm format. The enrolment last updated date must be greater than or equal to this datetime. Enrolments last updated prior to Nov 2018 may not appear. Time is optional and in clients current timezone. NOTE: lastUpdated_min & max must be used together and can be up to 90 days apart.
lastUpdated_max
	In YYYY-MM-DD hh:mm format. The enrolment last updated date must be less than or equal to this datetime. Enrolments last updated prior to Nov 2018 may not appear. Time is optional and in clients current timezone.
enrolmentDate_min
	In YYYY-MM-DD hh:mm format. The enrolment date must be greater than or equal to this datetime. Time is optional and in clients current timezone. NOTE: enrolmentDate_min & max must be used together and can be up to 30 days apart.
enrolmentDate_max
	In YYYY-MM-DD hh:mm format. The enrolment date must be less than or equal to this datetime. Time is optional and in clients current timezone.
type
	The type of the activity. w = workshop, p = accredited program, el = e-learning. Only p & w work at this time
filterType
	Filter related activities. s = show Subjects and related eLearning, el = show related eLearning only. Only s & el works at this time. To use this filter you must also pass a contactID for the student.
*/
func (s *CoursesService) GetEnrolments(parms map[string]string) ([]Enrolment, *Response, error) {
	var obj []Enrolment
	resp, err := do(s.client, "GET", Params{parms: parms, u: "/course/enrolments"}, obj)

	if err != nil {
		return obj, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_datetime", "2006-01-02 15:04:05")

	json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
