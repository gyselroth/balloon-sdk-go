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

// NewGetUserGroupsParams creates a new GetUserGroupsParams object
// with the default values initialized.
func NewGetUserGroupsParams() *GetUserGroupsParams {
	var ()
	return &GetUserGroupsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetUserGroupsParamsWithTimeout creates a new GetUserGroupsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetUserGroupsParamsWithTimeout(timeout time.Duration) *GetUserGroupsParams {
	var ()
	return &GetUserGroupsParams{

		timeout: timeout,
	}
}

// NewGetUserGroupsParamsWithContext creates a new GetUserGroupsParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetUserGroupsParamsWithContext(ctx context.Context) *GetUserGroupsParams {
	var ()
	return &GetUserGroupsParams{

		Context: ctx,
	}
}

// NewGetUserGroupsParamsWithHTTPClient creates a new GetUserGroupsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetUserGroupsParamsWithHTTPClient(client *http.Client) *GetUserGroupsParams {
	var ()
	return &GetUserGroupsParams{
		HTTPClient: client,
	}
}

/*GetUserGroupsParams contains all the parameters to send to the API endpoint
for the get user groups operation typically these are written to a http.Request
*/
type GetUserGroupsParams struct {

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
	/*User
	  Resource identifier

	*/
	User string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get user groups params
func (o *GetUserGroupsParams) WithTimeout(timeout time.Duration) *GetUserGroupsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get user groups params
func (o *GetUserGroupsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get user groups params
func (o *GetUserGroupsParams) WithContext(ctx context.Context) *GetUserGroupsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get user groups params
func (o *GetUserGroupsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get user groups params
func (o *GetUserGroupsParams) WithHTTPClient(client *http.Client) *GetUserGroupsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get user groups params
func (o *GetUserGroupsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAttributes adds the attributes to the get user groups params
func (o *GetUserGroupsParams) WithAttributes(attributes []string) *GetUserGroupsParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the get user groups params
func (o *GetUserGroupsParams) SetAttributes(attributes []string) {
	o.Attributes = attributes
}

// WithLimit adds the limit to the get user groups params
func (o *GetUserGroupsParams) WithLimit(limit *int64) *GetUserGroupsParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the get user groups params
func (o *GetUserGroupsParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOffset adds the offset to the get user groups params
func (o *GetUserGroupsParams) WithOffset(offset *int64) *GetUserGroupsParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get user groups params
func (o *GetUserGroupsParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithQuery adds the query to the get user groups params
func (o *GetUserGroupsParams) WithQuery(query *string) *GetUserGroupsParams {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the get user groups params
func (o *GetUserGroupsParams) SetQuery(query *string) {
	o.Query = query
}

// WithSort adds the sort to the get user groups params
func (o *GetUserGroupsParams) WithSort(sort *string) *GetUserGroupsParams {
	o.SetSort(sort)
	return o
}

// SetSort adds the sort to the get user groups params
func (o *GetUserGroupsParams) SetSort(sort *string) {
	o.Sort = sort
}

// WithUser adds the user to the get user groups params
func (o *GetUserGroupsParams) WithUser(user string) *GetUserGroupsParams {
	o.SetUser(user)
	return o
}

// SetUser adds the user to the get user groups params
func (o *GetUserGroupsParams) SetUser(user string) {
	o.User = user
}

// WriteToRequest writes these params to a swagger request
func (o *GetUserGroupsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user
	if err := r.SetPathParam("user", o.User); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
