COVERAGE_FILE = coverage.out
PREFIX=/usr/local

test: acceptance

install: build
	mkdir -p $(PREFIX)/bin
	cp -v bin/pomodoro $(PREFIX)/bin/pomodoro

uninstall:
	rm -vf $(PREFIX)/bin/pomodoro

coverage: unit
	go tool cover -html=$(COVERAGE_FILE)

acceptance: build
	bats test

build: unit
	go build -o bin/pomodoro

unit:
	go test -coverprofile=$(COVERAGE_FILE) -timeout 25ms

.PHONY: test coverage acceptance build unit
