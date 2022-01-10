// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// ApiResponse is an object.
type ApiResponse struct {
	// Code:
	Code int32 `json:"code,omitempty"`
	// Message:
	Message string `json:"message,omitempty"`
	// Type:
	Type string `json:"type,omitempty"`
}

// Validate implements basic validation for this model
func (m ApiResponse) Validate() error {
	return validation.Errors{}.Filter()
}

// GetCode returns the Code property
func (m ApiResponse) GetCode() int32 {
	return m.Code
}

// SetCode sets the Code property
func (m *ApiResponse) SetCode(val int32) {
	m.Code = val
}

// GetMessage returns the Message property
func (m ApiResponse) GetMessage() string {
	return m.Message
}

// SetMessage sets the Message property
func (m *ApiResponse) SetMessage(val string) {
	m.Message = val
}

// GetType returns the Type property
func (m ApiResponse) GetType() string {
	return m.Type
}

// SetType sets the Type property
func (m *ApiResponse) SetType(val string) {
	m.Type = val
}
