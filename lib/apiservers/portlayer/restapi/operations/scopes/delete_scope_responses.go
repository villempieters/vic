package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

/*DeleteScopeOK OK

swagger:response deleteScopeOK
*/
type DeleteScopeOK struct {
}

// NewDeleteScopeOK creates DeleteScopeOK with default headers values
func NewDeleteScopeOK() *DeleteScopeOK {
	return &DeleteScopeOK{}
}

// WriteResponse to the client
func (o *DeleteScopeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
}

/*DeleteScopeNotFound Not found

swagger:response deleteScopeNotFound
*/
type DeleteScopeNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteScopeNotFound creates DeleteScopeNotFound with default headers values
func NewDeleteScopeNotFound() *DeleteScopeNotFound {
	return &DeleteScopeNotFound{}
}

// WithPayload adds the payload to the delete scope not found response
func (o *DeleteScopeNotFound) WithPayload(payload *models.Error) *DeleteScopeNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete scope not found response
func (o *DeleteScopeNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteScopeNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteScopeInternalServerError Internal server error

swagger:response deleteScopeInternalServerError
*/
type DeleteScopeInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteScopeInternalServerError creates DeleteScopeInternalServerError with default headers values
func NewDeleteScopeInternalServerError() *DeleteScopeInternalServerError {
	return &DeleteScopeInternalServerError{}
}

// WithPayload adds the payload to the delete scope internal server error response
func (o *DeleteScopeInternalServerError) WithPayload(payload *models.Error) *DeleteScopeInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete scope internal server error response
func (o *DeleteScopeInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteScopeInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
