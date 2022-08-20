package config

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"log"
	"os"
	"sync"
)

var awsEndpoint = getEnvVariable("APP_AWS_ENDPOINT", "http://localhost:8000")
var awsRegion = getEnvVariable("APP_AWS_REGION", "eu-west-2")
var DaysLookAhead = getEnvVariable("APP_DAYS_LOOKAHEAD", "7, 14")
var NotificationTitle = getEnvVariable("APP_NOTIFICATION_TITLE", "New Upcoming Birthday")
var NotificationMessage = getEnvVariable("APP_NOTIFICATION_MESSAGE", "It's %s's birthday on %s")
var SnsTopic = getEnvVariable("APP_SNS_TOPIC", "")

var once sync.Once
var config aws.Config

func GetAwsConfig() aws.Config {
	once.Do(func() {
		endpointResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
			return aws.Endpoint{URL: awsEndpoint}, nil
		})

		cfg, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion(awsRegion),
			awsconfig.WithEndpointResolverWithOptions(endpointResolver))
		if err != nil {
			log.Fatalln(fmt.Sprintf("unable to initialise AWS config: %s", err))
		}
		config = cfg
	})

	return config
}

func getEnvVariable(key, def string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	} else {
		return def
	}
}
