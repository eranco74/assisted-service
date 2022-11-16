// Code generated by go-swagger; DO NOT EDIT.

package events

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/strfmt"
)

// V2EventsSubscriptionListURL generates an URL for the v2 events subscription list operation
type V2EventsSubscriptionListURL struct {
	ClusterID *strfmt.UUID

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *V2EventsSubscriptionListURL) WithBasePath(bp string) *V2EventsSubscriptionListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *V2EventsSubscriptionListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *V2EventsSubscriptionListURL) Build() (*url.URL, error) {
	var _result url.URL

	var _path = "/v2/events/subscription"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/api/assisted-install"
	}
	_result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var clusterIDQ string
	if o.ClusterID != nil {
		clusterIDQ = o.ClusterID.String()
	}
	if clusterIDQ != "" {
		qs.Set("cluster_id", clusterIDQ)
	}

	_result.RawQuery = qs.Encode()

	return &_result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *V2EventsSubscriptionListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *V2EventsSubscriptionListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *V2EventsSubscriptionListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on V2EventsSubscriptionListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on V2EventsSubscriptionListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *V2EventsSubscriptionListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
