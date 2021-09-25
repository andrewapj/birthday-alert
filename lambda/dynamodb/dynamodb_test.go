package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"reflect"
	"sort"
	"testing"
)

func TestBuildTable(t *testing.T) {

	ddb := GetSession()
	defer deleteTable(ddb, t)

	buildTable(ddb, t)
}

func TestGetKey(t *testing.T) {

	// Given: a session and a DB table
	ddb := GetSession()
	buildTable(ddb, t)
	defer deleteTable(ddb, t)

	// And: some test data in the DB
	item := Item{
		Date:  "01/12",
		Names: []string{"Person 1", "Person 2"},
	}
	data, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		t.Fatalf("Unable to marshal test data, %s", err)
	}

	_, err = ddb.PutItem(&dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(Table),
	})
	if err != nil {
		t.Fatalf("Unable to put test data into table. %s", err)
	}

	// When: We get the birthdays for 01/12
	results := GetKey("01/12")

	// Then: We should get the correct names
	if len(results.Names) != 2 {
		t.Fatal("Expected a slice containing 2 elements")
	}
	sort.Strings(results.Names)
	if !reflect.DeepEqual(results.Names, item.Names) {
		t.Fatalf("Result data incorrect. Expected %s, got %s", item.Names, results.Names)
	}
}

func TestGetKeyWithNoData(t *testing.T) {

	// Given: a session and a DB table, but no data
	ddb := GetSession()
	buildTable(ddb, t)
	defer deleteTable(ddb, t)

	// When: We get the birthdays for 01/12
	results := GetKey("01/12")

	// Then: The names for that date should be empty
	if !reflect.DeepEqual(results.Names, []string{}) {
		t.Fatalf("Result data incorrect. Expected %s, got %s", []string{}, results.Names)
	}
}

//Helper function to build a DynamoDB table that is usually created by the CDK
func buildTable(ddb *dynamodb.DynamoDB, t *testing.T) {
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
		TableName:   aws.String(Table),
	})
	if err != nil {
		t.Error("Unable to create Dynamo DB table")
		t.Error(err)
		t.FailNow()
	}
}

//Helper function to delete a DynamoDB table
func deleteTable(ddb *dynamodb.DynamoDB, t *testing.T) {
	_, err := ddb.DeleteTable(&dynamodb.DeleteTableInput{TableName: aws.String(Table)})
	if err != nil {
		t.Error("Unable to create Dynamo DB table")
		t.Error(err)
		t.FailNow()
	}
}
