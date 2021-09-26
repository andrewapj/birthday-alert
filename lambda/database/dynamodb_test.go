package database

import (
	"github.com/andrewapj/birthday-alert-cdk/lambda/test_helper"
	"reflect"
	"sort"
	"testing"
)

func TestBuildTable(t *testing.T) {

	ddb := GetSession()
	defer test_helper.DeleteTable(ddb, t)

	test_helper.BuildTable(ddb, t)
}

func TestGetKey(t *testing.T) {

	// Given: a session and a DB table
	ddb := GetSession()
	test_helper.BuildTable(ddb, t)
	defer test_helper.DeleteTable(ddb, t)

	// And: some test data in the DB
	item := test_helper.PutItem(t, ddb, Item{
		Date:  "01/12",
		Names: []string{"Person 1", "Person 2"},
	})

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
	test_helper.BuildTable(ddb, t)
	defer test_helper.DeleteTable(ddb, t)

	// When: We get the birthdays for 01/12
	results := GetKey("01/12")

	// Then: The names for that date should be empty
	if !reflect.DeepEqual(results.Names, []string{}) {
		t.Fatalf("Result data incorrect. Expected %s, got %s", []string{}, results.Names)
	}
}
