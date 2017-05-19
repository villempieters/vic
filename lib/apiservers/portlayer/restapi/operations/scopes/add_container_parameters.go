package scopes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/vmware/vic/lib/apiservers/portlayer/models"
)

// NewAddContainerParams creates a new AddContainerParams object
// with the default values initialized.
func NewAddContainerParams() AddContainerParams {
	var ()
	return AddContainerParams{}
}

// AddContainerParams contains all the bound params for the add container operation
// typically these are obtained from a http.Request
//
// swagger:parameters AddContainer
type AddContainerParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*
	  Required: true
	  In: body
	*/
	Config *models.ScopesAddContainerConfig
	/*
	  Required: true
	  In: path
	*/
	Scope string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *AddContainerParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	if runtime.HasBody(r) {
		defer r.Body.Close()
		var body models.ScopesAddContainerConfig
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

	rScope, rhkScope, _ := route.Params.GetOK("scope")
	if err := o.bindScope(rScope, rhkScope, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddContainerParams) bindScope(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Scope = raw

	return nil
}
