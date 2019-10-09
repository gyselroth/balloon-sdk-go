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

// NewDeleteNodeParams creates a new DeleteNodeParams object
// with the default values initialized.
func NewDeleteNodeParams() *DeleteNodeParams {
	var ()
	return &DeleteNodeParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteNodeParamsWithTimeout creates a new DeleteNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteNodeParamsWithTimeout(timeout time.Duration) *DeleteNodeParams {
	var ()
	return &DeleteNodeParams{

		timeout: timeout,
	}
}

// NewDeleteNodeParamsWithContext creates a new DeleteNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteNodeParamsWithContext(ctx context.Context) *DeleteNodeParams {
	var ()
	return &DeleteNodeParams{

		Context: ctx,
	}
}

// NewDeleteNodeParamsWithHTTPClient creates a new DeleteNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteNodeParamsWithHTTPClient(client *http.Client) *DeleteNodeParams {
	var ()
	return &DeleteNodeParams{
		HTTPClient: client,
	}
}

/*DeleteNodeParams contains all the parameters to send to the API endpoint
for the delete node operation typically these are written to a http.Request
*/
type DeleteNodeParams struct {

	/*At
	  Has to be a valid unix timestamp if so the node will destroy itself at this specified time instead immediatly

	*/
	At *float64
	/*Force
	  Force flag need to be set to delete a node from trash (node must have the deleted flag)

	*/
	Force *bool
	/*IgnoreFlag
	  If both ignore_flag and force_flag were set, the node will be deleted completely

	*/
	IgnoreFlag *bool
	/*Node
	  Node identifier

	*/
	Node string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete node params
func (o *DeleteNodeParams) WithTimeout(timeout time.Duration) *DeleteNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete node params
func (o *DeleteNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete node params
func (o *DeleteNodeParams) WithContext(ctx context.Context) *DeleteNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete node params
func (o *DeleteNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete node params
func (o *DeleteNodeParams) WithHTTPClient(client *http.Client) *DeleteNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete node params
func (o *DeleteNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAt adds the at to the delete node params
func (o *DeleteNodeParams) WithAt(at *float64) *DeleteNodeParams {
	o.SetAt(at)
	return o
}

// SetAt adds the at to the delete node params
func (o *DeleteNodeParams) SetAt(at *float64) {
	o.At = at
}

// WithForce adds the force to the delete node params
func (o *DeleteNodeParams) WithForce(force *bool) *DeleteNodeParams {
	o.SetForce(force)
	return o
}

// SetForce adds the force to the delete node params
func (o *DeleteNodeParams) SetForce(force *bool) {
	o.Force = force
}

// WithIgnoreFlag adds the ignoreFlag to the delete node params
func (o *DeleteNodeParams) WithIgnoreFlag(ignoreFlag *bool) *DeleteNodeParams {
	o.SetIgnoreFlag(ignoreFlag)
	return o
}

// SetIgnoreFlag adds the ignoreFlag to the delete node params
func (o *DeleteNodeParams) SetIgnoreFlag(ignoreFlag *bool) {
	o.IgnoreFlag = ignoreFlag
}

// WithNode adds the node to the delete node params
func (o *DeleteNodeParams) WithNode(node string) *DeleteNodeParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the delete node params
func (o *DeleteNodeParams) SetNode(node string) {
	o.Node = node
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.At != nil {

		// query param at
		var qrAt float64
		if o.At != nil {
			qrAt = *o.At
		}
		qAt := swag.FormatFloat64(qrAt)
		if qAt != "" {
			if err := r.SetQueryParam("at", qAt); err != nil {
				return err
			}
		}

	}

	if o.Force != nil {

		// query param force
		var qrForce bool
		if o.Force != nil {
			qrForce = *o.Force
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {
			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}

	}

	if o.IgnoreFlag != nil {

		// query param ignore_flag
		var qrIgnoreFlag bool
		if o.IgnoreFlag != nil {
			qrIgnoreFlag = *o.IgnoreFlag
		}
		qIgnoreFlag := swag.FormatBool(qrIgnoreFlag)
		if qIgnoreFlag != "" {
			if err := r.SetQueryParam("ignore_flag", qIgnoreFlag); err != nil {
				return err
			}
		}

	}

	// path param node
	if err := r.SetPathParam("node", o.Node); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}