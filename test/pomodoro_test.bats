load test_helper

@test "shows usage" {
  run $pomodoro

  [ $status -eq 0 ]
  echo "$output" | grep "pomodoro"
}

@test "starts a timer and checks the status" {
  run $pomodoro start

  [ $status -eq 0 ]

  run $pomodoro status

  [ $status -eq 0 ]
  echo "$output" | grep "^25🍅 $"
}
