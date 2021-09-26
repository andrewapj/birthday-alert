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
- Create a group and policy with admin permissions, or just enough permissions to perform the required actions.
- Create a user and add it to the group.
- Create an access key for the new user.

### Configure AWS CLI
- Configure the CLI for the new user
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

### Adding birthdays to AWS DynamoDB
- Birthdays can be added manually via the AWS DynamoDB console
- The hash key for the table is 'Date'.
  - In the format "DD/MM", left padded with zeros as necessary
- Each item contains only one other attribute called 'Names'
  - This is a list of strings that contains the names of people who's birthdays fall on the date
  specified by the key
- Below is a CSV example:
```
"Date","Names"
"01/09","[{""S"":""Person1, Person2""}]"
```

## Useful CDK Commands
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
* `go test`         run unit tests