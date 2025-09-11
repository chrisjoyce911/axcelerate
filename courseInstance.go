package axcelerate

import (
	"fmt"
	"time"

	jsontime "github.com/liamylian/jsontime/v2/v2"
)

// Instance of a course
type Instance struct {
	RowID                     int               `json:"ROWID"`
	Type                      string            `json:"TYPE"`
	CourseName                string            `json:"COURSENAME"`
	Cost                      int               `json:"COST"`
	CustomFieldWeekends       interface{}       `json:"CUSTOMFIELD_WEEKENDS"`
	Code                      string            `json:"CODE"`
	DateDescriptor            string            `json:"DATEDESCRIPTOR"`
	EnrolmentOpen             bool              `json:"ENROLMENTOPEN"`
	Public                    bool              `json:"PUBLIC"`
	CourseID                  int               `json:"ID"`
	InstanceID                int               `json:"INSTANCEID"`
	IsActive                  bool              `json:"ISACTIVE"`
	Location                  string            `json:"LOCATION"`
	Domain                    string            `json:"DOMAINNAME"`
	MaxParticipants           int               `json:"MAXPARTICIPANTS"`
	MinParticipants           int               `json:"MINPARTICIPANTS"`
	Name                      string            `json:"NAME"`
	Notices                   interface{}       `json:"NOTICES"`
	OwnerContactID            int               `json:"OWNERCONTACTID"`
	Participants              int               `json:"PARTICIPANTS"`
	ParticipantVacancy        int               `json:"PARTICIPANTVACANCY"`
	StartDate                 time.Time         `json:"STARTDATE" time_format:"axc_datetime"`
	FinishDate                time.Time         `json:"FINISHDATE" time_format:"axc_datetime"`
	TrainerContactID          IntOrZero         `json:"TRAINERCONTACTID"`
	VirtualClassroomID        interface{}       `json:"VIRTUALCLASSROOMID"`
	Count                     int               `json:"COUNT"`
	LinkedClassID             int               `json:"LINKEDCLASSID"`
	VenueContactID            int               `json:"VENUECONTACTID"`
	LastUpdatedUTC            time.Time         `json:"LASTUPDATEDUTC" time_format:"axc_date_hours"`
	LinkedElearning           []LinkedElearning `json:"LINKEDELEARNING"`
	ComplexDates              []ComplexDate     `json:"COMPLEXDATES"`
	SyncDateDescriptor        bool              `json:"SYNCDATEDESCRIPTOR"`
	Items                     []interface{}     `json:"ITEMS"`
	Duration                  string            `json:"DURATION"`
	DomainID                  int               `json:"DOMAINID"`
	State                     string            `json:"STATE"`
	GroupedCourseID           *int              `json:"GROUPEDCOURSEID"`
	GroupedCourseName         *string           `json:"GROUPEDCOURSENAME"`
	GroupedCourseSimultaneous bool              `json:"GROUPEDCOURSEISSIMULTANEOUS"`
	GroupedMaxParticipants    int               `json:"GROUPEDMAXPARTICIPANTS"`
	GroupedParticipants       int               `json:"GROUPEDPARTICIPANTS"`
	TrainingCategory          interface{}       `json:"TRAININGCATEGORY"`
	GSTType                   int               `json:"GST_TYPE"`
	Status                    string            `json:"STATUS"`
}

// LinkedElearning represents the structure for linked e-learning data
type LinkedElearning struct {
	EndDate    interface{} `json:"ENDDATE"`
	InstanceID int         `json:"INSTANCEID"`
	StartDate  interface{} `json:"STARTDATE"`
	Code       string      `json:"CODE"`
	Name       string      `json:"NAME"`
}

// ComplexDate represents the structure for complex date information
type ComplexDate struct {
	ComplexID        int         `json:"COMPLEXID"`
	Date             string      `json:"DATE"`      // Keep as string for parsing flexibility
	StartTime        string      `json:"STARTTIME"` // Keep as string for parsing flexibility
	EndTime          string      `json:"ENDTIME"`   // Keep as string for parsing flexibility
	TrainerContactID IntOrZero   `json:"TRAINERCONTACTID"`
	Location         string      `json:"LOCATION"`
	Address          string      `json:"ADDRESS"`
	RoomID           int         `json:"ROOMID"`
	VenueContactID   interface{} `json:"VENUECONTACTID"`  // Flexible to handle null, string, or int
	State            *string     `json:"STATE,omitempty"` // Pointer to handle null or missing values
	Trainers         []Trainers  `json:"TRAINERS"`
}

// Trainers models each entry in the COMPLEXDATES[].TRAINERS array
type Trainers struct {
	FinishDateAndTimeOverride interface{} `json:"FINISHDATEANDTIMEOVERRIDE"`
	FKComplexID               int         `json:"FKCOMPLEXID"`
	PackUp                    bool        `json:"PACKUP"`
	TrainerContactID          int         `json:"FKTRAINERCONTACTID"`
	SetUp                     bool        `json:"SETUP"`
	StartDateAndTimeOverride  interface{} `json:"STARTDATEANDTIMEOVERRIDE"`
}

/*
GetCoursesInstances returns a list of instances

activityType

	The type of the activity. w = workshop, p = accredited program, el = e-learning.

# Request Parameters

public

	Whether to include public courses only. If false, returns ALL course instances for type w and el. For type p, passing false will return ONLY non-public classes.

current

	Whether to include only current courses. A current course instance is a course that is currently running, or coming up. If false, returns all course instances.

isActive

	Whether to include active/inactive courses instances only. By default both will be included

lastUpdated_min	date time

	In YYYY-MM-DD hh:mm format. The course instance last updated date must be greater than or equal to this date time. Instances last updated prior to Nov 2018 may not appear. Time is optional and in clients current timezone. Only applicable to w or p types.

lastUpdated_max	date time

	In YYYY-MM-DD hh:mm format. The course instance last updated date must be less than or equal to this date time. Instances last updated prior to Nov 2018 may not appear. Time is optional and in clients current timezone. Only applicable to w or p types.
*/
func (s *CoursesService) GetCoursesInstances(coursesID int, activityType string, parms map[string]string) ([]Instance, *Response, error) {
	var obj []Instance

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	parms["ID"] = fmt.Sprintf("%d", coursesID)
	parms["type"] = activityType

	resp, err := do(s.client, "GET", Params{parms: parms, u: "/course/instances"}, obj)

	if err != nil {
		return obj, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat

	jsontime.AddTimeFormatAlias("axc_datetime", "2006-01-02 15:04:05")
	jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")

	err = json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}
