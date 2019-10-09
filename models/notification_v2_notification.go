// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// NotificationV2Notification A notification
// swagger:model notification.v2.Notification
type NotificationV2Notification struct {
	CoreV2Resource

	// Notification locale. A locale contains a static message which is of a given locale.
	Locale *string `json:"locale,omitempty"`

	// Message.
	Message string `json:"message,omitempty"`

	// node
	Node *NotificationV2NotificationAO1Node `json:"node,omitempty"`

	// sender
	Sender *NotificationV2NotificationAO1Sender `json:"sender,omitempty"`

	// Subject.
	Subject string `json:"subject,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *NotificationV2Notification) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 CoreV2Resource
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.CoreV2Resource = aO0

	// AO1
	var dataAO1 struct {
		Locale *string `json:"locale,omitempty"`

		Message string `json:"message,omitempty"`

		Node *NotificationV2NotificationAO1Node `json:"node,omitempty"`

		Sender *NotificationV2NotificationAO1Sender `json:"sender,omitempty"`

		Subject string `json:"subject,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.Locale = dataAO1.Locale

	m.Message = dataAO1.Message

	m.Node = dataAO1.Node

	m.Sender = dataAO1.Sender

	m.Subject = dataAO1.Subject

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m NotificationV2Notification) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.CoreV2Resource)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)

	var dataAO1 struct {
		Locale *string `json:"locale,omitempty"`

		Message string `json:"message,omitempty"`

		Node *NotificationV2NotificationAO1Node `json:"node,omitempty"`

		Sender *NotificationV2NotificationAO1Sender `json:"sender,omitempty"`

		Subject string `json:"subject,omitempty"`
	}

	dataAO1.Locale = m.Locale

	dataAO1.Message = m.Message

	dataAO1.Node = m.Node

	dataAO1.Sender = m.Sender

	dataAO1.Subject = m.Subject

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this notification v2 notification
func (m *NotificationV2Notification) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with CoreV2Resource
	if err := m.CoreV2Resource.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NotificationV2Notification) validateNode(formats strfmt.Registry) error {

	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *NotificationV2Notification) validateSender(formats strfmt.Registry) error {

	if swag.IsZero(m.Sender) { // not required
		return nil
	}

	if m.Sender != nil {
		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NotificationV2Notification) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationV2Notification) UnmarshalBinary(b []byte) error {
	var res NotificationV2Notification
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NotificationV2NotificationAO1Node A notification may point to a node. If there is no node given this is null.
// swagger:model NotificationV2NotificationAO1Node
type NotificationV2NotificationAO1Node struct {

	// The id of the node.
	ID string `json:"id,omitempty"`

	// The name of the node.
	Name string `json:"name,omitempty"`
}

// Validate validates this notification v2 notification a o1 node
func (m *NotificationV2NotificationAO1Node) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationV2NotificationAO1Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationV2NotificationAO1Node) UnmarshalBinary(b []byte) error {
	var res NotificationV2NotificationAO1Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// NotificationV2NotificationAO1Sender Points to the message. If there is no sender (For example the message came from the system itself) this is null.
// swagger:model NotificationV2NotificationAO1Sender
type NotificationV2NotificationAO1Sender struct {

	// The id of the sender.
	ID string `json:"id,omitempty"`

	// The username of the sender.
	Name string `json:"name,omitempty"`
}

// Validate validates this notification v2 notification a o1 sender
func (m *NotificationV2NotificationAO1Sender) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *NotificationV2NotificationAO1Sender) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NotificationV2NotificationAO1Sender) UnmarshalBinary(b []byte) error {
	var res NotificationV2NotificationAO1Sender
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}