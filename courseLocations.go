package axcelerate

import (
	"encoding/json"
	"fmt"
)

// Location for a workshop
type Location string

/*
GetCoursesLocations The locations are pulled from aXcelerate to match any and all locations used for workshops. Used to pass to the calendar service as a filter.
The list is returned as an alphabetical array of location strings and is different for every client.

Parameters

public
	(bool) Only show Public locations (any location attached to a public activity)
onlyFuture
	(bool) Only show locations with a future activity related to it.
*/
func (s *CoursesService) GetCoursesLocations(public, onlyFuture bool) ([]Location, *Response, error) {
	var obj []Location

	parms := map[string]string{"public": fmt.Sprintf("%t", public), "onlyFuture": fmt.Sprintf("%t", onlyFuture)}

	resp, err := do(s.client, "GET", Params{parms: parms, u: "/course/locations"}, obj)

	if err != nil {
		return obj, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	return obj, resp, err
}
