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

// NewUndeleteUserParams creates a new UndeleteUserParams object
// with the default values initialized.
func NewUndeleteUserParams() *UndeleteUserParams {
	var ()
	return &UndeleteUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUndeleteUserParamsWithTimeout creates a new UndeleteUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUndeleteUserParamsWithTimeout(timeout time.Duration) *UndeleteUserParams {
	var ()
	return &UndeleteUserParams{

		timeout: timeout,
	}
}

// NewUndeleteUserParamsWithContext creates a new UndeleteUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewUndeleteUserParamsWithContext(ctx context.Context) *UndeleteUserParams {
	var ()
	return &UndeleteUserParams{

		Context: ctx,
	}
}

// NewUndeleteUserParamsWithHTTPClient creates a new UndeleteUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUndeleteUserParamsWithHTTPClient(client *http.Client) *UndeleteUserParams {
	var ()
	return &UndeleteUserParams{
		HTTPClient: client,
	}
}

/*UndeleteUserParams contains all the parameters to send to the API endpoint
for the undelete user operation typically these are written to a http.Request
*/
type UndeleteUserParams struct {

	/*User
	  Resource identifier

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the undelete user params
func (o *UndeleteUserParams) WithTimeout(timeout time.Duration) *UndeleteUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the undelete user params
func (o *UndeleteUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the undelete user params
func (o *UndeleteUserParams) WithContext(ctx context.Context) *UndeleteUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the undelete user params
func (o *UndeleteUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the undelete user params
func (o *UndeleteUserParams) WithHTTPClient(client *http.Client) *UndeleteUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the undelete user params
func (o *UndeleteUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUser adds the user to the undelete user params
func (o *UndeleteUserParams) WithUser(user string) *UndeleteUserParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the undelete user params
func (o *UndeleteUserParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *UndeleteUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
