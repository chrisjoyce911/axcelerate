package axcelerate

import (
	"bytes"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContactService_VerifyUSI(t *testing.T) {
	type fields struct {
		StatusCode int
		Body       string
	}
	type args struct {
		contactID int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    USIstatus
		want1   *Response
		wantErr bool
	}{
		{
			name: "already",
			fields: fields{
				StatusCode: 200,
				Body:       LoadTestData("contact_verifyUSI_already.json"),
			},
			args: args{contactID: 13910132},
			want: USIstatus{
				UsiVerified: true,
				Msg:         "USI is already present and verified. ",
			},

			want1: &Response{
				StatusCode:    200,
				Body:          LoadTestData("contact_verifyUSI_already.json"),
				ContentLength: 0,
			},
		},
		{
			name: "fail",
			fields: fields{
				StatusCode: 200,
				Body:       LoadTestData("contact_verifyUSI_fail.json"),
			},
			args: args{contactID: 13933281},
			want: USIstatus{
				UsiVerified: false,
				Msg:         "USI Valid but the student's personal data does not match. This must be fixed before their USI can be marked as verified. Please check: Surname. ",
				Data: USIdata{
					UsiStatus:   "Valid",
					FamilyName:  "NO_MATCH",
					FirstName:   "MATCH",
					DateOfBirth: "MATCH",
				},
			},

			want1: &Response{
				StatusCode:    200,
				Body:          LoadTestData("contact_verifyUSI_fail.json"),
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
			s := &ContactService{
				client: client,
			}

			got, got1, err := s.VerifyUSI(tt.args.contactID)

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
