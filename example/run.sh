#!/bin/bash
# Simple script to run examples without including test files

cd "$(dirname "$0")"

go run main.go contactExamples.go courseExamples.go accountingExamples.go templateReportExamples.go venueMediaExamples.go trainerExamples.go"$@"
