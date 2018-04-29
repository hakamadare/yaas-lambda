#!/bin/bash
set -e

GO=$(which go 2>/dev/null)
ZIP=$(which zip 2>/dev/null)

if [ -z "$GO" ]; then
  echo "no 'go' binary found in path, exiting"
  exit 1
fi

if [ -z "$ZIP" ]; then
  echo "no 'zip' binary found in path, exiting"
  exit 1
fi

if [ -z "$GOOS" ]; then
  GOOS=linux
fi

if [ -z "$GOARCH" ]; then
  GOARCH=amd64
fi

$GO build -o main main.go && \
  $ZIP main.zip main
