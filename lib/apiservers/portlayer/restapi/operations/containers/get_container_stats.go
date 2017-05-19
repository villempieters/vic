package containers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetContainerStatsHandlerFunc turns a function with the right signature into a get container stats handler
type GetContainerStatsHandlerFunc func(GetContainerStatsParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetContainerStatsHandlerFunc) Handle(params GetContainerStatsParams) middleware.Responder {
	return fn(params)
}

// GetContainerStatsHandler interface for that can handle valid get container stats params
type GetContainerStatsHandler interface {
	Handle(GetContainerStatsParams) middleware.Responder
}

// NewGetContainerStats creates a new http.Handler for the get container stats operation
func NewGetContainerStats(ctx *middleware.Context, handler GetContainerStatsHandler) *GetContainerStats {
	return &GetContainerStats{Context: ctx, Handler: handler}
}

/*GetContainerStats swagger:route GET /containers/{id}/stats containers getContainerStats

Gets the container stats

Gets the container stats by id

*/
type GetContainerStats struct {
	Context *middleware.Context
	Handler GetContainerStatsHandler
}

func (o *GetContainerStats) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewGetContainerStatsParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
