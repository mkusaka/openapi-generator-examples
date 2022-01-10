// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// UploadFileQueryParameters is an object.
type UploadFileQueryParameters struct {
	// PetId: ID of pet to update
	PetId int64 `json:"petId"`
	// AdditionalMetadata: Additional Metadata
	AdditionalMetadata string `json:"additionalMetadata,omitempty"`
}

// Validate implements basic validation for this model
func (m UploadFileQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetPetId returns the PetId property
func (m UploadFileQueryParameters) GetPetId() int64 {
	return m.PetId
}

// SetPetId sets the PetId property
func (m *UploadFileQueryParameters) SetPetId(val int64) {
	m.PetId = val
}

// GetAdditionalMetadata returns the AdditionalMetadata property
func (m UploadFileQueryParameters) GetAdditionalMetadata() string {
	return m.AdditionalMetadata
}

// SetAdditionalMetadata sets the AdditionalMetadata property
func (m *UploadFileQueryParameters) SetAdditionalMetadata(val string) {
	m.AdditionalMetadata = val
}
