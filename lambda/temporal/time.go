package temporal

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

// GetLookaheadDays returns an array containing all the days we will check for birthdays
func GetLookaheadDays(t time.Time, lookaheadDays string) []string {

	lookaheadDayValues := convertLookaheadDays(lookaheadDays)

	modifiedDays := make([]string, 0)
	for _, value := range lookaheadDayValues {
		modifiedDays = append(modifiedDays, getDateString(t.AddDate(0, 0, value)))
	}
	return modifiedDays
}

// convertLookaheadDays converts the lookahead days csv string into an array of integers
func convertLookaheadDays(lookaheadDays string) []int {
	lookaheadStrings := strings.Split(lookaheadDays, ",")
	if len(lookaheadStrings) == 0 {
		log.Println("Found no lookahead values, returning an empty array")
		return []int{}
	}

	lookaheadValues := make([]int, 0)
	for _, lookaheadString := range lookaheadStrings {
		if v, err := strconv.Atoi(strings.TrimSpace(lookaheadString)); err != nil {
			log.Println(fmt.Sprintf("Unable to parse a lookahead value %s", lookaheadString))
			return []int{}
		} else {
			lookaheadValues = append(lookaheadValues, v)
		}
	}
	return lookaheadValues
}

// getDateString Gets a date string in the format DD/MM. This is the format used as the hash key in the dynamo db table
func getDateString(t time.Time) string {
	day := t.Day()
	month := t.Month()

	return fmt.Sprintf("%02d/%02d", day, month)
}
