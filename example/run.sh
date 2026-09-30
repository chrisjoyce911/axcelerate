#!/bin/bash
# Run an example by name. `go run ./example` builds the example package
# (main.go plus the files package it imports) without the test files.
set -e
cd "$(dirname "$0")"
go run ./example "$@"
