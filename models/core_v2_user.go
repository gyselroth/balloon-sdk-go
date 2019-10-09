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

// CoreV2User A user is used to access the server.
// swagger:model core.v2.User
type CoreV2User struct {
	CoreV2Resource

	// Admin user flag.
	Admin *bool `json:"admin,omitempty"`

	// Authentication source.
	// Enum: [internal external]
	Auth *string `json:"auth,omitempty"`

	// Available storage in bytes. Is -1 if there is no hard_quota.
	Available *float64 `json:"available,omitempty"`

	// ISO 8601 timestamp when the resource was changed.
	// Format: date-time
	Changed strfmt.DateTime `json:"changed,omitempty"`

	// Hard quota in bytes (Max. limit of storage usage). Note that external storage is not part of the quota. The default is no limit.
	HardQuota *float64 `json:"hard_quota,omitempty"`

	// Is true if the user has a local password set.
	HasPassword *bool `json:"has_password,omitempty"`

	// User locale.
	Locale *string `json:"locale,omitempty"`

	// Display name (Usuallly same as username).
	Name string `json:"name,omitempty"`

	// User namespace.
	Namespace string `json:"namespace,omitempty"`

	// Soft quota in bytes (Warning of high quota usage). The default is no limit.
	SoftQuota *float64 `json:"soft_quota,omitempty"`

	// Used storage in bytes.
	Used float64 `json:"used,omitempty"`

	// Unique username.
	Username string `json:"username,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CoreV2User) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CoreV2Resource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CoreV2Resource = aO0

	// AO1
	var dataAO1 struct {
		Admin *bool `json:"admin,omitempty"`

		Auth *string `json:"auth,omitempty"`

		Available *float64 `json:"available,omitempty"`

		Changed strfmt.DateTime `json:"changed,omitempty"`

		HardQuota *float64 `json:"hard_quota,omitempty"`

		HasPassword *bool `json:"has_password,omitempty"`

		Locale *string `json:"locale,omitempty"`

		Name string `json:"name,omitempty"`

		Namespace string `json:"namespace,omitempty"`

		SoftQuota *float64 `json:"soft_quota,omitempty"`

		Used float64 `json:"used,omitempty"`

		Username string `json:"username,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Admin = dataAO1.Admin

	m.Auth = dataAO1.Auth

	m.Available = dataAO1.Available

	m.Changed = dataAO1.Changed

	m.HardQuota = dataAO1.HardQuota

	m.HasPassword = dataAO1.HasPassword

	m.Locale = dataAO1.Locale

	m.Name = dataAO1.Name

	m.Namespace = dataAO1.Namespace

	m.SoftQuota = dataAO1.SoftQuota

	m.Used = dataAO1.Used

	m.Username = dataAO1.Username

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CoreV2User) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CoreV2Resource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Admin *bool `json:"admin,omitempty"`

		Auth *string `json:"auth,omitempty"`

		Available *float64 `json:"available,omitempty"`

		Changed strfmt.DateTime `json:"changed,omitempty"`

		HardQuota *float64 `json:"hard_quota,omitempty"`

		HasPassword *bool `json:"has_password,omitempty"`

		Locale *string `json:"locale,omitempty"`

		Name string `json:"name,omitempty"`

		Namespace string `json:"namespace,omitempty"`

		SoftQuota *float64 `json:"soft_quota,omitempty"`

		Used float64 `json:"used,omitempty"`

		Username string `json:"username,omitempty"`
	}

	dataAO1.Admin = m.Admin

	dataAO1.Auth = m.Auth

	dataAO1.Available = m.Available

	dataAO1.Changed = m.Changed

	dataAO1.HardQuota = m.HardQuota

	dataAO1.HasPassword = m.HasPassword

	dataAO1.Locale = m.Locale

	dataAO1.Name = m.Name

	dataAO1.Namespace = m.Namespace

	dataAO1.SoftQuota = m.SoftQuota

	dataAO1.Used = m.Used

	dataAO1.Username = m.Username

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this core v2 user
func (m *CoreV2User) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CoreV2Resource
	if err := m.CoreV2Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAuth(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanged(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var coreV2UserTypeAuthPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["internal","external"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		coreV2UserTypeAuthPropEnum = append(coreV2UserTypeAuthPropEnum, v)
	}
}

// property enum
func (m *CoreV2User) validateAuthEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, coreV2UserTypeAuthPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CoreV2User) validateAuth(formats strfmt.Registry) error {

	if swag.IsZero(m.Auth) { // not required
		return nil
	}

	// value enum
	if err := m.validateAuthEnum("auth", "body", *m.Auth); err != nil {
		return err
	}

	return nil
}

func (m *CoreV2User) validateChanged(formats strfmt.Registry) error {

	if swag.IsZero(m.Changed) { // not required
		return nil
	}

	if err := validate.FormatOf("changed", "body", "date-time", m.Changed.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2User) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2User) UnmarshalBinary(b []byte) error {
	var res CoreV2User
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}