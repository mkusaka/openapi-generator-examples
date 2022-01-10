// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// FindPetsByTagsQueryParameters is an object.
type FindPetsByTagsQueryParameters struct {
	// Tags: Tags to filter by
	Tags []string `json:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m FindPetsByTagsQueryParameters) Validate() error {
	return validation.Errors{
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetTags returns the Tags property
func (m FindPetsByTagsQueryParameters) GetTags() []string {
	return m.Tags
}

// SetTags sets the Tags property
func (m *FindPetsByTagsQueryParameters) SetTags(val []string) {
	m.Tags = val
}
