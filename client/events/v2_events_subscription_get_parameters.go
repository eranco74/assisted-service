// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewV2EventsSubscriptionGetParams creates a new V2EventsSubscriptionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2EventsSubscriptionGetParams() *V2EventsSubscriptionGetParams {
	return &V2EventsSubscriptionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2EventsSubscriptionGetParamsWithTimeout creates a new V2EventsSubscriptionGetParams object
// with the ability to set a timeout on a request.
func NewV2EventsSubscriptionGetParamsWithTimeout(timeout time.Duration) *V2EventsSubscriptionGetParams {
	return &V2EventsSubscriptionGetParams{
		timeout: timeout,
	}
}

// NewV2EventsSubscriptionGetParamsWithContext creates a new V2EventsSubscriptionGetParams object
// with the ability to set a context for a request.
func NewV2EventsSubscriptionGetParamsWithContext(ctx context.Context) *V2EventsSubscriptionGetParams {
	return &V2EventsSubscriptionGetParams{
		Context: ctx,
	}
}

// NewV2EventsSubscriptionGetParamsWithHTTPClient creates a new V2EventsSubscriptionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2EventsSubscriptionGetParamsWithHTTPClient(client *http.Client) *V2EventsSubscriptionGetParams {
	return &V2EventsSubscriptionGetParams{
		HTTPClient: client,
	}
}

/* V2EventsSubscriptionGetParams contains all the parameters to send to the API endpoint
   for the v2 events subscription get operation.

   Typically these are written to a http.Request.
*/
type V2EventsSubscriptionGetParams struct {

	/* SubscriptionID.

	   The subscription id.

	   Format: uuid
	*/
	SubscriptionID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 events subscription get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2EventsSubscriptionGetParams) WithDefaults() *V2EventsSubscriptionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 events subscription get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2EventsSubscriptionGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 events subscription get params
func (o *V2EventsSubscriptionGetParams) WithTimeout(timeout time.Duration) *V2EventsSubscriptionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 events subscription get params
func (o *V2EventsSubscriptionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 events subscription get params
func (o *V2EventsSubscriptionGetParams) WithContext(ctx context.Context) *V2EventsSubscriptionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 events subscription get params
func (o *V2EventsSubscriptionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 events subscription get params
func (o *V2EventsSubscriptionGetParams) WithHTTPClient(client *http.Client) *V2EventsSubscriptionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 events subscription get params
func (o *V2EventsSubscriptionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSubscriptionID adds the subscriptionID to the v2 events subscription get params
func (o *V2EventsSubscriptionGetParams) WithSubscriptionID(subscriptionID strfmt.UUID) *V2EventsSubscriptionGetParams {
	o.SetSubscriptionID(subscriptionID)
	return o
}

// SetSubscriptionID adds the subscriptionId to the v2 events subscription get params
func (o *V2EventsSubscriptionGetParams) SetSubscriptionID(subscriptionID strfmt.UUID) {
	o.SubscriptionID = subscriptionID
}

// WriteToRequest writes these params to a swagger request
func (o *V2EventsSubscriptionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param subscription_id
	if err := r.SetPathParam("subscription_id", o.SubscriptionID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
