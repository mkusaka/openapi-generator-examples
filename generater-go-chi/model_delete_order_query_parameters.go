// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// DeleteOrderQueryParameters is an object.
type DeleteOrderQueryParameters struct {
	// OrderId: ID of the order that needs to be deleted
	OrderId int64 `json:"orderId"`
}

// Validate implements basic validation for this model
func (m DeleteOrderQueryParameters) Validate() error {
	return validation.Errors{}.Filter()
}

// GetOrderId returns the OrderId property
func (m DeleteOrderQueryParameters) GetOrderId() int64 {
	return m.OrderId
}

// SetOrderId sets the OrderId property
func (m *DeleteOrderQueryParameters) SetOrderId(val int64) {
	m.OrderId = val
}
