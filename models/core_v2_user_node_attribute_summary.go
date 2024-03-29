// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CoreV2UserNodeAttributeSummary Request node fiel statistics.
// swagger:model core.v2.UserNodeAttributeSummary
type CoreV2UserNodeAttributeSummary map[string][]CoreV2UserNodeAttributeSummaryItems0

// Validate validates this core v2 user node attribute summary
func (m CoreV2UserNodeAttributeSummary) Validate(formats strfmt.Registry) error {
	var res []error

	for k := range m {

		if err := validate.Required(k, "body", m[k]); err != nil {
			return err
		}

		for i := 0; i < len(m[k]); i++ {

			if err := m[k][i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName(k + "." + strconv.Itoa(i))
				}
				return err
			}

		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// CoreV2UserNodeAttributeSummaryItems0 core v2 user node attribute summary items0
// swagger:model CoreV2UserNodeAttributeSummaryItems0
type CoreV2UserNodeAttributeSummaryItems0 struct {

	// id
	ID string `json:"_id,omitempty"`

	// sum
	Sum int64 `json:"sum,omitempty"`
}

// Validate validates this core v2 user node attribute summary items0
func (m *CoreV2UserNodeAttributeSummaryItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2UserNodeAttributeSummaryItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2UserNodeAttributeSummaryItems0) UnmarshalBinary(b []byte) error {
	var res CoreV2UserNodeAttributeSummaryItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
