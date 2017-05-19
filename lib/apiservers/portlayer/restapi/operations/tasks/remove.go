package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// RemoveHandlerFunc turns a function with the right signature into a remove handler
type RemoveHandlerFunc func(RemoveParams) middleware.Responder

// Handle executing the request and returning a response
func (fn RemoveHandlerFunc) Handle(params RemoveParams) middleware.Responder {
	return fn(params)
}

// RemoveHandler interface for that can handle valid remove params
type RemoveHandler interface {
	Handle(RemoveParams) middleware.Responder
}

// NewRemove creates a new http.Handler for the remove operation
func NewRemove(ctx *middleware.Context, handler RemoveHandler) *Remove {
	return &Remove{Context: ctx, Handler: handler}
}

/*Remove swagger:route DELETE /tasks tasks remove

Initiates an task remove operation

Initiates an task remove operation

*/
type Remove struct {
	Context *middleware.Context
	Handler RemoveHandler
}

func (o *Remove) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewRemoveParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
