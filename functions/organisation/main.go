package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/keithballdotnet/aws-serverless/functions/organisation/model"
	"github.com/keithballdotnet/aws-serverless/functions/organisation/store"
	uuid "github.com/satori/go.uuid"
)

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// Add a new org
	newOrganisation := &model.Organisation{OrganisationID: fmt.Sprintf("orgid:%s", uuid.NewV4().String()), Name: fmt.Sprintf("TS%s", time.Now().Format(time.RFC3339Nano))}
	err := orgStore.Store(newOrganisation)
	if err != nil {
		log.Printf("unable to store organisation: %v", err)
		return events.APIGatewayProxyResponse{}, err
	}

	orgs, err := orgStore.List()
	if err != nil {
		log.Printf("unable to get organisations: %v\n", err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusInternalServerError,
		}, err
	}

	for _, org := range orgs {
		log.Printf("Found org: %s - %s\n", org.OrganisationID, org.Name)
		//if org.OrganisationID ==
	}

	id := request.QueryStringParameters["id"]
	if id == "" {
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusBadRequest,
			Body:       "you must give an id",
		}, nil
	}

	organisation, err := orgStore.Get(id)
	if err != nil {
		log.Printf("unable to get organisation: %s : %v\n", id, err)
		return events.APIGatewayProxyResponse{
			StatusCode: http.StatusNotFound,
			Body:       fmt.Sprintf("'%s' not found", id),
		}, nil
	}

	return events.APIGatewayProxyResponse{
		Body:       "Got " + organisation.Name,
		StatusCode: http.StatusOK,
	}, nil
}

var orgStore store.OrganisationStore

func main() {
	log.Println("Starting organisation...")

	var err error
	orgStore, err = store.CreateOrganisationStore()
	if err != nil {
		log.Panicf("unable to start function, store creation failed: %v", err)
	}

	log.Println("Running lambda...")
	lambda.Start(Handler)
}
