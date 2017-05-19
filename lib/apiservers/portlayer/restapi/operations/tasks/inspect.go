package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// InspectHandlerFunc turns a function with the right signature into a inspect handler
type InspectHandlerFunc func(InspectParams) middleware.Responder

// Handle executing the request and returning a response
func (fn InspectHandlerFunc) Handle(params InspectParams) middleware.Responder {
	return fn(params)
}

// InspectHandler interface for that can handle valid inspect params
type InspectHandler interface {
	Handle(InspectParams) middleware.Responder
}

// NewInspect creates a new http.Handler for the inspect operation
func NewInspect(ctx *middleware.Context, handler InspectHandler) *Inspect {
	return &Inspect{Context: ctx, Handler: handler}
}

/*Inspect swagger:route GET /tasks tasks inspect

Initiates an task inspect operation

Initiates an task inspect operation

*/
type Inspect struct {
	Context *middleware.Context
	Handler InspectHandler
}

func (o *Inspect) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewInspectParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
