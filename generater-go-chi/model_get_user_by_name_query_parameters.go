// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// GetUserByNameQueryParameters is an object.
type GetUserByNameQueryParameters struct {
	// Username: The name that needs to be fetched. Use user1 for testing.
	Username string `json:"username"`
}

// Validate implements basic validation for this model
func (m GetUserByNameQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetUsername returns the Username property
func (m GetUserByNameQueryParameters) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *GetUserByNameQueryParameters) SetUsername(val string) {
	m.Username = val
}
