package database

import (
	"fmt"
	"github.com/andrewapj/birthday-alert-cdk/lambda/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"log"
	"sync"
)
import "github.com/aws/aws-sdk-go/aws/session"

type Item struct {
	Date  string
	Names []string
}

var d *dynamodb.DynamoDB
var once sync.Once
var Table = "Birthdays"
var Key = "Date"

func GetSession() *dynamodb.DynamoDB {
	once.Do(func() {
		s := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		d = dynamodb.New(s, &aws.Config{
			Endpoint: aws.String(config.AwsEndpoint),
			Region:   aws.String(config.AwsRegion),
		})
	})
	return d
}

func GetKey(key string) Item {
	ddb := GetSession()

	result, err := ddb.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			Key: {
				S: aws.String(key),
			},
		},
		TableName: aws.String(Table),
	})
	if err != nil {
		log.Println(fmt.Sprintf("Unable to get key %s from DynamoDB. Get item failed.%s", key, err))
		return emptyItem(key)
	}

	if result.Item == nil {
		return emptyItem(key)
	} else {
		item := Item{}
		err = dynamodbattribute.UnmarshalMap(result.Item, &item)
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
