load test_helper

@test "works" {
  run $pomodoro

  [ $status -eq 0 ]
}
