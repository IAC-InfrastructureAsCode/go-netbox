// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ipam

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

	"github.com/netbox-community/go-netbox/v3/netbox/models"
)

// NewIpamAggregatesCreateParams creates a new IpamAggregatesCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamAggregatesCreateParams() *IpamAggregatesCreateParams {
	return &IpamAggregatesCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamAggregatesCreateParamsWithTimeout creates a new IpamAggregatesCreateParams object
// with the ability to set a timeout on a request.
func NewIpamAggregatesCreateParamsWithTimeout(timeout time.Duration) *IpamAggregatesCreateParams {
	return &IpamAggregatesCreateParams{
		timeout: timeout,
	}
}

// NewIpamAggregatesCreateParamsWithContext creates a new IpamAggregatesCreateParams object
// with the ability to set a context for a request.
func NewIpamAggregatesCreateParamsWithContext(ctx context.Context) *IpamAggregatesCreateParams {
	return &IpamAggregatesCreateParams{
		Context: ctx,
	}
}

// NewIpamAggregatesCreateParamsWithHTTPClient creates a new IpamAggregatesCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamAggregatesCreateParamsWithHTTPClient(client *http.Client) *IpamAggregatesCreateParams {
	return &IpamAggregatesCreateParams{
		HTTPClient: client,
	}
}

/*
IpamAggregatesCreateParams contains all the parameters to send to the API endpoint

	for the ipam aggregates create operation.

	Typically these are written to a http.Request.
*/
type IpamAggregatesCreateParams struct {

	// Data.
	Data *models.WritableAggregate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam aggregates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesCreateParams) WithDefaults() *IpamAggregatesCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam aggregates create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamAggregatesCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam aggregates create params
func (o *IpamAggregatesCreateParams) WithTimeout(timeout time.Duration) *IpamAggregatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam aggregates create params
func (o *IpamAggregatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam aggregates create params
func (o *IpamAggregatesCreateParams) WithContext(ctx context.Context) *IpamAggregatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam aggregates create params
func (o *IpamAggregatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam aggregates create params
func (o *IpamAggregatesCreateParams) WithHTTPClient(client *http.Client) *IpamAggregatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam aggregates create params
func (o *IpamAggregatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam aggregates create params
func (o *IpamAggregatesCreateParams) WithData(data *models.WritableAggregate) *IpamAggregatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam aggregates create params
func (o *IpamAggregatesCreateParams) SetData(data *models.WritableAggregate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IpamAggregatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
