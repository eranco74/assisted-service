// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2EventsSubscriptionDeleteHandlerFunc turns a function with the right signature into a v2 events subscription delete handler
type V2EventsSubscriptionDeleteHandlerFunc func(V2EventsSubscriptionDeleteParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2EventsSubscriptionDeleteHandlerFunc) Handle(params V2EventsSubscriptionDeleteParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2EventsSubscriptionDeleteHandler interface for that can handle valid v2 events subscription delete params
type V2EventsSubscriptionDeleteHandler interface {
	Handle(V2EventsSubscriptionDeleteParams, interface{}) middleware.Responder
}

// NewV2EventsSubscriptionDelete creates a new http.Handler for the v2 events subscription delete operation
func NewV2EventsSubscriptionDelete(ctx *middleware.Context, handler V2EventsSubscriptionDeleteHandler) *V2EventsSubscriptionDelete {
	return &V2EventsSubscriptionDelete{Context: ctx, Handler: handler}
}

/* V2EventsSubscriptionDelete swagger:route DELETE /v2/events/subscription/{subscription_id} events v2EventsSubscriptionDelete

Delete events subscription.

*/
type V2EventsSubscriptionDelete struct {
	Context *middleware.Context
	Handler V2EventsSubscriptionDeleteHandler
}

func (o *V2EventsSubscriptionDelete) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2EventsSubscriptionDeleteParams()
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
