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

// NewGetDeltaParams creates a new GetDeltaParams object
// with the default values initialized.
func NewGetDeltaParams() *GetDeltaParams {
	var ()
	return &GetDeltaParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeltaParamsWithTimeout creates a new GetDeltaParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeltaParamsWithTimeout(timeout time.Duration) *GetDeltaParams {
	var ()
	return &GetDeltaParams{

		timeout: timeout,
	}
}

// NewGetDeltaParamsWithContext creates a new GetDeltaParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeltaParamsWithContext(ctx context.Context) *GetDeltaParams {
	var ()
	return &GetDeltaParams{

		Context: ctx,
	}
}

// NewGetDeltaParamsWithHTTPClient creates a new GetDeltaParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeltaParamsWithHTTPClient(client *http.Client) *GetDeltaParams {
	var ()
	return &GetDeltaParams{
		HTTPClient: client,
	}
}

/*GetDeltaParams contains all the parameters to send to the API endpoint
for the get delta operation typically these are written to a http.Request
*/
type GetDeltaParams struct {

	/*Attributes
	  Filter attributes, per default not all attributes would be returned

	*/
	Attributes []string
	/*Cursor
	  Set a cursor to rquest next nodes within delta processing

	*/
	Cursor *string
	/*Limit
	  Limit the number of delta entries, if too low you have to call this endpoint more often since has_more would be true more often

	*/
	Limit *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get delta params
func (o *GetDeltaParams) WithTimeout(timeout time.Duration) *GetDeltaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get delta params
func (o *GetDeltaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get delta params
func (o *GetDeltaParams) WithContext(ctx context.Context) *GetDeltaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get delta params
func (o *GetDeltaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get delta params
func (o *GetDeltaParams) WithHTTPClient(client *http.Client) *GetDeltaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get delta params
func (o *GetDeltaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributes adds the attributes to the get delta params
func (o *GetDeltaParams) WithAttributes(attributes []string) *GetDeltaParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get delta params
func (o *GetDeltaParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithCursor adds the cursor to the get delta params
func (o *GetDeltaParams) WithCursor(cursor *string) *GetDeltaParams {
	o.SetCursor(cursor)
	return o
}

// SetCursor adds the cursor to the get delta params
func (o *GetDeltaParams) SetCursor(cursor *string) {
	o.Cursor = cursor
}

// WithLimit adds the limit to the get delta params
func (o *GetDeltaParams) WithLimit(limit *int64) *GetDeltaParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get delta params
func (o *GetDeltaParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeltaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	valuesAttributes := o.Attributes

	joinedAttributes := swag.JoinByFormat(valuesAttributes, "")
	// query array param attributes
	if err := r.SetQueryParam("attributes", joinedAttributes...); err != nil {
		return err
	}

	if o.Cursor != nil {

		// query param cursor
		var qrCursor string
		if o.Cursor != nil {
			qrCursor = *o.Cursor
		}
		qCursor := qrCursor
		if qCursor != "" {
			if err := r.SetQueryParam("cursor", qCursor); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
