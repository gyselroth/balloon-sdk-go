// Code generated by go-swagger; DO NOT EDIT.

package core_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUserAvatarParams creates a new GetUserAvatarParams object
// with the default values initialized.
func NewGetUserAvatarParams() *GetUserAvatarParams {
	var ()
	return &GetUserAvatarParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserAvatarParamsWithTimeout creates a new GetUserAvatarParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserAvatarParamsWithTimeout(timeout time.Duration) *GetUserAvatarParams {
	var ()
	return &GetUserAvatarParams{

		timeout: timeout,
	}
}

// NewGetUserAvatarParamsWithContext creates a new GetUserAvatarParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserAvatarParamsWithContext(ctx context.Context) *GetUserAvatarParams {
	var ()
	return &GetUserAvatarParams{

		Context: ctx,
	}
}

// NewGetUserAvatarParamsWithHTTPClient creates a new GetUserAvatarParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserAvatarParamsWithHTTPClient(client *http.Client) *GetUserAvatarParams {
	var ()
	return &GetUserAvatarParams{
		HTTPClient: client,
	}
}

/*GetUserAvatarParams contains all the parameters to send to the API endpoint
for the get user avatar operation typically these are written to a http.Request
*/
type GetUserAvatarParams struct {

	/*User
	  Resource identifier

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user avatar params
func (o *GetUserAvatarParams) WithTimeout(timeout time.Duration) *GetUserAvatarParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user avatar params
func (o *GetUserAvatarParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user avatar params
func (o *GetUserAvatarParams) WithContext(ctx context.Context) *GetUserAvatarParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user avatar params
func (o *GetUserAvatarParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user avatar params
func (o *GetUserAvatarParams) WithHTTPClient(client *http.Client) *GetUserAvatarParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user avatar params
func (o *GetUserAvatarParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the get user avatar params
func (o *GetUserAvatarParams) WithUser(user string) *GetUserAvatarParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the get user avatar params
func (o *GetUserAvatarParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserAvatarParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}