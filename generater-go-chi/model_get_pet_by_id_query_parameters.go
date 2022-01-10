// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// GetPetByIdQueryParameters is an object.
type GetPetByIdQueryParameters struct {
	// PetId: ID of pet to return
	PetId int64 `json:"petId"`
}

// Validate implements basic validation for this model
func (m GetPetByIdQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPetId returns the PetId property
func (m GetPetByIdQueryParameters) GetPetId() int64 {
	return m.PetId
}

// SetPetId sets the PetId property
func (m *GetPetByIdQueryParameters) SetPetId(val int64) {
	m.PetId = val
}
