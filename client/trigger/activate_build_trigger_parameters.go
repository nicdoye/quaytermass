// Code generated by go-swagger; DO NOT EDIT.

package trigger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/nicdoye/quaytermass/models"
)

// NewActivateBuildTriggerParams creates a new ActivateBuildTriggerParams object
// with the default values initialized.
func NewActivateBuildTriggerParams() *ActivateBuildTriggerParams {
	var ()
	return &ActivateBuildTriggerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewActivateBuildTriggerParamsWithTimeout creates a new ActivateBuildTriggerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewActivateBuildTriggerParamsWithTimeout(timeout time.Duration) *ActivateBuildTriggerParams {
	var ()
	return &ActivateBuildTriggerParams{

		timeout: timeout,
	}
}

// NewActivateBuildTriggerParamsWithContext creates a new ActivateBuildTriggerParams object
// with the default values initialized, and the ability to set a context for a request
func NewActivateBuildTriggerParamsWithContext(ctx context.Context) *ActivateBuildTriggerParams {
	var ()
	return &ActivateBuildTriggerParams{

		Context: ctx,
	}
}

// NewActivateBuildTriggerParamsWithHTTPClient creates a new ActivateBuildTriggerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewActivateBuildTriggerParamsWithHTTPClient(client *http.Client) *ActivateBuildTriggerParams {
	var ()
	return &ActivateBuildTriggerParams{
		HTTPClient: client,
	}
}

/*ActivateBuildTriggerParams contains all the parameters to send to the API endpoint
for the activate build trigger operation typically these are written to a http.Request
*/
type ActivateBuildTriggerParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.BuildTriggerActivateRequest
	/*Repository
	  The full path of the repository. e.g. namespace/name

	*/
	Repository string
	/*TriggerUUID
	  The UUID of the build trigger

	*/
	TriggerUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the activate build trigger params
func (o *ActivateBuildTriggerParams) WithTimeout(timeout time.Duration) *ActivateBuildTriggerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the activate build trigger params
func (o *ActivateBuildTriggerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the activate build trigger params
func (o *ActivateBuildTriggerParams) WithContext(ctx context.Context) *ActivateBuildTriggerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the activate build trigger params
func (o *ActivateBuildTriggerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the activate build trigger params
func (o *ActivateBuildTriggerParams) WithHTTPClient(client *http.Client) *ActivateBuildTriggerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the activate build trigger params
func (o *ActivateBuildTriggerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the activate build trigger params
func (o *ActivateBuildTriggerParams) WithBody(body *models.BuildTriggerActivateRequest) *ActivateBuildTriggerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the activate build trigger params
func (o *ActivateBuildTriggerParams) SetBody(body *models.BuildTriggerActivateRequest) {
	o.Body = body
}

// WithRepository adds the repository to the activate build trigger params
func (o *ActivateBuildTriggerParams) WithRepository(repository string) *ActivateBuildTriggerParams {
	o.SetRepository(repository)
	return o
}

// SetRepository adds the repository to the activate build trigger params
func (o *ActivateBuildTriggerParams) SetRepository(repository string) {
	o.Repository = repository
}

// WithTriggerUUID adds the triggerUUID to the activate build trigger params
func (o *ActivateBuildTriggerParams) WithTriggerUUID(triggerUUID string) *ActivateBuildTriggerParams {
	o.SetTriggerUUID(triggerUUID)
	return o
}

// SetTriggerUUID adds the triggerUuid to the activate build trigger params
func (o *ActivateBuildTriggerParams) SetTriggerUUID(triggerUUID string) {
	o.TriggerUUID = triggerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *ActivateBuildTriggerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param repository
	if err := r.SetPathParam("repository", o.Repository); err != nil {
		return err
	}

	// path param trigger_uuid
	if err := r.SetPathParam("trigger_uuid", o.TriggerUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
