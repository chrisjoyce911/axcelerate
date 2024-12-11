package main

import (
	"testing"
	"time"

	"github.com/chrisjoyce911/axcelerate"
)

func TestCourseEnrol(t *testing.T) {
	// Initialize the mock client
	mockClient := axcelerate.NewMockClient()

	// Configure the mock CourseEnrol method
	mockClient.Courses.MockCourseEnrol = func(parms map[string]string) (*axcelerate.Enrol, *axcelerate.Response, error) {
		return &axcelerate.Enrol{
			InvoiceID: 12345,
			ContactID: 67890,
			LearnerID: 54321,
			Amount:    50,
		}, nil, nil
	}

	// Use the mock Courses service
	params := map[string]string{
		"contactID":       "67890",
		"instanceID":      "54321",
		"type":            "w",
		"generateInvoice": "true",
	}

	enrolment, _, err := mockClient.Courses.CourseEnrol(params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	// if enrolment.LearnerID != 12345 {
	// 	t.Errorf("expected enrolment ID 12345, got %d", enrolment.EnrolmentID)
	// }

	if enrolment.ContactID != 67890 {
		t.Errorf("expected contact ID 67890, got %d", enrolment.ContactID)
	}

	// if enrolment.InstanceID != 54321 {
	// 	t.Errorf("expected instance ID 54321, got %d", enrolment.InstanceID)
	// }
}

func TestCourseEnrolmentUpdate(t *testing.T) {
	// Initialize the mock client
	mockClient := axcelerate.NewMockClient()

	// Configure the mock CourseEnrolmentUpdate method
	mockClient.Courses.MockCourseEnrolmentUpdate = func(contactID, instanceID int, activityType string, parms map[string]string) (*axcelerate.EnrolmentUpdate, *axcelerate.Response, error) {
		return &axcelerate.EnrolmentUpdate{
			Data:     "Updated Successfully",
			Error:    false,
			Messages: "Update completed",
			Code:     "200",
			Details:  "Details about the update",
		}, nil, nil
	}

	// Use the mock Courses service
	params := map[string]string{"additionalParam": "value"}
	contactID := 12345
	instanceID := 67890
	activityType := "w"

	enrolmentUpdate, _, err := mockClient.Courses.CourseEnrolmentUpdate(contactID, instanceID, activityType, params)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if enrolmentUpdate.Data != "Updated Successfully" {
		t.Errorf("expected Data 'Updated Successfully', got %s", enrolmentUpdate.Data)
	}

	if enrolmentUpdate.Code != "200" {
		t.Errorf("expected Code '200', got %s", enrolmentUpdate.Code)
	}
}

func TestGetCoursesInstanceDetail(t *testing.T) {
	// Initialize the mock client
	mockClient := axcelerate.NewMockClient()

	// Configure the mock GetCoursesInstanceDetail method
	mockClient.Courses.MockGetCoursesInstanceDetail = func(instanceID int, activityType string) (axcelerate.InstanceDetail, *axcelerate.Response, error) {
		return axcelerate.InstanceDetail{
			InstanceID:         instanceID,
			Name:               "Mock Activity",
			StartDate:          time.Now(),
			FinishDate:         time.Now().Add(24 * time.Hour),
			Participants:       15,
			ParticipantVacancy: 10,
			Status:             "Active",
			Location:           "Mock Location",
		}, nil, nil
	}

	// Use the mock Courses service
	instanceID := 123
	activityType := "w"

	detail, _, err := mockClient.Courses.GetCoursesInstanceDetail(instanceID, activityType)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if detail.InstanceID != instanceID {
		t.Errorf("expected InstanceID %d, got %d", instanceID, detail.InstanceID)
	}

	if detail.Name != "Mock Activity" {
		t.Errorf("expected Name 'Mock Activity', got %s", detail.Name)
	}

	if detail.Participants != 15 {
		t.Errorf("expected Participants 15, got %d", detail.Participants)
	}

	if detail.Status != "Active" {
		t.Errorf("expected Status 'Active', got %s", detail.Status)
	}
}
