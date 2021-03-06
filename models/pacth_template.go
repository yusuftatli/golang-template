// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// PacthTemplate pacth template
//
// swagger:model PacthTemplate
type PacthTemplate struct {

	// created at
	CreatedAt string `json:"createdAt,omitempty"`

	// id
	ID uint64 `json:"id,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// size
	Size string `json:"size,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// updated at
	UpdatedAt string `json:"updatedAt,omitempty"`
}

// Validate validates this pacth template
func (m *PacthTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this pacth template based on context it is used
func (m *PacthTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *PacthTemplate) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PacthTemplate) UnmarshalBinary(b []byte) error {
	var res PacthTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
