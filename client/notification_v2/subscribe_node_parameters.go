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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewSubscribeNodeParams creates a new SubscribeNodeParams object
// with the default values initialized.
func NewSubscribeNodeParams() *SubscribeNodeParams {
	var ()
	return &SubscribeNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSubscribeNodeParamsWithTimeout creates a new SubscribeNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSubscribeNodeParamsWithTimeout(timeout time.Duration) *SubscribeNodeParams {
	var ()
	return &SubscribeNodeParams{

		timeout: timeout,
	}
}

// NewSubscribeNodeParamsWithContext creates a new SubscribeNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewSubscribeNodeParamsWithContext(ctx context.Context) *SubscribeNodeParams {
	var ()
	return &SubscribeNodeParams{

		Context: ctx,
	}
}

// NewSubscribeNodeParamsWithHTTPClient creates a new SubscribeNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSubscribeNodeParamsWithHTTPClient(client *http.Client) *SubscribeNodeParams {
	var ()
	return &SubscribeNodeParams{
		HTTPClient: client,
	}
}

/*SubscribeNodeParams contains all the parameters to send to the API endpoint
for the subscribe node operation typically these are written to a http.Request
*/
type SubscribeNodeParams struct {

	/*ExcludeMe
	  Exclude subscription owner (me) from change notifications

	*/
	ExcludeMe *bool
	/*Node
	  Node identifier

	*/
	Node string
	/*Recursive
	  Apply subscription to children (inclusive newly added children)

	*/
	Recursive *bool
	/*Subscribe
	  If true the subscription is active

	*/
	Subscribe *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the subscribe node params
func (o *SubscribeNodeParams) WithTimeout(timeout time.Duration) *SubscribeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the subscribe node params
func (o *SubscribeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the subscribe node params
func (o *SubscribeNodeParams) WithContext(ctx context.Context) *SubscribeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the subscribe node params
func (o *SubscribeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the subscribe node params
func (o *SubscribeNodeParams) WithHTTPClient(client *http.Client) *SubscribeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the subscribe node params
func (o *SubscribeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithExcludeMe adds the excludeMe to the subscribe node params
func (o *SubscribeNodeParams) WithExcludeMe(excludeMe *bool) *SubscribeNodeParams {
	o.SetExcludeMe(excludeMe)
	return o
}

// SetExcludeMe adds the excludeMe to the subscribe node params
func (o *SubscribeNodeParams) SetExcludeMe(excludeMe *bool) {
	o.ExcludeMe = excludeMe
}

// WithNode adds the node to the subscribe node params
func (o *SubscribeNodeParams) WithNode(node string) *SubscribeNodeParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the subscribe node params
func (o *SubscribeNodeParams) SetNode(node string) {
	o.Node = node
}

// WithRecursive adds the recursive to the subscribe node params
func (o *SubscribeNodeParams) WithRecursive(recursive *bool) *SubscribeNodeParams {
	o.SetRecursive(recursive)
	return o
}

// SetRecursive adds the recursive to the subscribe node params
func (o *SubscribeNodeParams) SetRecursive(recursive *bool) {
	o.Recursive = recursive
}

// WithSubscribe adds the subscribe to the subscribe node params
func (o *SubscribeNodeParams) WithSubscribe(subscribe *bool) *SubscribeNodeParams {
	o.SetSubscribe(subscribe)
	return o
}

// SetSubscribe adds the subscribe to the subscribe node params
func (o *SubscribeNodeParams) SetSubscribe(subscribe *bool) {
	o.Subscribe = subscribe
}

// WriteToRequest writes these params to a swagger request
func (o *SubscribeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ExcludeMe != nil {

		// query param exclude_me
		var qrExcludeMe bool
		if o.ExcludeMe != nil {
			qrExcludeMe = *o.ExcludeMe
		}
		qExcludeMe := swag.FormatBool(qrExcludeMe)
		if qExcludeMe != "" {
			if err := r.SetQueryParam("exclude_me", qExcludeMe); err != nil {
				return err
			}
		}

	}

	// path param node
	if err := r.SetPathParam("node", o.Node); err != nil {
		return err
	}

	if o.Recursive != nil {

		// query param recursive
		var qrRecursive bool
		if o.Recursive != nil {
			qrRecursive = *o.Recursive
		}
		qRecursive := swag.FormatBool(qrRecursive)
		if qRecursive != "" {
			if err := r.SetQueryParam("recursive", qRecursive); err != nil {
				return err
			}
		}

	}

	if o.Subscribe != nil {

		// query param subscribe
		var qrSubscribe bool
		if o.Subscribe != nil {
			qrSubscribe = *o.Subscribe
		}
		qSubscribe := swag.FormatBool(qrSubscribe)
		if qSubscribe != "" {
			if err := r.SetQueryParam("subscribe", qSubscribe); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
