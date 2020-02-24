// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_p_vm_instances

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
)

// NewPcloudPvminstancesConsolePostParams creates a new PcloudPvminstancesConsolePostParams object
// with the default values initialized.
func NewPcloudPvminstancesConsolePostParams() *PcloudPvminstancesConsolePostParams {
	var ()
	return &PcloudPvminstancesConsolePostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudPvminstancesConsolePostParamsWithTimeout creates a new PcloudPvminstancesConsolePostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPcloudPvminstancesConsolePostParamsWithTimeout(timeout time.Duration) *PcloudPvminstancesConsolePostParams {
	var ()
	return &PcloudPvminstancesConsolePostParams{

		timeout: timeout,
	}
}

// NewPcloudPvminstancesConsolePostParamsWithContext creates a new PcloudPvminstancesConsolePostParams object
// with the default values initialized, and the ability to set a context for a request
func NewPcloudPvminstancesConsolePostParamsWithContext(ctx context.Context) *PcloudPvminstancesConsolePostParams {
	var ()
	return &PcloudPvminstancesConsolePostParams{

		Context: ctx,
	}
}

// NewPcloudPvminstancesConsolePostParamsWithHTTPClient creates a new PcloudPvminstancesConsolePostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPcloudPvminstancesConsolePostParamsWithHTTPClient(client *http.Client) *PcloudPvminstancesConsolePostParams {
	var ()
	return &PcloudPvminstancesConsolePostParams{
		HTTPClient: client,
	}
}

/*PcloudPvminstancesConsolePostParams contains all the parameters to send to the API endpoint
for the pcloud pvminstances console post operation typically these are written to a http.Request
*/
type PcloudPvminstancesConsolePostParams struct {

	/*CloudInstanceID
	  Cloud Instance ID of a PCloud Instance

	*/
	CloudInstanceID string
	/*PvmInstanceID
	  PCloud PVM Instance ID

	*/
	PvmInstanceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) WithTimeout(timeout time.Duration) *PcloudPvminstancesConsolePostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) WithContext(ctx context.Context) *PcloudPvminstancesConsolePostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) WithHTTPClient(client *http.Client) *PcloudPvminstancesConsolePostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) WithCloudInstanceID(cloudInstanceID string) *PcloudPvminstancesConsolePostParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithPvmInstanceID adds the pvmInstanceID to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) WithPvmInstanceID(pvmInstanceID string) *PcloudPvminstancesConsolePostParams {
	o.SetPvmInstanceID(pvmInstanceID)
	return o
}

// SetPvmInstanceID adds the pvmInstanceId to the pcloud pvminstances console post params
func (o *PcloudPvminstancesConsolePostParams) SetPvmInstanceID(pvmInstanceID string) {
	o.PvmInstanceID = pvmInstanceID
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudPvminstancesConsolePostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	// path param pvm_instance_id
	if err := r.SetPathParam("pvm_instance_id", o.PvmInstanceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}