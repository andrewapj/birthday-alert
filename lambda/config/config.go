package config

import "os"

var AwsEndpoint = getEnvVariable("APP_AWS_ENDPOINT", "http://localhost:8000")
var AwsRegion = getEnvVariable("APP_AWS_REGION", "eu-west-1")
var DaysLookAhead = getEnvVariable("APP_DAYS_LOOKAHEAD", "7, 14")
var NotificationString = getEnvVariable("APP_NOTIFICATION_STRING", "It's %s's birthday on %s")
var SnsTopic = getEnvVariable("APP_SNS_TOPIC", "")

func getEnvVariable(key, def string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		return def
	}
}
