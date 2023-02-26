#!/bin/bash
cd ~/aws-lambda-return-from-json/cmd
GOOS=linux CGO_ENABLED=0 go build ~/aws-lambda-return-from-json/cmd/main.go
sleep 5
[ -e ~/aws-lambda-return-from-json/cmd/main.zip ] && rm ~/aws-lambda-return-from-json/cmd/main.zip
zip -jrm ~/aws-lambda-return-from-json/cmd/main.zip ~/aws-lambda-return-from-json/cmd/main && echo "INFO: Arquivo \"main.zip\" criado em ~/aws-lambda-return-from-json/cmd"