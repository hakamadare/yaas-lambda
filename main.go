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
  log.Printf("request ID: %s", request.RequestContext.RequestID)

  var ys yaas.YesString
  var err error

  // inspect the request path
  rpath := request.Path
  rbody := request.Body
  log.Printf("path: %s", rpath)
  log.Printf("body: %s", rbody)

  switch rpath {
  case "/yaas/invert":
    body,e := parseJsonBody(rbody)
    if e != nil {
      err = e
      log.Printf("error parsing JSON '%s': %s", body, err)
    } else {
      ys,err = yaas.Inverse(yaas.ParsedYes(body))
    }
  case "/yaas":
    ys,err = yaas.Yes()
  }

  if (err == nil) {
    marshaled,err := json.Marshal(string(ys))

    if (err != nil) {
      return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500}, err
    }

    body := string(marshaled)

    return events.APIGatewayProxyResponse{Body: body, StatusCode: 200, Headers: corsHeaders()}, err
  } else {
    return events.APIGatewayProxyResponse{Body: err.Error(), StatusCode: 500, Headers: corsHeaders()}, err
  }
}

func parseJsonBody(body string) (string, error) {
  var bs string
  err := json.Unmarshal([]byte(body), &bs)
  return bs,err
}

func corsHeaders() (map[string]string) {
  headers := map[string]string{
    "Access-Control-Allow-Methods": "GET",
    "Access-Control-Allow-Headers": "*",
    "Access-Control-Allow-Origin": "*",
  }

  return headers
}
