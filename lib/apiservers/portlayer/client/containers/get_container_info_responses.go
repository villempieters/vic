package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// GetContainerInfoReader is a Reader for the GetContainerInfo structure.
type GetContainerInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetContainerInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetContainerInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetContainerInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetContainerInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetContainerInfoOK creates a GetContainerInfoOK with default headers values
func NewGetContainerInfoOK() *GetContainerInfoOK {
	return &GetContainerInfoOK{}
}

/*GetContainerInfoOK handles this case with default header values.

OK
*/
type GetContainerInfoOK struct {
	Payload *models.ContainerInfo
}

func (o *GetContainerInfoOK) Error() string {
	return fmt.Sprintf("[GET /containers/info/{id}][%d] getContainerInfoOK  %+v", 200, o.Payload)
}

func (o *GetContainerInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContainerInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContainerInfoNotFound creates a GetContainerInfoNotFound with default headers values
func NewGetContainerInfoNotFound() *GetContainerInfoNotFound {
	return &GetContainerInfoNotFound{}
}

/*GetContainerInfoNotFound handles this case with default header values.

not found
*/
type GetContainerInfoNotFound struct {
	Payload *models.Error
}

func (o *GetContainerInfoNotFound) Error() string {
	return fmt.Sprintf("[GET /containers/info/{id}][%d] getContainerInfoNotFound  %+v", 404, o.Payload)
}

func (o *GetContainerInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetContainerInfoInternalServerError creates a GetContainerInfoInternalServerError with default headers values
func NewGetContainerInfoInternalServerError() *GetContainerInfoInternalServerError {
	return &GetContainerInfoInternalServerError{}
}

/*GetContainerInfoInternalServerError handles this case with default header values.

server error
*/
type GetContainerInfoInternalServerError struct {
	Payload *models.Error
}

func (o *GetContainerInfoInternalServerError) Error() string {
	return fmt.Sprintf("[GET /containers/info/{id}][%d] getContainerInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *GetContainerInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
