# tmux-pomodoro [![Build Status](https://travis-ci.org/justincampbell/tmux-pomodoro.svg?branch=conversion)](https://travis-ci.org/justincampbell/tmux-pomodoro)

## Installation

1. Download the latest package for your platform from the [Releases page](https://github.com/justincampbell/tmux-pomodoro/releases/latest).
2. Untar the package with `tar -zxvf tmux-pomodoro*.tar.gz`.
3. Move the extracted `pomodoro` file to a directory in your `$PATH` (for most systems, this will be `/usr/local/bin/`).

Or, if you have a [Go development environment](https://golang.org/doc/install):

```
go get github.com/justincampbell/tmux-pomodoro
```

## Usage

### Tmux Configuration

```tmux
# Place the current pomodoro status on the right side of your status bar
set -g status-right '#(pomodoro status)'

# Map a key to start a timer
bind-key p run-shell 'pomodoro start'
```

### Commands

* `start` Start a timer for 25 minutes
* `status` Show the remaining time, or an exclamation point if done
* `clear` Clear the timer

`start` and `clear` also call `tmux refresh-client -S`, which will instantly update your tmux status bar shell commands.
