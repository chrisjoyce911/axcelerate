package axcelerate

type MockCoursesService struct {
	MockCourseEnrol              func(parms map[string]string) (*Enrol, *Response, error)
	MockCourseEnrolmentUpdate    func(contactID, instanceID int, activityType string, parms map[string]string) (*EnrolmentUpdate, *Response, error)
	MockGetCoursesInstanceDetail func(instanceID int, activityType string) (InstanceDetail, *Response, error)
}

func (m *MockCoursesService) CourseEnrol(parms map[string]string) (*Enrol, *Response, error) {
	if m.MockCourseEnrol != nil {
		return m.MockCourseEnrol(parms)
	}
	return nil, nil, nil
}

func (m *MockCoursesService) CourseEnrolmentUpdate(contactID, instanceID int, activityType string, parms map[string]string) (*EnrolmentUpdate, *Response, error) {
	if m.MockCourseEnrolmentUpdate != nil {
		return m.MockCourseEnrolmentUpdate(contactID, instanceID, activityType, parms)
	}
	return nil, nil, nil
}

func (m *MockCoursesService) GetCoursesInstanceDetail(instanceID int, activityType string) (InstanceDetail, *Response, error) {
	if m.MockGetCoursesInstanceDetail != nil {
		return m.MockGetCoursesInstanceDetail(instanceID, activityType)
	}
	return InstanceDetail{}, nil, nil
}
