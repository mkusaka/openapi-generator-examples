// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// UpdateUserQueryParameters is an object.
type UpdateUserQueryParameters struct {
	// Username: name that need to be deleted
	Username string `json:"username"`
}

// Validate implements basic validation for this model
func (m UpdateUserQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetUsername returns the Username property
func (m UpdateUserQueryParameters) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *UpdateUserQueryParameters) SetUsername(val string) {
	m.Username = val
}
