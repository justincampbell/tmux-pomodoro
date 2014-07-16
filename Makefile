test: build
	bats test

build:
	go build -o bin/pomodoro

.PHONY: test build
