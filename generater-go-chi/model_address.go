// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Address is an object.
type Address struct {
	// City:
	City string `json:"city,omitempty"`
	// State:
	State string `json:"state,omitempty"`
	// Street:
	Street string `json:"street,omitempty"`
	// Zip:
	Zip string `json:"zip,omitempty"`
}

// Validate implements basic validation for this model
func (m Address) Validate() error {
	return validation.Errors{}.Filter()
}

// GetCity returns the City property
func (m Address) GetCity() string {
	return m.City
}

// SetCity sets the City property
func (m *Address) SetCity(val string) {
	m.City = val
}

// GetState returns the State property
func (m Address) GetState() string {
	return m.State
}

// SetState sets the State property
func (m *Address) SetState(val string) {
	m.State = val
}

// GetStreet returns the Street property
func (m Address) GetStreet() string {
	return m.Street
}

// SetStreet sets the Street property
func (m *Address) SetStreet(val string) {
	m.Street = val
}

// GetZip returns the Zip property
func (m Address) GetZip() string {
	return m.Zip
}

// SetZip sets the Zip property
func (m *Address) SetZip(val string) {
	m.Zip = val
}
