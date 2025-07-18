package axcelerate

import (
	"encoding/json"
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// Enrolment for a course
type Enrolment struct {
	EnrolmentID         int        `json:"ENROLID"`
	InstanceID          int        `json:"INSTANCEID"`
	ContactID           int        `json:"CONTACTID"`
	CourseID            int        `json:"ID"`
	LearnerID           int        `json:"LEARNERID"`
	OwnerID             int        `json:"OWNERID"`
	AmountPaid          *float32   `json:"AMOUNTPAID"` // Nullable
	Code                string     `json:"CODE"`
	Coupon              *string    `json:"CUSTOMFIELD_COUPON"`      // Nullable
	CompeteName         *string    `json:"CUSTOMFIELD_COMPETENAME"` // Nullable
	PriceBeat           []string   `json:"CUSTOMFIELD_PRICEBEAT"`
	ELA                 *string    `json:"CUSTOMFIELD_PFAQUIZ"`     // Nullable
	ELALink             *string    `json:"CUSTOMFIELD_PFAQUIZLINK"` // Nullable
	ELAcomplete         *string    `json:"CUSTOMFIELD_PFAQUIZDATE"` // Nullable
	AllowEmployer       *string    `json:"CUSTOMFIELD_THIRDPARTYEMPLOYEREMAIL"`
	Delivery            *string    `json:"DELIVERY"`
	Name                string     `json:"NAME"`
	EnrolmentDate       time.Time  `json:"ENROLMENTDATE" time_format:"axc_date_hours"`
	Startdate           time.Time  `json:"STARTDATE" time_format:"axc_date_hours"`
	Finishdate          *time.Time `json:"FINISHDATE" time_format:"axc_date_hours"` // Nullable
	MobilePhone         string     `json:"MOBILEPHONE"`
	GivenName           string     `json:"GIVENNAME"`
	Surname             string     `json:"SURNAME"`
	Email               string     `json:"EMAIL"`
	Status              string     `json:"STATUS"`
	Type                string     `json:"TYPE"`
	LastUpdateUTC       time.Time  `json:"LASTUPDATEDUTC" time_format:"axc_date_hours"`
	Activities          []Activity `json:"ACTIVITIES"`           // Nested activities array
	PhotoURL            *string    `json:"PHOTOURL"`             // Nullable
	CustomFieldWeekends *string    `json:"CUSTOMFIELD_WEEKENDS"` // Nullable
}

// Delivery struct for nested DELIVERY object
type Delivery struct {
	Code        int    `json:"CODE"`
	Description string `json:"DESCRIPTION"`
}

// Activity struct to represent each activity in ACTIVITIES array
type Activity struct {
	EnrolmentDate      time.Time  `json:"ENROLMENTDATE" time_format:"axc_date_hours"`
	OutcomeCode        *string    `json:"OUTCOMECODE"`
	MobilePhone        *string    `json:"MOBILEPHONE"`
	VirtualClassroomID *string    `json:"VIRTUALCLASSROOMID"` // Nullable
	VicEnrolmentID     *string    `json:"VICENROLMENTID"`
	Delivery           *Delivery  `json:"DELIVERY"`
	InstanceID         int        `json:"INSTANCEID"`
	ContactID          int        `json:"CONTACTID"`
	StartDate          time.Time  `json:"STARTDATE" time_format:"axc_date_hours"`
	Surname            string     `json:"SURNAME"`
	Activities         []Activity `json:"ACTIVITIES"`
	Status             string     `json:"STATUS"`
	AmountPaid         *float32   `json:"AMOUNTPAID"` // Nullable
	FinishDate         *time.Time `json:"FINISHDATE" time_format:"axc_date_hours"`
	Code               string     `json:"CODE"`
	EnrolID            int        `json:"ENROLID"`
	LearnerID          int        `json:"LEARNERID"`
	Name               string     `json:"NAME"`
	ID                 int        `json:"ID"`
	Type               string     `json:"TYPE"`
	GivenName          string     `json:"GIVENNAME"`
	Email              string     `json:"EMAIL"`
}

// Custom unmarshaler for Delivery
func (d *Delivery) UnmarshalJSON(data []byte) error {
	// First, try to unmarshal as a simple string
	var deliveryStr string
	if err := json.Unmarshal(data, &deliveryStr); err == nil {
		// If successful, set Description and leave Code as default
		d.Description = deliveryStr
		return nil
	}

	// If not a string, try to unmarshal as an object
	type Alias Delivery // Create an alias to prevent recursion
	var deliveryObj Alias
	if err := json.Unmarshal(data, &deliveryObj); err != nil {
		return err
	}

	// Assign values from the object
	*d = Delivery(deliveryObj)
	return nil
}

/*
GetEnrolments returns an array of struct containing the unique learnerID and contactID for each student's enrolment

# Request Parameters

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
	jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
