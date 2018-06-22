// Code generated by go-swagger; DO NOT EDIT.

package organization

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

// NewCreateOrganizationApplicationParams creates a new CreateOrganizationApplicationParams object
// with the default values initialized.
func NewCreateOrganizationApplicationParams() *CreateOrganizationApplicationParams {
	var ()
	return &CreateOrganizationApplicationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateOrganizationApplicationParamsWithTimeout creates a new CreateOrganizationApplicationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateOrganizationApplicationParamsWithTimeout(timeout time.Duration) *CreateOrganizationApplicationParams {
	var ()
	return &CreateOrganizationApplicationParams{

		timeout: timeout,
	}
}

// NewCreateOrganizationApplicationParamsWithContext creates a new CreateOrganizationApplicationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateOrganizationApplicationParamsWithContext(ctx context.Context) *CreateOrganizationApplicationParams {
	var ()
	return &CreateOrganizationApplicationParams{

		Context: ctx,
	}
}

// NewCreateOrganizationApplicationParamsWithHTTPClient creates a new CreateOrganizationApplicationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateOrganizationApplicationParamsWithHTTPClient(client *http.Client) *CreateOrganizationApplicationParams {
	var ()
	return &CreateOrganizationApplicationParams{
		HTTPClient: client,
	}
}

/*CreateOrganizationApplicationParams contains all the parameters to send to the API endpoint
for the create organization application operation typically these are written to a http.Request
*/
type CreateOrganizationApplicationParams struct {

	/*Body
	  Request body contents.

	*/
	Body *models.NewApp
	/*Orgname
	  The name of the organization

	*/
	Orgname string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create organization application params
func (o *CreateOrganizationApplicationParams) WithTimeout(timeout time.Duration) *CreateOrganizationApplicationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create organization application params
func (o *CreateOrganizationApplicationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create organization application params
func (o *CreateOrganizationApplicationParams) WithContext(ctx context.Context) *CreateOrganizationApplicationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create organization application params
func (o *CreateOrganizationApplicationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create organization application params
func (o *CreateOrganizationApplicationParams) WithHTTPClient(client *http.Client) *CreateOrganizationApplicationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create organization application params
func (o *CreateOrganizationApplicationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create organization application params
func (o *CreateOrganizationApplicationParams) WithBody(body *models.NewApp) *CreateOrganizationApplicationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create organization application params
func (o *CreateOrganizationApplicationParams) SetBody(body *models.NewApp) {
	o.Body = body
}

// WithOrgname adds the orgname to the create organization application params
func (o *CreateOrganizationApplicationParams) WithOrgname(orgname string) *CreateOrganizationApplicationParams {
	o.SetOrgname(orgname)
	return o
}

// SetOrgname adds the orgname to the create organization application params
func (o *CreateOrganizationApplicationParams) SetOrgname(orgname string) {
	o.Orgname = orgname
}

// WriteToRequest writes these params to a swagger request
func (o *CreateOrganizationApplicationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param orgname
	if err := r.SetPathParam("orgname", o.Orgname); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
