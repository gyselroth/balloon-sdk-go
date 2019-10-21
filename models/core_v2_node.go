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

// CoreV2Node A Node.
// swagger:model core.v2.Node
type CoreV2Node struct {
	CoreV2Resource

	// Access level.
	// Enum: [rw w m rp]
	Access *string `json:"access,omitempty"`

	// ISO 8601 timestamp when the resource was changed.
	// Format: date-time
	Changed strfmt.DateTime `json:"changed,omitempty"`

	// ISO 8601 timestamp when the resource was deleted.
	// Format: date-time
	Deleted strfmt.DateTime `json:"deleted,omitempty"`

	// If node is of type Collection this flag is true.
	Directory *bool `json:"directory,omitempty"`

	// meta
	Meta *CoreV2NodeMeta `json:"meta,omitempty"`

	// Mimetype.
	Mime string `json:"mime,omitempty"`

	// Node name.
	Name string `json:"name,omitempty"`

	// parent
	Parent *CoreV2NodeAO1Parent `json:"parent,omitempty"`

	// The path abstraction of the node.
	Path string `json:"path,omitempty"`

	// Readonly only affects the content of the node but not metadata.
	Readonly *bool `json:"readonly,omitempty"`

	// share
	Share *CoreV2NodeAO1Share `json:"share,omitempty"`

	// Is true if the node has a protected public sharelink.
	SharelinkHasPassword *bool `json:"sharelink_has_password,omitempty"`

	// shareowner
	Shareowner *CoreV2NodeAO1Shareowner `json:"shareowner,omitempty"`

	// The size in bytes. If the node is of type collection the size is the number of child nodes.
	Size int64 `json:"size,omitempty"`

	// Is true if the user has an active subscription on this node.
	Subscription *bool `json:"subscription,omitempty"`

	// Subscription excludes actions from the subscribed user itself.
	SubscriptionExcludeMe *bool `json:"subscription_exclude_me,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *CoreV2Node) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CoreV2Resource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CoreV2Resource = aO0

	// AO1
	var dataAO1 struct {
		Access *string `json:"access,omitempty"`

		Changed strfmt.DateTime `json:"changed,omitempty"`

		Deleted strfmt.DateTime `json:"deleted,omitempty"`

		Directory *bool `json:"directory,omitempty"`

		Meta *CoreV2NodeMeta `json:"meta,omitempty"`

		Mime string `json:"mime,omitempty"`

		Name string `json:"name,omitempty"`

		Parent *CoreV2NodeAO1Parent `json:"parent,omitempty"`

		Path string `json:"path,omitempty"`

		Readonly *bool `json:"readonly,omitempty"`

		Share *CoreV2NodeAO1Share `json:"share,omitempty"`

		SharelinkHasPassword *bool `json:"sharelink_has_password,omitempty"`

		Shareowner *CoreV2NodeAO1Shareowner `json:"shareowner,omitempty"`

		Size int64 `json:"size,omitempty"`

		Subscription *bool `json:"subscription,omitempty"`

		SubscriptionExcludeMe *bool `json:"subscription_exclude_me,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Access = dataAO1.Access

	m.Changed = dataAO1.Changed

	m.Deleted = dataAO1.Deleted

	m.Directory = dataAO1.Directory

	m.Meta = dataAO1.Meta

	m.Mime = dataAO1.Mime

	m.Name = dataAO1.Name

	m.Parent = dataAO1.Parent

	m.Path = dataAO1.Path

	m.Readonly = dataAO1.Readonly

	m.Share = dataAO1.Share

	m.SharelinkHasPassword = dataAO1.SharelinkHasPassword

	m.Shareowner = dataAO1.Shareowner

	m.Size = dataAO1.Size

	m.Subscription = dataAO1.Subscription

	m.SubscriptionExcludeMe = dataAO1.SubscriptionExcludeMe

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m CoreV2Node) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CoreV2Resource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Access *string `json:"access,omitempty"`

		Changed strfmt.DateTime `json:"changed,omitempty"`

		Deleted strfmt.DateTime `json:"deleted,omitempty"`

		Directory *bool `json:"directory,omitempty"`

		Meta *CoreV2NodeMeta `json:"meta,omitempty"`

		Mime string `json:"mime,omitempty"`

		Name string `json:"name,omitempty"`

		Parent *CoreV2NodeAO1Parent `json:"parent,omitempty"`

		Path string `json:"path,omitempty"`

		Readonly *bool `json:"readonly,omitempty"`

		Share *CoreV2NodeAO1Share `json:"share,omitempty"`

		SharelinkHasPassword *bool `json:"sharelink_has_password,omitempty"`

		Shareowner *CoreV2NodeAO1Shareowner `json:"shareowner,omitempty"`

		Size int64 `json:"size,omitempty"`

		Subscription *bool `json:"subscription,omitempty"`

		SubscriptionExcludeMe *bool `json:"subscription_exclude_me,omitempty"`
	}

	dataAO1.Access = m.Access

	dataAO1.Changed = m.Changed

	dataAO1.Deleted = m.Deleted

	dataAO1.Directory = m.Directory

	dataAO1.Meta = m.Meta

	dataAO1.Mime = m.Mime

	dataAO1.Name = m.Name

	dataAO1.Parent = m.Parent

	dataAO1.Path = m.Path

	dataAO1.Readonly = m.Readonly

	dataAO1.Share = m.Share

	dataAO1.SharelinkHasPassword = m.SharelinkHasPassword

	dataAO1.Shareowner = m.Shareowner

	dataAO1.Size = m.Size

	dataAO1.Subscription = m.Subscription

	dataAO1.SubscriptionExcludeMe = m.SubscriptionExcludeMe

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this core v2 node
func (m *CoreV2Node) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CoreV2Resource
	if err := m.CoreV2Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateChanged(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDeleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeta(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShare(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateShareowner(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var coreV2NodeTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rw","w","m","rp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		coreV2NodeTypeAccessPropEnum = append(coreV2NodeTypeAccessPropEnum, v)
	}
}

// property enum
func (m *CoreV2Node) validateAccessEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, coreV2NodeTypeAccessPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CoreV2Node) validateAccess(formats strfmt.Registry) error {

	if swag.IsZero(m.Access) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccessEnum("access", "body", *m.Access); err != nil {
		return err
	}

	return nil
}

func (m *CoreV2Node) validateChanged(formats strfmt.Registry) error {

	if swag.IsZero(m.Changed) { // not required
		return nil
	}

	if err := validate.FormatOf("changed", "body", "date-time", m.Changed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CoreV2Node) validateDeleted(formats strfmt.Registry) error {

	if swag.IsZero(m.Deleted) { // not required
		return nil
	}

	if err := validate.FormatOf("deleted", "body", "date-time", m.Deleted.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *CoreV2Node) validateMeta(formats strfmt.Registry) error {

	if swag.IsZero(m.Meta) { // not required
		return nil
	}

	if m.Meta != nil {
		if err := m.Meta.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("meta")
			}
			return err
		}
	}

	return nil
}

func (m *CoreV2Node) validateParent(formats strfmt.Registry) error {

	if swag.IsZero(m.Parent) { // not required
		return nil
	}

	if m.Parent != nil {
		if err := m.Parent.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("parent")
			}
			return err
		}
	}

	return nil
}

