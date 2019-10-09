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

// GetNodeAttributeSummaryReader is a Reader for the GetNodeAttributeSummary structure.
type GetNodeAttributeSummaryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodeAttributeSummaryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNodeAttributeSummaryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetNodeAttributeSummaryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetNodeAttributeSummaryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetNodeAttributeSummaryOK creates a GetNodeAttributeSummaryOK with default headers values
func NewGetNodeAttributeSummaryOK() *GetNodeAttributeSummaryOK {
	return &GetNodeAttributeSummaryOK{}
}

/*GetNodeAttributeSummaryOK handles this case with default header values.

User
*/
type GetNodeAttributeSummaryOK struct {
	Payload models.CoreV2UserNodeAttributeSummary
}

func (o *GetNodeAttributeSummaryOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user}/node-attribute-summary][%d] getNodeAttributeSummaryOK  %+v", 200, o.Payload)
}

func (o *GetNodeAttributeSummaryOK) GetPayload() models.CoreV2UserNodeAttributeSummary {
	return o.Payload
}

func (o *GetNodeAttributeSummaryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodeAttributeSummaryForbidden creates a GetNodeAttributeSummaryForbidden with default headers values
func NewGetNodeAttributeSummaryForbidden() *GetNodeAttributeSummaryForbidden {
	return &GetNodeAttributeSummaryForbidden{}
}

/*GetNodeAttributeSummaryForbidden handles this case with default header values.

Access denied
*/
type GetNodeAttributeSummaryForbidden struct {
}

func (o *GetNodeAttributeSummaryForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user}/node-attribute-summary][%d] getNodeAttributeSummaryForbidden ", 403)
}

func (o *GetNodeAttributeSummaryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetNodeAttributeSummaryNotFound creates a GetNodeAttributeSummaryNotFound with default headers values
func NewGetNodeAttributeSummaryNotFound() *GetNodeAttributeSummaryNotFound {
	return &GetNodeAttributeSummaryNotFound{}
}

/*GetNodeAttributeSummaryNotFound handles this case with default header values.

Resource does not exists
*/
type GetNodeAttributeSummaryNotFound struct {
}

func (o *GetNodeAttributeSummaryNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/users/{user}/node-attribute-summary][%d] getNodeAttributeSummaryNotFound ", 404)
}

func (o *GetNodeAttributeSummaryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
