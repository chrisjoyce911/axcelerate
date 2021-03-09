package axcelerate

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"time"
)

// Enrolments for a course
type Enrolments []struct {
	Activities []struct {
		Activities []struct {
			Activities []interface{} `json:"ACTIVITIES"`
			Amountpaid int64         `json:"AMOUNTPAID"`
			Code       string        `json:"CODE"`
			Contactid  int64         `json:"CONTACTID"`
			Delivery   struct {
				Code        int64  `json:"CODE"`
				Description string `json:"DESCRIPTION"`
			} `json:"DELIVERY"`
			Email         string      `json:"EMAIL"`
			Enrolmentdate string      `json:"ENROLMENTDATE"`
			Finishdate    string      `json:"FINISHDATE"`
			Givenname     string      `json:"GIVENNAME"`
			ID            int64       `json:"ID"`
			Instanceid    int64       `json:"INSTANCEID"`
			Learnerid     int64       `json:"LEARNERID"`
			Mobilephone   interface{} `json:"MOBILEPHONE"`
			Name          string      `json:"NAME"`
			Startdate     string      `json:"STARTDATE"`
			Surname       string      `json:"SURNAME"`
			Type          string      `json:"TYPE"`
		} `json:"ACTIVITIES"`
		Amountpaid interface{} `json:"AMOUNTPAID"`
		Code       string      `json:"CODE"`
		Contactid  int64       `json:"CONTACTID"`
		Delivery   struct {
			Code        int64  `json:"CODE"`
			Description string `json:"DESCRIPTION"`
		} `json:"DELIVERY"`
		Email         string      `json:"EMAIL"`
		Enrolmentdate string      `json:"ENROLMENTDATE"`
		Finishdate    string      `json:"FINISHDATE"`
		Givenname     string      `json:"GIVENNAME"`
		ID            int64       `json:"ID"`
		Instanceid    int64       `json:"INSTANCEID"`
		Learnerid     int64       `json:"LEARNERID"`
		Mobilephone   interface{} `json:"MOBILEPHONE"`
		Name          string      `json:"NAME"`
		Outcomecode   string      `json:"OUTCOMECODE"`
		Startdate     string      `json:"STARTDATE"`
		Status        string      `json:"STATUS"`
		Surname       string      `json:"SURNAME"`
		Type          string      `json:"TYPE"`
	} `json:"ACTIVITIES"`
	Amountpaid int64  `json:"AMOUNTPAID"`
	Code       string `json:"CODE"`
	Contactid  int64  `json:"CONTACTID"`
	Delivery   struct {
		Code        int64  `json:"CODE"`
		Description string `json:"DESCRIPTION"`
	} `json:"DELIVERY"`
	Email         string      `json:"EMAIL"`
	Enrolmentdate string      `json:"ENROLMENTDATE"`
	Finishdate    interface{} `json:"FINISHDATE"`
	Givenname     string      `json:"GIVENNAME"`
	ID            int64       `json:"ID"`
	Instanceid    int64       `json:"INSTANCEID"`
	Learnerid     int64       `json:"LEARNERID"`
	Mobilephone   interface{} `json:"MOBILEPHONE"`
	Name          string      `json:"NAME"`
	Schoolorgid   int64       `json:"SCHOOLORGID"`
	Schooltypeid  string      `json:"SCHOOLTYPEID"`
	Startdate     string      `json:"STARTDATE"`
	Surname       string      `json:"SURNAME"`
	Type          string      `json:"TYPE"`
}

// GetenrolmentsOptions used to query enrolements
// Header					Type		Required	Default	Description
// contactID				numeric		false				The ID of the Contact.
// orgID					numeric		false				The ID of the Organisation.
// instanceID				numeric		false				The ID of the Activity Instance.
// ID						numeric		false				The ID Activity Type - Use this in combination with enrolmentDateOlderThan for speed.
// lastUpdated_min			datetime	false				In 'YYYY-MM-DD hh:mm' format. The enrolment last updated date must be greater than or equal to this datetime. Enrolments last updated prior to Nov 2018 may not appear. Time is optional and in client's current timezone. NOTE: lastUpdated_min & max must be used together and can be up to 90 days apart.
// lastUpdated_max			datetime	false				In 'YYYY-MM-DD hh:mm' format. The enrolment last updated date must be less than or equal to this datetime. Enrolments last updated prior to Nov 2018 may not appear. Time is optional and in client's current timezone.
// enrolmentDate_min		datetime	false				In 'YYYY-MM-DD hh:mm' format. The enrolment date must be greater than or equal to this datetime. Time is optional and in client's current timezone. NOTE: enrolmentDate_min & max must be used together and can be up to 30 days apart.
// enrolmentDate_max		datetime	false				In 'YYYY-MM-DD hh:mm' format. The enrolment date must be less than or equal to this datetime. Time is optional and in client's current timezone.
// type						string		false		p		The type of the activity. w = workshop, p = accredited program, el = e-learning. Only p & w work at this time
// filterType				string		false				Filter related activities. s = show Subjects and related eLearning, el = show related eLearning only. Only s & el works at this time. To use this filter you must also pass a contactID for the student.
type GetenrolmentsOptions struct {
	ContactID        int       `url:"contactID,omitempty"`
	OrgID            int       `url:"orgID,omitempty"`
	InstanceID       int       `url:"instanceID,omitempty"`
	ID               int       `url:"ID,omitempty"`
	LastUpdatedMin   time.Time `url:"lastUpdated_min,omitempty"`
	LastUpdatedMax   time.Time `url:"lastUpdated_max,omitempty"`
	EnrolmentDateMin time.Time `url:"enrolmentDate_min,omitempty"`
	EnrolmentDateMax time.Time `url:"enrolmentDate_max,omitempty"`
	ActivityType     string    `url:"type,omitempty"`
	FilterType       string    `url:"filterType,omitempty"`
}

// GetEnrolments returns an array of structs containing the unique learnerID and contactID for each student's enrolment
func (c *CoursesService) GetEnrolments(opts *GetenrolmentsOptions) ([]Course, error) {
	var courses []Course
	tmp := opts
	if tmp == nil {
		tmp = &GetenrolmentsOptions{}
	}

	u, err := addOptions("/api/course/enrolments", tmp)
	if err != nil {
		log.Print("An error occurs while passing options ", err)
		return courses, err
	}

	URL, err := url.Parse(c.client.baseURL.String())
	URL.Path = u
	if err != nil {
		log.Print("An error occurs while handling url ", err)
	}

	req, err := http.NewRequest("GET", URL.String(), nil)
	if err != nil {
		log.Print("Debug 1", err)
		return courses, err
	}

	resp, err := c.client.do(req, &courses)
	if err != nil {
		if err != nil {
			log.Print("Debug 2", err)
			return courses, err
		}
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&courses)
	return courses, err
}
