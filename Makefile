SHELL=/bin/bash
APP=circle-k
APP_COMMIT=$(shell git rev-parse HEAD)
ALL_PACKAGES=$(shell go list ./... | grep -v "vendor")
SOURCE_DIRS=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
COVERAGE_MIN=80

test:
	@echo "> running test and creating coverage report"
	go test -race -p=1 -cover -coverprofile=coverage.txt -covermode=atomic $(ALL_PACKAGES)
	@go tool cover -html=coverage.txt -o coverage.html
	@go tool cover -func=coverage.txt | grep -i total:
	@go tool cover -func=coverage.txt | gawk '/total:.*statements/ {if (strtonum($$3) < $(COVERAGE_MIN)) {print "ERR: coverage is lower than $(COVERAGE_MIN)"; exit 1}}'

