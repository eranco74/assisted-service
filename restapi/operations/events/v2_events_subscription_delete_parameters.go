// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// NewV2EventsSubscriptionDeleteParams creates a new V2EventsSubscriptionDeleteParams object
//
// There are no default values defined in the spec.
func NewV2EventsSubscriptionDeleteParams() V2EventsSubscriptionDeleteParams {

	return V2EventsSubscriptionDeleteParams{}
}

// V2EventsSubscriptionDeleteParams contains all the bound params for the v2 events subscription delete operation
// typically these are obtained from a http.Request
//
// swagger:parameters v2EventsSubscriptionDelete
type V2EventsSubscriptionDeleteParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*The subscription id.
	  Required: true
	  In: path
	*/
	SubscriptionID strfmt.UUID
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls.
//
// To ensure default values, the struct must have been initialized with NewV2EventsSubscriptionDeleteParams() beforehand.
func (o *V2EventsSubscriptionDeleteParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error

	o.HTTPRequest = r

	rSubscriptionID, rhkSubscriptionID, _ := route.Params.GetOK("subscription_id")
	if err := o.bindSubscriptionID(rSubscriptionID, rhkSubscriptionID, route.Formats); err != nil {
		res = append(res, err)
	}
	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindSubscriptionID binds and validates parameter SubscriptionID from path.
func (o *V2EventsSubscriptionDeleteParams) bindSubscriptionID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	// Required: true
	// Parameter is provided by construction from the route

	// Format: uuid
	value, err := formats.Parse("uuid", raw)
	if err != nil {
		return errors.InvalidType("subscription_id", "path", "strfmt.UUID", raw)
	}
	o.SubscriptionID = *(value.(*strfmt.UUID))

	if err := o.validateSubscriptionID(formats); err != nil {
		return err
	}

	return nil
}

// validateSubscriptionID carries on validations for parameter SubscriptionID
func (o *V2EventsSubscriptionDeleteParams) validateSubscriptionID(formats strfmt.Registry) error {

	if err := validate.FormatOf("subscription_id", "path", "uuid", o.SubscriptionID.String(), formats); err != nil {
		return err
	}
	return nil
}
