package main

import (
	"net/http"
	"testing"

	"github.com/aws/aws-lambda-go/events"
	"github.com/stretchr/testify/assert"
)

func TestHandler(t *testing.T) {

	tests := []struct {
		request      events.APIGatewayProxyRequest
		expectBody   string
		expectStatus int
		err          error
	}{
		{
			request:      events.APIGatewayProxyRequest{Body: "Steve"},
			expectBody:   "Hello Steve",
			expectStatus: http.StatusOK,
			err:          nil,
		},
	}

	for _, test := range tests {
		response, err := Handler(test.request)
		assert.IsType(t, test.err, err)
		assert.Equal(t, test.expectBody, response.Body)
		assert.Equal(t, test.expectStatus, response.StatusCode)
	}

}
