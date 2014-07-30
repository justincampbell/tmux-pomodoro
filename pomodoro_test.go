package main

import (
	"os"
	"strings"
	"testing"
	"time"

	"github.com/bmizerany/assert"
)

var emptyArgs []string

var newTime = make(chan time.Time)
var output = make(chan string, 1)

func Test_main(t *testing.T) {
	os.Setenv("HOME", os.TempDir())

	main()
}

func Test_parseCommand_start(t *testing.T) {
	now, _ := time.Parse(time.Kitchen, "10:00AM")

	newTime, output, _ := parseCommand(noTime, now, []string{"start"})

	expected, _ := time.Parse(time.Kitchen, "10:25AM")

	assert.T(t, expected.Equal(newTime))
	assert.Equal(t, "Timer started, 25 minutes remaining", output)
}

func Test_parseCommand_status(t *testing.T) {
	existingTime, _ := time.Parse(time.Kitchen, "10:25AM")
	now, _ := time.Parse(time.Kitchen, "10:05AM")

	newTime, output, _ := parseCommand(existingTime, now, []string{"status"})

	assert.Equal(t, noTime, newTime)
	assert.Equal(t, "20", output)
}

func Test_parseCommand_done(t *testing.T) {
	existingTime, _ := time.Parse(time.Kitchen, "10:25AM")
	now, _ := time.Parse(time.Kitchen, "10:25AM")

	newTime, output, _ := parseCommand(existingTime, now, []string{"status"})

	assert.Equal(t, noTime, newTime)
	assert.Equal(t, "0", output)
}

func Test_parseCommand_past(t *testing.T) {
	existingTime, _ := time.Parse(time.Kitchen, "10:25AM")
	now, _ := time.Parse(time.Kitchen, "10:35AM")

	newTime, output, _ := parseCommand(existingTime, now, []string{"status"})

	assert.Equal(t, noTime, newTime)
	assert.Equal(t, "❗️", output)
}

func Test_readExistingTime(t *testing.T) {
	os.Setenv("HOME", os.TempDir())

	readExistingTime()
}

func Test_filePath(t *testing.T) {
	assert.T(t, strings.HasSuffix(filePath(), "/.pomodoro"))
}
