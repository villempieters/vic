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

// ContainerRenameReader is a Reader for the ContainerRename structure.
type ContainerRenameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContainerRenameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContainerRenameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewContainerRenameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 409:
		result := NewContainerRenameConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewContainerRenameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContainerRenameOK creates a ContainerRenameOK with default headers values
func NewContainerRenameOK() *ContainerRenameOK {
	return &ContainerRenameOK{}
}

/*ContainerRenameOK handles this case with default header values.

OK
*/
type ContainerRenameOK struct {
	Payload string
}

func (o *ContainerRenameOK) Error() string {
	return fmt.Sprintf("[PATCH /containers/{handle}/name][%d] containerRenameOK  %+v", 200, o.Payload)
}

func (o *ContainerRenameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRenameNotFound creates a ContainerRenameNotFound with default headers values
func NewContainerRenameNotFound() *ContainerRenameNotFound {
	return &ContainerRenameNotFound{}
}

/*ContainerRenameNotFound handles this case with default header values.

no such container
*/
type ContainerRenameNotFound struct {
	Payload *models.Error
}

func (o *ContainerRenameNotFound) Error() string {
	return fmt.Sprintf("[PATCH /containers/{handle}/name][%d] containerRenameNotFound  %+v", 404, o.Payload)
}

func (o *ContainerRenameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRenameConflict creates a ContainerRenameConflict with default headers values
func NewContainerRenameConflict() *ContainerRenameConflict {
	return &ContainerRenameConflict{}
}

/*ContainerRenameConflict handles this case with default header values.

conflict, name already assigned
*/
type ContainerRenameConflict struct {
	Payload *models.Error
}

func (o *ContainerRenameConflict) Error() string {
	return fmt.Sprintf("[PATCH /containers/{handle}/name][%d] containerRenameConflict  %+v", 409, o.Payload)
}

func (o *ContainerRenameConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContainerRenameInternalServerError creates a ContainerRenameInternalServerError with default headers values
func NewContainerRenameInternalServerError() *ContainerRenameInternalServerError {
	return &ContainerRenameInternalServerError{}
}

/*ContainerRenameInternalServerError handles this case with default header values.

server error
*/
type ContainerRenameInternalServerError struct {
	Payload *models.Error
}

func (o *ContainerRenameInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /containers/{handle}/name][%d] containerRenameInternalServerError  %+v", 500, o.Payload)
}

func (o *ContainerRenameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
