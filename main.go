package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/lubualo/gambit-user/awsgo"
	"github.com/lubualo/gambit-user/db"
	"github.com/lubualo/gambit-user/models"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.AWSInit()
	if !IsValid() {
		const errorMsg = "error in params: 'SecretName' missing"
		fmt.Println(errorMsg)
		err := errors.New(errorMsg)
		return event, err
	}

	var data models.SignUp
	for row, att := range event.Request.UserAttributes {
		switch row {
			case "email":
				data.UserEmail = att
				fmt.Println("Email = " + data.UserEmail)
			case "sub":
				data.UserUUID = att
				fmt.Println("Sub = " + data.UserUUID)
		}
	}

	err := db.ReadSecret()
	if err != nil {
		fmt.Println("Error while reading secret: " + err.Error())
		return event, err
	}


	err = db.SignUp(data)
	return event, err
}

func IsValid() bool {
	var hasParam bool
	_, hasParam = os.LookupEnv("SecretName")
	return hasParam
}
