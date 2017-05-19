package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewRemoveContainerParams creates a new RemoveContainerParams object
// with the default values initialized.
func NewRemoveContainerParams() RemoveContainerParams {
	var ()
	return RemoveContainerParams{}
}

// RemoveContainerParams contains all the bound params for the remove container operation
// typically these are obtained from a http.Request
//
// swagger:parameters RemoveContainer
type RemoveContainerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: path
	*/
	Handle string
	/*
	  Required: true
	  In: path
	*/
	Scope string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *RemoveContainerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rHandle, rhkHandle, _ := route.Params.GetOK("handle")
	if err := o.bindHandle(rHandle, rhkHandle, route.Formats); err != nil {
		res = append(res, err)
	}

	rScope, rhkScope, _ := route.Params.GetOK("scope")
	if err := o.bindScope(rScope, rhkScope, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *RemoveContainerParams) bindHandle(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Handle = raw

	return nil
}

func (o *RemoveContainerParams) bindScope(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Scope = raw

	return nil
}
