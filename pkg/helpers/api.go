package helpers

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

// Response will marshal the response into json
func Response(data interface{}, code int, err error) (events.APIGatewayProxyResponse, error) {
	body, jerr := json.Marshal(data)
	if jerr != nil {
		// TODO:  Explode!!!
	}
	return events.APIGatewayProxyResponse{
		Body:       string(body),
		StatusCode: code,
	}, err
}

// GetErrorData will wrap an error so you can embed an err in a response
func GetErrorData(err error) map[string]string {
	return map[string]string{
		"err": err.Error(),
	}
}

// ParseBody takes the body from the request, parses the json to a given struct pointer
func ParseBody(request events.APIGatewayProxyRequest, castTo interface{}) error {
	return json.Unmarshal([]byte(request.Body), &castTo)
}
