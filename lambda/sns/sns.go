package sns

import (
	"context"
	"fmt"
	"github.com/andrewapj/birthday-alert-cdk/lambda/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sns"
	"log"
)

func PublishMessage(message string) {

	client := sns.NewFromConfig(config.GetAwsConfig())

	_, err := client.Publish(context.TODO(), &sns.PublishInput{
		Message:  aws.String(message),
		Subject:  aws.String(config.NotificationTitle),
		TopicArn: aws.String(config.SnsTopic),
	})

	if err != nil {
		log.Fatalf(fmt.Sprintf("Error publishing message to %s. %s", config.SnsTopic, err))
	}
}
