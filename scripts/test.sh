#!/bin/bash
green() {
  "$@" | GREP_COLORS='mt=01;32' grep --color .
}

red() {
  "$@" | GREP_COLORS='mt=01;31' grep --color .
}

yellow() {
  "$@" | GREP_COLORS='mt=01;93' grep --color .
}

check_command_success() {
  CODE_TO_COMPARE_TO=$2
  RETURNED_CODE=$1
  if [ $RETURNED_CODE -ne $CODE_TO_COMPARE_TO ]; then
    if [[ $2 != "" ]]; then
      red echo "$3"
    fi
    exit 1
  fi
}


green echo "Starting building testing container"
docker build -t testing_binary_chop:test . -f ./build/package/testing/Dockerfile
RESULT=$?
check_command_success $RESULT 0 "Could not build testing container"

green echo "Starting tests within Docker container"
docker container create --name testing_chop testing_binary_chop:test
RESULT=$?
check_command_success $RESULT 0 "Could not start testing container"

green echo "Extracting cover HTML page"
docker container cp testing_chop:/go/src/github.com/clnbs/karateChop/cover.html ./cover.html
RESULT=$?
check_command_success $RESULT 0 "Could not extract cover HTML page"

green echo "Deleting testing container"
docker container rm -f testing_chop
RESULT=$?
check_command_success $RESULT 0 "Could not remove testing container"
