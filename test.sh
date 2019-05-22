#!/bin/bash

STACK_NAME=test-1

aws cloudformation create-stack \
  --stack-name $STACK_NAME \
  --template-body file://test.cfn.yml \
  --capabilities CAPABILITY_NAMED_IAM \

