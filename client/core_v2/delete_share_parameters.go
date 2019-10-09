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

// NewDeleteShareParams creates a new DeleteShareParams object
// with the default values initialized.
func NewDeleteShareParams() *DeleteShareParams {
	var ()
	return &DeleteShareParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteShareParamsWithTimeout creates a new DeleteShareParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteShareParamsWithTimeout(timeout time.Duration) *DeleteShareParams {
	var ()
	return &DeleteShareParams{

		timeout: timeout,
	}
}

// NewDeleteShareParamsWithContext creates a new DeleteShareParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteShareParamsWithContext(ctx context.Context) *DeleteShareParams {
	var ()
	return &DeleteShareParams{

		Context: ctx,
	}
}

// NewDeleteShareParamsWithHTTPClient creates a new DeleteShareParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteShareParamsWithHTTPClient(client *http.Client) *DeleteShareParams {
	var ()
	return &DeleteShareParams{
		HTTPClient: client,
	}
}

/*DeleteShareParams contains all the parameters to send to the API endpoint
for the delete share operation typically these are written to a http.Request
*/
type DeleteShareParams struct {

	/*Collection
	  Collection identifier

	*/
	Collection string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete share params
func (o *DeleteShareParams) WithTimeout(timeout time.Duration) *DeleteShareParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete share params
func (o *DeleteShareParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete share params
func (o *DeleteShareParams) WithContext(ctx context.Context) *DeleteShareParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete share params
func (o *DeleteShareParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete share params
func (o *DeleteShareParams) WithHTTPClient(client *http.Client) *DeleteShareParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete share params
func (o *DeleteShareParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollection adds the collection to the delete share params
func (o *DeleteShareParams) WithCollection(collection string) *DeleteShareParams {
	o.SetCollection(collection)
	return o
}

// SetCollection adds the collection to the delete share params
func (o *DeleteShareParams) SetCollection(collection string) {
	o.Collection = collection
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteShareParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collection
	if err := r.SetPathParam("collection", o.Collection); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}