package config

import "os"

var AwsEndpoint = getEnvVariable("APP_AWS_ENDPOINT", "http://localhost:8000")
var AwsRegion = getEnvVariable("APP_AWS_REGION", "eu-west-1")

func getEnvVariable(key, def string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		return def
	}
}
