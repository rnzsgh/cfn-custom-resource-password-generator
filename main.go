package main

import (
	"context"
	"flag"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-lambda-go/lambdacontext"
	log "github.com/golang/glog"
	"github.com/sethvargo/go-password/password"
)

func init() {
	flag.Parse()
	flag.Lookup("logtostderr").Value.Set("true")
}

func main() {
	lambda.Start(cfn.LambdaWrap(handler))
}

func handler(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {

	event.ResourceProperties["PhysicalResourceID"] = lambdacontext.LogStreamName

	data = map[string]interface{}{}

	var p string
	if event.RequestType == "Create" {
		if p, err = password.Generate(39, 10, 0, false, true); err != nil {
			log.Errorf("Failed to generate password - reason: %v", err)
		} else {
			data["Password"] = p
		}
	}

	return
}
