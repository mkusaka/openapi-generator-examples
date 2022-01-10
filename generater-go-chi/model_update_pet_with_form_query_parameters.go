// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// UpdatePetWithFormQueryParameters is an object.
type UpdatePetWithFormQueryParameters struct {
	// PetId: ID of pet that needs to be updated
	PetId int64 `json:"petId"`
	// Name: Name of pet that needs to be updated
	Name string `json:"name,omitempty"`
	// Status: Status of pet that needs to be updated
	Status string `json:"status,omitempty"`
}

// Validate implements basic validation for this model
func (m UpdatePetWithFormQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPetId returns the PetId property
func (m UpdatePetWithFormQueryParameters) GetPetId() int64 {
	return m.PetId
}

// SetPetId sets the PetId property
func (m *UpdatePetWithFormQueryParameters) SetPetId(val int64) {
	m.PetId = val
}

// GetName returns the Name property
func (m UpdatePetWithFormQueryParameters) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *UpdatePetWithFormQueryParameters) SetName(val string) {
	m.Name = val
}

// GetStatus returns the Status property
func (m UpdatePetWithFormQueryParameters) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *UpdatePetWithFormQueryParameters) SetStatus(val string) {
	m.Status = val
}
