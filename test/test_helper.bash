setup() {
  PMDR_TMPDIR=$BATS_TMPDIR/pmdr

  rm -rf $PMDR_TMPDIR
  mkdir -p $PMDR_TMPDIR

  pomodoro=bin/pomodoro

  HOME=$PMDR_TMPDIR
}

teardown() {
  echo $status: $output
}
