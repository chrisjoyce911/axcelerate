package axcelerate

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"

	jsontime "github.com/liamylian/jsontime/v2/v2"
	"github.com/stretchr/testify/assert"
)

func TestCoursesService_GetEnrolments(t *testing.T) {

	t.Skip("skipping testing need new data examples")

	type fields struct {
		StatusCode int
		Body       string
	}
	type args struct {
		parms map[string]string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Enrolment
		want1   *Response
		wantErr bool
	}{
		{
			name: "Single",
			fields: fields{
				StatusCode: 200,
				Body:       LoadTestData("course_enrolments.json"),
			},
			args: args{
				parms: map[string]string{"instanceID": "422662", "contactID": "1111", "type": "w"},
			},
			want: []Enrolment{},

			want1: &Response{
				StatusCode:    200,
				Body:          LoadTestData("course_enrolments.json"),
				ContentLength: 0,
			},
		},
		{
			name: "already",
			fields: fields{
				StatusCode: 200,
				Body:       LoadTestData("course_enrolments_multi.json"),
			},
			args: args{
				parms: map[string]string{"instanceID": "422662", "contactID": "1111", "type": "w"},
			},
			want: []Enrolment{},

			want1: &Response{
				StatusCode:    200,
				Body:          LoadTestData("course_enrolments_multi.json"),
				ContentLength: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			tclient := NewTestClient(func(req *http.Request) *http.Response {
				return &http.Response{
					StatusCode: tt.fields.StatusCode,
					Body:       io.NopCloser(bytes.NewBufferString(tt.fields.Body)),
					Header:     make(http.Header),
				}
			})

			client, _ := NewClient("", "", HttpClient(tclient))
			s := &CoursesService{
				client: client,
			}

			got, got1, err := s.GetEnrolments(tt.args.parms)
			if (err != nil) != tt.wantErr {
				t.Errorf("CoursesService.GetEnrolments() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var obj []Enrolment
			var json = jsontime.ConfigWithCustomTimeFormat
			jsontime.AddTimeFormatAlias("axc_date_hours", "2006-01-02 15:04")

			_ = json.Unmarshal([]byte(got1.Body), &obj)

			if !reflect.DeepEqual(got, obj) {
				t.Errorf("CoursesService.GetEnrolments() got = %v, want %v", got, obj)
			}

			assert.Equal(t, "face-to-face", got[0].Delivery, "Response")

			// assert.Equal(t, 10, got[0].Delivery.Code, "Response")
			// assert.Equal(t, "face-to-face", got[0].Delivery.Description, "Response")

			// if !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("CoursesService.GetEnrolments() got = %v, want %v", got, tt.want)
			// }

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CoursesService.GetEnrolments() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
