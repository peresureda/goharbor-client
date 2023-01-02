// Code generated by go-swagger; DO NOT EDIT.

package scan_data_export

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/mittwald/goharbor-client/v5/apiv2/model"
)

// NewExportScanDataParams creates a new ExportScanDataParams object
// with the default values initialized.
func NewExportScanDataParams() *ExportScanDataParams {
	var ()
	return &ExportScanDataParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExportScanDataParamsWithTimeout creates a new ExportScanDataParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExportScanDataParamsWithTimeout(timeout time.Duration) *ExportScanDataParams {
	var ()
	return &ExportScanDataParams{

		timeout: timeout,
	}
}

// NewExportScanDataParamsWithContext creates a new ExportScanDataParams object
// with the default values initialized, and the ability to set a context for a request
func NewExportScanDataParamsWithContext(ctx context.Context) *ExportScanDataParams {
	var ()
	return &ExportScanDataParams{

		Context: ctx,
	}
}

// NewExportScanDataParamsWithHTTPClient creates a new ExportScanDataParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExportScanDataParamsWithHTTPClient(client *http.Client) *ExportScanDataParams {
	var ()
	return &ExportScanDataParams{
		HTTPClient: client,
	}
}

/*ExportScanDataParams contains all the parameters to send to the API endpoint
for the export scan data operation typically these are written to a http.Request
*/
type ExportScanDataParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*XScanDataType
	  The type of scan data to export

	*/
	XScanDataType string
	/*Criteria
	  The criteria for the export

	*/
	Criteria *model.ScanDataExportRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the export scan data params
func (o *ExportScanDataParams) WithTimeout(timeout time.Duration) *ExportScanDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export scan data params
func (o *ExportScanDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export scan data params
func (o *ExportScanDataParams) WithContext(ctx context.Context) *ExportScanDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export scan data params
func (o *ExportScanDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export scan data params
func (o *ExportScanDataParams) WithHTTPClient(client *http.Client) *ExportScanDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export scan data params
func (o *ExportScanDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the export scan data params
func (o *ExportScanDataParams) WithXRequestID(xRequestID *string) *ExportScanDataParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the export scan data params
func (o *ExportScanDataParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithXScanDataType adds the xScanDataType to the export scan data params
func (o *ExportScanDataParams) WithXScanDataType(xScanDataType string) *ExportScanDataParams {
	o.SetXScanDataType(xScanDataType)
	return o
}

// SetXScanDataType adds the xScanDataType to the export scan data params
func (o *ExportScanDataParams) SetXScanDataType(xScanDataType string) {
	o.XScanDataType = xScanDataType
}

// WithCriteria adds the criteria to the export scan data params
func (o *ExportScanDataParams) WithCriteria(criteria *model.ScanDataExportRequest) *ExportScanDataParams {
	o.SetCriteria(criteria)
	return o
}

// SetCriteria adds the criteria to the export scan data params
func (o *ExportScanDataParams) SetCriteria(criteria *model.ScanDataExportRequest) {
	o.Criteria = criteria
}

// WriteToRequest writes these params to a swagger request
func (o *ExportScanDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	// header param X-Scan-Data-Type
	if err := r.SetHeaderParam("X-Scan-Data-Type", o.XScanDataType); err != nil {
		return err
	}

	if o.Criteria != nil {
		if err := r.SetBodyParam(o.Criteria); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}