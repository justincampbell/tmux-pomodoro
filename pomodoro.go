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

func init() {
	flag.Parse()
}

func main() {
	existingTime := readExistingTime()
	now := time.Now()
	args := flag.Args()

	newTime, output := parseCommand(existingTime, now, args)

	if newTime != noTime {
		writeTime(newTime)
	}

	fmt.Println(output)
}

func parseCommand(existingTime time.Time, now time.Time, args []string) (newTime time.Time, output string) {
	if len(args) > 0 && args[0] == "status" {
		remaining := existingTime.Sub(now)
		remainingMinutes := remaining.Minutes()
		output = strconv.FormatFloat(remainingMinutes, 'f', 0, 64)
	} else {
		duration, err := time.ParseDuration("25m")
		if err != nil {
			panic(err)
		}
		newTime = now.Add(duration)
		output = "25"
	}

	return
}

func writeTime(t time.Time) {
	bytes := []byte(t.Format(timeFormat))
	err := ioutil.WriteFile(filePath(), bytes, 0644)
	if err != nil {
		panic(err)
	}
}

func readExistingTime() time.Time {
	var noTime time.Time

	bytes, err := ioutil.ReadFile(filePath())
	if err != nil {
		return noTime
	}

	contents := string(bytes[:])
	contents = strings.TrimSpace(contents)

	result, err := time.Parse(timeFormat, contents)
	if err != nil {
		panic(err)
	}

	return result
}

func filePath() string {
	return homeDir() + "/.pomodoro"
}

func homeDir() string {
	return os.Getenv("HOME")
}
