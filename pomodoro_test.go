package main

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

var emptyArgs []string

var aTime, _ = time.Parse(time.Kitchen, "10:00AM")
var emptyState = State{
	endTime: noTime,
	now:     aTime,
}

func Test_main(t *testing.T) {
	os.Setenv("HOME", os.TempDir())

	main()
}

func Test_parseCommand_start(t *testing.T) {
	newState, output := parseCommand(emptyState, "start")

	expected, _ := time.Parse(time.Kitchen, "10:25AM")

	assert.T(t, expected.Equal(newState.endTime))
	assert.Equal(t, "Timer started, 25 minutes remaining", output.text)
}

func Test_parseCommand_status(t *testing.T) {
	endTime, _ := time.Parse(time.Kitchen, "10:25AM")
	now, _ := time.Parse(time.Kitchen, "10:05AM")

	state := State{endTime: endTime, now: now}

	newState, output := parseCommand(state, "status")

	assert.Equal(t, state, newState)
	assert.Equal(t, "20🍅 ", output.text)
}

func Test_parseCommand_done(t *testing.T) {
	endTime, _ := time.Parse(time.Kitchen, "10:25AM")
	now, _ := time.Parse(time.Kitchen, "10:25AM")

	state := State{endTime: endTime, now: now}

	newState, output := parseCommand(state, "status")

	assert.Equal(t, state, newState)
	assert.Equal(t, "0🍅 ", output.text)
}

func Test_parseCommand_past(t *testing.T) {
	endTime, _ := time.Parse(time.Kitchen, "10:25AM")
	now, _ := time.Parse(time.Kitchen, "10:35AM")

	state := State{endTime: endTime, now: now}

	newState, output := parseCommand(state, "status")

	assert.Equal(t, state, newState)
	assert.Equal(t, "❗️ 🍅 ", output.text)

}

func Test_parseCommand_bad(t *testing.T) {
	newState, output := parseCommand(emptyState, "foobar")

	assert.Equal(t, emptyState, newState)
	assert.Equal(t, usage, output.text)
	assert.Equal(t, 1, output.returnCode)
}

func Test_readExistingTime(t *testing.T) {
	os.Setenv("HOME", os.TempDir())

	readExistingTime()
}

func Test_filePath(t *testing.T) {
	assert.T(t, strings.HasSuffix(filePath(), "/.pomodoro"))
}
