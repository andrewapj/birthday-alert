# Birthday Alert

An AWS Lambda that alerts about upcoming birthdays, deployed using the CDK

## Prerequisites

- AWS Account
- AWS CLI
- Node
- Go

## Setup

### Install CDK
```
npm install -g aws-cdk
cdk --version
```

### Create User and Group

- Login to the AWS console as the root user and select IAM.
- Create a group with admin permissions, or just enough permissions to perform the required actions.
- Create a user and add it to the group.
- Create an access key for the user.

### Configure AWS CLI
```
aws configure
```

### Bootstrap the CDK
```
cdk bootstrap aws://{AWS_ACCOUNT}/{REGION}
```

## Testing locally
- In the 'lambda' directory
```
docker-compose up -d
go test ./... -v
docker-compose down
```

## Deploying the lambda

### Build the lambda
- In the 'lambda' directory
- Ensure the build script is executable

```
./build.sh
```

### Deploy the lambda
- In the 'cdk' directory
```
cdk deploy
```

## Useful CDK Commands
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
* `go test`         run unit tests