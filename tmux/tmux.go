package tmux

import "os/exec"

func Command(subcommand string, args ...string) error {
	command := []string{subcommand}
	for _, arg := range args {
		command = append(command, arg)
	}
	return exec.Command("tmux", command...).Start()
}

func DisplayMessage(message string) error {
	return Command("display-message", message)
}

func RefreshClient(args ...string) error {
	return Command("refresh-client", args...)
}
