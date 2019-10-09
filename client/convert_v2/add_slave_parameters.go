// Code generated by go-swagger; DO NOT EDIT.

package convert_v2

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

	models "github.com/gyselroth/balloon-sdk-go/models"
)

// NewAddSlaveParams creates a new AddSlaveParams object
// with the default values initialized.
func NewAddSlaveParams() *AddSlaveParams {
	var ()
	return &AddSlaveParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddSlaveParamsWithTimeout creates a new AddSlaveParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddSlaveParamsWithTimeout(timeout time.Duration) *AddSlaveParams {
	var ()
	return &AddSlaveParams{

		timeout: timeout,
	}
}

// NewAddSlaveParamsWithContext creates a new AddSlaveParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddSlaveParamsWithContext(ctx context.Context) *AddSlaveParams {
	var ()
	return &AddSlaveParams{

		Context: ctx,
	}
}

// NewAddSlaveParamsWithHTTPClient creates a new AddSlaveParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddSlaveParamsWithHTTPClient(client *http.Client) *AddSlaveParams {
	var ()
	return &AddSlaveParams{
		HTTPClient: client,
	}
}

/*AddSlaveParams contains all the parameters to send to the API endpoint
for the add slave operation typically these are written to a http.Request
*/
type AddSlaveParams struct {

	/*Body
	  Add new conversion slave

	*/
	Body *models.ConvertV2Slave
	/*File
	  File identifier

	*/
	File string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add slave params
func (o *AddSlaveParams) WithTimeout(timeout time.Duration) *AddSlaveParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add slave params
func (o *AddSlaveParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add slave params
func (o *AddSlaveParams) WithContext(ctx context.Context) *AddSlaveParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add slave params
func (o *AddSlaveParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add slave params
func (o *AddSlaveParams) WithHTTPClient(client *http.Client) *AddSlaveParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add slave params
func (o *AddSlaveParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add slave params
func (o *AddSlaveParams) WithBody(body *models.ConvertV2Slave) *AddSlaveParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add slave params
func (o *AddSlaveParams) SetBody(body *models.ConvertV2Slave) {
	o.Body = body
}

// WithFile adds the file to the add slave params
func (o *AddSlaveParams) WithFile(file string) *AddSlaveParams {
	o.SetFile(file)
	return o
}

// SetFile adds the file to the add slave params
func (o *AddSlaveParams) SetFile(file string) {
	o.File = file
}

// WriteToRequest writes these params to a swagger request
func (o *AddSlaveParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param file
	if err := r.SetPathParam("file", o.File); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}