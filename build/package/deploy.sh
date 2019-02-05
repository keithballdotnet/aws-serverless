
#!/usr/bin/env bash

set -u
set -e 

GREEN="\033[0;32m"
PS_CLEAR="\033[0m"

function echo_g() {
    echo "\n$GREEN$1$PS_CLEAR"
}

echo_g 'Running tests...'
go test github.com/keithballdotnet/aws-serverless/functions/organisation

echo_g Building...
GOOS=linux GOARCH=amd64 go build -o bin/main github.com/keithballdotnet/aws-serverless/functions/organisation

echo_g Zipping...
zip bin/deployment.zip bin/main

echo_g Deploying...
aws lambda update-function-code --function-name organisation --zip-file fileb://./bin/deployment.zip --region eu-central-1