package axcelerate

import "encoding/json"

type Enrol struct {
	InvoiceID int `json:"INVOICEID"`
	ContactID int `json:"CONTACTID"`
	LearnerID int `json:"LEARNERID"`
	Amount    int `json:"AMOUNT"`
}

/*
CourseEnrol Enrols a Contact in an Activity Instance. This method sends notifications to the student, payer, and administrator. Enrolments into linked E-Learning will be scheduled and processed by the system later.

Request Parameters (passed as a map[string]string):
--------------------------------------------
Header                      Type        Required    Default     Description
contactID                   numeric     true                    The ID of the Contact you are enrolling.
instanceID                  numeric     true                    The ID of the Activity Instance you are enrolling the contact into.
type                        string      true                    The type of the activity. w = workshop, p = accredited program, el = e-learning.
tentative                   boolean     false                   Enrol the Contact as Tentative.
payerID                     numeric     false                   The ID of the Contact that is paying for the course. Defaults to the student (contactID) if omitted.
invoiceID                   numeric     false       0           The ID of the invoice created for this Contact (during payment) or the ID of the invoice to add this enrolment to. Defaults to 0 if omitted.
PONumber                    string      false                   The Purchase Order Number for a new invoice generated during this booking.
generateInvoice             boolean     false       true        Determines if a new invoice should be generated for this booking (only works if the course cost > $0).
lockInvoiceItems            boolean     false       true        If a new invoice is generated, determines if items are locked (required before payments can be applied).
archiveInvoice              boolean     false       false       Determines if the new invoice is completely locked and archived (prevents future modifications).
forceBooking                boolean     false       false       Forces the enrolment even if it is closed.
bookOnDefaultWorkshops      boolean     false       true        (type=p) Automatically book into workshops linked to the program’s units. Errors if multiple workshops are linked.
syncDatesWithWorkshop       boolean     false       true        (type=p) Sync unit start/end dates with linked workshops.
syncUOCdates                boolean     false       false       (type=p) Sync Unit of Competency dates with Unit of Study dates (only if syncDatesWithWorkshop is false).
syncWithClassSchedule       boolean     false       false       (type=p) Offset unit start/end dates based on commencement date and class schedule.
applyGST (DEPRECATED)       boolean     false       false       (type=p) Determines if cost includes GST (normally GST-free for accredited training).
GST_type                    numeric     false                   Specifies GST type: 0 = no GST, 1 = GST included.
autoGrantCT                 boolean     false       false       Automatically set subject outcome to CT for existing competencies in accredited bookings.
dateCommenced               date        false                   (type=p) The enrolment commencement date.
dateCompletionExpected      date        false                   (type=p) The expected completion date (ignored if using syncWithClassSchedule).
suppressNotifications       boolean     false       false       Suppresses all booking notification emails.
sendAdminNotification       boolean     false       false       Sends admin notification even if suppressNotifications is true.
blockAdminNotification      boolean     false       false       Explicitly blocks admin notifications.
useRegistrationFormDefaults boolean     false       false       (type=p) Applies default registration form values for accredited program bookings.
StudyReasonID               numeric     false                   AVETMISS Study Reason (for accredited enrolments).
FundingNational             numeric     false                   AVETMISS Funding Source - National (for accredited enrolments).
marketingAgentContactID     numeric     false                   The ID of the marketing agent associated with the enrolment.
serviceDate                 string      false                   The service date for the invoice item.
cost                        numeric     false                   Discounted cost for the enrolment.
commencingProgramCohortIdentifiers string false                Commencing program cohort ID (first 6 characters considered).
discountIDList              string      false                   A list of valid discountIDs for this booking (required if 'cost' is provided).
customField_[variableName]  string      false                   Custom field value for the enrolment (comma-delimited or JSON array for multiple values).
FundingState                string      false                   AVETMISS Funding Source - State (for accredited enrolments).
PSTACDateVIC                date        false                   (type=p) Program Supervised Teaching Activity Completion Date (for Victorian reporting).
commencedWhileAtSchool      boolean     false                   (type=p) Indicates if the enrolment commenced while at school (for Victorian reporting).

Returns:
--------------------------------------------
- Enrol: Struct containing the enrolment details.
- *Response: API response metadata.
- error: Error object if any issue occurs during the request.

Example Usage:
--------------------------------------------
params := map[string]string{
    "contactID":    "123",
    "instanceID":   "456",
    "type":         "w",
    "payerID":      "789",
    "generateInvoice": "true",
}

enrolment, resp, err := coursesService.CourseEnrol(params)
if err != nil {
    log.Fatalf("Error enrolling contact: %v", err)
}
fmt.Printf("Enrolment ID: %d\n", enrolment.InvoiceID)
*/

func (s *CoursesService) CourseEnrol(parms map[string]string) (Enrol, *Response, error) {
	var obj Enrol

	if len(parms) == 0 {
		parms = map[string]string{}
	}

	resp, err := do(s.client, "POST", Params{parms: parms, u: "/course/enrol"}, obj)

	if err != nil {
		return obj, resp, err
	}

	err = json.Unmarshal([]byte(resp.Body), &obj)

	return obj, resp, err
}
