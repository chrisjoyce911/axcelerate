package axcelerate

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoursesService_UpdateInstanceCost(t *testing.T) {
	type fields struct {
		StatusCode int
		Body       string
	}
	type args struct {
		instanceID   int
		activityType string
		cost         int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    UpdateInstanceDetail
		want1   *Response
		wantErr bool
	}{
		{
			name: "OK",
			fields: fields{
				StatusCode: 200,
				Body:       LoadTestData("course_instance_update.json"),
			},
			args: args{instanceID: 422662, activityType: "w", cost: 123},
			want: UpdateInstanceDetail{
				Status:     "success",
				Message:    "Workshop Instance Updated.",
				InstanceID: 1901499,
			},

			want1: &Response{
				StatusCode:    200,
				Body:          LoadTestData("course_instance_update.json"),
				ContentLength: 0,
			},
		},
		{
			name: "Error",
			fields: fields{
				StatusCode: 200,
				Body:       LoadTestData("course_instance_update_error.json"),
			},
			args: args{instanceID: 422662, activityType: "w"},
			want: UpdateInstanceDetail{
				Data:     "",
				Error:    true,
				Code:     "0",
				Messages: "key [TYPE] doesn't exist",
			},

			want1: &Response{
				StatusCode:    200,
				Body:          LoadTestData("course_instance_update_error.json"),
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

			got, got1, err := s.UpdateInstanceCost(tt.args.instanceID, tt.args.activityType, tt.args.cost)

			if err == nil {
				assert.Equal(t, tt.fields.StatusCode, got1.StatusCode, "HTTPStatus did not match")
				assert.Equal(t, tt.fields.Body, got1.Body, "Body did not match")
				assert.Equal(t, tt.want, got, "Response")
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("CoursesService.UpdateInstanceCost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CoursesService.UpdateInstanceCost() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CoursesService.UpdateInstanceCost() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
