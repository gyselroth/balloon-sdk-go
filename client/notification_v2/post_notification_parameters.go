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

// NewPostNotificationParams creates a new PostNotificationParams object
// with the default values initialized.
func NewPostNotificationParams() *PostNotificationParams {
	var ()
	return &PostNotificationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostNotificationParamsWithTimeout creates a new PostNotificationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostNotificationParamsWithTimeout(timeout time.Duration) *PostNotificationParams {
	var ()
	return &PostNotificationParams{

		timeout: timeout,
	}
}

// NewPostNotificationParamsWithContext creates a new PostNotificationParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostNotificationParamsWithContext(ctx context.Context) *PostNotificationParams {
	var ()
	return &PostNotificationParams{

		Context: ctx,
	}
}

// NewPostNotificationParamsWithHTTPClient creates a new PostNotificationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostNotificationParamsWithHTTPClient(client *http.Client) *PostNotificationParams {
	var ()
	return &PostNotificationParams{
		HTTPClient: client,
	}
}

/*PostNotificationParams contains all the parameters to send to the API endpoint
for the post notification operation typically these are written to a http.Request
*/
type PostNotificationParams struct {

	/*Body*/
	Body *models.NotificationV2Notification

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post notification params
func (o *PostNotificationParams) WithTimeout(timeout time.Duration) *PostNotificationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post notification params
func (o *PostNotificationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post notification params
func (o *PostNotificationParams) WithContext(ctx context.Context) *PostNotificationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post notification params
func (o *PostNotificationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post notification params
func (o *PostNotificationParams) WithHTTPClient(client *http.Client) *PostNotificationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post notification params
func (o *PostNotificationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post notification params
func (o *PostNotificationParams) WithBody(body *models.NotificationV2Notification) *PostNotificationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post notification params
func (o *PostNotificationParams) SetBody(body *models.NotificationV2Notification) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostNotificationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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