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

// NewGetNodeEventLogParams creates a new GetNodeEventLogParams object
// with the default values initialized.
func NewGetNodeEventLogParams() *GetNodeEventLogParams {
	var ()
	return &GetNodeEventLogParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetNodeEventLogParamsWithTimeout creates a new GetNodeEventLogParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetNodeEventLogParamsWithTimeout(timeout time.Duration) *GetNodeEventLogParams {
	var ()
	return &GetNodeEventLogParams{

		timeout: timeout,
	}
}

// NewGetNodeEventLogParamsWithContext creates a new GetNodeEventLogParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetNodeEventLogParamsWithContext(ctx context.Context) *GetNodeEventLogParams {
	var ()
	return &GetNodeEventLogParams{

		Context: ctx,
	}
}

// NewGetNodeEventLogParamsWithHTTPClient creates a new GetNodeEventLogParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetNodeEventLogParamsWithHTTPClient(client *http.Client) *GetNodeEventLogParams {
	var ()
	return &GetNodeEventLogParams{
		HTTPClient: client,
	}
}

/*GetNodeEventLogParams contains all the parameters to send to the API endpoint
for the get node event log operation typically these are written to a http.Request
*/
type GetNodeEventLogParams struct {

	/*Attributes
	  Filter attributes

	*/
	Attributes []string
	/*Limit
	  Objects limit, per default 20 objects will get returned

	*/
	Limit *float64
	/*Node
	  Node identifier

	*/
	Node string
	/*Offset
	  Objects offset, per default it starts from 0. You may also request a negative offset which will return results from the end [total - offset].

	*/
	Offset *float64
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

// WithTimeout adds the timeout to the get node event log params
func (o *GetNodeEventLogParams) WithTimeout(timeout time.Duration) *GetNodeEventLogParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get node event log params
func (o *GetNodeEventLogParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get node event log params
func (o *GetNodeEventLogParams) WithContext(ctx context.Context) *GetNodeEventLogParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get node event log params
func (o *GetNodeEventLogParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get node event log params
func (o *GetNodeEventLogParams) WithHTTPClient(client *http.Client) *GetNodeEventLogParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get node event log params
func (o *GetNodeEventLogParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributes adds the attributes to the get node event log params
func (o *GetNodeEventLogParams) WithAttributes(attributes []string) *GetNodeEventLogParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get node event log params
func (o *GetNodeEventLogParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithLimit adds the limit to the get node event log params
func (o *GetNodeEventLogParams) WithLimit(limit *float64) *GetNodeEventLogParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get node event log params
func (o *GetNodeEventLogParams) SetLimit(limit *float64) {
	o.Limit = limit
}

// WithNode adds the node to the get node event log params
func (o *GetNodeEventLogParams) WithNode(node string) *GetNodeEventLogParams {
	o.SetNode(node)
	return o
}

// SetNode adds the node to the get node event log params
func (o *GetNodeEventLogParams) SetNode(node string) {
	o.Node = node
}

// WithOffset adds the offset to the get node event log params
func (o *GetNodeEventLogParams) WithOffset(offset *float64) *GetNodeEventLogParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get node event log params
func (o *GetNodeEventLogParams) SetOffset(offset *float64) {
	o.Offset = offset
}

// WithQuery adds the query to the get node event log params
func (o *GetNodeEventLogParams) WithQuery(query *string) *GetNodeEventLogParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get node event log params
func (o *GetNodeEventLogParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the get node event log params
func (o *GetNodeEventLogParams) WithSort(sort *string) *GetNodeEventLogParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get node event log params
func (o *GetNodeEventLogParams) SetSort(sort *string) {
	o.Sort = sort
}

// WriteToRequest writes these params to a swagger request
func (o *GetNodeEventLogParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
		var qrLimit float64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatFloat64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param node
	if err := r.SetPathParam("node", o.Node); err != nil {
		return err
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset float64
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatFloat64(qrOffset)
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