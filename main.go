package main

import (
	"context"
	"github/lubualo/gambit-user/awsgo"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/lubualo/gambit-user/awsgo"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, evemt events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AWSInit()
}
