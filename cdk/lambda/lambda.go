package lambda

import (
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// CreateBirthdayLambda create the lambda with an extra policy attached that allows it to access the DynamoDB table.
func CreateBirthdayLambda(parent constructs.Construct, env *awscdk.Environment, table awsdynamodb.Table, topic awssns.Topic) awslambda.Function {
	l := awslambda.NewFunction(parent, jsii.String("BirthdayAlertLambda"), &awslambda.FunctionProps{
		Description:  jsii.String("Lambda function to alert about upcoming birthdays"),
		FunctionName: jsii.String("birthday-alert"),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(10)),
		Code:         awslambda.Code_FromAsset(jsii.String("../lambda/main.zip"), &awss3assets.AssetOptions{}),
		Handler:      jsii.String("main"),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Architecture: awslambda.Architecture_X86_64(),
		Environment: &map[string]*string{
			"APP_PROFILE":              jsii.String("aws"),
			"APP_AWS_REGION":           env.Region,
			"APP_DAYS_LOOKAHEAD":       jsii.String("7, 14"),
			"APP_NOTIFICATION_MESSAGE": jsii.String("It's %s's birthday on %s"),
			"APP_SNS_TOPIC":            topic.TopicArn(),
		},
	})

	l.AddToRolePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions:   &[]*string{jsii.String("*")},
		Effect:    awsiam.Effect_ALLOW,
		Resources: &[]*string{table.TableArn()},
	}))

	return l
}
