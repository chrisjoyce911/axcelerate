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

func TestCoursesService_GetCoursesInstanceDetail(t *testing.T) {
	type fields struct {
		StatusCode int
		Body       string
	}
	type args struct {
		instanceID   int
		activityType string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    InstanceDetail
		want1   *Response
		wantErr bool
	}{
		{
			name: "Full Body",
			fields: fields{
				StatusCode: 200,
				Body:       LoadTestData("course_instance_detail.json"),
			},
			args: args{instanceID: 422662, activityType: "p"},
			want: InstanceDetail{
				DateDescriptor:  "22/05/2014 - 22/05/2014",
				EnrolmentOpen:   false,
				CourseID:        4109,
				Location:        "VM Learning Offices",
				MinParticipants: 0,
				Name:            "AAAAA test",
				OwnerContactID:  1100635,
				InstanceID:      67569,
				FinishDate:      time.Date(2014, 05, 22, 0, 0, 0, 0, time.Local),
				StartDate:       time.Date(2014, 05, 22, 0, 0, 0, 0, time.Local),
			},

			want1: &Response{
				StatusCode:    200,
				Body:          LoadTestData("course_instance_detail.json"),
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

			s := &CoursesService{
				client: NewClient("", "", nil, tclient),
			}

			got, got1, err := s.GetCoursesInstanceDetail(tt.args.instanceID, tt.args.activityType)

			if err == nil {
				assert.Equal(t, tt.fields.StatusCode, got1.StatusCode, "HTTPStatus did not match")
				assert.Equal(t, tt.fields.Body, got1.Body, "Body did not match")
				assert.Equal(t, tt.want, got, "Response")
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("CoursesService.GetCoursesInstanceDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CoursesService.GetCoursesInstanceDetail() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CoursesService.GetCoursesInstanceDetail() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
