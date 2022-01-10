// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Customer is an object.
type Customer struct {
	// Address:
	Address []Address `json:"address,omitempty"`
	// Id:
	Id int64 `json:"id,omitempty"`
	// Username:
	Username string `json:"username,omitempty"`
}

// Validate implements basic validation for this model
func (m Customer) Validate() error {
	return validation.Errors{
		"address": validation.Validate(
			m.Address,
		),
	}.Filter()
}

// GetAddress returns the Address property
func (m Customer) GetAddress() []Address {
	return m.Address
}

// SetAddress sets the Address property
func (m *Customer) SetAddress(val []Address) {
	m.Address = val
}

// GetId returns the Id property
func (m Customer) GetId() int64 {
	return m.Id
}

// SetId sets the Id property
func (m *Customer) SetId(val int64) {
	m.Id = val
}

// GetUsername returns the Username property
func (m Customer) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *Customer) SetUsername(val string) {
	m.Username = val
}
