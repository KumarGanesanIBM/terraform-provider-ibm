// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/IBM-Cloud/power-go-client/power/models"
)

// NewPcloudNetworksPostParams creates a new PcloudNetworksPostParams object
// with the default values initialized.
func NewPcloudNetworksPostParams() *PcloudNetworksPostParams {
	var ()
	return &PcloudNetworksPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudNetworksPostParamsWithTimeout creates a new PcloudNetworksPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudNetworksPostParamsWithTimeout(timeout time.Duration) *PcloudNetworksPostParams {
	var ()
	return &PcloudNetworksPostParams{

		timeout: timeout,
	}
}

// NewPcloudNetworksPostParamsWithContext creates a new PcloudNetworksPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudNetworksPostParamsWithContext(ctx context.Context) *PcloudNetworksPostParams {
	var ()
	return &PcloudNetworksPostParams{

		Context: ctx,
	}
}

// NewPcloudNetworksPostParamsWithHTTPClient creates a new PcloudNetworksPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudNetworksPostParamsWithHTTPClient(client *http.Client) *PcloudNetworksPostParams {
	var ()
	return &PcloudNetworksPostParams{
		HTTPClient: client,
	}
}

/*PcloudNetworksPostParams contains all the parameters to send to the API endpoint
for the pcloud networks post operation typically these are written to a http.Request
*/
type PcloudNetworksPostParams struct {

	/*Body
	  Parameters for the creation of a new network

	*/
	Body *models.NetworkCreate
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud networks post params
func (o *PcloudNetworksPostParams) WithTimeout(timeout time.Duration) *PcloudNetworksPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud networks post params
func (o *PcloudNetworksPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud networks post params
func (o *PcloudNetworksPostParams) WithContext(ctx context.Context) *PcloudNetworksPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud networks post params
func (o *PcloudNetworksPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud networks post params
func (o *PcloudNetworksPostParams) WithHTTPClient(client *http.Client) *PcloudNetworksPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud networks post params
func (o *PcloudNetworksPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud networks post params
func (o *PcloudNetworksPostParams) WithBody(body *models.NetworkCreate) *PcloudNetworksPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud networks post params
func (o *PcloudNetworksPostParams) SetBody(body *models.NetworkCreate) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud networks post params
func (o *PcloudNetworksPostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudNetworksPostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud networks post params
func (o *PcloudNetworksPostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudNetworksPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
