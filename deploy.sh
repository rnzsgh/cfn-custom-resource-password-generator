#!/bin/bash

rm -Rf cfn-custom-resource-password-generator.zip main

BUCKET_NAME=public-aws-serverless-repo
GOOS=linux go build main.go

zip cfn-custom-resource-password-generator.zip ./main

aws s3 cp cfn-custom-resource-password-generator.zip s3://$BUCKET_NAME/cfn-custom-resource-password-generator.zip

aws s3api put-object-tagging --bucket $BUCKET_NAME --key cfn-custom-resource-password-generator.zip --tagging 'TagSet={Key=public,Value=yes}'
