// Code generated by go-swagger; DO NOT EDIT.

package wopi_v2

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
)

// NewGetWopiDocumentParams creates a new GetWopiDocumentParams object
// with the default values initialized.
func NewGetWopiDocumentParams() *GetWopiDocumentParams {
	var ()
	return &GetWopiDocumentParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWopiDocumentParamsWithTimeout creates a new GetWopiDocumentParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWopiDocumentParamsWithTimeout(timeout time.Duration) *GetWopiDocumentParams {
	var ()
	return &GetWopiDocumentParams{

		timeout: timeout,
	}
}

// NewGetWopiDocumentParamsWithContext creates a new GetWopiDocumentParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWopiDocumentParamsWithContext(ctx context.Context) *GetWopiDocumentParams {
	var ()
	return &GetWopiDocumentParams{

		Context: ctx,
	}
}

// NewGetWopiDocumentParamsWithHTTPClient creates a new GetWopiDocumentParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWopiDocumentParamsWithHTTPClient(client *http.Client) *GetWopiDocumentParams {
	var ()
	return &GetWopiDocumentParams{
		HTTPClient: client,
	}
}

/*GetWopiDocumentParams contains all the parameters to send to the API endpoint
for the get wopi document operation typically these are written to a http.Request
*/
type GetWopiDocumentParams struct {

	/*AccessToken
	  An access token to access the document

	*/
	AccessToken string
	/*Document
	  The document id

	*/
	Document string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get wopi document params
func (o *GetWopiDocumentParams) WithTimeout(timeout time.Duration) *GetWopiDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get wopi document params
func (o *GetWopiDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get wopi document params
func (o *GetWopiDocumentParams) WithContext(ctx context.Context) *GetWopiDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get wopi document params
func (o *GetWopiDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get wopi document params
func (o *GetWopiDocumentParams) WithHTTPClient(client *http.Client) *GetWopiDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get wopi document params
func (o *GetWopiDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessToken adds the accessToken to the get wopi document params
func (o *GetWopiDocumentParams) WithAccessToken(accessToken string) *GetWopiDocumentParams {
	o.SetAccessToken(accessToken)
	return o
}

// SetAccessToken adds the accessToken to the get wopi document params
func (o *GetWopiDocumentParams) SetAccessToken(accessToken string) {
	o.AccessToken = accessToken
}

// WithDocument adds the document to the get wopi document params
func (o *GetWopiDocumentParams) WithDocument(document string) *GetWopiDocumentParams {
	o.SetDocument(document)
	return o
}

// SetDocument adds the document to the get wopi document params
func (o *GetWopiDocumentParams) SetDocument(document string) {
	o.Document = document
}

// WriteToRequest writes these params to a swagger request
func (o *GetWopiDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param access_token
	qrAccessToken := o.AccessToken
	qAccessToken := qrAccessToken
	if qAccessToken != "" {
		if err := r.SetQueryParam("access_token", qAccessToken); err != nil {
			return err
		}
	}

	// path param document
	if err := r.SetPathParam("document", o.Document); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
