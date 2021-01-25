package main

import (
	"aws-billing-notify/handler"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	fmt.Println("start")
	lambda.Start(handler.Run)
}