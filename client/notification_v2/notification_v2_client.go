// Code generated by go-swagger; DO NOT EDIT.

package notification_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new notification v2 API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for notification v2 API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteNotification deletes a notification
*/
func (a *Client) DeleteNotification(params *DeleteNotificationParams) (*DeleteNotificationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteNotification",
		Method:             "DELETE",
		PathPattern:        "/api/v2/notifications/{notification}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteNotificationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteNotificationNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetMyNotifications gets my nofitifications
*/
func (a *Client) GetMyNotifications(params *GetMyNotificationsParams) (*GetMyNotificationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetMyNotificationsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getMyNotifications",
		Method:             "GET",
		PathPattern:        "/api/v2/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetMyNotificationsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetMyNotificationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getMyNotifications: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetNotification gets a single notification
*/
func (a *Client) GetNotification(params *GetNotificationParams) (*GetNotificationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNotification",
		Method:             "GET",
		PathPattern:        "/api/v2/notifications/{notification}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNotificationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNotificationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostNotification sends notification
*/
func (a *Client) PostNotification(params *PostNotificationParams) (*PostNotificationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostNotificationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postNotification",
		Method:             "POST",
		PathPattern:        "/api/v2/notifications",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostNotificationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostNotificationAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for postNotification: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SendBroadcast sends notification to all users
*/
func (a *Client) SendBroadcast(params *SendBroadcastParams) (*SendBroadcastAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendBroadcastParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendBroadcast",
		Method:             "POST",
		PathPattern:        "/api/v2/notifications/broadcast",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendBroadcastReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SendBroadcastAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendBroadcast: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SendMail sends mail
*/
func (a *Client) SendMail(params *SendMailParams) (*SendMailAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSendMailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "sendMail",
		Method:             "POST",
		PathPattern:        "/api/v2/notifications/mail",
		ProducesMediaTypes: []string{""},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SendMailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SendMailAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for sendMail: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
SubscribeNode receives node updates
*/
func (a *Client) SubscribeNode(params *SubscribeNodeParams) (*SubscribeNodeAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSubscribeNodeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "subscribeNode",
		Method:             "POST",
		PathPattern:        "/api/v2/nodes/{node}/subscription",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &SubscribeNodeReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SubscribeNodeAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for subscribeNode: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
