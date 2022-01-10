// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// DeletePetQueryParameters is an object.
type DeletePetQueryParameters struct {
	// ApiKey:
	ApiKey string `json:"api_key,omitempty"`
	// PetId: Pet id to delete
	PetId int64 `json:"petId"`
}

// Validate implements basic validation for this model
func (m DeletePetQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetApiKey returns the ApiKey property
func (m DeletePetQueryParameters) GetApiKey() string {
	return m.ApiKey
}

// SetApiKey sets the ApiKey property
func (m *DeletePetQueryParameters) SetApiKey(val string) {
	m.ApiKey = val
}

// GetPetId returns the PetId property
func (m DeletePetQueryParameters) GetPetId() int64 {
	return m.PetId
}

// SetPetId sets the PetId property
func (m *DeletePetQueryParameters) SetPetId(val int64) {
	m.PetId = val
}
