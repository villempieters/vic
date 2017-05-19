package logging

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

/*LoggingUnbindOK OK

swagger:response loggingUnbindOK
*/
type LoggingUnbindOK struct {

	/*
	  In: Body
	*/
	Payload *models.LoggingUnbindResponse `json:"body,omitempty"`
}

// NewLoggingUnbindOK creates LoggingUnbindOK with default headers values
func NewLoggingUnbindOK() *LoggingUnbindOK {
	return &LoggingUnbindOK{}
}

// WithPayload adds the payload to the logging unbind o k response
func (o *LoggingUnbindOK) WithPayload(payload *models.LoggingUnbindResponse) *LoggingUnbindOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logging unbind o k response
func (o *LoggingUnbindOK) SetPayload(payload *models.LoggingUnbindResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoggingUnbindOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*LoggingUnbindNotFound VirtualDevice not found

swagger:response loggingUnbindNotFound
*/
type LoggingUnbindNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLoggingUnbindNotFound creates LoggingUnbindNotFound with default headers values
func NewLoggingUnbindNotFound() *LoggingUnbindNotFound {
	return &LoggingUnbindNotFound{}
}

// WithPayload adds the payload to the logging unbind not found response
func (o *LoggingUnbindNotFound) WithPayload(payload *models.Error) *LoggingUnbindNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logging unbind not found response
func (o *LoggingUnbindNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoggingUnbindNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*LoggingUnbindInternalServerError Disconnecting VirtualDevice failed

swagger:response loggingUnbindInternalServerError
*/
type LoggingUnbindInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewLoggingUnbindInternalServerError creates LoggingUnbindInternalServerError with default headers values
func NewLoggingUnbindInternalServerError() *LoggingUnbindInternalServerError {
	return &LoggingUnbindInternalServerError{}
}

// WithPayload adds the payload to the logging unbind internal server error response
func (o *LoggingUnbindInternalServerError) WithPayload(payload *models.Error) *LoggingUnbindInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the logging unbind internal server error response
func (o *LoggingUnbindInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *LoggingUnbindInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
