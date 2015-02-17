package tmux

import "os/exec"

// Command runs a tmux command with given arguments
func Command(subcommand string, args ...string) error {
	command := []string{subcommand}
	for _, arg := range args {
		command = append(command, arg)
	}
	return exec.Command("tmux", command...).Start()
}

// DisplayMessage runs tmux display-message with the given string
func DisplayMessage(message string) error {
	return Command("display-message", message)
}

// RefreshClient tells tmux to refresh the display, including the status-bar
func RefreshClient(args ...string) error {
	return Command("refresh-client", args...)
}
