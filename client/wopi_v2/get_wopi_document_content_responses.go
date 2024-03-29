// Code generated by go-swagger; DO NOT EDIT.

package wopi_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gyselroth/balloon-sdk-go/models"
)

// GetWopiDocumentContentReader is a Reader for the GetWopiDocumentContent structure.
type GetWopiDocumentContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWopiDocumentContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWopiDocumentContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWopiDocumentContentBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWopiDocumentContentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWopiDocumentContentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWopiDocumentContentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewGetWopiDocumentContentUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWopiDocumentContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWopiDocumentContentOK creates a GetWopiDocumentContentOK with default headers values
func NewGetWopiDocumentContentOK() *GetWopiDocumentContentOK {
	return &GetWopiDocumentContentOK{}
}

/*GetWopiDocumentContentOK handles this case with default header values.

Contents
*/
type GetWopiDocumentContentOK struct {
}

func (o *GetWopiDocumentContentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}/contents][%d] getWopiDocumentContentOK ", 200)
}

func (o *GetWopiDocumentContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWopiDocumentContentBadRequest creates a GetWopiDocumentContentBadRequest with default headers values
func NewGetWopiDocumentContentBadRequest() *GetWopiDocumentContentBadRequest {
	return &GetWopiDocumentContentBadRequest{}
}

/*GetWopiDocumentContentBadRequest handles this case with default header values.

Bad Reqeust
*/
type GetWopiDocumentContentBadRequest struct {
	Payload *models.CoreV2Error
}

func (o *GetWopiDocumentContentBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}/contents][%d] getWopiDocumentContentBadRequest  %+v", 400, o.Payload)
}

func (o *GetWopiDocumentContentBadRequest) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetWopiDocumentContentBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWopiDocumentContentUnauthorized creates a GetWopiDocumentContentUnauthorized with default headers values
func NewGetWopiDocumentContentUnauthorized() *GetWopiDocumentContentUnauthorized {
	return &GetWopiDocumentContentUnauthorized{}
}

/*GetWopiDocumentContentUnauthorized handles this case with default header values.

Unauthorized
*/
type GetWopiDocumentContentUnauthorized struct {
	Payload *models.CoreV2Error
}

func (o *GetWopiDocumentContentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}/contents][%d] getWopiDocumentContentUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWopiDocumentContentUnauthorized) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetWopiDocumentContentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWopiDocumentContentForbidden creates a GetWopiDocumentContentForbidden with default headers values
func NewGetWopiDocumentContentForbidden() *GetWopiDocumentContentForbidden {
	return &GetWopiDocumentContentForbidden{}
}

/*GetWopiDocumentContentForbidden handles this case with default header values.

Forbidden
*/
type GetWopiDocumentContentForbidden struct {
	Payload *models.CoreV2Error
}

func (o *GetWopiDocumentContentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}/contents][%d] getWopiDocumentContentForbidden  %+v", 403, o.Payload)
}

func (o *GetWopiDocumentContentForbidden) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetWopiDocumentContentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWopiDocumentContentNotFound creates a GetWopiDocumentContentNotFound with default headers values
func NewGetWopiDocumentContentNotFound() *GetWopiDocumentContentNotFound {
	return &GetWopiDocumentContentNotFound{}
}

/*GetWopiDocumentContentNotFound handles this case with default header values.

The specified resource was not found
*/
type GetWopiDocumentContentNotFound struct {
	Payload *models.CoreV2Error
}

func (o *GetWopiDocumentContentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}/contents][%d] getWopiDocumentContentNotFound  %+v", 404, o.Payload)
}

func (o *GetWopiDocumentContentNotFound) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetWopiDocumentContentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWopiDocumentContentUnprocessableEntity creates a GetWopiDocumentContentUnprocessableEntity with default headers values
func NewGetWopiDocumentContentUnprocessableEntity() *GetWopiDocumentContentUnprocessableEntity {
	return &GetWopiDocumentContentUnprocessableEntity{}
}

/*GetWopiDocumentContentUnprocessableEntity handles this case with default header values.

Unauthorized
*/
type GetWopiDocumentContentUnprocessableEntity struct {
	Payload *models.CoreV2Error
}

func (o *GetWopiDocumentContentUnprocessableEntity) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}/contents][%d] getWopiDocumentContentUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *GetWopiDocumentContentUnprocessableEntity) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetWopiDocumentContentUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWopiDocumentContentInternalServerError creates a GetWopiDocumentContentInternalServerError with default headers values
func NewGetWopiDocumentContentInternalServerError() *GetWopiDocumentContentInternalServerError {
	return &GetWopiDocumentContentInternalServerError{}
}

/*GetWopiDocumentContentInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetWopiDocumentContentInternalServerError struct {
	Payload *models.CoreV2Error
}

func (o *GetWopiDocumentContentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}/contents][%d] getWopiDocumentContentInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWopiDocumentContentInternalServerError) GetPayload() *models.CoreV2Error {
	return o.Payload
}

func (o *GetWopiDocumentContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
