// Code generated by go-swagger; DO NOT EDIT.

package elasticsearch_v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/gyselroth/balloon-sdk-go/models"
)

// QueryElasticsearchReader is a Reader for the QueryElasticsearch structure.
type QueryElasticsearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryElasticsearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryElasticsearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQueryElasticsearchOK creates a QueryElasticsearchOK with default headers values
func NewQueryElasticsearchOK() *QueryElasticsearchOK {
	return &QueryElasticsearchOK{}
}

/*QueryElasticsearchOK handles this case with default header values.

Nodes
*/
type QueryElasticsearchOK struct {
	Payload *models.CoreV2Nodes
}

func (o *QueryElasticsearchOK) Error() string {
	return fmt.Sprintf("[GET /api/v2/nodes/search][%d] queryElasticsearchOK  %+v", 200, o.Payload)
}

func (o *QueryElasticsearchOK) GetPayload() *models.CoreV2Nodes {
	return o.Payload
}

func (o *QueryElasticsearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CoreV2Nodes)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
