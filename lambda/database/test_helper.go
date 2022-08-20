package database

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"testing"
)

// PutItem puts an item into an existing dynamoDB table for tests
func PutItem(t *testing.T, ddb *dynamodb.Client, item Item) Item {
	data, err := attributevalue.MarshalMap(item)
	if err != nil {
		t.Fatalf("Unable to marshal test data, %s", err)
	}

	_, err = ddb.PutItem(context.TODO(), &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(DynamoDBTable),
	})

	if err != nil {
		t.Fatalf("Unable to put test data into table. %s", err)
	}
	return item
}

// BuildTable builds a DynamoDB table for tests that is usually created by the CDK
func BuildTable(ddb *dynamodb.Client, t *testing.T) {

	_, err := ddb.CreateTable(context.TODO(), &dynamodb.CreateTableInput{
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String(DynamoDBKey),
				AttributeType: types.ScalarAttributeTypeS,
			},
		},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String(DynamoDBKey),
				KeyType:       types.KeyTypeHash,
			},
		},
		TableName:   aws.String(DynamoDBTable),
		BillingMode: types.BillingModePayPerRequest,
		TableClass:  "",
	})

	if err != nil {
		t.Error("Unable to create Dynamo DB table")
		t.Error(err)
		t.FailNow()
	}
}

// DeleteTable deletes a DynamoDB table for tests
func DeleteTable(ddb *dynamodb.Client, t *testing.T) {
	_, err := ddb.DeleteTable(context.TODO(), &dynamodb.DeleteTableInput{TableName: aws.String(DynamoDBTable)})
	if err != nil {
		t.Error("Unable to create Dynamo DB table")
		t.Error(err)
		t.FailNow()
	}
}