func (m *CoreV2Node) validateShare(formats strfmt.Registry) error {

	if swag.IsZero(m.Share) { // not required
		return nil
	}

	if m.Share != nil {
		if err := m.Share.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("share")
			}
			return err
		}
	}

	return nil
}

func (m *CoreV2Node) validateShareowner(formats strfmt.Registry) error {

	if swag.IsZero(m.Shareowner) { // not required
		return nil
	}

	if m.Shareowner != nil {
		if err := m.Shareowner.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("shareowner")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2Node) UnmarshalBinary(b []byte) error {
	var res CoreV2Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CoreV2NodeAO1Parent Points to the parent collection. If the nodes is in the root this is null.
// swagger:model CoreV2NodeAO1Parent
type CoreV2NodeAO1Parent struct {

	// The id of the parent collection
	ID string `json:"id,omitempty"`

	// The name of the parent collection
	Name string `json:"name,omitempty"`
}

// Validate validates this core v2 node a o1 parent
func (m *CoreV2NodeAO1Parent) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2NodeAO1Parent) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2NodeAO1Parent) UnmarshalBinary(b []byte) error {
	var res CoreV2NodeAO1Parent
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CoreV2NodeAO1Share Points to the shared node (or share reference). If the node is not part of any share this is null.
// swagger:model CoreV2NodeAO1Share
type CoreV2NodeAO1Share struct {

	// The id of the share collection. If the share is a reference (incoming share) it will point to the share reference and not the id of the shared collection itself.
	ID string `json:"id,omitempty"`

	// The name of the share (or share reference).
	Name string `json:"name,omitempty"`
}

// Validate validates this core v2 node a o1 share
func (m *CoreV2NodeAO1Share) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2NodeAO1Share) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2NodeAO1Share) UnmarshalBinary(b []byte) error {
	var res CoreV2NodeAO1Share
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// CoreV2NodeAO1Shareowner Points to the share owner. If the node is not part of any share this is null.
// swagger:model CoreV2NodeAO1Shareowner
type CoreV2NodeAO1Shareowner struct {

	// The id of the share owner.
	ID string `json:"id,omitempty"`

	// The username of the share owner.
	Name string `json:"name,omitempty"`
}

// Validate validates this core v2 node a o1 shareowner
func (m *CoreV2NodeAO1Shareowner) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2NodeAO1Shareowner) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2NodeAO1Shareowner) UnmarshalBinary(b []byte) error {
	var res CoreV2NodeAO1Shareowner
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
