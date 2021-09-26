package main

import (
	"errors"
	"fmt"
	"github.com/andrewapj/birthday-alert-cdk/lambda/config"
	"github.com/andrewapj/birthday-alert-cdk/lambda/database"
	"github.com/andrewapj/birthday-alert-cdk/lambda/temporal"
	"github.com/aws/aws-lambda-go/lambda"
	"log"
	"time"
)

func HandleRequest() error {
	log.Println("Lambda started")

	messages, err := GetBirthdayMessages(time.Now(), config.DaysLookAhead)
	if err != nil {
		log.Println("Error getting new birthday messages")
		return err
	}

	if len(messages) > 0 {
		log.Println("Generated the following birthday messages")
		log.Println(fmt.Sprintf("%v", messages))
	}

	return nil
}

func GetBirthdayMessages(t time.Time, daysLookahead string) ([]string, error) {
	log.Println(fmt.Sprintf("Using the following look ahead values. %s", config.DaysLookAhead))
	keys := temporal.GetLookaheadDays(t, daysLookahead)
	if len(keys) == 0 {
		log.Println("Found no dates to search for")
		return nil, errors.New("no dates found")
	}

	messages := make([]string, 0)
	for _, date := range keys {
		log.Println(fmt.Sprintf("Checking for birthdays on %s", date))
		item := database.GetKey(date)

		if len(item.Names) > 0 {
			log.Println(fmt.Sprintf("Found birthdays for %v on %s", item.Names, date))
			for _, name := range item.Names {
				messages = append(messages, fmt.Sprintf(config.NotificationString, name, date))
			}
		} else {
			log.Println(fmt.Sprintf("Found no birthdays on %s", date))
		}
	}
	return messages, nil
}

func main() {
	lambda.Start(HandleRequest)
}
