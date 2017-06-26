#!/usr/bin/env bash

get_tmux_option() {
  local option="$1"
  local default_value="$2"
  local option_value=$(tmux show-option -gqv "$option")

  if [[ -z $option_value ]]; then
    echo "$default_value"
  else
    echo "$option_value"
  fi
}

cmd_exists() {
  local cmd="$1"

  type "$cmd" >/dev/null 2>&1
}

lowercase(){
    echo "$1" | sed "y/ABCDEFGHIJKLMNOPQRSTUVWXYZ/abcdefghijklmnopqrstuvwxyz/"
}

install_pomodoro() {
  local target_dir=$1
  mkdir -p "$target_dir"

  local os=$(lowercase $(uname))
  local mach=$(uname -m)
  if [[ "$mach" == "x86_64" ]]; then
    suffix=amd64
  elif [[ "$mach" == "i?86" ]]; then
    suffix=386
  elif [[ "$mach" == "arm*" ]]; then
    suffix=arm
  fi

  (cd $target_dir && curl -fsSL "https://github.com/justincampbell/tmux-pomodoro/releases/download/v1.2.1/tmux-pomodoro_${os}_${suffix}.tar.gz" | tar xzf -)
}

get_pomodoro_cmd() {
  local bin_name=tmux-pomodoro
  local cur_dir="$(cd "$(dirname "$0")" && pwd)"
  local bin_dir="$cur_dir/bin"
  local PATH="$bin_dir":$PATH
  if ! cmd_exists "$bin_name"; then
    install_pomodoro "$bin_dir"
  fi
  which "$bin_name"
}

# Local Variables:
# mode: Shell-Script
# sh-indentation: 2
# indent-tabs-mode: nil
# sh-basic-offset: 2
# End:
# vim: ft=sh sw=2 ts=2 et
