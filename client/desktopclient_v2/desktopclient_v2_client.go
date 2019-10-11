// Code generated by go-swagger; DO NOT EDIT.

package desktopclient_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new desktopclient v2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for desktopclient v2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetDesktopClient downloads balloon desktop client
*/
func (a *Client) GetDesktopClient(params *GetDesktopClientParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*GetDesktopClientOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDesktopClientParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getDesktopClient",
		Method:             "GET",
		PathPattern:        "/api/v2/desktop-clients/{format}/content",
		ProducesMediaTypes: []string{"octet/stream"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetDesktopClientReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDesktopClientOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDesktopClient: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
