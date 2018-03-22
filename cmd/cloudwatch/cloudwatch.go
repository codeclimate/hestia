package main

import (
	"context"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/bugsnag/bugsnag-go"
	"github.com/codeclimate/hestia/internal/config"
	"github.com/codeclimate/hestia/internal/notifiers"
	"log"
	"os"
)

func init() {
	api_key := config.Fetch("bugsnag_api_key")

	bugsnag.Configure(bugsnag.Configuration{
		APIKey:          api_key,
		ReleaseStage:    os.Getenv("BUGSNAG_RELEASE_STAGE"),
		ProjectPackages: []string{"github.com/codeclimate/hestia"},
		Synchronous:     true,
	})
}

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, event events.CloudWatchEvent) (string, error) {
	defer bugsnag.AutoNotify()

	log.Println("Processing CloudWatch event request")

	notifier := notifiers.Slack{Channel: "#slacktest"}
	notifier.Log("test")

	return "done", nil
}
