// Code generated by go-swagger; DO NOT EDIT.

package sharelink_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new sharelink v2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for sharelink v2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateLink creates a unique sharing link of a node global accessible a possible existing link will be deleted if this method will be called
*/
func (a *Client) CreateLink(params *CreateLinkParams) (*CreateLinkCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createLink",
		Method:             "POST",
		PathPattern:        "/api/v2/nodes/{node}/share-link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateLinkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateLinkCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteLink deletes an existing sharing link
*/
func (a *Client) DeleteLink(params *DeleteLinkParams) (*DeleteLinkNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLinkParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLink",
		Method:             "DELETE",
		PathPattern:        "/api/v2/nodes/{node}/share-link",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteLinkReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLinkNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteLink: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
