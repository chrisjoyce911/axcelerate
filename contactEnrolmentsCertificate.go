package axcelerate

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"
)

type Certificate struct {
	Certificate string `json:"CERTIFICATE"`
}

type Media struct {
	FileName    string `db:"file_name" json:"file_name,omitempty"`
	FileType    string `db:"file_type" json:"file_type,omitempty"`
	ContentType string `db:"content_type" json:"content_type,omitempty"`
	Data        []byte `db:"file_data" json:"file_data,omitempty"`
	Size        int    `db:"file_size" json:"file_size,omitempty"`
}

// ContactEnrolmentsCertificate Returns a Certificate
func (s *ContactService) ContactEnrolmentsCertificate(enrolID int) (Media, *Response, error) {

	var obj Certificate

	parms := map[string]string{"enrolID": fmt.Sprintf("%d", enrolID)}
	url := "/contact/enrolment/certificate"

	resp, err := do(s.client, "GET", Params{parms: parms, u: url}, obj)

	if err != nil {
		return Media{}, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)
	if err != nil {
		return Media{}, resp, err
	}

	// Decode the base64 encoded certificate
	pdfData, err := base64.StdEncoding.DecodeString(obj.Certificate)
	if err != nil {
		return Media{}, resp, err
	}

	fileName := fmt.Sprintf("%d-Certificate.pdf", enrolID)

	media := createMedia(pdfData, fileName)

	// Output the media struct
	println("FileName:", media.FileName)
	println("FileType:", media.FileType)
	println("ContentType:", media.ContentType)
	println("Data size:", len(media.Data))
	println("Size:", media.Size)

	return media, resp, err
}

func createMedia(fileData []byte, fileName string) Media {
	return Media{
		FileName:    fileName,
		FileType:    filepath.Ext(fileName),
		ContentType: getContentType(fileName),
		Data:        fileData,
		Size:        len(fileData),
	}
}

func getContentType(fileName string) string {
	switch strings.ToLower(filepath.Ext(fileName)) {
	case ".pdf":
		return "application/pdf"
	default:
		return "application/octet-stream"
	}
}
