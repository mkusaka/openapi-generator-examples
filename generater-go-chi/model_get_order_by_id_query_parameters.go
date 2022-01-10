// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// GetOrderByIdQueryParameters is an object.
type GetOrderByIdQueryParameters struct {
	// OrderId: ID of order that needs to be fetched
	OrderId int64 `json:"orderId"`
}

// Validate implements basic validation for this model
func (m GetOrderByIdQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetOrderId returns the OrderId property
func (m GetOrderByIdQueryParameters) GetOrderId() int64 {
	return m.OrderId
}

// SetOrderId sets the OrderId property
func (m *GetOrderByIdQueryParameters) SetOrderId(val int64) {
	m.OrderId = val
}
