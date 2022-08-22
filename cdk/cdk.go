package main

import (
	"cdk/dynamodb"
	"cdk/event"
	"cdk/lambda"
	"cdk/sns"
	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/aws/jsii-runtime-go"
	"os"
)

func main() {
	app := awscdk.NewApp(nil)

	stack := awscdk.NewStack(app, jsii.String("BirthdayAlertStack"), &awscdk.StackProps{
		Env: env(),
	})

	table := dynamodb.CreateBirthdayTable(stack)
	topic := sns.CreateSNSAndSubscription(stack)
	l := lambda.CreateBirthdayLambda(stack, env(), table, topic)
	sns.AddSNSPermissions(topic, l)
	event.CreateEventRule(stack, l)

	app.Synth(nil)
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
