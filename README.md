# AWS Serverless
An experiment with aws serverless tech

<a href="https://travis-ci.org/keithballdotnet/aws-serverless"><img src="https://travis-ci.org/keithballdotnet/aws-serverless.svg?branch=master" alt="Build"></a>

## AWS Serverless Application Model 

### Prerequisites

[Docker](https://www.docker.com/products/docker-desktop)<br>
[AWS Cli](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html)<br>
[AWS SAM cli](https://docs.aws.amazon.com/serverless-application-model/latest/developerguide/serverless-sam-cli-install-mac.html)

```bash
brew update
brew upgrade
brew tap aws/tap
brew install aws-sam-cli
```

### Run the hello world example

Test, build and then run the function locally
```bash
./src/github.com/keithballdotnet/aws-serverless/build/package/helloWorldLocalTest.sh
```

You can test the function using a curl command
```bash
$ curl -XPUT -d "Steve" http://127.0.0.1:3000/helloWorld
Hello Steve
```

