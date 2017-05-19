package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetStateHandlerFunc turns a function with the right signature into a get state handler
type GetStateHandlerFunc func(GetStateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetStateHandlerFunc) Handle(params GetStateParams) middleware.Responder {
	return fn(params)
}

// GetStateHandler interface for that can handle valid get state params
type GetStateHandler interface {
	Handle(GetStateParams) middleware.Responder
}

// NewGetState creates a new http.Handler for the get state operation
func NewGetState(ctx *middleware.Context, handler GetStateHandler) *GetState {
	return &GetState{Context: ctx, Handler: handler}
}

/*GetState swagger:route GET /containers/{handle}/state containers getState

Get the current state of the a container

*/
type GetState struct {
	Context *middleware.Context
	Handler GetStateHandler
}

func (o *GetState) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetStateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
