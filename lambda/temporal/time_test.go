package temporal

import (
	"reflect"
	"testing"
	"time"
)

func TestGetLookaheadDays(t *testing.T) {

	times := time.Date(2021, 1, 1, 0, 1, 0, 0, &time.Location{})

	dates := GetLookaheadDays(times, "7, 14")
	if !reflect.DeepEqual(dates, []string{"08/01", "15/01"}) {
		t.Fatalf("Expected %v, got %v", []string{"08/01", "15/01"}, dates)
	}
}

func TestConvertLookaheadDays(t *testing.T) {

	testData := map[string][]int{
		"":      {},
		"1,2":   {1, 2},
		"10,20": {10, 20},
		"30,30": {30, 30},
	}

	for csvString, expectedDays := range testData {
		resultsDays := convertLookaheadDays(csvString)

		if !reflect.DeepEqual(resultsDays, expectedDays) {
			t.Fatalf("Expected %v, got %v", expectedDays, resultsDays)
		}
	}
}

func TestGetDateString(t *testing.T) {

	times := map[string]time.Time{
		"01/01": time.Date(2021, 1, 1, 0, 1, 0, 0, &time.Location{}),
		"10/01": time.Date(2021, 1, 10, 0, 1, 0, 0, &time.Location{}),
		"01/10": time.Date(2021, 10, 1, 0, 1, 0, 0, &time.Location{}),
		"10/10": time.Date(2021, 10, 10, 0, 1, 0, 0, &time.Location{}),
	}

	for expectedDate, testTime := range times {
		resultDate := getDateString(testTime)

		if resultDate != expectedDate {
			t.Fatalf("Expected %s, got %s", expectedDate, resultDate)
		}
	}
}
