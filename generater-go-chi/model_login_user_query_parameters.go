// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// LoginUserQueryParameters is an object.
type LoginUserQueryParameters struct {
	// Username: The user name for login
	Username string `json:"username,omitempty"`
	// Password: The password for login in clear text
	Password string `json:"password,omitempty"`
}

// Validate implements basic validation for this model
func (m LoginUserQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetUsername returns the Username property
func (m LoginUserQueryParameters) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *LoginUserQueryParameters) SetUsername(val string) {
	m.Username = val
}

// GetPassword returns the Password property
func (m LoginUserQueryParameters) GetPassword() string {
	return m.Password
}

// SetPassword sets the Password property
func (m *LoginUserQueryParameters) SetPassword(val string) {
	m.Password = val
}
