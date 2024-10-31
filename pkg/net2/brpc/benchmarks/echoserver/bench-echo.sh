#!/bin/bash

program_name=$0

function print_usage {
    echo ""
    echo "Usage: $program_name [connections] [duration]"
    echo ""
    echo "connections:  Connections to keep open to the destinations"
    echo "duration:     Exit after the specified amount of time"
    echo ""
    echo "--- EXAMPLE ---"
    echo ""
    echo "$program_name 1000 30"
    echo ""
    exit 1
}

# if less than two arguments supplied, display usage
if [  $# -le 1 ]; then
  print_usage
  exit 1
fi

# check whether user had supplied -h or --help . If yes display usage
if [[ ( $# == "--help") ||  $# == "-h" ]]; then
  print_usage
  exit 0
fi

set -e

echo ""
echo "--- BENCH ECHO START ---"
echo ""

function cleanup() {
  echo "--- BENCH ECHO DONE ---"
  # shellcheck disable=SC2046
  kill -9 $(jobs -rp)
  # shellcheck disable=SC2046
  wait $(jobs -rp) 2>/dev/null
}
trap cleanup EXIT

mkdir -p bin

eval "$(pkill -9 -f echoserver || printf "")"

conn_num=$1
test_duration=$2

function go_bench() {
  echo "--- $1 ---"
  echo ""
  go build -tags=poll_opt -gcflags="-l=4" -ldflags="-s -w" -o "$2" "$3"

  ./$2 --port "$4" --multicore="$5" &

  echo "Warming up for 1 seconds..."
  sleep 1
  echo ""

  echo "--- BENCHMARK START ---"
  printf "*** %d connections, %d seconds\n" "$conn_num" "$test_duration"
  echo ""

  tcpkali --connections "$conn_num" --connect-rate "$conn_num" --duration "$test_duration"'s' -f ./testdata/test.data 127.0.0.1:"$4"
  echo ""
  echo "--- BENCHMARK DONE ---"
  echo ""
}

go_bench "GNET" echoserver echo_server.go 9300 true
