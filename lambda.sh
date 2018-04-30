#!/bin/bash
set -e

if [ -z "$LAMBDA_FUNCTION" ]; then
  echo 'set $LAMBDA_FUNCTION before proceeding'
  exit 1
fi

GO=$(which go 2>/dev/null)
ZIP=$(which zip 2>/dev/null)
AWS=$(which aws 2>/dev/null)

if [ -z "$GO" ]; then
  echo "no 'go' executable found in path, exiting"
  exit 1
fi

if [ -z "$ZIP" ]; then
  echo "no 'zip' executable found in path, exiting"
  exit 1
fi

if [ -z "$AWS" ]; then
  echo "no 'aws' executable found in path, exiting"
  exit 1
fi

if [ -z "$GOOS" ]; then
  GOOS=linux
fi

if [ -z "$GOARCH" ]; then
  GOARCH=amd64
fi

$GO build -o main main.go && \
  $ZIP main.zip main && \
  $AWS lambda update-function-code \
    --function-name $LAMBDA_FUNCTION \
    --zip-file fileb://main.zip \
    --publish
