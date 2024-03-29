// Code generated by go-swagger; DO NOT EDIT.

package ballon_app_elasticsearch_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new ballon app elasticsearch v2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ballon app elasticsearch v2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
QueryElasticsearch extendeds search query using elasticsearch
*/
func (a *Client) QueryElasticsearch(params *QueryElasticsearchParams, authInfo runtime.ClientAuthInfoWriter) (*QueryElasticsearchOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewQueryElasticsearchParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "queryElasticsearch",
		Method:             "GET",
		PathPattern:        "/api/v2/nodes/search",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &QueryElasticsearchReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*QueryElasticsearchOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for queryElasticsearch: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
