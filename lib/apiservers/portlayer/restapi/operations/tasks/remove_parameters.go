package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// NewRemoveParams creates a new RemoveParams object
// with the default values initialized.
func NewRemoveParams() RemoveParams {
	var ()
	return RemoveParams{}
}

// RemoveParams contains all the bound params for the remove operation
// typically these are obtained from a http.Request
//
// swagger:parameters Remove
type RemoveParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: body
	*/
	Config *models.TaskRemoveConfig
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *RemoveParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.TaskRemoveConfig
		if err := route.Consumer.Consume(r.Body, &body); err != nil {
			if err == io.EOF {
				res = append(res, errors.Required("config", "body"))
			} else {
				res = append(res, errors.NewParseError("config", "body", "", err))
			}

		} else {
			if err := body.Validate(route.Formats); err != nil {
				res = append(res, err)
			}

			if len(res) == 0 {
				o.Config = &body
			}
		}

	} else {
		res = append(res, errors.Required("config", "body"))
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
