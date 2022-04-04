// Code generated by go-swagger; DO NOT EDIT.

package installer

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

// NewInstallClusterParams creates a new InstallClusterParams object
// with the default values initialized.
func NewInstallClusterParams() *InstallClusterParams {
	var ()
	return &InstallClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewInstallClusterParamsWithTimeout creates a new InstallClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewInstallClusterParamsWithTimeout(timeout time.Duration) *InstallClusterParams {
	var ()
	return &InstallClusterParams{

		timeout: timeout,
	}
}

// NewInstallClusterParamsWithContext creates a new InstallClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewInstallClusterParamsWithContext(ctx context.Context) *InstallClusterParams {
	var ()
	return &InstallClusterParams{

		Context: ctx,
	}
}

// NewInstallClusterParamsWithHTTPClient creates a new InstallClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewInstallClusterParamsWithHTTPClient(client *http.Client) *InstallClusterParams {
	var ()
	return &InstallClusterParams{
		HTTPClient: client,
	}
}

/*InstallClusterParams contains all the parameters to send to the API endpoint
for the install cluster operation typically these are written to a http.Request
*/
type InstallClusterParams struct {

	/*ClusterID
	  The cluster to be installed.

	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the install cluster params
func (o *InstallClusterParams) WithTimeout(timeout time.Duration) *InstallClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the install cluster params
func (o *InstallClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the install cluster params
func (o *InstallClusterParams) WithContext(ctx context.Context) *InstallClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the install cluster params
func (o *InstallClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the install cluster params
func (o *InstallClusterParams) WithHTTPClient(client *http.Client) *InstallClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the install cluster params
func (o *InstallClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the install cluster params
func (o *InstallClusterParams) WithClusterID(clusterID strfmt.UUID) *InstallClusterParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the install cluster params
func (o *InstallClusterParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *InstallClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}