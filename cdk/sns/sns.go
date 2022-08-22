package sns

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsiam"
	"github.com/aws/aws-cdk-go/awscdk/v2/awslambda"
	"github.com/aws/aws-cdk-go/awscdk/v2/awssns"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
	"os"
)

// CreateSNSAndSubscription creates an SNS topic and a subscription
func CreateSNSAndSubscription(construct constructs.Construct) awssns.Topic {
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

// AddSNSPermissions adds a resource permission to SNS
func AddSNSPermissions(t awssns.Topic, function awslambda.Function) {
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
