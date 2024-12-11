package main

import (
	"testing"

	"github.com/chrisjoyce911/axcelerate"
)

func TestTemplateEmail(t *testing.T) {
	// Initialize the mock service
	mockService := &axcelerate.MockTemplateService{
		MockTemplateEmail: func(params axcelerate.TemplateEmailParams) (*axcelerate.EmailResponse, *axcelerate.Response, error) {
			return &axcelerate.EmailResponse{
				FailedCount:    0,
				AttemptedCount: 1,
				SuccessCount:   1,
				Message:        "Email sent successfully.",
			}, &axcelerate.Response{Status: "200 OK"}, nil
		},
	}

	// Define test inputs
	params := axcelerate.TemplateEmailParams{
		ContactID: 123,
		PlanID:    456,
		Subject:   "Test Email Subject",
	}

	// Execute the method
	emailResponse, resp, err := mockService.TemplateEmail(params)

	// Validate the result
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Status != "200 OK" {
		t.Errorf("expected status '200 OK', got '%s'", resp.Status)
	}

	if emailResponse.Message != "Email sent successfully." {
		t.Errorf("expected Message 'Email sent successfully.', got '%s'", emailResponse.Message)
	}

	if emailResponse.AttemptedCount != 1 {
		t.Errorf("expected AttemptedCount 1, got %d", emailResponse.AttemptedCount)
	}

	if emailResponse.SuccessCount != 1 {
		t.Errorf("expected SuccessCount 1, got %d", emailResponse.SuccessCount)
	}

	if emailResponse.FailedCount != 0 {
		t.Errorf("expected FailedCount 0, got %d", emailResponse.FailedCount)
	}
}
