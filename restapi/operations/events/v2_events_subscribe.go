// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2EventsSubscribeHandlerFunc turns a function with the right signature into a v2 events subscribe handler
type V2EventsSubscribeHandlerFunc func(V2EventsSubscribeParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2EventsSubscribeHandlerFunc) Handle(params V2EventsSubscribeParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2EventsSubscribeHandler interface for that can handle valid v2 events subscribe params
type V2EventsSubscribeHandler interface {
	Handle(V2EventsSubscribeParams, interface{}) middleware.Responder
}

// NewV2EventsSubscribe creates a new http.Handler for the v2 events subscribe operation
func NewV2EventsSubscribe(ctx *middleware.Context, handler V2EventsSubscribeHandler) *V2EventsSubscribe {
	return &V2EventsSubscribe{Context: ctx, Handler: handler}
}

/* V2EventsSubscribe swagger:route POST /v2/events/subscription events v2EventsSubscribe

Subscribe a URL to be called upon events for a cluster.

*/
type V2EventsSubscribe struct {
	Context *middleware.Context
	Handler V2EventsSubscribeHandler
}

func (o *V2EventsSubscribe) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2EventsSubscribeParams()
	uprinc, aCtx, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	if aCtx != nil {
		*r = *aCtx
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc.(interface{}) // this is really a interface{}, I promise
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}
