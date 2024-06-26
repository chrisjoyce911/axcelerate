package axcelerate

import (
	jsontime "github.com/liamylian/jsontime/v2/v2"
)

/*
ContactSearch used as a contact search. This method will return an array of contacts matching all the search parameters.

The field values must be an exact match OR start with the supplied parameters (i.e. wildcard on the right of the string).

You can choose to search on almost any field or use the param 'q' or 'search' to perform search across these fields: GIVENNAME, PREFERREDNAME, MIDDLENAME, SURNAME, EMAILADDRESS and MOBILEPHONE.

Request Parameters
q

	A Search String. i.e. q='Nathan Gordon'

offset

	Record to start at.

displayLength

	Maximum number of records to return (up to a system maximum of 100)

contactEntryDate

	yyyy-mm-dd formatted date

lastUpdated

	yyyy-mm-dd formatted date

givenName

	Given (first) name

surname

	Surname (last name)

emailAddress

	Email address. Must be a valid email address

contactRoleID

	Filter contacts who are in a particular Contact Role.

contactIDs

	A comma-delimited list of contactIDs to filter the result set by. This can be used in conjunction with other filters.

contactID

	Supplying a contact ID will perform an exact match search.

optionalID

	Supplying an Optional ID will perform an exact match search.

DOB

	yyyy-mm-dd formatted date

//
*/
func (s *ContactService) ContactSearch(parms map[string]string) ([]Contact, *Response, error) {
	var obj []Contact

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	resp, err := do(s.client, "GET", Params{parms: parms, u: "/contacts/search"}, obj)

	if err != nil {
		return obj, resp, err
	}

	var json = jsontime.ConfigWithCustomTimeFormat
	jsontime.AddTimeFormatAlias("axc_time", "15:04")
	jsontime.AddTimeFormatAlias("axc_time_long", "15:04:05")
	jsontime.AddTimeFormatAlias("axc_date", "2006-01-02")
	jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")

	err = json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}
