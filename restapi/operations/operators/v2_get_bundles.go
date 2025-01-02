// Code generated by go-swagger; DO NOT EDIT.

package operators

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// V2GetBundlesHandlerFunc turns a function with the right signature into a v2 get bundles handler
type V2GetBundlesHandlerFunc func(V2GetBundlesParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn V2GetBundlesHandlerFunc) Handle(params V2GetBundlesParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// V2GetBundlesHandler interface for that can handle valid v2 get bundles params
type V2GetBundlesHandler interface {
	Handle(V2GetBundlesParams, interface{}) middleware.Responder
}

// NewV2GetBundles creates a new http.Handler for the v2 get bundles operation
func NewV2GetBundles(ctx *middleware.Context, handler V2GetBundlesHandler) *V2GetBundles {
	return &V2GetBundles{Context: ctx, Handler: handler}
}

/*
	V2GetBundles swagger:route GET /v2/operators/bundles operators v2GetBundles

# Get list of bundle names

Retrieves a list of bundle names.
*/
type V2GetBundles struct {
	Context *middleware.Context
	Handler V2GetBundlesHandler
}

func (o *V2GetBundles) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewV2GetBundlesParams()
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
