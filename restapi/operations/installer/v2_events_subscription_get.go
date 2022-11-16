// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2EventsSubscriptionGetHandlerFunc turns a function with the right signature into a v2 events subscription get handler
type V2EventsSubscriptionGetHandlerFunc func(V2EventsSubscriptionGetParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2EventsSubscriptionGetHandlerFunc) Handle(params V2EventsSubscriptionGetParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2EventsSubscriptionGetHandler interface for that can handle valid v2 events subscription get params
type V2EventsSubscriptionGetHandler interface {
	Handle(V2EventsSubscriptionGetParams, interface{}) middleware.Responder
}

// NewV2EventsSubscriptionGet creates a new http.Handler for the v2 events subscription get operation
func NewV2EventsSubscriptionGet(ctx *middleware.Context, handler V2EventsSubscriptionGetHandler) *V2EventsSubscriptionGet {
	return &V2EventsSubscriptionGet{Context: ctx, Handler: handler}
}

/* V2EventsSubscriptionGet swagger:route GET /v2/events/subscription/{subscription_id} installer v2EventsSubscriptionGet

Get events subscription.

*/
type V2EventsSubscriptionGet struct {
	Context *middleware.Context
	Handler V2EventsSubscriptionGetHandler
}

func (o *V2EventsSubscriptionGet) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2EventsSubscriptionGetParams()
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
