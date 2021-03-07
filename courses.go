package axcelerate

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

// CoursesService handles all interactions with Contact
type CoursesService service

// Course object with the full contact information
type Course struct {
	Code             string      `json:"CODE"`
	Cost             int64       `json:"COST"`
	Count            int64       `json:"COUNT"`
	Delivery         string      `json:"DELIVERY"`
	Duration         float32     `json:"DURATION"`
	Durationtype     interface{} `json:"DURATIONTYPE"`
	ID               int64       `json:"ID"`
	Isactive         bool        `json:"ISACTIVE"`
	Name             string      `json:"NAME"`
	Primaryimage     interface{} `json:"PRIMARYIMAGE"`
	Rowid            int64       `json:"ROWID"`
	Secondaryimage   interface{} `json:"SECONDARYIMAGE"`
	Shortdescription interface{} `json:"SHORTDESCRIPTION"`
	Streamname       interface{} `json:"STREAMNAME"`
	Type             string      `json:"TYPE"`
}

// Instance of a course
type Instance struct {
	Cost                int64       `json:"COST"`
	CustomfieldWeekends interface{} `json:"CUSTOMFIELD_WEEKENDS"`
	Datedescriptor      string      `json:"DATEDESCRIPTOR"`
	Enrolmentopen       bool        `json:"ENROLMENTOPEN"`
	Finishdate          string      `json:"FINISHDATE"`
	CoursesID           int64       `json:"ID"`
	InstanceID          int64       `json:"INSTANCEID"`
	Isactive            bool        `json:"ISACTIVE"`
	Location            string      `json:"LOCATION"`
	Maxparticipants     int64       `json:"MAXPARTICIPANTS"`
	Minparticipants     int64       `json:"MINPARTICIPANTS"`
	Name                string      `json:"NAME"`
	Notices             interface{} `json:"NOTICES"`
	Ownercontactid      int64       `json:"OWNERCONTACTID"`
	Participants        int64       `json:"PARTICIPANTS"`
	Participantvacancy  int64       `json:"PARTICIPANTVACANCY"`
	Startdate           string      `json:"STARTDATE"`
	Trainercontactid    int64       `json:"TRAINERCONTACTID"`
	Virtualclassroomid  interface{} `json:"VIRTUALCLASSROOMID"`
}

// Attended for a Course Instance.
type Attended struct {
	Attendanceid     interface{} `json:"ATTENDANCEID"`
	Attendedduration interface{} `json:"ATTENDEDDURATION"`
	Attendedflag     interface{} `json:"ATTENDEDFLAG"`
	Comment          string      `json:"COMMENT"`
	Completedflag    bool        `json:"COMPLETEDFLAG"`
	Contactid        int64       `json:"CONTACTID"`
	Enrolid          int64       `json:"ENROLID"`
	Finishtime       string      `json:"FINISHTIME"`
	Givenname        string      `json:"GIVENNAME"`
	Starttime        string      `json:"STARTTIME"`
	Surname          string      `json:"SURNAME"`
}

// Attendeds for a Course Instance.
type Attendeds struct {
	Complexid  interface{} `json:"COMPLEXID"`
	Date       string      `json:"DATE"`
	Duration   interface{} `json:"DURATION"`
	Enrollees  []Attended  `json:"ENROLLEES"`
	Finishtime string      `json:"FINISHTIME"`
	Instanceid int64       `json:"INSTANCEID"`
	Starttime  string      `json:"STARTTIME"`
	Type       string      `json:"TYPE"`
}

// CoursesOptions for Updateing
type CoursesOptions struct {
	CoursesID      int    `url:"ID"`              // The ID of the Course to filter.
	SearchTerm     string `url:"searchTerm"`      // The term to use when filtering activities.
	CourseType     string `url:"type"`            // The course type to return. w = workshop, p = accredited program, el = e-learning, all = All types.
	TrainingArea   string `url:"trainingArea"`    // The Training Area to Search
	Offset         int    `url:"offset"`          // Used for paging - start at record.
	DisplayLength  int    `url:"displayLength"`   // Used for paging - total records to retrieve.
	SortColumn     int    `url:"sortColumn"`      // The column index to sort by.
	SortDirection  string `url:"sortDirection"`   // The sort by direction 'ASC' OR 'DESC'.
	Current        bool   `url:"current"`         // Current courses flag. True to show only current courses
	Public         bool   `url:"public"`          // Whether to include public courses only. If false, returns all couse types regardless of public settings.
	LastUpdatedMin bool   `url:"lastUpdated_min"` // In 'YYYY-MM-DD hh:mm' format. The course last updated date must be greater than or equal to this datetime. Courses last updated prior to Nov 2018 may not appear. Time is optional and in client's current timezone. Only applicable to w or p types.
	LastUpdatedMax bool   `url:"lastUpdated_max"` // In 'YYYY-MM-DD hh:mm' format. The course last updated date must be less than or equal to this datetime. Courses last updated prior to Nov 2018 may not appear. Time is optional and in client's current timezone. Only applicable to w or p types.
	IsActive       bool   `url:"givenName"`       // Whether to include active/inactive courses only. By default both will be included
}

// GetCourses returns a list of courses. Returns accredited, Non-accredited and e-learning courses seperately or returns all together
func (c *CoursesService) GetCourses() ([]Course, error) {
	var courses []Course

	URL, error := url.Parse(c.client.baseURL.String())
	URL.Path = "/api/courses/"
	if error != nil {
		log.Fatal("An error occurs while handling url", error)
	}
	query := URL.Query()
	query.Set("current", "true")
	query.Set("isActive", "true")
	query.Set("displayLength", "25")
	URL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return courses, err
	}

	resp, err := c.client.do(req, &courses)
	if err != nil {
		return courses, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&courses)
	return courses, err
}

// GetCoursesInstances returns a list of instances
func (c *CoursesService) GetCoursesInstances(coursesID int, active bool) ([]Instance, error) {
	// lastUpdatedMin, lastUpdatedMax string
	var instances []Instance

	URL, error := url.Parse(c.client.baseURL.String())
	URL.Path = "/api/course/instances"
	if error != nil {
		log.Fatal("An error occurs while handling url", error)
	}
	query := URL.Query()
	query.Set("ID", fmt.Sprintf("%d", coursesID))
	if active {
		query.Set("current", "1")
	} else {
		query.Set("current", "0")
	}

	query.Set("type", "w")

	URL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return instances, err
	}

	resp, err := c.client.do(req, &instances)
	if err != nil {
		return instances, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&instances)
	return instances, err
}

// GetCoursesAttendeds for a Course Instance.
func (c *CoursesService) GetCoursesAttendeds(instanceID int) ([]Attendeds, error) {
	var attendeds []Attendeds

	URL, error := url.Parse(c.client.baseURL.String())
	URL.Path = "/api/course/instance/attendance"
	if error != nil {
		log.Fatal("An error occurs while handling url", error)
	}
	query := URL.Query()
	query.Set("instanceID", fmt.Sprintf("%d", instanceID))

	URL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return attendeds, err
	}

	resp, err := c.client.do(req, &attendeds)
	if err != nil {
		return attendeds, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&attendeds)
	return attendeds, err
}
