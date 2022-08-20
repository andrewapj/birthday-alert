package main

import (
	"github.com/andrewapj/birthday-alert-cdk/lambda/database"
	"reflect"
	"sort"
	"testing"
	"time"
)

func TestGetBirthdayMessages(t *testing.T) {

	//Given: A database table
	ddb := database.GetClient()
	database.BuildTable(ddb, t)
	defer database.DeleteTable(ddb, t)

	//And: The current time
	theTime := time.Date(2021, 1, 1, 0, 1, 0, 0, &time.Location{})

	//And: The days to look ahead
	lookahead := "7, 14"

	//And: Two birthdays in the DB
	database.PutItem(t, ddb, database.Item{
		Date:  "08/01",
		Names: []string{"Bob", "Sue"},
	})

	//When: We get the messages
	messages, err := GetBirthdayMessages(theTime, lookahead)
	if err != nil {
		t.Fatalf("Unexpected error %s", err)
	}

	//Then: There should be the correct number
	if len(messages) != 2 {
		t.Fatalf("Expected 2 messages, got %d", len(messages))
	}

	//And: They should be correctly formatted
	sort.Strings(messages)
	expected := []string{"It's Bob's birthday on 08/01", "It's Sue's birthday on 08/01"}
	if !reflect.DeepEqual(messages, expected) {
		t.Fatalf("Messages are not equal. Expected %v, got %v", expected, messages)
	}
}
