// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_volumes

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

// NewPcloudCloudinstancesVolumesPutParams creates a new PcloudCloudinstancesVolumesPutParams object
// with the default values initialized.
func NewPcloudCloudinstancesVolumesPutParams() *PcloudCloudinstancesVolumesPutParams {
	var ()
	return &PcloudCloudinstancesVolumesPutParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesVolumesPutParamsWithTimeout creates a new PcloudCloudinstancesVolumesPutParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudCloudinstancesVolumesPutParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesVolumesPutParams {
	var ()
	return &PcloudCloudinstancesVolumesPutParams{

		timeout: timeout,
	}
}

// NewPcloudCloudinstancesVolumesPutParamsWithContext creates a new PcloudCloudinstancesVolumesPutParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudCloudinstancesVolumesPutParamsWithContext(ctx context.Context) *PcloudCloudinstancesVolumesPutParams {
	var ()
	return &PcloudCloudinstancesVolumesPutParams{

		Context: ctx,
	}
}

// NewPcloudCloudinstancesVolumesPutParamsWithHTTPClient creates a new PcloudCloudinstancesVolumesPutParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudCloudinstancesVolumesPutParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesVolumesPutParams {
	var ()
	return &PcloudCloudinstancesVolumesPutParams{
		HTTPClient: client,
	}
}

/*PcloudCloudinstancesVolumesPutParams contains all the parameters to send to the API endpoint
for the pcloud cloudinstances volumes put operation typically these are written to a http.Request
*/
type PcloudCloudinstancesVolumesPutParams struct {

	/*Body
	  Parameters to update a cloud instance volume

	*/
	Body *models.UpdateVolume
	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*VolumeID
	  Volume ID

	*/
	VolumeID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesVolumesPutParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) WithContext(ctx context.Context) *PcloudCloudinstancesVolumesPutParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesVolumesPutParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) WithBody(body *models.UpdateVolume) *PcloudCloudinstancesVolumesPutParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) SetBody(body *models.UpdateVolume) {
	o.Body = body
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesVolumesPutParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithVolumeID adds the volumeID to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) WithVolumeID(volumeID string) *PcloudCloudinstancesVolumesPutParams {
	o.SetVolumeID(volumeID)
	return o
}

// SetVolumeID adds the volumeId to the pcloud cloudinstances volumes put params
func (o *PcloudCloudinstancesVolumesPutParams) SetVolumeID(volumeID string) {
	o.VolumeID = volumeID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesVolumesPutParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param volume_id
	if err := r.SetPathParam("volume_id", o.VolumeID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
