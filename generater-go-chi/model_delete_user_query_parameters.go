// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// DeleteUserQueryParameters is an object.
type DeleteUserQueryParameters struct {
	// Username: The name that needs to be deleted
	Username string `json:"username"`
}

// Validate implements basic validation for this model
func (m DeleteUserQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetUsername returns the Username property
func (m DeleteUserQueryParameters) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *DeleteUserQueryParameters) SetUsername(val string) {
	m.Username = val
}
