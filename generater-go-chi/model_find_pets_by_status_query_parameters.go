// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// FindPetsByStatusQueryParameters is an object.
type FindPetsByStatusQueryParameters struct {
	// Status: Status values that need to be considered for filter
	Status string `json:"status,omitempty"`
}

// Validate implements basic validation for this model
func (m FindPetsByStatusQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetStatus returns the Status property
func (m FindPetsByStatusQueryParameters) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *FindPetsByStatusQueryParameters) SetStatus(val string) {
	m.Status = val
}
