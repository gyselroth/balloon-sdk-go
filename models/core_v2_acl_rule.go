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

// CoreV2ACLRule ACL rule.
// swagger:model core.v2.AclRule
type CoreV2ACLRule struct {

	// ACL rules.
	ID string `json:"id,omitempty"`

	// Privilege.
	// Enum: [rw w m w+ d]
	Privilege string `json:"privilege,omitempty"`

	// role
	Role *CoreV2ACLRuleRole `json:"role,omitempty"`

	// The type of the resource.
	// Enum: [user group]
	Type string `json:"type,omitempty"`
}

// Validate validates this core v2 Acl rule
func (m *CoreV2ACLRule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrivilege(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var coreV2AclRuleTypePrivilegePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rw","w","m","w+","d"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		coreV2AclRuleTypePrivilegePropEnum = append(coreV2AclRuleTypePrivilegePropEnum, v)
	}
}

const (

	// CoreV2ACLRulePrivilegeRw captures enum value "rw"
	CoreV2ACLRulePrivilegeRw string = "rw"

	// CoreV2ACLRulePrivilegeW captures enum value "w"
	CoreV2ACLRulePrivilegeW string = "w"

	// CoreV2ACLRulePrivilegeM captures enum value "m"
	CoreV2ACLRulePrivilegeM string = "m"

	// CoreV2ACLRulePrivilegeW captures enum value "w+"
	CoreV2ACLRulePrivilegeW string = "w+"

	// CoreV2ACLRulePrivilegeD captures enum value "d"
	CoreV2ACLRulePrivilegeD string = "d"
)

// prop value enum
func (m *CoreV2ACLRule) validatePrivilegeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, coreV2AclRuleTypePrivilegePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CoreV2ACLRule) validatePrivilege(formats strfmt.Registry) error {

	if swag.IsZero(m.Privilege) { // not required
		return nil
	}

	// value enum
	if err := m.validatePrivilegeEnum("privilege", "body", m.Privilege); err != nil {
		return err
	}

	return nil
}

func (m *CoreV2ACLRule) validateRole(formats strfmt.Registry) error {

	if swag.IsZero(m.Role) { // not required
		return nil
	}

	if m.Role != nil {
		if err := m.Role.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("role")
			}
			return err
		}
	}

	return nil
}

var coreV2AclRuleTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["user","group"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		coreV2AclRuleTypeTypePropEnum = append(coreV2AclRuleTypeTypePropEnum, v)
	}
}

const (

	// CoreV2ACLRuleTypeUser captures enum value "user"
	CoreV2ACLRuleTypeUser string = "user"

	// CoreV2ACLRuleTypeGroup captures enum value "group"
	CoreV2ACLRuleTypeGroup string = "group"
)

// prop value enum
func (m *CoreV2ACLRule) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, coreV2AclRuleTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CoreV2ACLRule) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2ACLRule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2ACLRule) UnmarshalBinary(b []byte) error {
	var res CoreV2ACLRule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CoreV2ACLRuleRole Points to the share owner. If the node is not part of any share this is null.
// swagger:model CoreV2ACLRuleRole
type CoreV2ACLRuleRole struct {

	// The id of the role resource.
	ID string `json:"id,omitempty"`

	// The name of role resource.
	Name string `json:"name,omitempty"`
}

// Validate validates this core v2 ACL rule role
func (m *CoreV2ACLRuleRole) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2ACLRuleRole) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2ACLRuleRole) UnmarshalBinary(b []byte) error {
	var res CoreV2ACLRuleRole
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}