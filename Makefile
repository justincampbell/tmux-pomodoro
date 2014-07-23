COVERAGE_FILE = coverage.out

test: acceptance

coverage: unit
	go tool cover -html=$(COVERAGE_FILE)

acceptance: build
	bats test

build: unit
	go build -o bin/pomodoro

unit:
	go test -coverprofile=$(COVERAGE_FILE) -timeout 25ms

.PHONY: test coverage acceptance build unit
