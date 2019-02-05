package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/keithballdotnet/aws-serverless/functions/organisation/model"
	"github.com/keithballdotnet/aws-serverless/functions/organisation/store"
	"github.com/keithballdotnet/aws-serverless/pkg/helpers"
	uuid "github.com/satori/go.uuid"
)

// ListHandler will return a collection of organisations
func ListHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("ListHandler() RequestID %s\n", request.RequestContext.RequestID)
	orgs, err := orgStore.List()
	if err != nil {
		log.Printf("unable to get organisations: %v\n", err)
		return helpers.Response(helpers.GetErrorData(err), http.StatusInternalServerError, nil)
	}

	return helpers.Response(orgs, http.StatusOK, nil)
}

// CreateHandler will handle the creation of a new organisation
func CreateHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("CreateHandler() RequestID %s\n", request.RequestContext.RequestID)
	// Add a new org
	newOrganisation := &model.Organisation{OrganisationID: fmt.Sprintf("orgid:%s", uuid.NewV4().String()), Name: fmt.Sprintf("TS%s", time.Now().Format(time.RFC3339Nano))}
	err := orgStore.Store(newOrganisation)
	if err != nil {
		log.Printf("unable to store organisation: %v", err)
		return helpers.Response(helpers.GetErrorData(err), http.StatusInternalServerError, nil)
	}
	return helpers.Response(newOrganisation, http.StatusCreated, nil)
}

// GetHandler will get a specific organisation
func GetHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	log.Printf("GetHandler() RequestID %s\n", request.RequestContext.RequestID)

	id := request.QueryStringParameters["id"]
	if id == "" {
		return helpers.Response(helpers.GetErrorData(errors.New("no id passed")), http.StatusBadRequest, nil)
	}

	organisation, err := orgStore.Get(id)
	if err != nil {
		log.Printf("unable to get organisation: %s : %v\n", id, err)
		return helpers.Response(helpers.GetErrorData(fmt.Errorf("'%s' not found", id)), http.StatusNotFound, nil)
	}

	return helpers.Response(organisation, http.StatusOK, nil)
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
	lambda.Start(Router())
}

// Router routes endpoints to the correct method
// GET without an ID in the QueryString, calls the List method,
// GET with an ID calls the Get method,
// POST calls the Store method.
func Router() func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return func(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
		switch request.HTTPMethod {
		case "GET":
			id := request.QueryStringParameters["id"]
			if id != "" {
				return GetHandler(request)
			}
			return ListHandler(request)
		case "POST":
			return CreateHandler(request)
		default:
			return helpers.Response(helpers.GetErrorData(errors.New("method not allowed")), http.StatusMethodNotAllowed, nil)
		}
	}
}
