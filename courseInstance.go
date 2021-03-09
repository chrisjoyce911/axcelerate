package axcelerate

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
)

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

// InstanceDetail
type InstanceDetail struct {
	Complexdates        []interface{} `json:"COMPLEXDATES"`
	Cost                int64         `json:"COST"`
	CustomfieldWeekends interface{}   `json:"CUSTOMFIELD_WEEKENDS"`
	Datedescriptor      string        `json:"DATEDESCRIPTOR"`
	Enrolmentopen       bool          `json:"ENROLMENTOPEN"`
	Finishdate          string        `json:"FINISHDATE"`
	ID                  int64         `json:"ID"`
	Instanceid          int64         `json:"INSTANCEID"`
	Items               []interface{} `json:"ITEMS"`
	Linkedclassid       int64         `json:"LINKEDCLASSID"`
	Linkedelearning     []struct {
		Code       string      `json:"CODE"`
		Enddate    interface{} `json:"ENDDATE"`
		Instanceid int64       `json:"INSTANCEID"`
		Name       string      `json:"NAME"`
		Startdate  interface{} `json:"STARTDATE"`
	} `json:"LINKEDELEARNING"`
	Location           string      `json:"LOCATION"`
	Maxparticipants    int64       `json:"MAXPARTICIPANTS"`
	Minparticipants    int64       `json:"MINPARTICIPANTS"`
	Name               string      `json:"NAME"`
	Notices            interface{} `json:"NOTICES"`
	Ownercontactid     int64       `json:"OWNERCONTACTID"`
	Participants       int64       `json:"PARTICIPANTS"`
	Participantvacancy int64       `json:"PARTICIPANTVACANCY"`
	Public             bool        `json:"PUBLIC"`
	Startdate          string      `json:"STARTDATE"`
	Trainercontactid   int64       `json:"TRAINERCONTACTID"`
	Venuecontactid     int64       `json:"VENUECONTACTID"`
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

// GetCoursesInstanceDetail gets details of an Instance
func (c *CoursesService) GetCoursesInstanceDetail(instanceID int, activityType string) (InstanceDetail, error) {
	// lastUpdatedMin, lastUpdatedMax string
	var details InstanceDetail

	URL, error := url.Parse(c.client.baseURL.String())
	URL.Path = "/api/course/instance/detail"
	if error != nil {
		log.Fatal("An error occurs while handling url", error)
	}
	query := URL.Query()
	query.Set("instanceID", fmt.Sprintf("%d", instanceID))
	query.Set("type", activityType)

	URL.RawQuery = query.Encode()

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		return details, err
	}

	fmt.Println(req)
	resp, err := c.client.do(req, &details)
	if err != nil {
		return details, err
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&details)
	return details, err
}
