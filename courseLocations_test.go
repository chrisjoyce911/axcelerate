package axcelerate

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCoursesService_GetCoursesLocations(t *testing.T) {
	type fields struct {
		StatusCode int
		Body       string
	}
	type args struct {
		public     bool
		onlyFuture bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []Location
		want1   *Response
		wantErr bool
	}{
		{
			name: "Full Body",
			fields: fields{
				StatusCode: 200,
				Body:       `["108 Fairley St","Adelaide","Armidale"]`,
			},
			args: args{public: true,
				onlyFuture: true},
			want: []Location{"108 Fairley St", "Adelaide", "Armidale"},
			want1: &Response{
				StatusCode:    200,
				Body:          `["108 Fairley St","Adelaide","Armidale"]`,
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

			got, got1, err := s.GetCoursesLocations(tt.args.public, tt.args.onlyFuture)

			if err == nil {
				assert.Equal(t, tt.fields.StatusCode, got1.StatusCode, "HTTPStatus did not match")
				assert.Equal(t, tt.fields.Body, got1.Body, "Body did not match")
				assert.Equal(t, tt.want, got, "Response")
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CoursesService.GetCoursesLocations() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("CoursesService.GetCoursesLocations() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
