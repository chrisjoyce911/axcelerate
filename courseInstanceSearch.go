package axcelerate

import jsontime "github.com/liamylian/jsontime/v2/v2"

/*

GetCoursesInstanceSearch Advanced Course Instance Search - Returns instances.

Request Parameters

ID
	The Activity Type ID.
InstanceID
	The Instance ID.
type
	The type of the activity. w = workshop, p = accredited program, el = e-learning, all = workshops, accredited programs and e-learning.
trainingCategory
	The Training Category to Search - Uses Like operator %name%
location
	The course location to search- Uses Like operator %name%. Only works with type 'w' instances.
state
	The course State to search - works with type 'w' & 'p' instances.
code
	The course code to search - Uses Like operator: code%
name
	The course name to search - Uses Like operator %name%
searchTerm
	For a general search use this param
enrolmentOpen
	Return Course Instances that are open for enrolment.
startDate_min
	The course start date must be greater than this date. Null values will also be returned for el and p types. Ignored if instanceID is passed.
startDate_max
	The course start date must be less than this date. Null values will also be returned for el and p types. Ignored if instanceID is passed.
finishDate_min
	The course finish date must be greater than this date. Null values will also be returned for el and p types.
finishDate_max
	The course finish date must be less than this date. Null values will also be returned for el and p types.
lastUpdated_min
	In 'YYYY-MM-DD hh:mm' format with time optional. The course instance last updated date must be greater than or equal to this datetime. NOTE: lastUpdated_min & max must be used together (unless ID is passed) and can be up to 90 days apart. These fields are mutually exclusive with start and finish date min/max searches and are both ignored if instanceID is passed.
lastUpdated_max
	In 'YYYY-MM-DD hh:mm' format with time optional. The course instance last updated date must be less than or equal to this datetime.
trainerContactID
	The ContactID of the Trainer/Consultant the instance is assigned to.
domainID
	The DomainID the instance belongs to (the domainID of the user).
deliveryLocationID
	For type = p only. The unique ID of an accredited delivery location, reported to NCVER. Refers to locations listed under the course/deliveryLocations endpoint.
orgID
	The organisation ID of the Client Contact of the course.
orgIDTree
	The Client Contact of the course is either this Organisation ID or a child organisation of this Organisation ID.
offset
	Used for paging - start at record.
displayLength
	Used for paging - total records to retrieve.
sortColumn
	The column index to sort by.
sortDirection
	The sort by direction 'ASC' OR 'DESC'.
public
	Whether to include public courses. If false, returns only In-House course instances.
isActive
	You can chose to include or exclude Deleted / Archived and Inactive courses.
purgeCache
	Currently the API will cache the query for 30 seconds - Setting this flag to true gets the latest data.
groupedCourseName
	If the Grouped Workshop Data flag is on in your account, you can search by the Grouped Course Name (Type W Only)
groupedCourseID
	If the Grouped Workshop Data flag is on in your account, you can search by the Grouped Course ID (Type W Only)
*/
func (s *CoursesService) GetCoursesInstanceSearch(parms map[string]string) ([]Instance, *Response, error) {
	var obj []Instance

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	resp, err := do(s.client, "POST", Params{parms: parms, u: "/course/instance/search"}, obj)

	if err != nil {
		return obj, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_datetime", "2006-01-02 15:04:05")

	err = json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}

// 2020-11-30 13:00:00
