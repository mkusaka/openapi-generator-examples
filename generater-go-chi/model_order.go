// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"time"
)

// Order is an object.
type Order struct {
	// Complete:
	Complete bool `json:"complete,omitempty"`
	// Id:
	Id int64 `json:"id,omitempty"`
	// PetId:
	PetId int64 `json:"petId,omitempty"`
	// Quantity:
	Quantity int32 `json:"quantity,omitempty"`
	// ShipDate:
	ShipDate time.Time `json:"shipDate,omitempty"`
	// Status: Order Status
	Status string `json:"status,omitempty"`
}

// Validate implements basic validation for this model
func (m Order) Validate() error {
	return validation.Errors{}.Filter()
}

// GetComplete returns the Complete property
func (m Order) GetComplete() bool {
	return m.Complete
}

// SetComplete sets the Complete property
func (m *Order) SetComplete(val bool) {
	m.Complete = val
}

// GetId returns the Id property
func (m Order) GetId() int64 {
	return m.Id
}

// SetId sets the Id property
func (m *Order) SetId(val int64) {
	m.Id = val
}

// GetPetId returns the PetId property
func (m Order) GetPetId() int64 {
	return m.PetId
}

// SetPetId sets the PetId property
func (m *Order) SetPetId(val int64) {
	m.PetId = val
}

// GetQuantity returns the Quantity property
func (m Order) GetQuantity() int32 {
	return m.Quantity
}

// SetQuantity sets the Quantity property
func (m *Order) SetQuantity(val int32) {
	m.Quantity = val
}

// GetShipDate returns the ShipDate property
func (m Order) GetShipDate() time.Time {
	return m.ShipDate
}

// SetShipDate sets the ShipDate property
func (m *Order) SetShipDate(val time.Time) {
	m.ShipDate = val
}

// GetStatus returns the Status property
func (m Order) GetStatus() string {
	return m.Status
}

// SetStatus sets the Status property
func (m *Order) SetStatus(val string) {
	m.Status = val
}
