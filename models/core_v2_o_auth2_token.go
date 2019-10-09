// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// CoreV2OAuth2Token OAUTH2 access token response according RFC6749 https://tools.ietf.org/html/rfc6749#section-4.3.3
// swagger:model core.v2.OAuth2Token
type CoreV2OAuth2Token struct {

	// Resource access_token.
	AccessToken string `json:"access_token,omitempty"`

	// The number of seconds the issues acceess_token is valid for. This is usually 1hour.
	ExpiresIn *float64 `json:"expires_in,omitempty"`

	// An optional refresh token to get new access_tokens via the refresh_token grant type.
	RefreshToken string `json:"refresh_token,omitempty"`

	// The scopes.
	Scope string `json:"scope,omitempty"`

	// The type of token. This is usually a Bearer token.
	TokenType *string `json:"token_type,omitempty"`
}

// Validate validates this core v2 o auth2 token
func (m *CoreV2OAuth2Token) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CoreV2OAuth2Token) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CoreV2OAuth2Token) UnmarshalBinary(b []byte) error {
	var res CoreV2OAuth2Token
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}