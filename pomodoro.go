package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var (
	noTime     time.Time
	timeFormat = time.RFC3339
)

var duration, _ = time.ParseDuration("25m")

var usage = `tmux-pomodoro
github.com/justincampbell/tmux-pomodoro

  pomodoro start   Start a timer for 25 minutes
  pomodoro status  Show the remaining time, or an exclamation point if done
  pomodoro clear   Clear the timer
`

type State struct {
	endTime time.Time
	now     time.Time
}

type Output struct {
	text       string
	returnCode int
}

func init() {
	flag.Parse()
}

func main() {
	state := State{
		endTime: readExistingTime(),
		now:     time.Now(),
	}

	args := flag.Args()
	var command string
	if len(args) == 0 {
		command = ""
	} else {
		command = args[0]
	}

	newState, output := parseCommand(state, command)

	if newState.endTime != state.endTime {
		writeTime(newState.endTime)
	}

	fmt.Println(output.text)

	if output.returnCode != 0 {
		os.Exit(output.returnCode)
	}
}

func refreshTmux() {
	_ = exec.Command("tmux", "refresh-client", "-S").Start()
}

func parseCommand(state State, command string) (newState State, output Output) {
	newState = state

	switch command {
	case "start":
		newState.endTime = state.now.Add(duration)
		output.text = "Timer started, 25 minutes remaining"
		killRunningBeepers()
		startBeeper()
		refreshTmux()
	case "status":
		if state.endTime == noTime {
			return
		}
		output.text = formatRemainingTime(state.endTime, state.now) + " 🍅 "
	case "clear":
		newState.endTime = noTime
		output.text = "Pomodoro cleared!"
		killRunningBeepers()
		refreshTmux()
	case "beep":
		<-time.NewTicker(duration).C
		_ = exec.Command("tmux", "display-message", "Pomodoro done, take a break!").Run()
		refreshTmux()
	case "":
		output.text = usage
	default:
		output.text = usage
		output.returnCode = 1
	}

	return
}

func startBeeper() {
	command := exec.Command("pomodoro", "beep")
	err := command.Start()
	if err != nil {
		log.Println(err)
	}

	bytes := []byte(strconv.Itoa(command.Process.Pid))
	err = ioutil.WriteFile(pidFilePath(), bytes, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func killRunningBeepers() {
	bytes, err := ioutil.ReadFile(pidFilePath())
	if err != nil {
		return
	}
	pid, err := strconv.Atoi(string(bytes[:]))
	if err != nil {
		return
	}
	process, err := os.FindProcess(pid)
	if err != nil {
		return
	}
	_ = process.Kill()
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
	var bytes []byte
	if t != noTime {
		bytes = []byte(t.Format(timeFormat))
	}
	err := ioutil.WriteFile(filePath(), bytes, 0644)
	if err != nil {
		log.Fatal(err)
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

func pidFilePath() string {
	return homeDir() + "/.pomodoro.pid"
}

func homeDir() string {
	return os.Getenv("HOME")
}
