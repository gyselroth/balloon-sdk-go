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

	models "github.com/gyselroth/balloon-sdk-go/models"
)

// NewCreateRootCollectionParams creates a new CreateRootCollectionParams object
// with the default values initialized.
func NewCreateRootCollectionParams() *CreateRootCollectionParams {
	var ()
	return &CreateRootCollectionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRootCollectionParamsWithTimeout creates a new CreateRootCollectionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRootCollectionParamsWithTimeout(timeout time.Duration) *CreateRootCollectionParams {
	var ()
	return &CreateRootCollectionParams{

		timeout: timeout,
	}
}

// NewCreateRootCollectionParamsWithContext creates a new CreateRootCollectionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRootCollectionParamsWithContext(ctx context.Context) *CreateRootCollectionParams {
	var ()
	return &CreateRootCollectionParams{

		Context: ctx,
	}
}

// NewCreateRootCollectionParamsWithHTTPClient creates a new CreateRootCollectionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRootCollectionParamsWithHTTPClient(client *http.Client) *CreateRootCollectionParams {
	var ()
	return &CreateRootCollectionParams{
		HTTPClient: client,
	}
}

/*CreateRootCollectionParams contains all the parameters to send to the API endpoint
for the create root collection operation typically these are written to a http.Request
*/
type CreateRootCollectionParams struct {

	/*Body*/
	Body *models.CoreV2Collection

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create root collection params
func (o *CreateRootCollectionParams) WithTimeout(timeout time.Duration) *CreateRootCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create root collection params
func (o *CreateRootCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create root collection params
func (o *CreateRootCollectionParams) WithContext(ctx context.Context) *CreateRootCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create root collection params
func (o *CreateRootCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create root collection params
func (o *CreateRootCollectionParams) WithHTTPClient(client *http.Client) *CreateRootCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create root collection params
func (o *CreateRootCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create root collection params
func (o *CreateRootCollectionParams) WithBody(body *models.CoreV2Collection) *CreateRootCollectionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create root collection params
func (o *CreateRootCollectionParams) SetBody(body *models.CoreV2Collection) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRootCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
