package j_group

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// NewPostRemoteAPIJGroupCreateParams creates a new PostRemoteAPIJGroupCreateParams object
// with the default values initialized.
func NewPostRemoteAPIJGroupCreateParams() *PostRemoteAPIJGroupCreateParams {
	var ()
	return &PostRemoteAPIJGroupCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJGroupCreateParamsWithTimeout creates a new PostRemoteAPIJGroupCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJGroupCreateParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJGroupCreateParams {
	var ()
	return &PostRemoteAPIJGroupCreateParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJGroupCreateParamsWithContext creates a new PostRemoteAPIJGroupCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJGroupCreateParamsWithContext(ctx context.Context) *PostRemoteAPIJGroupCreateParams {
	var ()
	return &PostRemoteAPIJGroupCreateParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJGroupCreateParams contains all the parameters to send to the API endpoint
for the post remote API j group create operation typically these are written to a http.Request
*/
type PostRemoteAPIJGroupCreateParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j group create params
func (o *PostRemoteAPIJGroupCreateParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJGroupCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j group create params
func (o *PostRemoteAPIJGroupCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j group create params
func (o *PostRemoteAPIJGroupCreateParams) WithContext(ctx context.Context) *PostRemoteAPIJGroupCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j group create params
func (o *PostRemoteAPIJGroupCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API j group create params
func (o *PostRemoteAPIJGroupCreateParams) WithBody(body *models.DefaultSelector) *PostRemoteAPIJGroupCreateParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API j group create params
func (o *PostRemoteAPIJGroupCreateParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJGroupCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if o.Body == nil {
		o.Body = new(models.DefaultSelector)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
