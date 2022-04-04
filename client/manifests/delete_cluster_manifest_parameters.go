// Code generated by go-swagger; DO NOT EDIT.

package manifests

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

// NewDeleteClusterManifestParams creates a new DeleteClusterManifestParams object
// with the default values initialized.
func NewDeleteClusterManifestParams() *DeleteClusterManifestParams {
	var (
		folderDefault = string("manifests")
	)
	return &DeleteClusterManifestParams{
		Folder: &folderDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteClusterManifestParamsWithTimeout creates a new DeleteClusterManifestParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteClusterManifestParamsWithTimeout(timeout time.Duration) *DeleteClusterManifestParams {
	var (
		folderDefault = string("manifests")
	)
	return &DeleteClusterManifestParams{
		Folder: &folderDefault,

		timeout: timeout,
	}
}

// NewDeleteClusterManifestParamsWithContext creates a new DeleteClusterManifestParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteClusterManifestParamsWithContext(ctx context.Context) *DeleteClusterManifestParams {
	var (
		folderDefault = string("manifests")
	)
	return &DeleteClusterManifestParams{
		Folder: &folderDefault,

		Context: ctx,
	}
}

// NewDeleteClusterManifestParamsWithHTTPClient creates a new DeleteClusterManifestParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteClusterManifestParamsWithHTTPClient(client *http.Client) *DeleteClusterManifestParams {
	var (
		folderDefault = string("manifests")
	)
	return &DeleteClusterManifestParams{
		Folder:     &folderDefault,
		HTTPClient: client,
	}
}

/*DeleteClusterManifestParams contains all the parameters to send to the API endpoint
for the delete cluster manifest operation typically these are written to a http.Request
*/
type DeleteClusterManifestParams struct {

	/*ClusterID
	  The cluster whose manifest should be deleted.

	*/
	ClusterID strfmt.UUID
	/*FileName
	  The manifest file name to delete from the cluster.

	*/
	FileName string
	/*Folder
	  The folder that contains the files. Manifests can be placed in 'manifests' or 'openshift' directories.

	*/
	Folder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete cluster manifest params
func (o *DeleteClusterManifestParams) WithTimeout(timeout time.Duration) *DeleteClusterManifestParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete cluster manifest params
func (o *DeleteClusterManifestParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete cluster manifest params
func (o *DeleteClusterManifestParams) WithContext(ctx context.Context) *DeleteClusterManifestParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete cluster manifest params
func (o *DeleteClusterManifestParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete cluster manifest params
func (o *DeleteClusterManifestParams) WithHTTPClient(client *http.Client) *DeleteClusterManifestParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete cluster manifest params
func (o *DeleteClusterManifestParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the delete cluster manifest params
func (o *DeleteClusterManifestParams) WithClusterID(clusterID strfmt.UUID) *DeleteClusterManifestParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the delete cluster manifest params
func (o *DeleteClusterManifestParams) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WithFileName adds the fileName to the delete cluster manifest params
func (o *DeleteClusterManifestParams) WithFileName(fileName string) *DeleteClusterManifestParams {
	o.SetFileName(fileName)
	return o
}

// SetFileName adds the fileName to the delete cluster manifest params
func (o *DeleteClusterManifestParams) SetFileName(fileName string) {
	o.FileName = fileName
}

// WithFolder adds the folder to the delete cluster manifest params
func (o *DeleteClusterManifestParams) WithFolder(folder *string) *DeleteClusterManifestParams {
	o.SetFolder(folder)
	return o
}

// SetFolder adds the folder to the delete cluster manifest params
func (o *DeleteClusterManifestParams) SetFolder(folder *string) {
	o.Folder = folder
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteClusterManifestParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	// query param file_name
	qrFileName := o.FileName
	qFileName := qrFileName
	if qFileName != "" {
		if err := r.SetQueryParam("file_name", qFileName); err != nil {
			return err
		}
	}

	if o.Folder != nil {

		// query param folder
		var qrFolder string
		if o.Folder != nil {
			qrFolder = *o.Folder
		}
		qFolder := qrFolder
		if qFolder != "" {
			if err := r.SetQueryParam("folder", qFolder); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}