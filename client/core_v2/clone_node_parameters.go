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

// NewCloneNodeParams creates a new CloneNodeParams object
// with the default values initialized.
func NewCloneNodeParams() *CloneNodeParams {
	var (
		conflictDefault = float64(0)
	)
	return &CloneNodeParams{
		Conflict: &conflictDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCloneNodeParamsWithTimeout creates a new CloneNodeParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCloneNodeParamsWithTimeout(timeout time.Duration) *CloneNodeParams {
	var (
		conflictDefault = float64(0)
	)
	return &CloneNodeParams{
		Conflict: &conflictDefault,

		timeout: timeout,
	}
}

// NewCloneNodeParamsWithContext creates a new CloneNodeParams object
// with the default values initialized, and the ability to set a context for a request
func NewCloneNodeParamsWithContext(ctx context.Context) *CloneNodeParams {
	var (
		conflictDefault = float64(0)
	)
	return &CloneNodeParams{
		Conflict: &conflictDefault,

		Context: ctx,
	}
}

// NewCloneNodeParamsWithHTTPClient creates a new CloneNodeParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCloneNodeParamsWithHTTPClient(client *http.Client) *CloneNodeParams {
	var (
		conflictDefault = float64(0)
	)
	return &CloneNodeParams{
		Conflict:   &conflictDefault,
		HTTPClient: client,
	}
}

/*CloneNodeParams contains all the parameters to send to the API endpoint
for the clone node operation typically these are written to a http.Request
*/
type CloneNodeParams struct {

	/*Conflict
	  Conflict resolution

	*/
	Conflict *float64
	/*Destid
	  Destination collection, if this is null root is taken

	*/
	Destid *string
	/*Node
	  Node identifier

	*/
	Node string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the clone node params
func (o *CloneNodeParams) WithTimeout(timeout time.Duration) *CloneNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the clone node params
func (o *CloneNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the clone node params
func (o *CloneNodeParams) WithContext(ctx context.Context) *CloneNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the clone node params
func (o *CloneNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the clone node params
func (o *CloneNodeParams) WithHTTPClient(client *http.Client) *CloneNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the clone node params
func (o *CloneNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConflict adds the conflict to the clone node params
func (o *CloneNodeParams) WithConflict(conflict *float64) *CloneNodeParams {
	o.SetConflict(conflict)
	return o
}

// SetConflict adds the conflict to the clone node params
func (o *CloneNodeParams) SetConflict(conflict *float64) {
	o.Conflict = conflict
}

// WithDestid adds the destid to the clone node params
func (o *CloneNodeParams) WithDestid(destid *string) *CloneNodeParams {
	o.SetDestid(destid)
	return o
}

// SetDestid adds the destid to the clone node params
func (o *CloneNodeParams) SetDestid(destid *string) {
	o.Destid = destid
}

// WithNode adds the node to the clone node params
func (o *CloneNodeParams) WithNode(node string) *CloneNodeParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the clone node params
func (o *CloneNodeParams) SetNode(node string) {
	o.Node = node
}

// WriteToRequest writes these params to a swagger request
func (o *CloneNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Conflict != nil {

		// query param conflict
		var qrConflict float64
		if o.Conflict != nil {
			qrConflict = *o.Conflict
		}
		qConflict := swag.FormatFloat64(qrConflict)
		if qConflict != "" {
			if err := r.SetQueryParam("conflict", qConflict); err != nil {
				return err
			}
		}

	}

	if o.Destid != nil {

		// query param destid
		var qrDestid string
		if o.Destid != nil {
			qrDestid = *o.Destid
		}
		qDestid := qrDestid
		if qDestid != "" {
			if err := r.SetQueryParam("destid", qDestid); err != nil {
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