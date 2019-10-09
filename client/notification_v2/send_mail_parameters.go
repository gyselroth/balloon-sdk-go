// Code generated by go-swagger; DO NOT EDIT.

package notification_v2

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

// NewSendMailParams creates a new SendMailParams object
// with the default values initialized.
func NewSendMailParams() *SendMailParams {
	var ()
	return &SendMailParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendMailParamsWithTimeout creates a new SendMailParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendMailParamsWithTimeout(timeout time.Duration) *SendMailParams {
	var ()
	return &SendMailParams{

		timeout: timeout,
	}
}

// NewSendMailParamsWithContext creates a new SendMailParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendMailParamsWithContext(ctx context.Context) *SendMailParams {
	var ()
	return &SendMailParams{

		Context: ctx,
	}
}

// NewSendMailParamsWithHTTPClient creates a new SendMailParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendMailParamsWithHTTPClient(client *http.Client) *SendMailParams {
	var ()
	return &SendMailParams{
		HTTPClient: client,
	}
}

/*SendMailParams contains all the parameters to send to the API endpoint
for the send mail operation typically these are written to a http.Request
*/
type SendMailParams struct {

	/*Body
	  Mail

	*/
	Body *models.NotificationV2Mail

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send mail params
func (o *SendMailParams) WithTimeout(timeout time.Duration) *SendMailParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send mail params
func (o *SendMailParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send mail params
func (o *SendMailParams) WithContext(ctx context.Context) *SendMailParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send mail params
func (o *SendMailParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send mail params
func (o *SendMailParams) WithHTTPClient(client *http.Client) *SendMailParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send mail params
func (o *SendMailParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the send mail params
func (o *SendMailParams) WithBody(body *models.NotificationV2Mail) *SendMailParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the send mail params
func (o *SendMailParams) SetBody(body *models.NotificationV2Mail) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *SendMailParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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