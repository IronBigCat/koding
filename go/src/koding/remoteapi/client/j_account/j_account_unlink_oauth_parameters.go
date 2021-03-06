package j_account

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

	"koding/remoteapi/models"
)

// NewJAccountUnlinkOauthParams creates a new JAccountUnlinkOauthParams object
// with the default values initialized.
func NewJAccountUnlinkOauthParams() *JAccountUnlinkOauthParams {
	var ()
	return &JAccountUnlinkOauthParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewJAccountUnlinkOauthParamsWithTimeout creates a new JAccountUnlinkOauthParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewJAccountUnlinkOauthParamsWithTimeout(timeout time.Duration) *JAccountUnlinkOauthParams {
	var ()
	return &JAccountUnlinkOauthParams{

		timeout: timeout,
	}
}

// NewJAccountUnlinkOauthParamsWithContext creates a new JAccountUnlinkOauthParams object
// with the default values initialized, and the ability to set a context for a request
func NewJAccountUnlinkOauthParamsWithContext(ctx context.Context) *JAccountUnlinkOauthParams {
	var ()
	return &JAccountUnlinkOauthParams{

		Context: ctx,
	}
}

/*JAccountUnlinkOauthParams contains all the parameters to send to the API endpoint
for the j account unlink oauth operation typically these are written to a http.Request
*/
type JAccountUnlinkOauthParams struct {

	/*Body
	  body of the request

	*/
	Body models.DefaultSelector
	/*ID
	  Mongo ID of target instance

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the j account unlink oauth params
func (o *JAccountUnlinkOauthParams) WithTimeout(timeout time.Duration) *JAccountUnlinkOauthParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the j account unlink oauth params
func (o *JAccountUnlinkOauthParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the j account unlink oauth params
func (o *JAccountUnlinkOauthParams) WithContext(ctx context.Context) *JAccountUnlinkOauthParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the j account unlink oauth params
func (o *JAccountUnlinkOauthParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the j account unlink oauth params
func (o *JAccountUnlinkOauthParams) WithBody(body models.DefaultSelector) *JAccountUnlinkOauthParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the j account unlink oauth params
func (o *JAccountUnlinkOauthParams) SetBody(body models.DefaultSelector) {
	o.Body = body
}

// WithID adds the id to the j account unlink oauth params
func (o *JAccountUnlinkOauthParams) WithID(id string) *JAccountUnlinkOauthParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the j account unlink oauth params
func (o *JAccountUnlinkOauthParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *JAccountUnlinkOauthParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
