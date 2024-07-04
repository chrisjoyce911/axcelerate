package axcelerate

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestContactService_ContactEnrolments(t *testing.T) {

	t.Skip("skipping testing need new data examples")

	type fields struct {
		StatusCode int
		Body       string
	}
	type args struct {
		contactID int
		parms     map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []ContactEnrolment
		want1   *Response
		wantErr bool
	}{
		{
			name: "single",
			fields: fields{
				StatusCode: 200,
				Body:       LoadTestData("contact_enrolments_13928472.json"),
			},
			args: args{contactID: 13910132},
			want: []ContactEnrolment{
				{
					RowID:          1,
					Type:           "w",
					ID:             92956,
					InstanceID:     1856676,
					EnrolID:        10219553,
					VicenrolmentID: "",
					InvoiceID:      3042841,
					InvoicePaid:    true,
					LearnerID:      10219553,
					Code:           "PFA-011",
					Location:       "MITTAGONG",
					Delivery:       "Face-to-Face",
					DeliveryMode:   "",
					Activitytype:   "Provide First Aid including CPR - HLTAID011",
					Name:           "Provide First Aid (AM)",
					CommencedDate:  time.Date(1900, time.January, 1, 0, 0, 0, 0, time.Local),
					StartDate:      time.Date(2022, time.April, 12, 0, 0, 0, 0, time.Local),
					FinishDate:     time.Date(2022, time.April, 12, 0, 0, 0, 0, time.Local),
					CompletionDate: time.Date(2022, time.April, 12, 0, 0, 0, 0, time.Local),
					Mandatory:      false,
					Status:         "Completed",
					Count:          1,
					// Complexdates: []EnrolmentComplexdates{
					// 	{
					// 		Date:             time.Date(2022, time.April, 12, 0, 0, 0, 0, time.Local),
					// 		StartTime:        time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					// 		EndTime:          time.Date(1, time.January, 1, 0, 0, 0, 0, time.UTC),
					// 		TrainerContactID: 0,
					// 		Location:         "",
					// 		RoomID:           0,
					// 	},
					// },
				},
			},

			want1: &Response{
				StatusCode:    200,
				Body:          LoadTestData("contact_enrolments_13928472.json"),
				ContentLength: 0,
			},
		},

		{
			name: "paid",
			fields: fields{
				StatusCode: 200,
				Body:       LoadTestData("contact_enrolments_paid.json"),
			},
			args: args{contactID: 13910132},
			want: []ContactEnrolment{
				{
					RowID:          1,
					Type:           "w",
					ID:             92956,
					InstanceID:     1856676,
					EnrolID:        10219553,
					VicenrolmentID: "",
					InvoiceID:      3042841,
					InvoicePaid:    true,
				},
			},

			want1: &Response{
				StatusCode:    200,
				Body:          LoadTestData("contact_enrolments_paid.json"),
				ContentLength: 0,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tclient := NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: tt.fields.StatusCode,
					Body:       ioutil.NopCloser(bytes.NewBufferString(tt.fields.Body)),
					Header:     make(http.Header),
				}
			})

			s := &ContactService{
				client: NewClient("", "", nil, tclient),
			}

			got, got1, err := s.ContactEnrolments(tt.args.contactID, tt.args.parms)

			if err == nil {
				assert.Equal(t, tt.fields.StatusCode, got1.StatusCode, "HTTPStatus did not match")
				assert.Equal(t, tt.fields.Body, got1.Body, "Body did not match")
				assert.Equal(t, tt.want, got, "Response")
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("ContactService.VerifyUSI() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ContactService.VerifyUSI() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("ContactService.VerifyUSI() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
