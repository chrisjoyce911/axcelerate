package axcelerate

import (
	"encoding/json"
)

// CoursesService handles all interactions with Contact
type CoursesService struct {
	client *Client
}

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

// GetCourses returns a list of courses. Returns accredited, Non-accredited and e-learning courses seperately or returns all together
// Header			Type		Required	Default	Description
// ID				numeric		false				The ID of the Course to filter.
// searchTerm		string		false				The term to use when filtering activities.
// type				string		false		all		The course type to return. w = workshop, p = accredited program, el = e-learning, all = All types.
// trainingArea		string		false				The Training Area to Search
// offset			numeric		false		0		Used for paging - start at record.
// displayLength	numeric		false		10		Used for paging - total records to retrieve.
// sortColumn		numeric		false		1		The column index to sort by.
// sortDirection	string		false		ASC		The sort by direction 'ASC' OR 'DESC'.
// current			boolean		false		true	Current courses flag. True to show only current courses
// public			boolean		false		true	Whether to include public courses only. If false, returns all couse types regardless of public settings.
// lastUpdated_min	datetime	false				In 'YYYY-MM-DD hh:mm' format. The course last updated date must be greater than or equal to this datetime. Courses last updated prior to Nov 2018 may not appear. Time is optional and in client's current timezone. Only applicable to w or p types.
// lastUpdated_max	datetime	false				In 'YYYY-MM-DD hh:mm' format. The course last updated date must be less than or equal to this datetime. Courses last updated prior to Nov 2018 may not appear. Time is optional and in client's current timezone. Only applicable to w or p types.
// isActive			boolean		false				Whether to include active/inactive courses only. By default both will be included
func (s *CoursesService) GetCourses(parms map[string]string) (*[]Course, *Response, error) {
	a := new([]Course)
	resp, err := do(s.client, "GET", Params{parms: parms, u: "/courses/"}, a)

	if err != nil {
		return nil, resp, err
	}

	json.Unmarshal([]byte(resp.Body), &a)
	return a, resp, err
}
