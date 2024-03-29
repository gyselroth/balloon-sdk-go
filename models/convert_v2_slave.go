// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// ConvertV2Slave A convert slave
// swagger:model convert.v2.Slave
type ConvertV2Slave struct {
	CoreV2Resource

	// Slave format.
	Format string `json:"format,omitempty"`

	// master
	Master *ConvertV2SlaveAO1Master `json:"master,omitempty"`

	// slave
	Slave *ConvertV2SlaveAO1Slave `json:"slave,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *ConvertV2Slave) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CoreV2Resource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CoreV2Resource = aO0

	// AO1
	var dataAO1 struct {
		Format string `json:"format,omitempty"`

		Master *ConvertV2SlaveAO1Master `json:"master,omitempty"`

		Slave *ConvertV2SlaveAO1Slave `json:"slave,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Format = dataAO1.Format

	m.Master = dataAO1.Master

	m.Slave = dataAO1.Slave

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m ConvertV2Slave) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CoreV2Resource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Format string `json:"format,omitempty"`

		Master *ConvertV2SlaveAO1Master `json:"master,omitempty"`

		Slave *ConvertV2SlaveAO1Slave `json:"slave,omitempty"`
	}

	dataAO1.Format = m.Format

	dataAO1.Master = m.Master

	dataAO1.Slave = m.Slave

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this convert v2 slave
func (m *ConvertV2Slave) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CoreV2Resource
	if err := m.CoreV2Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMaster(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSlave(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConvertV2Slave) validateMaster(formats strfmt.Registry) error {

	if swag.IsZero(m.Master) { // not required
		return nil
	}

	if m.Master != nil {
		if err := m.Master.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("master")
			}
			return err
		}
	}

	return nil
}

func (m *ConvertV2Slave) validateSlave(formats strfmt.Registry) error {

	if swag.IsZero(m.Slave) { // not required
		return nil
	}

	if m.Slave != nil {
		if err := m.Slave.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("slave")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConvertV2Slave) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConvertV2Slave) UnmarshalBinary(b []byte) error {
	var res ConvertV2Slave
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConvertV2SlaveAO1Master Master node from which the slave is from.
// swagger:model ConvertV2SlaveAO1Master
type ConvertV2SlaveAO1Master struct {

	// The id of the node.
	ID string `json:"id,omitempty"`

	// The name of the node.
	Name string `json:"name,omitempty"`
}

// Validate validates this convert v2 slave a o1 master
func (m *ConvertV2SlaveAO1Master) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConvertV2SlaveAO1Master) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConvertV2SlaveAO1Master) UnmarshalBinary(b []byte) error {
	var res ConvertV2SlaveAO1Master
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// ConvertV2SlaveAO1Slave The generated slave node. As long as the slave has not yet been created this is null.
// swagger:model ConvertV2SlaveAO1Slave
type ConvertV2SlaveAO1Slave struct {

	// The id of the node.
	ID string `json:"id,omitempty"`

	// The name of the node.
	Name string `json:"name,omitempty"`
}

// Validate validates this convert v2 slave a o1 slave
func (m *ConvertV2SlaveAO1Slave) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ConvertV2SlaveAO1Slave) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConvertV2SlaveAO1Slave) UnmarshalBinary(b []byte) error {
	var res ConvertV2SlaveAO1Slave
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
