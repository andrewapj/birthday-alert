package event

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsevents"
	"github.com/aws/aws-cdk-go/awscdk/v2/awseventstargets"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// CreateEventRule creates an event bridge rule
func CreateEventRule(construct constructs.Construct, function awslambda.Function) {
	awsevents.NewRule(construct, jsii.String("BirthdayAlertEvent"), &awsevents.RuleProps{
		Description:  jsii.String("The event that triggers the Birthday Alert Lambda"),
		Enabled:      jsii.Bool(true),
		EventBus:     nil,
		EventPattern: nil,
		RuleName:     jsii.String("BirthdayAlertEvent"),
		Schedule: awsevents.Schedule_Cron(&awsevents.CronOptions{
			Hour:   jsii.String("12"),
			Minute: jsii.String("0"),
		}),
		Targets: &[]awsevents.IRuleTarget{
			awseventstargets.NewLambdaFunction(function, &awseventstargets.LambdaFunctionProps{})},
	})
}
