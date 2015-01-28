load test_helper

@test "shows usage" {
  run $pomodoro

  [ $status -eq 0 ]
  echo "$output" | grep "pomodoro"
}
