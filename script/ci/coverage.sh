#!/bin/sh
set -e

# Go get our dependencies
go get github.com/pierrre/gotestcover

# Run our Test Coverage Report
gotestcover -coverprofile=cover.out github.com/dukeluke16/sample-golang-webservice/...

# Prettify Our Test Coverage Report
go tool cover -html=cover.out -o=cover.html
