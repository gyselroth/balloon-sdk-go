// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CoreV2Link core v2 link
// swagger:model core.v2.Link
type CoreV2Link struct {

	// href
	Href string `json:"href,omitempty"`
}

// Validate validates this core v2 link
func (m *CoreV2Link) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2Link) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2Link) UnmarshalBinary(b []byte) error {
	var res CoreV2Link
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
