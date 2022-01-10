// This file is auto-generated, DO NOT EDIT.
//
// Source:
//     Title: Swagger Petstore - OpenAPI 3.0
//     Version: 1.0.9
package api

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

// User is an object.
type User struct {
	// Email:
	Email string `json:"email,omitempty"`
	// FirstName:
	FirstName string `json:"firstName,omitempty"`
	// Id:
	Id int64 `json:"id,omitempty"`
	// LastName:
	LastName string `json:"lastName,omitempty"`
	// Password:
	Password string `json:"password,omitempty"`
	// Phone:
	Phone string `json:"phone,omitempty"`
	// UserStatus: User Status
	UserStatus int32 `json:"userStatus,omitempty"`
	// Username:
	Username string `json:"username,omitempty"`
}

// Validate implements basic validation for this model
func (m User) Validate() error {
	return validation.Errors{}.Filter()
}

// GetEmail returns the Email property
func (m User) GetEmail() string {
	return m.Email
}

// SetEmail sets the Email property
func (m *User) SetEmail(val string) {
	m.Email = val
}

// GetFirstName returns the FirstName property
func (m User) GetFirstName() string {
	return m.FirstName
}

// SetFirstName sets the FirstName property
func (m *User) SetFirstName(val string) {
	m.FirstName = val
}

// GetId returns the Id property
func (m User) GetId() int64 {
	return m.Id
}

// SetId sets the Id property
func (m *User) SetId(val int64) {
	m.Id = val
}

// GetLastName returns the LastName property
func (m User) GetLastName() string {
	return m.LastName
}

// SetLastName sets the LastName property
func (m *User) SetLastName(val string) {
	m.LastName = val
}

// GetPassword returns the Password property
func (m User) GetPassword() string {
	return m.Password
}

// SetPassword sets the Password property
func (m *User) SetPassword(val string) {
	m.Password = val
}

// GetPhone returns the Phone property
func (m User) GetPhone() string {
	return m.Phone
}

// SetPhone sets the Phone property
func (m *User) SetPhone(val string) {
	m.Phone = val
}

// GetUserStatus returns the UserStatus property
func (m User) GetUserStatus() int32 {
	return m.UserStatus
}

// SetUserStatus sets the UserStatus property
func (m *User) SetUserStatus(val int32) {
	m.UserStatus = val
}

// GetUsername returns the Username property
func (m User) GetUsername() string {
	return m.Username
}

// SetUsername sets the Username property
func (m *User) SetUsername(val string) {
	m.Username = val
}
