package dynamodb

import (
	"github.com/aws/aws-cdk-go/awscdk/v2/awsdynamodb"
	"github.com/aws/constructs-go/constructs/v10"
	"github.com/aws/jsii-runtime-go"
)

// CreateBirthdayTable create the DynamoDB table
func CreateBirthdayTable(parent constructs.Construct) awsdynamodb.Table {
	return awsdynamodb.NewTable(parent, jsii.String("BirthdayAlertDDBTable"), &awsdynamodb.TableProps{
		PartitionKey: &awsdynamodb.Attribute{
			Name: jsii.String("Date"),
			Type: awsdynamodb.AttributeType_STRING,
		},
		SortKey:     nil,
		BillingMode: awsdynamodb.BillingMode_PAY_PER_REQUEST,
		TableName:   jsii.String("Birthdays"),
	})
}
