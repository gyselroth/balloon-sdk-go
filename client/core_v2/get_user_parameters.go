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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetUserParams creates a new GetUserParams object
// with the default values initialized.
func NewGetUserParams() *GetUserParams {
	var ()
	return &GetUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserParamsWithTimeout creates a new GetUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserParamsWithTimeout(timeout time.Duration) *GetUserParams {
	var ()
	return &GetUserParams{

		timeout: timeout,
	}
}

// NewGetUserParamsWithContext creates a new GetUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserParamsWithContext(ctx context.Context) *GetUserParams {
	var ()
	return &GetUserParams{

		Context: ctx,
	}
}

// NewGetUserParamsWithHTTPClient creates a new GetUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserParamsWithHTTPClient(client *http.Client) *GetUserParams {
	var ()
	return &GetUserParams{
		HTTPClient: client,
	}
}

/*GetUserParams contains all the parameters to send to the API endpoint
for the get user operation typically these are written to a http.Request
*/
type GetUserParams struct {

	/*Attributes
	  Filter attributes

	*/
	Attributes []string
	/*User
	  Resource identifier

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user params
func (o *GetUserParams) WithTimeout(timeout time.Duration) *GetUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user params
func (o *GetUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user params
func (o *GetUserParams) WithContext(ctx context.Context) *GetUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user params
func (o *GetUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) WithHTTPClient(client *http.Client) *GetUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user params
func (o *GetUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributes adds the attributes to the get user params
func (o *GetUserParams) WithAttributes(attributes []string) *GetUserParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get user params
func (o *GetUserParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithUser adds the user to the get user params
func (o *GetUserParams) WithUser(user string) *GetUserParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the get user params
func (o *GetUserParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAttributes := o.Attributes

	joinedAttributes := swag.JoinByFormat(valuesAttributes, "")
	// query array param attributes
	if err := r.SetQueryParam("attributes", joinedAttributes...); err != nil {
		return err
	}

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
