package main

import (
	"github.com/aws/aws-cdk-go/awscdk"
	"github.com/aws/aws-cdk-go/awscdk/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/awss3assets"
	"github.com/aws/constructs-go/constructs/v3"
	"github.com/aws/jsii-runtime-go"
	"os"
)

func main() {
	app := awscdk.NewApp(nil)
	stack := awscdk.NewStack(app, jsii.String("BirthdayAlertStack"), &awscdk.StackProps{
		Env: env(),
	})

	createLambda(stack)

	app.Synth(nil)
}

func createLambda(construct constructs.Construct) {
	awslambda.NewFunction(construct, jsii.String("BirthdayAlertLambda"), &awslambda.FunctionProps{
		Description:  jsii.String("Lambda function to alert about upcoming birthdays"),
		FunctionName: jsii.String("alert"),
		MemorySize:   jsii.Number(128),
		Timeout:      awscdk.Duration_Seconds(jsii.Number(10)),
		Code:         awslambda.Code_FromAsset(jsii.String("../lambda/main.zip"), &awss3assets.AssetOptions{}),
		Handler:      jsii.String("main"),
		Runtime:      awslambda.Runtime_GO_1_X(),
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
