#!/usr/bin/env bash

CURRENT_DIR="$(cd "$(dirname "$0")" && pwd)"
source "${CURRENT_DIR}/scripts/helpers.sh"

readonly pomodoro_start_key="$(get_tmux_option "@pomodoro-start-key" "p")"
readonly pomodoro_clear_key="$(get_tmux_option "@pomodoro-clear-key" "P")"
readonly pomodoro_cmd="$(get_pomodoro_cmd)"

append_option() {
  local dir=$1
  local options=$2
  local appendix=$3
  if ! echo "$options" | grep -qF "$appendix"; then
    case "$dir" in
      "left")
        options="$appendix $options"
        ;;
      "right")
        options="$options $appendix"
        ;;
    esac
  fi
  echo "$options"
}

update_tmux_option() {
  local option=$1
  local option_value
  option_value=$(get_tmux_option "$option")
  local new_option_value
  new_option_value=$(append_option left "$option_value" '#('"$pomodoro_cmd"' status)')
  tmux set -g "$option" "$new_option_value"
}

main() {
  update_tmux_option "status-right"
  tmux bind-key "$pomodoro_start_key" run "$pomodoro_cmd start > /dev/null"
  tmux bind-key "$pomodoro_clear_key" run "$pomodoro_cmd clear > /dev/null"
}
main

# Local Variables:
# mode: Shell-Script
# sh-indentation: 2
# indent-tabs-mode: nil
# sh-basic-offset: 2
# End:
# vim: ft=sh sw=2 ts=2 et
