#!/bin/bash
GOOS=linux CGO_ENABLED=0 go build ~/aws-lambda-return-from-json/cmd/main.go
sleep 5
zip -jrm ~/aws-lambda-return-from-json/cmd/main.zip ~/aws-lambda-return-from-json/cmd/main
echo "INFO: Arquivo \"main.zip\" criado em ~/aws-lambda-return-from-json/cmd"