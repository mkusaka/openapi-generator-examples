// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Category is an object.
type Category struct {
	// Id:
	Id int64 `json:"id,omitempty"`
	// Name:
	Name string `json:"name,omitempty"`
}

// Validate implements basic validation for this model
func (m Category) Validate() error {
	return validation.Errors{}.Filter()
}

// GetId returns the Id property
func (m Category) GetId() int64 {
	return m.Id
}

// SetId sets the Id property
func (m *Category) SetId(val int64) {
	m.Id = val
}

// GetName returns the Name property
func (m Category) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Category) SetName(val string) {
	m.Name = val
}
