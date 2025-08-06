package main

import (
	"fmt"
	"os"

	"github.com/chrisjoyce911/axcelerate"
	"github.com/joho/godotenv"
)

// EmailResponse struct for handling email responses
type EmailResponse struct {
	FailedCount    int      `json:"FAILEDCOUNT"`
	Message        string   `json:"MESSAGE"`
	Errors         []string `json:"ERRORS,omitempty"`
	AttemptedCount int      `json:"ATTEMPTEDCOUNT"`
	SuccessCount   int      `json:"SUCCESSCOUNT"`
}

var client *axcelerate.Client

func main() {
	_ = godotenv.Load()

	var apitoken string = os.Getenv("AXCELERATE_APITOKEN")
	var wstoken string = os.Getenv("AXCELERATE_WSTOKEN")
	var baseURL string = os.Getenv("AXCELERATE_BASEURL")

	fmt.Println("API Token:", apitoken)
	fmt.Println("WS Token:", wstoken)
	fmt.Println("Base URL:", baseURL)

	client, _ = axcelerate.NewClient(apitoken, wstoken, axcelerate.RateLimit(10), axcelerate.BaseURL(baseURL))

	// Check for command line argument to run specific example
	if len(os.Args) > 1 {
		example := os.Args[1]
		runExample(example)
		return
	}

	// Default: run contactNoteAddExample if no argument provided
	fmt.Println("No example specified. Running default contactNoteAddExample...")
	fmt.Println("Use: go run main.go <example_name> to run specific examples")
	fmt.Println("Run: go run main.go help for list of available examples")
	contactNoteAddExample()
}

// runExample runs the specified example function
func runExample(example string) {
	switch example {
	case "help":
		showHelp()

	// Contact Examples
	case "contactNoteAdd":
		contactNoteAddExample()
	case "contactSearch":
		contactSearch()
	case "contactEnrolments":
		contactEnrolments(14446094) // Default contact ID
	case "findME":
		findME(client)
	case "findMEandVerifyUSI":
		findMEandVerifyUSI(client)

	// Course Examples
	case "courseEnrolmentStatus":
		courseEnrolmentStatus()
	case "courseEnrolments":
		courseEnrolments(10148651) // Default contact ID
	case "courseEnrolment":
		courseEnrolment()
	case "getCoursesInstanceDetail":
		getCoursesInstanceDetail()
	case "getCoursesInstanceSearch":
		getCoursesInstanceSearch()
	case "updateInstanceMaxParticipants":
		updateInstanceMaxParticipants()
	case "updateFinCode":
		updateFinCode()

	// Accounting Examples
	case "transact":
		transact()
	case "invoiceVoid":
		invoiceVoid()
	case "paymentVerify":
		paymentVerify()
	case "getInvoices":
		getInvoices()

	// Template & Report Examples
	case "templateEmail":
		templateEmail()
	case "savedReport":
		savedReport()
	case "savedReportList":
		savedReportList()

	// Venue & Media Examples
	case "getVenueDetail":
		getVenueDetail()
	case "contactCertificate":
		contactCertificate()

	default:
		fmt.Printf("Unknown example: %s\n", example)
		fmt.Println("Run: go run main.go help for list of available examples")
	}
}

// showHelp displays all available examples
func showHelp() {
	fmt.Println("Available Examples:")
	fmt.Println()
	fmt.Println("Contact Examples:")
	fmt.Println("  contactNoteAdd           - Add a note to a contact")
	fmt.Println("  contactSearch            - Search for contacts by email")
	fmt.Println("  contactEnrolments        - Get contact enrollments")
	fmt.Println("  findME                   - Basic contact search")
	fmt.Println("  findMEandVerifyUSI       - Contact search with USI verification")
	fmt.Println()
	fmt.Println("Course Examples:")
	fmt.Println("  courseEnrolmentStatus    - Update course enrollment status")
	fmt.Println("  courseEnrolments         - Get course enrollments")
	fmt.Println("  courseEnrolment          - Update course enrollment with custom fields")
	fmt.Println("  getCoursesInstanceDetail - Get course instance details")
	fmt.Println("  getCoursesInstanceSearch - Search course instances")
	fmt.Println("  updateInstanceMaxParticipants - Update max participants")
	fmt.Println("  updateFinCode            - Bulk update financial codes")
	fmt.Println()
	fmt.Println("Accounting Examples:")
	fmt.Println("  transact                 - Create transactions")
	fmt.Println("  invoiceVoid              - Void invoices")
	fmt.Println("  paymentVerify            - Verify payments")
	fmt.Println("  getInvoices              - Get invoices for a contact")
	fmt.Println()
	fmt.Println("Template & Report Examples:")
	fmt.Println("  templateEmail            - Send template emails")
	fmt.Println("  savedReport              - Run saved reports")
	fmt.Println("  savedReportList          - Get list of saved reports")
	fmt.Println()
	fmt.Println("Venue & Media Examples:")
	fmt.Println("  getVenueDetail           - Get venue details")
	fmt.Println("  contactCertificate       - Get contact certificates")
	fmt.Println()
	fmt.Println("Usage: go run main.go <example_name>")
	fmt.Println("Example: go run main.go contactNoteAdd")
}
