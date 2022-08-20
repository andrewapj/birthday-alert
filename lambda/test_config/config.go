package test_config

import (
	"log"
	"os"
)

// SetAWSCredentials sets the AWS credentials for the tests, using environment variables
func SetAWSCredentials() {
	err := os.Setenv("AWS_ACCESS_KEY_ID", "key")
	if err != nil {
		log.Fatalln("Unable to set AWS_ACCESS_KEY_ID")
	}

	err = os.Setenv("AWS_SECRET_ACCESS_KEY", "key")
	if err != nil {
		log.Fatalln("Unable to set AWS_SECRET_ACCESS_KEY")
	}
}

// UnsetAWSCredentials unsets the AWS Credentials for the tests, using environment variables
func UnsetAWSCredentials() {
	err := os.Unsetenv("AWS_ACCESS_KEY_ID")
	if err != nil {
		log.Fatalln("Unable to unset AWS_ACCESS_KEY_ID")
	}

	err = os.Unsetenv("AWS_ACCESS_KEY_ID")
	if err != nil {
		log.Fatalln("Unable to unset AWS_ACCESS_KEY_ID")
	}
}
