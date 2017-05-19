package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// UnbindHandlerFunc turns a function with the right signature into a unbind handler
type UnbindHandlerFunc func(UnbindParams) middleware.Responder

// Handle executing the request and returning a response
func (fn UnbindHandlerFunc) Handle(params UnbindParams) middleware.Responder {
	return fn(params)
}

// UnbindHandler interface for that can handle valid unbind params
type UnbindHandler interface {
	Handle(UnbindParams) middleware.Responder
}

// NewUnbind creates a new http.Handler for the unbind operation
func NewUnbind(ctx *middleware.Context, handler UnbindHandler) *Unbind {
	return &Unbind{Context: ctx, Handler: handler}
}

/*Unbind swagger:route DELETE /tasks/binding tasks unbind

Deactivate an existing task

Deactivate a task that exists within the context of the provided handle

*/
type Unbind struct {
	Context *middleware.Context
	Handler UnbindHandler
}

func (o *Unbind) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewUnbindParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
