package interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

/*ContainerCloseStdinOK OK

swagger:response containerCloseStdinOK
*/
type ContainerCloseStdinOK struct {
}

// NewContainerCloseStdinOK creates ContainerCloseStdinOK with default headers values
func NewContainerCloseStdinOK() *ContainerCloseStdinOK {
	return &ContainerCloseStdinOK{}
}

// WriteResponse to the client
func (o *ContainerCloseStdinOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*ContainerCloseStdinNotFound Container not found

swagger:response containerCloseStdinNotFound
*/
type ContainerCloseStdinNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewContainerCloseStdinNotFound creates ContainerCloseStdinNotFound with default headers values
func NewContainerCloseStdinNotFound() *ContainerCloseStdinNotFound {
	return &ContainerCloseStdinNotFound{}
}

// WithPayload adds the payload to the container close stdin not found response
func (o *ContainerCloseStdinNotFound) WithPayload(payload *models.Error) *ContainerCloseStdinNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container close stdin not found response
func (o *ContainerCloseStdinNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerCloseStdinNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*ContainerCloseStdinInternalServerError Failed to Close stdin

swagger:response containerCloseStdinInternalServerError
*/
type ContainerCloseStdinInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewContainerCloseStdinInternalServerError creates ContainerCloseStdinInternalServerError with default headers values
func NewContainerCloseStdinInternalServerError() *ContainerCloseStdinInternalServerError {
	return &ContainerCloseStdinInternalServerError{}
}

// WithPayload adds the payload to the container close stdin internal server error response
func (o *ContainerCloseStdinInternalServerError) WithPayload(payload *models.Error) *ContainerCloseStdinInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the container close stdin internal server error response
func (o *ContainerCloseStdinInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ContainerCloseStdinInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
