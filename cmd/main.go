package main

import (
	"handlers/pkg/handlers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	lambda.Start(handler)
}

func handler(req events.LambdaFunctionURLRequest) (*events.LambdaFunctionURLResponse, error) {
  switch req.RequestContext.HTTP.Method {
    case "GET":    
      return handlers.BuscaAluno(req)
    default:
		  return handlers.UnhandledMethod()
  }
}

//#############################
//https://stackoverflow.com/a/72796824 #AWS Lambda Function URL
//http://www.inanzzz.com/index.php/post/1rwm/including-and-reading-static-files-with-embed-directive-at-compile-time-in-golang
