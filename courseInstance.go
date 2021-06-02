package axcelerate

import (
	"encoding/json"
	"fmt"
)

// Instance of a course
type Instance struct {
	Cost                int         `json:"COST"`
	CustomfieldWeekends interface{} `json:"CUSTOMFIELD_WEEKENDS"`
	Datedescriptor      string      `json:"DATEDESCRIPTOR"`
	EnrolmentOpen       bool        `json:"ENROLMENTOPEN"`
	Finishdate          string      `json:"FINISHDATE"`
	CourseID            int         `json:"ID"`
	InstanceID          int         `json:"INSTANCEID"`
	Isactive            bool        `json:"ISACTIVE"`
	Location            string      `json:"LOCATION"`
	Maxparticipants     int         `json:"MAXPARTICIPANTS"`
	Minparticipants     int         `json:"MINPARTICIPANTS"`
	Name                string      `json:"NAME"`
	Notices             interface{} `json:"NOTICES"`
	Ownercontactid      int         `json:"OWNERCONTACTID"`
	Participants        int         `json:"PARTICIPANTS"`
	Participantvacancy  int         `json:"PARTICIPANTVACANCY"`
	Startdate           string      `json:"STARTDATE"`
	TrainercontactID    int         `json:"TRAINERCONTACTID"`
	VirtualclassroomID  interface{} `json:"VIRTUALCLASSROOMID"`
	COUNT               int         `json:"COUNT"`
}

/*

GetCoursesInstances returns a list of instances

activityType
	The type of the activity. w = workshop, p = accredited program, el = e-learning.

Request Parameters

public
	Whether to include public courses only. If false, returns ALL course instances for type w and el. For type p, passing false will return ONLY non-public classes.
current
	Whether to include only current courses. A current course instance is a course that is currently running, or coming up. If false, returns all course instances.
isActive
	Whether to include active/inactive courses instances only. By default both will be included
lastUpdated_min	datetime
	In YYYY-MM-DD hh:mm format. The course instance last updated date must be greater than or equal to this datetime. Instances last updated prior to Nov 2018 may not appear. Time is optional and in clients current timezone. Only applicable to w or p types.
lastUpdated_max	datetime
	In YYYY-MM-DD hh:mm format. The course instance last updated date must be less than or equal to this datetime. Instances last updated prior to Nov 2018 may not appear. Time is optional and in clients current timezone. Only applicable to w or p types.
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

	json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}
