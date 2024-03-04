// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Error error
//
// swagger:model Error
type Error struct {

	// codigo
	// Required: true
	Codigo *int32 `json:"codigo"`

	// mensaje
	// Required: true
	Mensaje *string `json:"mensaje"`
}

// Validate validates this error
func (m *Error) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCodigo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMensaje(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Error) validateCodigo(formats strfmt.Registry) error {

	if err := validate.Required("codigo", "body", m.Codigo); err != nil {
		return err
	}

	return nil
}

func (m *Error) validateMensaje(formats strfmt.Registry) error {

	if err := validate.Required("mensaje", "body", m.Mensaje); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Error) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Error) UnmarshalBinary(b []byte) error {
	var res Error
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
