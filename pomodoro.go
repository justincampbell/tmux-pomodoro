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

	"github.com/justincampbell/tmux-pomodoro/tmux"
)

const timeFormat = time.RFC3339

var duration, _ = time.ParseDuration("25m")
var noTime time.Time

const usage = `
github.com/justincampbell/tmux-pomodoro

  pomodoro start   Start a timer for 25 minutes
  pomodoro status  Show the remaining time, or an exclamation point if done
  pomodoro clear   Clear the timer
`
const version = "v1.2.1"

// State is the state of the world passed through the functions to determine
// side-effects.
type State struct {
	endTime time.Time
	now     time.Time
}

// Output has fields for functions to set/append when they intend to output to
// the user.
type Output struct {
	text       string
	returnCode int
}

func init() {
	flag.Usage = func() {
		fmt.Printf("tmux-pomodoro %s\n", version)
		fmt.Printf("%s\n", strings.TrimSpace(usage))
	}

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

	if output.text != "" {
		fmt.Println(output.text)
	}

	if output.returnCode != 0 {
		os.Exit(output.returnCode)
	}
}

func refreshTmux() {
	_ = tmux.RefreshClient("-S")
}

func parseCommand(state State, command string) (newState State, output Output) {
	newState = state

	switch command {
	case "start":
		newState.endTime = state.now.Add(duration)
		output.text = "Timer started, 25 minutes remaining"
		killRunningBeepers()
		_ = startBeeper()
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
		_ = tmux.DisplayMessage("Pomodoro done, take a break!")
		refreshTmux()
	case "":
		flag.Usage()
	default:
		flag.Usage()
		output.returnCode = 1
	}

	return
}

func startBeeper() (err error) {
	ex, err := os.Executable()
	if err != nil {
		panic(err)
	}
	command := exec.Command(ex, "beep")
	err = command.Start()
	if err != nil {
		log.Println(err)
		return
	}
	bytes := []byte(strconv.Itoa(command.Process.Pid))
	err = ioutil.WriteFile(pidFilePath(), bytes, 0644)
	if err != nil {
		log.Println(err)
	}

	return
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
	}

	return "❗️"
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
