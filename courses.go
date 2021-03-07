package axcelerate

import (
	"encoding/json"
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
