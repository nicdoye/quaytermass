// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// TokenPermission Description of a token permission
// swagger:model TokenPermission
type TokenPermission struct {

	// Role to use for the token
	// Required: true
	// Enum: [read write admin]
	Role *string `json:"role"`
}

// Validate validates this token permission
func (m *TokenPermission) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var tokenPermissionTypeRolePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["read","write","admin"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		tokenPermissionTypeRolePropEnum = append(tokenPermissionTypeRolePropEnum, v)
	}
}

const (

	// TokenPermissionRoleRead captures enum value "read"
	TokenPermissionRoleRead string = "read"

	// TokenPermissionRoleWrite captures enum value "write"
	TokenPermissionRoleWrite string = "write"

	// TokenPermissionRoleAdmin captures enum value "admin"
	TokenPermissionRoleAdmin string = "admin"
)

// prop value enum
func (m *TokenPermission) validateRoleEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, tokenPermissionTypeRolePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *TokenPermission) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	// value enum
	if err := m.validateRoleEnum("role", "body", *m.Role); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *TokenPermission) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TokenPermission) UnmarshalBinary(b []byte) error {
	var res TokenPermission
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
