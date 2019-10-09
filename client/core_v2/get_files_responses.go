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

// GetFilesReader is a Reader for the GetFiles structure.
type GetFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetFilesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetFilesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFilesOK creates a GetFilesOK with default headers values
func NewGetFilesOK() *GetFilesOK {
	return &GetFilesOK{}
}

/*GetFilesOK handles this case with default header values.

List of files
*/
type GetFilesOK struct {
	Payload *models.CoreV2Files
}

func (o *GetFilesOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/files][%d] getFilesOK  %+v", 200, o.Payload)
}

func (o *GetFilesOK) GetPayload() *models.CoreV2Files {
	return o.Payload
}

func (o *GetFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Files)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFilesForbidden creates a GetFilesForbidden with default headers values
func NewGetFilesForbidden() *GetFilesForbidden {
	return &GetFilesForbidden{}
}

/*GetFilesForbidden handles this case with default header values.

Access denied
*/
type GetFilesForbidden struct {
}

func (o *GetFilesForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v2/files][%d] getFilesForbidden ", 403)
}

func (o *GetFilesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetFilesNotFound creates a GetFilesNotFound with default headers values
func NewGetFilesNotFound() *GetFilesNotFound {
	return &GetFilesNotFound{}
}

/*GetFilesNotFound handles this case with default header values.

Resource does not exists
*/
type GetFilesNotFound struct {
}

func (o *GetFilesNotFound) Error() string {
	return fmt.Sprintf("[GET /api/v2/files][%d] getFilesNotFound ", 404)
}

func (o *GetFilesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
