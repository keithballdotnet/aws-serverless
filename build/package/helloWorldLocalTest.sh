
#!/usr/bin/env bash

set -u
set -e 

GREEN="\033[0;32m"
PS_CLEAR="\033[0m"

function echo_g() {
    echo "\n$GREEN$1$PS_CLEAR"
}

echo_g 'Running tests...'
go test github.com/keithballdotnet/aws-serverless/functions/helloWorld

cd src/github.com/keithballdotnet/aws-serverless/functions/helloWorld

echo_g 'Building...'
GOOS=linux GOARCH=amd64 go build -o main

echo_g 'Running local...'
sam local start-api