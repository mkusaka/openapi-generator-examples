// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// Pet is an object.
type Pet struct {
	// Category:
	Category Category `json:"category,omitempty"`
	// Id:
	Id int64 `json:"id,omitempty"`
	// Name:
	Name string `json:"name"`
	// PhotoUrls:
	PhotoUrls []string `json:"photoUrls"`
	// Status: pet status in the store
	Status string `json:"status,omitempty"`
	// Tags:
	Tags []Tag `json:"tags,omitempty"`
}

// Validate implements basic validation for this model
func (m Pet) Validate() error {
	return validation.Errors{
		"category": validation.Validate(
			m.Category,
		),
		"photoUrls": validation.Validate(
			m.PhotoUrls, validation.NotNil,
		),
		"tags": validation.Validate(
			m.Tags,
		),
	}.Filter()
}

// GetCategory returns the Category property
func (m Pet) GetCategory() Category {
	return m.Category
}

// SetCategory sets the Category property
func (m *Pet) SetCategory(val Category) {
	m.Category = val
}

// GetId returns the Id property
func (m Pet) GetId() int64 {
	return m.Id
}

// SetId sets the Id property
func (m *Pet) SetId(val int64) {
	m.Id = val
}

// GetName returns the Name property
func (m Pet) GetName() string {
	return m.Name
}

// SetName sets the Name property
func (m *Pet) SetName(val string) {
	m.Name = val
}

// GetPhotoUrls returns the PhotoUrls property
func (m Pet) GetPhotoUrls() []string {
	return m.PhotoUrls
}

// SetPhotoUrls sets the PhotoUrls property
func (m *Pet) SetPhotoUrls(val []string) {
	m.PhotoUrls = val
}

// GetStatus returns the Status property
func (m Pet) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *Pet) SetStatus(val string) {
	m.Status = val
}

// GetTags returns the Tags property
func (m Pet) GetTags() []Tag {
	return m.Tags
}

// SetTags sets the Tags property
func (m *Pet) SetTags(val []Tag) {
	m.Tags = val
}
