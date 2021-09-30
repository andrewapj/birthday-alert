# Birthday Alert

An AWS Lambda that alerts about upcoming birthdays, deployed using the CDK.

## Architecture

The following AWS resources are used:
- AWS EventBridge Rule
  - Runs on a schedule of once per day and invokes the lambda
- AWS Lambda
  - Invoked by AWS EventBridge (or manually).
  - Takes the current date (in the format DD/MM) and adds the look ahead days to it.
  - Queries DynamoDB for each of these dates.
  - If a name is found, generates a message and sends it to SNS.
  - Configured by environment variables, set by the CDK.
- AWS SNS
  - Contains an email subscription.
  - Messages sent to this topic are sent via email.
- AWS DynamoDB
  - Contains the birthdays.
  - The Hash Key is the date of the birthday in the format (DD/MM).
  - Each key should contain a list of strings called 'Names'.
- AWS IAM
  - The required AWS roles and policies are managed by the CDK (excluding the permissions required to run the CDK)

## Prerequisites

- AWS Account
- AWS CLI
- Node
- Go

## Testing locally
- In the 'lambda' directory
```
docker-compose up -d
go test ./... -v
docker-compose down
```

## Deploying to AWS

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

### Building the lambda
- In the 'lambda' directory
- Ensure the build script is executable
```
./build.sh
```

### Running the CDK
- Set the required environment variables:
  - E.G. ```export CDK_EMAIL_SUBSCRIPTION=user@domain.com```
  
| Variable | Default Value | Description | Required |
|------|---------------|----------|---------------|
| CDK_DEFAULT_ACCOUNT | Value configured in the AWS Client | The account id of your AWS account| No |
| CDK_DEFAULT_REGION | Value configured in the AWS Client | The region to deploy to | No |
| CDK_EMAIL_SUBSCRIPTION | "" | The email address that should subscribe to SNS (receive email notifications) | Yes (If missing, no emails will be sent) |

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

## Running the lambda

### Automatically
- The lambda is set up to automatically run every day.
  - Specified by the AWS EventBridge Rule in the CDK

### Manually
- The lambda can be run manually in the AWS Lambda console.
- Find the function named 'alert'
- Select the 'Test' tab and click the 'Test' button

## Useful CDK Commands
* `cdk deploy`      deploy this stack to your default AWS account/region
* `cdk diff`        compare deployed stack with current state
* `cdk synth`       emits the synthesized CloudFormation template
* `go test`         run unit tests