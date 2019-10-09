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

// NewRollbackFileParams creates a new RollbackFileParams object
// with the default values initialized.
func NewRollbackFileParams() *RollbackFileParams {
	var ()
	return &RollbackFileParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRollbackFileParamsWithTimeout creates a new RollbackFileParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRollbackFileParamsWithTimeout(timeout time.Duration) *RollbackFileParams {
	var ()
	return &RollbackFileParams{

		timeout: timeout,
	}
}

// NewRollbackFileParamsWithContext creates a new RollbackFileParams object
// with the default values initialized, and the ability to set a context for a request
func NewRollbackFileParamsWithContext(ctx context.Context) *RollbackFileParams {
	var ()
	return &RollbackFileParams{

		Context: ctx,
	}
}

// NewRollbackFileParamsWithHTTPClient creates a new RollbackFileParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRollbackFileParamsWithHTTPClient(client *http.Client) *RollbackFileParams {
	var ()
	return &RollbackFileParams{
		HTTPClient: client,
	}
}

/*RollbackFileParams contains all the parameters to send to the API endpoint
for the rollback file operation typically these are written to a http.Request
*/
type RollbackFileParams struct {

	/*File
	  File identifier

	*/
	File string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the rollback file params
func (o *RollbackFileParams) WithTimeout(timeout time.Duration) *RollbackFileParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rollback file params
func (o *RollbackFileParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rollback file params
func (o *RollbackFileParams) WithContext(ctx context.Context) *RollbackFileParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rollback file params
func (o *RollbackFileParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rollback file params
func (o *RollbackFileParams) WithHTTPClient(client *http.Client) *RollbackFileParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rollback file params
func (o *RollbackFileParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFile adds the file to the rollback file params
func (o *RollbackFileParams) WithFile(file string) *RollbackFileParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the rollback file params
func (o *RollbackFileParams) SetFile(file string) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *RollbackFileParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param file
	if err := r.SetPathParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
