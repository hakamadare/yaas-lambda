package main

import (
  "context"
  "log"

  "github.com/hakamadare/yaas-go/yaas"
  "github.com/aws/aws-lambda-go/events"
  "github.com/aws/aws-lambda-go/lambda"
)

func main() {
  lambda.Start(Handler)
}

func Handler(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
  log.Print("request ID: %s", request.RequestContext.RequestID)

  ys,e := yaas.Yes()

  if (e == nil) {
    return events.APIGatewayProxyResponse{Body: string(ys), StatusCode: 200}, e
  } else {
    return events.APIGatewayProxyResponse{Body: e.Error(), StatusCode: 500}, e
  }
}
