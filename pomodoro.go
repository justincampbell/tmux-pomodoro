package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	noTime     time.Time
	timeFormat = time.RFC3339
)

var usage = `tmux-pomodoro
github.com/justincampbell/tmux-pomodoro

  pomodoro start   Start a timer for 25 minutes
  pomodoro status  Show the remaining time, or an exclamation point if done
  pomodoro clear   Clear the timer
`

func init() {
	flag.Parse()
}

func main() {
	existingTime := readExistingTime()
	now := time.Now()
	args := flag.Args()

	newTime, output, returnCode := parseCommand(existingTime, now, args)

	if newTime != noTime {
		writeTime(newTime)
	}

	fmt.Println(output)

	if returnCode != 0 {
		os.Exit(returnCode)
	}
}

func parseCommand(existingTime time.Time, now time.Time, args []string) (newTime time.Time, output string, returnCode int) {
	if len(args) > 0 {
		switch args[0] {
		case "start":
			duration, _ := time.ParseDuration("25m")
			newTime = now.Add(duration)
			output = "Timer started, 25 minutes remaining"
		case "status":
			output = formatRemainingTime(existingTime, now)
		default:
			output = usage
			returnCode = 1
		}
	} else {
		output = usage
		returnCode = 0
	}

	return
}

func formatRemainingTime(existingTime time.Time, now time.Time) string {
	remaining := existingTime.Sub(now)
	remainingMinutes := remaining.Minutes()

	if remainingMinutes >= 0 {
		return strconv.FormatFloat(remainingMinutes, 'f', 0, 64)
	} else {
		return "❗️"
	}
}

func writeTime(t time.Time) {
	bytes := []byte(t.Format(timeFormat))
	err := ioutil.WriteFile(filePath(), bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func readExistingTime() time.Time {
	bytes, err := ioutil.ReadFile(filePath())
	if err != nil {
		return noTime
	}

	contents := string(bytes[:])
	contents = strings.TrimSpace(contents)

	result, err := time.Parse(timeFormat, contents)
	if err != nil {
		return noTime
	}

	return result
}

func filePath() string {
	return homeDir() + "/.pomodoro"
}

func homeDir() string {
	return os.Getenv("HOME")
}
