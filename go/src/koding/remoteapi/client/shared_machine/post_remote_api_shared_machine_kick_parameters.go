package shared_machine

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

// NewPostRemoteAPISharedMachineKickParams creates a new PostRemoteAPISharedMachineKickParams object
// with the default values initialized.
func NewPostRemoteAPISharedMachineKickParams() *PostRemoteAPISharedMachineKickParams {
	var ()
	return &PostRemoteAPISharedMachineKickParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISharedMachineKickParamsWithTimeout creates a new PostRemoteAPISharedMachineKickParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISharedMachineKickParamsWithTimeout(timeout time.Duration) *PostRemoteAPISharedMachineKickParams {
	var ()
	return &PostRemoteAPISharedMachineKickParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISharedMachineKickParamsWithContext creates a new PostRemoteAPISharedMachineKickParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISharedMachineKickParamsWithContext(ctx context.Context) *PostRemoteAPISharedMachineKickParams {
	var ()
	return &PostRemoteAPISharedMachineKickParams{

		Context: ctx,
	}
}

/*PostRemoteAPISharedMachineKickParams contains all the parameters to send to the API endpoint
for the post remote API shared machine kick operation typically these are written to a http.Request
*/
type PostRemoteAPISharedMachineKickParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API shared machine kick params
func (o *PostRemoteAPISharedMachineKickParams) WithTimeout(timeout time.Duration) *PostRemoteAPISharedMachineKickParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API shared machine kick params
func (o *PostRemoteAPISharedMachineKickParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API shared machine kick params
func (o *PostRemoteAPISharedMachineKickParams) WithContext(ctx context.Context) *PostRemoteAPISharedMachineKickParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API shared machine kick params
func (o *PostRemoteAPISharedMachineKickParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API shared machine kick params
func (o *PostRemoteAPISharedMachineKickParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISharedMachineKickParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API shared machine kick params
func (o *PostRemoteAPISharedMachineKickParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISharedMachineKickParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
