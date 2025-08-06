package main

import (
	"fmt"

	"github.com/chrisjoyce911/axcelerate"
)

// templateEmail demonstrates sending template emails
func templateEmail() {
	p := axcelerate.TemplateEmailParams{
		PlanID:                  16247,
		ContactID:               11300044,
		InstanceID:              1977505,
		InvoiceID:               3378756,
		Subject:                 "Booking Confirmation - Australia Wide First Aid",
		Type:                    "w",
		InvoiceAttachmentPlanID: 3440,
	}

	eUpdate, reps, err := client.Template.TemplateEmail(p)

	if err != nil {
		fmt.Printf("Body: %s", reps.Body)
		fmt.Print(err.Error())
		return
	}

	fmt.Printf("eUpdate%+v", eUpdate)
}

// savedReport demonstrates running a saved report
func savedReport() {
	offsetRows := 0

	parms := map[string]string{}
	parms["offsetRows"] = fmt.Sprintf("%d", offsetRows)
	parms["filterOverride"] = ` [
 {
     "VALUE2": "0",
     "OPERATOR": "BETWEEN N Days",
     "DISPLAY": "Workshop Start Date",
     "NAME": "workshops.pstartdate",
     "VALUE": "0"
 }]`

	savedReport, _, err := client.Report.SavedReportRun(85957, parms)

	if err != nil {
		fmt.Print(err)
	}

	fmt.Print(savedReport.Data)
}

// savedReportList demonstrates getting a list of saved reports
func savedReportList() {
	cert, _, err := client.Report.SavedReportList()

	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Printf("%+v", cert)
}
