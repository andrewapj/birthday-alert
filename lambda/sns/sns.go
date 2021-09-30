package sns

import (
	"fmt"
	"github.com/andrewapj/birthday-alert-cdk/lambda/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
	"log"
)

func PublishMessage(message string) {

	s := session.Must(session.NewSessionWithOptions(session.Options{}))
	client := sns.New(s, aws.NewConfig().WithRegion(config.AwsRegion))

	_, err := client.Publish(&sns.PublishInput{
		Message:  aws.String(message),
		Subject:  aws.String("New Upcoming Birthday"),
		TopicArn: aws.String(config.SnsTopic),
	})
	if err != nil {
		log.Println(fmt.Sprintf("Error publishing message to %s. %s", config.SnsTopic, err))
	}
}
