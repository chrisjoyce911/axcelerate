# Axcelerate Go SDK Examples

This directory contains organized examples for using the Axcelerate Go SDK. The examples have been refactored into separate files based on functionality.

## File Structure

### Main Entry Point

- `main.go` - Main application entry point with example function calls

### Example Categories

#### Contact Examples (`contactExamples.go`)

- `contactNoteAddExample()` - Demonstrates adding notes to contacts
- `contactSearch()` - Search for contacts by email
- `contactEnrolments(contactID)` - Get contact enrollments
- `contactCertificate()` - Verify contact certificates
- `findME(client)` - Basic contact search
- `findMEandVerifyUSI(client)` - Contact search with USI verification

#### Course Examples (`courseExamples.go`)

- `courseEnrolmentStatus()` - Update course enrollment status
- `courseEnrolments(contactID)` - Get course enrollments
- `courseEnrolment()` - Update course enrollment with custom fields
- `getCoursesInstanceDetail()` - Get course instance details
- `getCoursesInstanceSearch()` - Search course instances
- `updateInstanceMaxParticipants()` - Update max participants for workshops
- `updateFinCode()` - Bulk update financial codes

#### Accounting Examples (`accountingExamples.go`)

- `transact()` - Create transactions
- `invoiceVoid()` - Void invoices
- `paymentVerify()` - Verify payments
- `getInvoices()` - Get invoices for a contact

#### Template & Report Examples (`templateReportExamples.go`)

- `templateEmail()` - Send template emails
- `savedReport()` - Run saved reports
- `savedReportList()` - Get list of saved reports

#### Venue & Media Examples (`venueMediaExamples.go`)

- `getVenueDetail()` - Get venue details
- `contactCertificate()` - Get contact certificates (media)
- `saveMediaToDisk()` - Helper function to save media files

## Usage

1. Set up your environment variables in a `.env` file:

   ```
   AXCELERATE_APITOKEN=your_api_token
   AXCELERATE_WSTOKEN=your_ws_token
   AXCELERATE_BASEURL=your_base_url
   ```

2. Run specific examples using one of these methods:

   **Option A: Using the convenience script (recommended):**

   ```bash
   ./run.sh <example_name>
   ./run.sh help
   ```

   **Option B: Using go run with all files:**

   ```bash
   go run main.go contactExamples.go courseExamples.go accountingExamples.go templateReportExamples.go venueMediaExamples.go <example_name>
   ```

   **Option C: Build and run:**

   ```bash
   go build -o examples .
   ./examples <example_name>
   ```

3. Get a list of all available examples:

   ```bash
   ./run.sh help
   ```

4. Run the default example (contactNoteAdd):
   ```bash
   ./run.sh
   ```

### Example Commands

```bash
# Contact examples
./run.sh contactNoteAdd
./run.sh contactSearch
./run.sh contactEnrolments

# Course examples
./run.sh courseEnrolmentStatus
./run.sh getCoursesInstanceDetail

# Accounting examples
./run.sh transact
./run.sh paymentVerify

# Template & Report examples
./run.sh templateEmail
./run.sh savedReport

# Venue & Media examples
./run.sh getVenueDetail
./run.sh contactCertificate
```

### Why not `go run *.go`?

The `go run *.go` command includes test files (`*_test.go`) which cannot be run directly with `go run`. The convenience script `run.sh` solves this by only including the necessary source files.

## Note Add Example

The `contactNoteAddExample()` function demonstrates how to use the new `NoteAdd` function:

```go
contactID := 11300044
params := map[string]string{
    "contactNote": "test note by example",
    "noteTypeID":  "27938",
}

response, httpResp, err := client.Contact.NoteAdd(contactID, params)
```

This example includes comprehensive error handling for both Go errors and API errors returned in the response.

## File Organization Benefits

- **Modularity**: Each category of examples is in its own file
- **Maintainability**: Easier to find and update specific examples
- **Clarity**: Main.go is clean and shows all available examples
- **Extensibility**: Easy to add new examples to appropriate categories
