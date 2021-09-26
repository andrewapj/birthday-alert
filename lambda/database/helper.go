package database

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"testing"
)

//PutItem puts an item into an existing dynamoDB table for tests
func PutItem(t *testing.T, ddb *dynamodb.DynamoDB, item Item) Item {
	data, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		t.Fatalf("Unable to marshal test data, %s", err)
	}

	_, err = ddb.PutItem(&dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(DynamoDBTable),
	})
	if err != nil {
		t.Fatalf("Unable to put test data into table. %s", err)
	}
	return item
}

//BuildTable builds a DynamoDB table for tests that is usually created by the CDK
func BuildTable(ddb *dynamodb.DynamoDB, t *testing.T) {
	_, err := ddb.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Date"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Date"),
				KeyType:       aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName:   aws.String(DynamoDBTable),
	})
	if err != nil {
		t.Error("Unable to create Dynamo DB table")
		t.Error(err)
		t.FailNow()
	}
}

//DeleteTable deletes a DynamoDB table for tests
func DeleteTable(ddb *dynamodb.DynamoDB, t *testing.T) {
	_, err := ddb.DeleteTable(&dynamodb.DeleteTableInput{TableName: aws.String(DynamoDBTable)})
	if err != nil {
		t.Error("Unable to create Dynamo DB table")
		t.Error(err)
		t.FailNow()
	}
}
