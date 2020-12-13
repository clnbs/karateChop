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

build_recursive() {
  green echo "Start building binary chop recursive container"
  docker build -t recursive:build . -f ./build/package/recursive/Dockerfile
  RESULT=$?
  check_command_success $RESULT 0 "Could not build binary chop recursive container"

  green echo "Creating container to build binary chop recursive"
  docker container create --name extract recursive:build
  RESULT=$?
  check_command_success $RESULT 0 "Could not start builder container"

  green echo "Extracting binary from builder container"
  docker container cp extract:/go/src/github.com/clnbs/karateChop/recursive.bin ./recursive.bin
  RESULT=$?
  check_command_success $RESULT 0 "Could not extract binary from builder image"

  green echo "Deleting builder container"
  docker container rm -f extract
  RESULT=$?
  check_command_success $RESULT 0 "Could not remove builder container"
}

OPTION=$1

if [ -z "$OPTION"  ]; then
  green echo "Compiling all implementation"
  build_recursive
  exit 0
elif [[ "$OPTION" == "recursive" ]]; then
  green echo "Compiling recursive implementation"
  build_recursive
fi