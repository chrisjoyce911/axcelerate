package axcelerate

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContactService_GetContact(t *testing.T) {
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
		want    Contact
		want1   *Response
		wantErr bool
	}{
		// {
		// 	name: "Full Body",
		// 	fields: fields{
		// 		StatusCode: 200,
		// 		Body:       `{"CONTACTID": 708861,"GIVENNAME": "Joe","SURNAME": "Bloggs","EMAILADDRESS": "joe.bloggs@fakeemail.com.au","SEX": "M","DOB": "1986-05-02","USI": "FAKE054321","USI_VERIFIED": true,"USI_EXEMPTION": false,"TITLE": null,"MIDDLENAME": null,"PREFERREDNAME": "Joey","LUI": "1234567890","TFN_RECORDED": true,"OPTIONALID": null,"POSITION": "Web Guru","SECTION": null,"DIVISION": null,"ORGANISATION": "VM Learning","ADDRESS1": "123 Fake Street","ADDRESS2": null,"CITY": "BRISBANE","STATE": "QLD","POSTCODE": "4000","COUNTRY": "Australia","SADDRESS1": "PO Box 1234","SADDRESS2": null,"SCITY": "BRISBANE","SSTATE": "QLD","SPOSTCODE": "4000","SCOUNTRY": "Australia","PHONE": "(07) 9876 5432","MOBILEPHONE": null,"WORKPHONE": "(07) 3215 8888","FAX": null,"OTHERPHONE": null,"COMMENT": "This is an example contact for the API","WEBSITE": "www.joebloggs.com.au","CITIZENSTATUSID": 1,"CITIZENSTATUSNAME": "Australian Citizen","COUNTRYOFBIRTHID": 8104,"COUNTRYOFBIRTHNAME": "United States of America","CITYOFBIRTH": "New York","COUNTRYOFCITIZENID": 1101,"COUNTRYOFCITIZENNAME": "Australia","INDIGENOUSSTATUSID": 4,"INDIGENOUSSTATUSNAME": "No, Neither Aboriginal Nor Torres Strait Islander","MAINLANGUAGEID": 1201,"MAINLANGUAGENAME": "English","ENGLISHPROFICIENCYID": 1,"ENGLISHASSISTANCEFLAG": false,"HIGHESTSCHOOLLEVELID": 12,"HIGHESTSCHOOLLEVELYEAR": "2004","CURRENTSCHOOLLEVEL": null,"ATSCHOOLFLAG": false,"ATSCHOOLNAME": null,"PRIOREDUCATIONIDS": [410],"PRIOREDUCATIONNAMES": ["Advanced Diploma or Associate Degree Level"],"DISABILITYFLAG": true,"DISABILITYTYPEIDS": [17],"DISABILITYTYPENAMES": ["Vision"],"LABOURFORCEID": 1,"LABOURFORCENAME": "Full-time employee","EMERGENCYCONTACT": "Jane Bloggs","EMERGENCYCONTACTRELATION": "Mother","EMERGENCYCONTACTPHONE": "0433 123 123","ANZSCOCODE": null,"ANZSICCODE": null,"EMPLOYERCONTACTID": 4321,"PAYERCONTACTID": 1234,"SUPERVISORCONTACTID": null,"COACHCONTACTID": null,"AGENTCONTACTID": null,"SACESTUDENTID": "123456A","CUSTOMFIELD_MYFIELD": "Test Value","CUSTOMFIELD_MYMULTIFIELD": ["Test Value 1","Test Value 2",],"CATEGORYIDS": [267,516,648,5868],"DOMAINIDS": [123,456]}`,
		// 	},
		// 	args: args{contactID: 708861},
		// 	want: Contact{
		// 		ContactID:                708861,
		// 		CountryID:                0,
		// 		Address1:                 "123 Fake Street",
		// 		Citizenstatusid:          float64(1),
		// 		Citizenstatusname:        "Australian Citizen",
		// 		City:                     "BRISBANE",
		// 		CityOfBirth:              "New York",
		// 		Comment:                  "This is an example contact for the API",
		// 		Country:                  "Australia",
		// 		CountryofBirthID:         float64(8104),
		// 		Countryofbirthname:       "United States of America",
		// 		Countryofcitizenname:     "Australia",
		// 		Countryofcitizenid:       float64(1101),
		// 		Disabilityflag:           true,
		// 		Emailaddress:             "joe.bloggs@fakeemail.com.au",
		// 		Emergencycontact:         "Jane Bloggs",
		// 		Emergencycontactrelation: "Mother",
		// 		Employercontactid:        float64(4321),
		// 		Englishassistanceflag:    false,
		// 		Englishproficiencyid:     float64(1),
		// 		Indigenousstatusname:     "No, Neither Aboriginal Nor Torres Strait Islander",
		// 		Labourforcename:          "Full-time employee",
		// 		Givenname:                "Joe",
		// 		HighestschoollevelID:     float64(12),
		// 		Highestschoollevelyear:   "2004",
		// 	},
		// 	want1: &Response{
		// 		StatusCode:    200,
		// 		Body:          `{"CONTACTID": 708861,"GIVENNAME": "Joe","SURNAME": "Bloggs","EMAILADDRESS": "joe.bloggs@fakeemail.com.au","SEX": "M","DOB": "1986-05-02","USI": "FAKE054321","USI_VERIFIED": true,"USI_EXEMPTION": false,"TITLE": null,"MIDDLENAME": null,"PREFERREDNAME": "Joey","LUI": "1234567890","TFN_RECORDED": true,"OPTIONALID": null,"POSITION": "Web Guru","SECTION": null,"DIVISION": null,"ORGANISATION": "VM Learning","ADDRESS1": "123 Fake Street","ADDRESS2": null,"CITY": "BRISBANE","STATE": "QLD","POSTCODE": "4000","COUNTRY": "Australia","SADDRESS1": "PO Box 1234","SADDRESS2": null,"SCITY": "BRISBANE","SSTATE": "QLD","SPOSTCODE": "4000","SCOUNTRY": "Australia","PHONE": "(07) 9876 5432","MOBILEPHONE": null,"WORKPHONE": "(07) 3215 8888","FAX": null,"OTHERPHONE": null,"COMMENT": "This is an example contact for the API","WEBSITE": "www.joebloggs.com.au","CITIZENSTATUSID": 1,"CITIZENSTATUSNAME": "Australian Citizen","COUNTRYOFBIRTHID": 8104,"COUNTRYOFBIRTHNAME": "United States of America","CITYOFBIRTH": "New York","COUNTRYOFCITIZENID": 1101,"COUNTRYOFCITIZENNAME": "Australia","INDIGENOUSSTATUSID": 4,"INDIGENOUSSTATUSNAME": "No, Neither Aboriginal Nor Torres Strait Islander","MAINLANGUAGEID": 1201,"MAINLANGUAGENAME": "English","ENGLISHPROFICIENCYID": 1,"ENGLISHASSISTANCEFLAG": false,"HIGHESTSCHOOLLEVELID": 12,"HIGHESTSCHOOLLEVELYEAR": "2004","CURRENTSCHOOLLEVEL": null,"ATSCHOOLFLAG": false,"ATSCHOOLNAME": null,"PRIOREDUCATIONIDS": [410],"PRIOREDUCATIONNAMES": ["Advanced Diploma or Associate Degree Level"],"DISABILITYFLAG": true,"DISABILITYTYPEIDS": [17],"DISABILITYTYPENAMES": ["Vision"],"LABOURFORCEID": 1,"LABOURFORCENAME": "Full-time employee","EMERGENCYCONTACT": "Jane Bloggs","EMERGENCYCONTACTRELATION": "Mother","EMERGENCYCONTACTPHONE": "0433 123 123","ANZSCOCODE": null,"ANZSICCODE": null,"EMPLOYERCONTACTID": 4321,"PAYERCONTACTID": 1234,"SUPERVISORCONTACTID": null,"COACHCONTACTID": null,"AGENTCONTACTID": null,"SACESTUDENTID": "123456A","CUSTOMFIELD_MYFIELD": "Test Value","CUSTOMFIELD_MYMULTIFIELD": ["Test Value 1","Test Value 2",],"CATEGORYIDS": [267,516,648,5868],"DOMAINIDS": [123,456]}`,
		// 		ContentLength: 0,
		// 	},
		// },
		// {
		// 	name: "student has full AVETMISS 7.0",
		// 	fields: fields{
		// 		StatusCode: 200,
		// 		Body:       `{"CONTACTID": 708861,"GIVENNAME": "Joe","SURNAME": "Bloggs","EMAILADDRESS": "joe.bloggs@fakeemail.com.au","SEX": "M","DOB": "1986-05-02","USI": "FAKE054321","USI_VERIFIED": true,"USI_EXEMPTION": false,"TITLE": null,"MIDDLENAME": null,"PREFERREDNAME": "Joey","LUI": "1234567890","TFN_RECORDED": true,"OPTIONALID": null,"POSITION": "Web Guru","SECTION": null,"DIVISION": null,"ORGANISATION": "VM Learning","ADDRESS1": "123 Fake Street","ADDRESS2": null,"CITY": "BRISBANE","STATE": "QLD","POSTCODE": "4000","COUNTRY": "Australia","SADDRESS1": "PO Box 1234","SADDRESS2": null,"SCITY": "BRISBANE","SSTATE": "QLD","SPOSTCODE": "4000","SCOUNTRY": "Australia","PHONE": "(07) 9876 5432","MOBILEPHONE": null,"WORKPHONE": "(07) 3215 8888","FAX": null,"OTHERPHONE": null,"COMMENT": "This is an example contact for the API","WEBSITE": "www.joebloggs.com.au","CITIZENSTATUSID": 1,"CITIZENSTATUSNAME": "Australian Citizen","COUNTRYOFBIRTHID": 8104,"COUNTRYOFBIRTHNAME": "United States of America","CITYOFBIRTH": "New York","COUNTRYOFCITIZENID": 1101,"COUNTRYOFCITIZENNAME": "Australia","INDIGENOUSSTATUSID": 4,"INDIGENOUSSTATUSNAME": "No, Neither Aboriginal Nor Torres Strait Islander","MAINLANGUAGEID": 1201,"MAINLANGUAGENAME": "English","ENGLISHPROFICIENCYID": 1,"ENGLISHASSISTANCEFLAG": false,"HIGHESTSCHOOLLEVELID": 12,"HIGHESTSCHOOLLEVELYEAR": "2004","CURRENTSCHOOLLEVEL": null,"ATSCHOOLFLAG": false,"ATSCHOOLNAME": null,"PRIOREDUCATIONIDS": [410],"PRIOREDUCATIONNAMES": ["Advanced Diploma or Associate Degree Level"],"DISABILITYFLAG": true,"DISABILITYTYPEIDS": [17],"DISABILITYTYPENAMES": ["Vision"],"LABOURFORCEID": 1,"LABOURFORCENAME": "Full-time employee","EMERGENCYCONTACT": "Jane Bloggs","EMERGENCYCONTACTRELATION": "Mother","EMERGENCYCONTACTPHONE": "0433 123 123","ANZSCOCODE": null,"ANZSICCODE": null,"EMPLOYERCONTACTID": 4321,"PAYERCONTACTID": 1234,"SUPERVISORCONTACTID": null,"COACHCONTACTID": null,"AGENTCONTACTID": null,"SACESTUDENTID": "123456A","CUSTOMFIELD_MYFIELD": "Test Value","CUSTOMFIELD_MYMULTIFIELD": ["Test Value 1","Test Value 2",],"CATEGORYIDS": [267,516,648,5868],"DOMAINIDS": [123,456]}`,
		// 	},
		// 	args: args{contactID: 708861},
		// 	want: Contact{
		// 		Address1: "123 Fake Street",
		// 	},
		// 	want1: &Response{
		// 		StatusCode:    200,
		// 		Body:          `{"CONTACTID": 708861,"GIVENNAME": "Joe","SURNAME": "Bloggs","EMAILADDRESS": "joe.bloggs@fakeemail.com.au","SEX": "M","DOB": "1986-05-02","USI": "FAKE054321","USI_VERIFIED": true,"USI_EXEMPTION": false,"TITLE": null,"MIDDLENAME": null,"PREFERREDNAME": "Joey","LUI": "1234567890","TFN_RECORDED": true,"OPTIONALID": null,"POSITION": "Web Guru","SECTION": null,"DIVISION": null,"ORGANISATION": "VM Learning","ADDRESS1": "123 Fake Street","ADDRESS2": null,"CITY": "BRISBANE","STATE": "QLD","POSTCODE": "4000","COUNTRY": "Australia","SADDRESS1": "PO Box 1234","SADDRESS2": null,"SCITY": "BRISBANE","SSTATE": "QLD","SPOSTCODE": "4000","SCOUNTRY": "Australia","PHONE": "(07) 9876 5432","MOBILEPHONE": null,"WORKPHONE": "(07) 3215 8888","FAX": null,"OTHERPHONE": null,"COMMENT": "This is an example contact for the API","WEBSITE": "www.joebloggs.com.au","CITIZENSTATUSID": 1,"CITIZENSTATUSNAME": "Australian Citizen","COUNTRYOFBIRTHID": 8104,"COUNTRYOFBIRTHNAME": "United States of America","CITYOFBIRTH": "New York","COUNTRYOFCITIZENID": 1101,"COUNTRYOFCITIZENNAME": "Australia","INDIGENOUSSTATUSID": 4,"INDIGENOUSSTATUSNAME": "No, Neither Aboriginal Nor Torres Strait Islander","MAINLANGUAGEID": 1201,"MAINLANGUAGENAME": "English","ENGLISHPROFICIENCYID": 1,"ENGLISHASSISTANCEFLAG": false,"HIGHESTSCHOOLLEVELID": 12,"HIGHESTSCHOOLLEVELYEAR": "2004","CURRENTSCHOOLLEVEL": null,"ATSCHOOLFLAG": false,"ATSCHOOLNAME": null,"PRIOREDUCATIONIDS": [410],"PRIOREDUCATIONNAMES": ["Advanced Diploma or Associate Degree Level"],"DISABILITYFLAG": true,"DISABILITYTYPEIDS": [17],"DISABILITYTYPENAMES": ["Vision"],"LABOURFORCEID": 1,"LABOURFORCENAME": "Full-time employee","EMERGENCYCONTACT": "Jane Bloggs","EMERGENCYCONTACTRELATION": "Mother","EMERGENCYCONTACTPHONE": "0433 123 123","ANZSCOCODE": null,"ANZSICCODE": null,"EMPLOYERCONTACTID": 4321,"PAYERCONTACTID": 1234,"SUPERVISORCONTACTID": null,"COACHCONTACTID": null,"AGENTCONTACTID": null,"SACESTUDENTID": "123456A","CUSTOMFIELD_MYFIELD": "Test Value","CUSTOMFIELD_MYMULTIFIELD": ["Test Value 1","Test Value 2",],"CATEGORYIDS": [267,516,648,5868],"DOMAINIDS": [123,456]}`,
		// 		ContentLength: 0,
		// 	},
		// },
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

			client, _ := NewClient("", "", HttpClient(tclient))
			s := &ContactService{
				client: client,
			}
			got, got1, err := s.GetContact(tt.args.contactID)

			if err == nil {
				assert.Equal(t, tt.fields.StatusCode, got1.StatusCode, "HTTPStatus did not match")
				assert.Equal(t, tt.fields.Body, got1.Body, "Body did not match")
				assert.Equal(t, tt.want, got, "Body did not match")
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("ContactService.GetContact() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

		})
	}
}
