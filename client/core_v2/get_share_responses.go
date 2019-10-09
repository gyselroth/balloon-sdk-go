// Code generated by go-swagger; DO NOT EDIT.

package core_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gyselroth/balloon-sdk-go/models"
)

// GetShareReader is a Reader for the GetShare structure.
type GetShareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetShareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetShareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetShareForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetShareNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetShareOK creates a GetShareOK with default headers values
func NewGetShareOK() *GetShareOK {
	return &GetShareOK{}
}

/*GetShareOK handles this case with default header values.

Share
*/
type GetShareOK struct {
	Payload *models.CoreV2Share
}

func (o *GetShareOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/collections/{collection}/share][%d] getShareOK  %+v", 200, o.Payload)
}

func (o *GetShareOK) GetPayload() *models.CoreV2Share {
	return o.Payload
}

func (o *GetShareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Share)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetShareForbidden creates a GetShareForbidden with default headers values
func NewGetShareForbidden() *GetShareForbidden {
	return &GetShareForbidden{}
}

/*GetShareForbidden handles this case with default header values.

Access denied
*/
type GetShareForbidden struct {
}

func (o *GetShareForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/collections/{collection}/share][%d] getShareForbidden ", 403)
}

func (o *GetShareForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetShareNotFound creates a GetShareNotFound with default headers values
func NewGetShareNotFound() *GetShareNotFound {
	return &GetShareNotFound{}
}

/*GetShareNotFound handles this case with default header values.

Resource does not exists
*/
type GetShareNotFound struct {
}

func (o *GetShareNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/collections/{collection}/share][%d] getShareNotFound ", 404)
}

func (o *GetShareNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}