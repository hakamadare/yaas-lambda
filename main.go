package main

import (
  "github.com/hakamadare/yaas-go/yaas"
  "github.com/aws/aws-lambda-go/lambda"
)

func main() {
  lambda.Start(yaasHandler)
}

func yaasHandler() (string, error) {
  ys,e := yaas.Yes()

  return string(ys), e
}
