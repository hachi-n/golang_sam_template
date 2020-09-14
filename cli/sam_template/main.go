package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler() error {
	fmt.Println("Hello World.")
	return nil
}

func main() {
	lambda.Start(handler)
}
