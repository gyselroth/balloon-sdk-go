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

// NewGetGroupsParams creates a new GetGroupsParams object
// with the default values initialized.
func NewGetGroupsParams() *GetGroupsParams {
	var ()
	return &GetGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetGroupsParamsWithTimeout creates a new GetGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetGroupsParamsWithTimeout(timeout time.Duration) *GetGroupsParams {
	var ()
	return &GetGroupsParams{

		timeout: timeout,
	}
}

// NewGetGroupsParamsWithContext creates a new GetGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetGroupsParamsWithContext(ctx context.Context) *GetGroupsParams {
	var ()
	return &GetGroupsParams{

		Context: ctx,
	}
}

// NewGetGroupsParamsWithHTTPClient creates a new GetGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetGroupsParamsWithHTTPClient(client *http.Client) *GetGroupsParams {
	var ()
	return &GetGroupsParams{
		HTTPClient: client,
	}
}

/*GetGroupsParams contains all the parameters to send to the API endpoint
for the get groups operation typically these are written to a http.Request
*/
type GetGroupsParams struct {

	/*Attributes
	  Filter attributes

	*/
	Attributes []string
	/*Limit
	  Objects limit, per default 20 objects will get returned

	*/
	Limit *int64
	/*Offset
	  Objects offset, per default it starts from 0. You may also request a negative offset which will return results from the end [total - offset].

	*/
	Offset *int64
	/*Query
	  Specify a MongoDB based resource query (https://docs.mongodb.com/manual/tutorial/query-documents) using JSON (For example: {"name": {$regex: 'foo.*'}}).

	*/
	Query *string
	/*Sort
	  Specify a MongoDB sort operation (https://docs.mongodb.com/manual/reference/method/cursor.sort/) using JSON (For example: {"name": -1}).

	*/
	Sort *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get groups params
func (o *GetGroupsParams) WithTimeout(timeout time.Duration) *GetGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get groups params
func (o *GetGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get groups params
func (o *GetGroupsParams) WithContext(ctx context.Context) *GetGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get groups params
func (o *GetGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get groups params
func (o *GetGroupsParams) WithHTTPClient(client *http.Client) *GetGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get groups params
func (o *GetGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributes adds the attributes to the get groups params
func (o *GetGroupsParams) WithAttributes(attributes []string) *GetGroupsParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get groups params
func (o *GetGroupsParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithLimit adds the limit to the get groups params
func (o *GetGroupsParams) WithLimit(limit *int64) *GetGroupsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get groups params
func (o *GetGroupsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get groups params
func (o *GetGroupsParams) WithOffset(offset *int64) *GetGroupsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get groups params
func (o *GetGroupsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQuery adds the query to the get groups params
func (o *GetGroupsParams) WithQuery(query *string) *GetGroupsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get groups params
func (o *GetGroupsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the get groups params
func (o *GetGroupsParams) WithSort(sort *string) *GetGroupsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get groups params
func (o *GetGroupsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Offset != nil {

		// query param offset
		var qrOffset int64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Query != nil {

		// query param query
		var qrQuery string
		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {
			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}

	}

	if o.Sort != nil {

		// query param sort
		var qrSort string
		if o.Sort != nil {
			qrSort = *o.Sort
		}
		qSort := qrSort
		if qSort != "" {
			if err := r.SetQueryParam("sort", qSort); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
