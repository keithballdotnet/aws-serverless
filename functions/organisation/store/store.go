package store

import (
	"github.com/keithballdotnet/aws-serverless/functions/organisation/model"
	"github.com/keithballdotnet/aws-serverless/pkg/storage"
	"github.com/pkg/errors"
)

// OrganisationStore is an interface for storing organisations
type OrganisationStore interface {
	// List gets a collection of resources
	List() ([]*model.Organisation, error)
	// Store an item
	Store(*model.Organisation) error
	// Get will return an item
	Get(key string) (*model.Organisation, error)
}

var instance OrganisationStore

type store struct {
	db *storage.DynamoDB
}

// CreateOrganisationStore will create an instance of the organisation store
func CreateOrganisationStore() (OrganisationStore, error) {
	if instance != nil {
		return instance, nil
	}

	// Create a connection
	conn, err := storage.CreateConnection("eu-central-1")
	if err != nil {
		return nil, errors.Wrap(err, "unable to create storage connection")
	}

	// Create a new Dynamodb Table instance
	db := storage.NewDynamoDB(conn, "organisations")

	instance := &store{db: db}

	return instance, nil
}

// List gets a collection
func (s *store) List() ([]*model.Organisation, error) {
	var list []*model.Organisation

	err := s.db.List(&list)
	if err != nil {
		return nil, errors.Wrap(err, "unable to list organisations")
	}

	return list, nil
}

// Store an item
func (s *store) Store(organisation *model.Organisation) error {
	err := s.db.Store(organisation)
	if err != nil {
		return errors.Wrap(err, "unable to store organisation")
	}

	return nil
}

// Get will return an item
func (s *store) Get(key string) (*model.Organisation, error) {
	var organisation model.Organisation
	err := s.db.Get("OrganisationID", key, &organisation)
	if err != nil {
		return nil, errors.Wrap(err, "unable to get organisation")
	}
	if organisation.Name == "" {
		return nil, errors.New("not found")
	}
	return &organisation, nil
}
