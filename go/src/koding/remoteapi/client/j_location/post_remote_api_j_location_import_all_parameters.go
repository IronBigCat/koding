package j_location

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostRemoteAPIJLocationImportAllParams creates a new PostRemoteAPIJLocationImportAllParams object
// with the default values initialized.
func NewPostRemoteAPIJLocationImportAllParams() *PostRemoteAPIJLocationImportAllParams {

	return &PostRemoteAPIJLocationImportAllParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostRemoteAPIJLocationImportAllParamsWithTimeout creates a new PostRemoteAPIJLocationImportAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostRemoteAPIJLocationImportAllParamsWithTimeout(timeout time.Duration) *PostRemoteAPIJLocationImportAllParams {

	return &PostRemoteAPIJLocationImportAllParams{

		timeout: timeout,
	}
}

// NewPostRemoteAPIJLocationImportAllParamsWithContext creates a new PostRemoteAPIJLocationImportAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostRemoteAPIJLocationImportAllParamsWithContext(ctx context.Context) *PostRemoteAPIJLocationImportAllParams {

	return &PostRemoteAPIJLocationImportAllParams{

		Context: ctx,
	}
}

/*PostRemoteAPIJLocationImportAllParams contains all the parameters to send to the API endpoint
for the post remote API j location import all operation typically these are written to a http.Request
*/
type PostRemoteAPIJLocationImportAllParams struct {
	timeout time.Duration
	Context context.Context
}

// WithTimeout adds the timeout to the post remote API j location import all params
func (o *PostRemoteAPIJLocationImportAllParams) WithTimeout(timeout time.Duration) *PostRemoteAPIJLocationImportAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post remote API j location import all params
func (o *PostRemoteAPIJLocationImportAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post remote API j location import all params
func (o *PostRemoteAPIJLocationImportAllParams) WithContext(ctx context.Context) *PostRemoteAPIJLocationImportAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post remote API j location import all params
func (o *PostRemoteAPIJLocationImportAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WriteToRequest writes these params to a swagger request
func (o *PostRemoteAPIJLocationImportAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	r.SetTimeout(o.timeout)
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
