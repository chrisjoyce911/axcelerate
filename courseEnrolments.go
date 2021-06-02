package axcelerate

import (
	"encoding/json"
)

// Enrolment for a course
type Enrolment struct {
	Activities []struct {
		Activities []interface{} `json:"ACTIVITIES"`
		Amountpaid interface{}   `json:"AMOUNTPAID"`
		Code       string        `json:"CODE"`
		Contactid  int64         `json:"CONTACTID"`
		Delivery   struct {
			Code        int64  `json:"CODE"`
			Description string `json:"DESCRIPTION"`
		} `json:"DELIVERY"`
		Email              string      `json:"EMAIL"`
		EnrolID            int64       `json:"ENROLID"`
		Enrolmentdate      string      `json:"ENROLMENTDATE"`
		Finishdate         interface{} `json:"FINISHDATE"`
		Givenname          string      `json:"GIVENNAME"`
		ID                 int64       `json:"ID"`
		InstanceID         int64       `json:"INSTANCEID"`
		LearnerID          int64       `json:"LEARNERID"`
		Mobilephone        string      `json:"MOBILEPHONE"`
		Name               string      `json:"NAME"`
		Outcomecode        string      `json:"OUTCOMECODE"`
		Startdate          string      `json:"STARTDATE"`
		Status             string      `json:"STATUS"`
		Surname            string      `json:"SURNAME"`
		Type               string      `json:"TYPE"`
		Vicenrolmentid     string      `json:"VICENROLMENTID"`
		Virtualclassroomid interface{} `json:"VIRTUALCLASSROOMID"`
	} `json:"ACTIVITIES"`
	Amountpaid                           int64       `json:"AMOUNTPAID"`
	Code                                 string      `json:"CODE"`
	ContactID                            int64       `json:"CONTACTID"`
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
	Delivery                             string      `json:"DELIVERY"`
	Email                                string      `json:"EMAIL"`
	EnrolID                              int64       `json:"ENROLID"`
	Enrolmentdate                        string      `json:"ENROLMENTDATE"`
	Finishdate                           interface{} `json:"FINISHDATE"`
	Givenname                            string      `json:"GIVENNAME"`
	ID                                   int64       `json:"ID"`
	InstanceID                           int64       `json:"INSTANCEID"`
	LearnerID                            int64       `json:"LEARNERID"`
	Mobilephone                          string      `json:"MOBILEPHONE"`
	Name                                 string      `json:"NAME"`
	OwnerID                              int64       `json:"OWNERID"`
	Pstacdatevic                         interface{} `json:"PSTACDATEVIC"`
	Schooldeliverylocationid             interface{} `json:"SCHOOLDELIVERYLOCATIONID"`
	SchoolorgID                          interface{} `json:"SCHOOLORGID"`
	SchooltypeID                         string      `json:"SCHOOLTYPEID"`
	Startdate                            string      `json:"STARTDATE"`
	Status                               string      `json:"STATUS"`
	Surname                              string      `json:"SURNAME"`
	Type                                 string      `json:"TYPE"`
	VicprogramenrolmentID                string      `json:"VICPROGRAMENROLMENTID"`
	VicprogramstatusID                   string      `json:"VICPROGRAMSTATUSID"`
}

/*
GetEnrolments returns an array of structs containing the unique learnerID and contactID for each student's enrolment

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

	json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
