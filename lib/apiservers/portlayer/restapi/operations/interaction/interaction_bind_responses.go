package interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

/*InteractionBindOK OK

swagger:response interactionBindOK
*/
type InteractionBindOK struct {

	/*
	  In: Body
	*/
	Payload *models.InteractionBindResponse `json:"body,omitempty"`
}

// NewInteractionBindOK creates InteractionBindOK with default headers values
func NewInteractionBindOK() *InteractionBindOK {
	return &InteractionBindOK{}
}

// WithPayload adds the payload to the interaction bind o k response
func (o *InteractionBindOK) WithPayload(payload *models.InteractionBindResponse) *InteractionBindOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the interaction bind o k response
func (o *InteractionBindOK) SetPayload(payload *models.InteractionBindResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InteractionBindOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*InteractionBindNotFound VirtualDevice not found

swagger:response interactionBindNotFound
*/
type InteractionBindNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewInteractionBindNotFound creates InteractionBindNotFound with default headers values
func NewInteractionBindNotFound() *InteractionBindNotFound {
	return &InteractionBindNotFound{}
}

// WithPayload adds the payload to the interaction bind not found response
func (o *InteractionBindNotFound) WithPayload(payload *models.Error) *InteractionBindNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the interaction bind not found response
func (o *InteractionBindNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InteractionBindNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*InteractionBindInternalServerError Connecting VirtualDevice failed

swagger:response interactionBindInternalServerError
*/
type InteractionBindInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewInteractionBindInternalServerError creates InteractionBindInternalServerError with default headers values
func NewInteractionBindInternalServerError() *InteractionBindInternalServerError {
	return &InteractionBindInternalServerError{}
}

// WithPayload adds the payload to the interaction bind internal server error response
func (o *InteractionBindInternalServerError) WithPayload(payload *models.Error) *InteractionBindInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the interaction bind internal server error response
func (o *InteractionBindInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *InteractionBindInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
