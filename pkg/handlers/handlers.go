package handlers

import (
	"handlers/pkg/aluno"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
)

var (
	ErrorMethodNotAllowed  = "method not allowed"
	ErrorInvalidAccessData = "access code is invalid"
)

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func BuscaAluno(req events.LambdaFunctionURLRequest) (
	*events.LambdaFunctionURLResponse, error,
) {

	acesso := req.QueryStringParameters["acesso"]
	if len(acesso) > 0 {
		result, err := aluno.BuscaAluno(acesso)
		if err != nil {
			return urlResponse(http.StatusNotFound, ErrorBody{aws.String(err.Error())})
		}
		return urlResponse(http.StatusOK, result)
	}
	return urlResponse(http.StatusBadRequest, ErrorInvalidAccessData)
}

func UnhandledMethod() (*events.LambdaFunctionURLResponse, error) {
	return urlResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
