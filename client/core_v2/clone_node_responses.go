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

// CloneNodeReader is a Reader for the CloneNode structure.
type CloneNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCloneNodeCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 207:
		result := NewCloneNodeMultiStatus()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewCloneNodeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCloneNodeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCloneNodeCreated creates a CloneNodeCreated with default headers values
func NewCloneNodeCreated() *CloneNodeCreated {
	return &CloneNodeCreated{}
}

/*CloneNodeCreated handles this case with default header values.

Node
*/
type CloneNodeCreated struct {
	Payload *models.CoreV2Node
}

func (o *CloneNodeCreated) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/clone][%d] cloneNodeCreated  %+v", 201, o.Payload)
}

func (o *CloneNodeCreated) GetPayload() *models.CoreV2Node {
	return o.Payload
}

func (o *CloneNodeCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Node)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCloneNodeMultiStatus creates a CloneNodeMultiStatus with default headers values
func NewCloneNodeMultiStatus() *CloneNodeMultiStatus {
	return &CloneNodeMultiStatus{}
}

/*CloneNodeMultiStatus handles this case with default header values.

Multi status if batch request
*/
type CloneNodeMultiStatus struct {
}

func (o *CloneNodeMultiStatus) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/clone][%d] cloneNodeMultiStatus ", 207)
}

func (o *CloneNodeMultiStatus) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneNodeForbidden creates a CloneNodeForbidden with default headers values
func NewCloneNodeForbidden() *CloneNodeForbidden {
	return &CloneNodeForbidden{}
}

/*CloneNodeForbidden handles this case with default header values.

Access denied
*/
type CloneNodeForbidden struct {
}

func (o *CloneNodeForbidden) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/clone][%d] cloneNodeForbidden ", 403)
}

func (o *CloneNodeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCloneNodeNotFound creates a CloneNodeNotFound with default headers values
func NewCloneNodeNotFound() *CloneNodeNotFound {
	return &CloneNodeNotFound{}
}

/*CloneNodeNotFound handles this case with default header values.

Resource does not exists
*/
type CloneNodeNotFound struct {
}

func (o *CloneNodeNotFound) Error() string {
	return fmt.Sprintf("[POST /api/v2/nodes/{node}/clone][%d] cloneNodeNotFound ", 404)
}

func (o *CloneNodeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}