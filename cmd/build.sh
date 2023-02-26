#!/bin/bash
[ -e ~/aws-lambda-return-from-json/cmd/main.zip ] && rm ~/aws-lambda-return-from-json/cmd/main.zip
GOOS=linux CGO_ENABLED=0 go build ~/aws-lambda-return-from-json/cmd/main.go
sleep 10
zip -jrm ~/aws-lambda-return-from-json/cmd/main.zip ~/aws-lambda-return-from-json/cmd/main
