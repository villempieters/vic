package interaction

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewContainerCloseStdinParams creates a new ContainerCloseStdinParams object
// with the default values initialized.
func NewContainerCloseStdinParams() *ContainerCloseStdinParams {
	var ()
	return &ContainerCloseStdinParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewContainerCloseStdinParamsWithTimeout creates a new ContainerCloseStdinParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewContainerCloseStdinParamsWithTimeout(timeout time.Duration) *ContainerCloseStdinParams {
	var ()
	return &ContainerCloseStdinParams{

		timeout: timeout,
	}
}

// NewContainerCloseStdinParamsWithContext creates a new ContainerCloseStdinParams object
// with the default values initialized, and the ability to set a context for a request
func NewContainerCloseStdinParamsWithContext(ctx context.Context) *ContainerCloseStdinParams {
	var ()
	return &ContainerCloseStdinParams{

		Context: ctx,
	}
}

// NewContainerCloseStdinParamsWithHTTPClient creates a new ContainerCloseStdinParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewContainerCloseStdinParamsWithHTTPClient(client *http.Client) *ContainerCloseStdinParams {
	var ()
	return &ContainerCloseStdinParams{
		HTTPClient: client,
	}
}

/*ContainerCloseStdinParams contains all the parameters to send to the API endpoint
for the container close stdin operation typically these are written to a http.Request
*/
type ContainerCloseStdinParams struct {

	/*OpID*/
	OpID *string
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the container close stdin params
func (o *ContainerCloseStdinParams) WithTimeout(timeout time.Duration) *ContainerCloseStdinParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the container close stdin params
func (o *ContainerCloseStdinParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the container close stdin params
func (o *ContainerCloseStdinParams) WithContext(ctx context.Context) *ContainerCloseStdinParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the container close stdin params
func (o *ContainerCloseStdinParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the container close stdin params
func (o *ContainerCloseStdinParams) WithHTTPClient(client *http.Client) *ContainerCloseStdinParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the container close stdin params
func (o *ContainerCloseStdinParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOpID adds the opID to the container close stdin params
func (o *ContainerCloseStdinParams) WithOpID(opID *string) *ContainerCloseStdinParams {
	o.SetOpID(opID)
	return o
}

// SetOpID adds the opId to the container close stdin params
func (o *ContainerCloseStdinParams) SetOpID(opID *string) {
	o.OpID = opID
}

// WithID adds the id to the container close stdin params
func (o *ContainerCloseStdinParams) WithID(id string) *ContainerCloseStdinParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the container close stdin params
func (o *ContainerCloseStdinParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ContainerCloseStdinParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.OpID != nil {

		// header param Op-ID
		if err := r.SetHeaderParam("Op-ID", *o.OpID); err != nil {
			return err
		}

	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
