package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

/*UnbindContainerOK OK

swagger:response unbindContainerOK
*/
type UnbindContainerOK struct {

	/*
	  In: Body
	*/
	Payload *models.UnbindContainerResponse `json:"body,omitempty"`
}

// NewUnbindContainerOK creates UnbindContainerOK with default headers values
func NewUnbindContainerOK() *UnbindContainerOK {
	return &UnbindContainerOK{}
}

// WithPayload adds the payload to the unbind container o k response
func (o *UnbindContainerOK) WithPayload(payload *models.UnbindContainerResponse) *UnbindContainerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unbind container o k response
func (o *UnbindContainerOK) SetPayload(payload *models.UnbindContainerResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnbindContainerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UnbindContainerNotFound Not found

swagger:response unbindContainerNotFound
*/
type UnbindContainerNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUnbindContainerNotFound creates UnbindContainerNotFound with default headers values
func NewUnbindContainerNotFound() *UnbindContainerNotFound {
	return &UnbindContainerNotFound{}
}

// WithPayload adds the payload to the unbind container not found response
func (o *UnbindContainerNotFound) WithPayload(payload *models.Error) *UnbindContainerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unbind container not found response
func (o *UnbindContainerNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnbindContainerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UnbindContainerInternalServerError error

swagger:response unbindContainerInternalServerError
*/
type UnbindContainerInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUnbindContainerInternalServerError creates UnbindContainerInternalServerError with default headers values
func NewUnbindContainerInternalServerError() *UnbindContainerInternalServerError {
	return &UnbindContainerInternalServerError{}
}

// WithPayload adds the payload to the unbind container internal server error response
func (o *UnbindContainerInternalServerError) WithPayload(payload *models.Error) *UnbindContainerInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the unbind container internal server error response
func (o *UnbindContainerInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UnbindContainerInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
