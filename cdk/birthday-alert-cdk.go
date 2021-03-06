package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awsdynamodb"
	"github.com/aws/aws-cdk-go/awscdk/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/aws-cdk-go/awscdk/awssns"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"os"
)

func main() {
	app := awscdk.NewApp(nil)
	stack := awscdk.NewStack(app, jsii.String("BirthdayAlertStack"), &awscdk.StackProps{
		Env: env(),
	})

	table := createDynamoDbTable(stack)
	topic := createSNSAndSubscription(stack)
	lambda := createLambda(stack, table, topic)
	addSNSPermissions(topic, lambda)
	createEventRule(stack, lambda)

	app.Synth(nil)
}

//Create the DynamoDB table
func createDynamoDbTable(construct constructs.Construct) awsdynamodb.Table {
	return awsdynamodb.NewTable(construct, jsii.String("BirthdayAlertDDBTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("Date"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey:     nil,
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TableName:   jsii.String("Birthdays"),
	})
}

//Create the lambda with an extra policy attached that allows it to access the DynamoDB table.
func createLambda(construct constructs.Construct, table awsdynamodb.Table, topic awssns.Topic) awslambda.Function {
	l := awslambda.NewFunction(construct, jsii.String("BirthdayAlertLambda"), &awslambda.FunctionProps{
		Description:  jsii.String("Lambda function to alert about upcoming birthdays"),
		FunctionName: jsii.String("alert"),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(10)),
		Code:         awslambda.Code_FromAsset(jsii.String("../lambda/main.zip"), &awss3assets.AssetOptions{}),
		Handler:      jsii.String("main"),
		Runtime:      awslambda.Runtime_GO_1_X(),
		Environment: &map[string]*string{
			"APP_AWS_ENDPOINT":        jsii.String(""),
			"APP_AWS_REGION":          env().Region,
			"APP_DAYS_LOOKAHEAD":      jsii.String("7, 14"),
			"APP_NOTIFICATION_STRING": jsii.String("It's %s's birthday on %s"),
			"APP_SNS_TOPIC":           topic.TopicArn(),
		},
	})

	l.AddToRolePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Actions:   &[]*string{jsii.String("*")},
		Effect:    awsiam.Effect_ALLOW,
		Resources: &[]*string{table.TableArn()},
	}))

	return l
}

func createSNSAndSubscription(construct constructs.Construct) awssns.Topic {
	t := awssns.NewTopic(construct, jsii.String("BirthdayAlertTopic"), &awssns.TopicProps{
		DisplayName: jsii.String("Birthday alert topic"),
		Fifo:        jsii.Bool(false),
		TopicName:   jsii.String("BIRTHDAY_ALERT"),
	})

	if email, exists := os.LookupEnv("CDK_EMAIL_SUBSCRIPTION"); exists == true {
		awssns.NewSubscription(construct, jsii.String("BirthdayAlertTopicSubscription"), &awssns.SubscriptionProps{
			Endpoint: jsii.String(email),
			Protocol: awssns.SubscriptionProtocol_EMAIL,
			Topic:    t,
		})
	}
	return t
}

func addSNSPermissions(t awssns.Topic, function awslambda.Function) {
	t.AddToResourcePolicy(awsiam.NewPolicyStatement(&awsiam.PolicyStatementProps{
		Principals: &[]awsiam.IPrincipal{function.GrantPrincipal()},
		Actions: &[]*string{
			jsii.String("sns:GetTopicAttributes"),
			jsii.String("sns:ListSubscriptionsByTopic"),
			jsii.String("sns:Publish"),
			jsii.String("sns:SetTopicAttributes"),
			jsii.String("sns:Subscribe"),
		},
		Effect:    awsiam.Effect_ALLOW,
		Resources: &[]*string{jsii.String("*")},
	}))
}

func createEventRule(construct constructs.Construct, function awslambda.Function) {
	awsevents.NewRule(construct, jsii.String("BirthdayAlertLambdaEvent"), &awsevents.RuleProps{
		Description:  jsii.String("The event that triggers the Birthday Alert Lambda"),
		Enabled:      jsii.Bool(true),
		EventBus:     nil,
		EventPattern: nil,
		RuleName:     jsii.String("BirthdayAlertLambdaEvent"),
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
			Hour:   jsii.String("12"),
			Minute: jsii.String("0"),
		}),
		Targets: &[]awsevents.IRuleTarget{
			awseventstargets.NewLambdaFunction(function, &awseventstargets.LambdaFunctionProps{})},
	})
}

func env() *awscdk.Environment {
	account, accountEnvOk := os.LookupEnv("CDK_DEFAULT_ACCOUNT")
	region, regionEnvOk := os.LookupEnv("CDK_DEFAULT_REGION")

	if accountEnvOk && regionEnvOk {
		return &awscdk.Environment{
			Account: jsii.String(account),
			Region:  jsii.String(region),
		}
	} else {
		return nil
	}
}
