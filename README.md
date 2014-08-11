# tmux-pomodoro [![Build Status](https://travis-ci.org/justincampbell/tmux-pomodoro.svg?branch=conversion)](https://travis-ci.org/justincampbell/tmux-pomodoro)

## Installation

### [Homebrew](http://brew.sh)

    brew tap justincampbell/formulae
    brew install tmux-pomodoro

### Package

    wget -O tmux-pomodoro-latest.tar.gz https://github.com/justincampbell/tmux-pomodoro/archive/latest.tar.gz tmux-pomodoro
    tar -zxvf tmux-pomodoro-latest.tar.gz
    cd tmux-pomodoro-latest/
    make install

## Usage

### Commands

* `start` Start a timer for 25 minutes
* `status` Show the remaining time, or an exclamation point if done
* `clear` Clear the timer

`start` and `clear` also call `tmux refresh-client -S`, which will instantly update your tmux status bar shell commands.
