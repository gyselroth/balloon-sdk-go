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

// GetWopiDocumentReader is a Reader for the GetWopiDocument structure.
type GetWopiDocumentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWopiDocumentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWopiDocumentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetWopiDocumentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWopiDocumentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWopiDocumentOK creates a GetWopiDocumentOK with default headers values
func NewGetWopiDocumentOK() *GetWopiDocumentOK {
	return &GetWopiDocumentOK{}
}

/*GetWopiDocumentOK handles this case with default header values.

WOPI document information
*/
type GetWopiDocumentOK struct {
	Payload *models.OfficeV2WopiDocument
}

func (o *GetWopiDocumentOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}][%d] getWopiDocumentOK  %+v", 200, o.Payload)
}

func (o *GetWopiDocumentOK) GetPayload() *models.OfficeV2WopiDocument {
	return o.Payload
}

func (o *GetWopiDocumentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OfficeV2WopiDocument)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWopiDocumentForbidden creates a GetWopiDocumentForbidden with default headers values
func NewGetWopiDocumentForbidden() *GetWopiDocumentForbidden {
	return &GetWopiDocumentForbidden{}
}

/*GetWopiDocumentForbidden handles this case with default header values.

Access denied
*/
type GetWopiDocumentForbidden struct {
}

func (o *GetWopiDocumentForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}][%d] getWopiDocumentForbidden ", 403)
}

func (o *GetWopiDocumentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetWopiDocumentNotFound creates a GetWopiDocumentNotFound with default headers values
func NewGetWopiDocumentNotFound() *GetWopiDocumentNotFound {
	return &GetWopiDocumentNotFound{}
}

/*GetWopiDocumentNotFound handles this case with default header values.

Resource does not exists
*/
type GetWopiDocumentNotFound struct {
}

func (o *GetWopiDocumentNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/office/wopi/document/{document}][%d] getWopiDocumentNotFound ", 404)
}

func (o *GetWopiDocumentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
