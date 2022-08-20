package database

import (
	"context"
	"fmt"
	"github.com/andrewapj/birthday-alert-cdk/lambda/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"log"
	"sync"
)

type Item struct {
	Date  string
	Names []string
}

var once sync.Once
var d *dynamodb.Client
var DynamoDBTable = "Birthdays"
var DynamoDBKey = "Date"

func GetClient() *dynamodb.Client {
	once.Do(func() {
		d = dynamodb.NewFromConfig(config.GetAwsConfig())
	})
	return d
}

func GetKey(key string) Item {
	ddb := GetClient()

	getItemInput := &dynamodb.GetItemInput{
		TableName: aws.String(DynamoDBTable),
		Key: map[string]types.AttributeValue{
			DynamoDBKey: &types.AttributeValueMemberS{Value: key},
		},
	}

	result, err := ddb.GetItem(context.TODO(), getItemInput)
	if err != nil {
		log.Println(fmt.Sprintf("Unable to get key %s from DynamoDB. Get item failed.%s", key, err))
		return emptyItem(key)
	}

	if result.Item == nil {
		return emptyItem(key)
	} else {
		item := Item{}
		err = attributevalue.UnmarshalMap(result.Item, &item)
		if err != nil {
			log.Println(fmt.Sprintf("Unable to unmarshal response. %s", err))
			return emptyItem(key)
		}
		return item
	}
}

func emptyItem(key string) Item {
	return Item{
		Date:  key,
		Names: []string{},
	}
}
