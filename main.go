package main

import (
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	// set lambda handler function
	lambda.Start(handleRequest)
}
