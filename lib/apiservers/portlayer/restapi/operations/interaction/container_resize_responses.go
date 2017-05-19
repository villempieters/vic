package interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

/*ContainerResizeOK OK

swagger:response containerResizeOK
*/
type ContainerResizeOK struct {
}

// NewContainerResizeOK creates ContainerResizeOK with default headers values
func NewContainerResizeOK() *ContainerResizeOK {
	return &ContainerResizeOK{}
}

// WriteResponse to the client
func (o *ContainerResizeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*ContainerResizeNotFound Container not found

swagger:response containerResizeNotFound
*/
type ContainerResizeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewContainerResizeNotFound creates ContainerResizeNotFound with default headers values
func NewContainerResizeNotFound() *ContainerResizeNotFound {
	return &ContainerResizeNotFound{}
}

// WithPayload adds the payload to the container resize not found response
func (o *ContainerResizeNotFound) WithPayload(payload *models.Error) *ContainerResizeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container resize not found response
func (o *ContainerResizeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerResizeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ContainerResizeInternalServerError Container resize failed

swagger:response containerResizeInternalServerError
*/
type ContainerResizeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewContainerResizeInternalServerError creates ContainerResizeInternalServerError with default headers values
func NewContainerResizeInternalServerError() *ContainerResizeInternalServerError {
	return &ContainerResizeInternalServerError{}
}

// WithPayload adds the payload to the container resize internal server error response
func (o *ContainerResizeInternalServerError) WithPayload(payload *models.Error) *ContainerResizeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container resize internal server error response
func (o *ContainerResizeInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerResizeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
