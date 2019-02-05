package model

import "errors"

// Organisation is an organisation or company
type Organisation struct {
	OrganisationID string `json:"OrganisationID,omitempty"`
	Name           string `json:"Name,omitempty"`
}

// Validate will check if a model is ok to be used
func (m *Organisation) Validate() error {
	if m.Name == "" {
		return errors.New("you must give a name")
	}
	return nil
}
