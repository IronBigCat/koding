package social_channel

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

// NewPostRemoteAPISocialChannelFetchActivityCountParams creates a new PostRemoteAPISocialChannelFetchActivityCountParams object
// with the default values initialized.
func NewPostRemoteAPISocialChannelFetchActivityCountParams() *PostRemoteAPISocialChannelFetchActivityCountParams {
	var ()
	return &PostRemoteAPISocialChannelFetchActivityCountParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPISocialChannelFetchActivityCountParamsWithTimeout creates a new PostRemoteAPISocialChannelFetchActivityCountParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPISocialChannelFetchActivityCountParamsWithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelFetchActivityCountParams {
	var ()
	return &PostRemoteAPISocialChannelFetchActivityCountParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPISocialChannelFetchActivityCountParamsWithContext creates a new PostRemoteAPISocialChannelFetchActivityCountParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPISocialChannelFetchActivityCountParamsWithContext(ctx context.Context) *PostRemoteAPISocialChannelFetchActivityCountParams {
	var ()
	return &PostRemoteAPISocialChannelFetchActivityCountParams{

		Context: ctx,
	}
}

/*PostRemoteAPISocialChannelFetchActivityCountParams contains all the parameters to send to the API endpoint
for the post remote API social channel fetch activity count operation typically these are written to a http.Request
*/
type PostRemoteAPISocialChannelFetchActivityCountParams struct {

	/*Body
	  body of the request

	*/
	Body *models.DefaultSelector

	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API social channel fetch activity count params
func (o *PostRemoteAPISocialChannelFetchActivityCountParams) WithTimeout(timeout time.Duration) *PostRemoteAPISocialChannelFetchActivityCountParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API social channel fetch activity count params
func (o *PostRemoteAPISocialChannelFetchActivityCountParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API social channel fetch activity count params
func (o *PostRemoteAPISocialChannelFetchActivityCountParams) WithContext(ctx context.Context) *PostRemoteAPISocialChannelFetchActivityCountParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API social channel fetch activity count params
func (o *PostRemoteAPISocialChannelFetchActivityCountParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithBody adds the body to the post remote API social channel fetch activity count params
func (o *PostRemoteAPISocialChannelFetchActivityCountParams) WithBody(body *models.DefaultSelector) *PostRemoteAPISocialChannelFetchActivityCountParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post remote API social channel fetch activity count params
func (o *PostRemoteAPISocialChannelFetchActivityCountParams) SetBody(body *models.DefaultSelector) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPISocialChannelFetchActivityCountParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
