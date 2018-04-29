package main

import (
  "context"
  "encoding/json"
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
    marshaled,err := json.Marshal(string(ys))

    if (err != nil) {
      return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, e
    }

    body := string(marshaled)

    return events.APIGatewayProxyResponse{Body: body, StatusCode: 200}, e
  } else {
    return events.APIGatewayProxyResponse{Body: e.Error(), StatusCode: 500}, e
  }
}
