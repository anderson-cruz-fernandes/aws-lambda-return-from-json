package handlers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

func urlResponse(status int, body interface{}) (*events.LambdaFunctionURLResponse, error) {
	stringBody, _ := json.Marshal(body)
  resp := events.LambdaFunctionURLResponse{
		StatusCode:      status,
		Headers:         map[string]string{"Content-Type": "application/json"},
		Body:            string(stringBody),
		IsBase64Encoded: false,
		Cookies:         []string{},
	}
	return &resp, nil
}
