test: acceptance

acceptance: build
	bats test

build: unit
	go build -o bin/pomodoro

unit:
	go test

.PHONY: test acceptance build unit
