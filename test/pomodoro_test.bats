load test_helper

@test "starts a timer and checks the status" {
  run $pomodoro

  [ $status -eq 0 ]

  run $pomodoro status

  [ $status -eq 0 ]
  echo $output | grep "^25$"
}
