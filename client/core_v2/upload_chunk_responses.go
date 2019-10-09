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

// UploadChunkReader is a Reader for the UploadChunk structure.
type UploadChunkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadChunkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadChunkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 201:
		result := NewUploadChunkCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 206:
		result := NewUploadChunkPartialContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewUploadChunkForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadChunkNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadChunkOK creates a UploadChunkOK with default headers values
func NewUploadChunkOK() *UploadChunkOK {
	return &UploadChunkOK{}
}

/*UploadChunkOK handles this case with default header values.

File updated
*/
type UploadChunkOK struct {
	Payload *models.CoreV2File
}

func (o *UploadChunkOK) Error() string {
	return fmt.Sprintf("[PUT /api/v2/files/chunk][%d] uploadChunkOK  %+v", 200, o.Payload)
}

func (o *UploadChunkOK) GetPayload() *models.CoreV2File {
	return o.Payload
}

func (o *UploadChunkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2File)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadChunkCreated creates a UploadChunkCreated with default headers values
func NewUploadChunkCreated() *UploadChunkCreated {
	return &UploadChunkCreated{}
}

/*UploadChunkCreated handles this case with default header values.

File newly created
*/
type UploadChunkCreated struct {
	Payload *models.CoreV2File
}

func (o *UploadChunkCreated) Error() string {
	return fmt.Sprintf("[PUT /api/v2/files/chunk][%d] uploadChunkCreated  %+v", 201, o.Payload)
}

func (o *UploadChunkCreated) GetPayload() *models.CoreV2File {
	return o.Payload
}

func (o *UploadChunkCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2File)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadChunkPartialContent creates a UploadChunkPartialContent with default headers values
func NewUploadChunkPartialContent() *UploadChunkPartialContent {
	return &UploadChunkPartialContent{}
}

/*UploadChunkPartialContent handles this case with default header values.

Chunk uploaded
*/
type UploadChunkPartialContent struct {
	Payload *models.CoreV2ChunkSession
}

func (o *UploadChunkPartialContent) Error() string {
	return fmt.Sprintf("[PUT /api/v2/files/chunk][%d] uploadChunkPartialContent  %+v", 206, o.Payload)
}

func (o *UploadChunkPartialContent) GetPayload() *models.CoreV2ChunkSession {
	return o.Payload
}

func (o *UploadChunkPartialContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2ChunkSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadChunkForbidden creates a UploadChunkForbidden with default headers values
func NewUploadChunkForbidden() *UploadChunkForbidden {
	return &UploadChunkForbidden{}
}

/*UploadChunkForbidden handles this case with default header values.

Access denied
*/
type UploadChunkForbidden struct {
}

func (o *UploadChunkForbidden) Error() string {
	return fmt.Sprintf("[PUT /api/v2/files/chunk][%d] uploadChunkForbidden ", 403)
}

func (o *UploadChunkForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUploadChunkNotFound creates a UploadChunkNotFound with default headers values
func NewUploadChunkNotFound() *UploadChunkNotFound {
	return &UploadChunkNotFound{}
}

/*UploadChunkNotFound handles this case with default header values.

Resource does not exists
*/
type UploadChunkNotFound struct {
}

func (o *UploadChunkNotFound) Error() string {
	return fmt.Sprintf("[PUT /api/v2/files/chunk][%d] uploadChunkNotFound ", 404)
}

func (o *UploadChunkNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
