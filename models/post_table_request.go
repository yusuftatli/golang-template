// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostTableRequest Parameters for table request
//
// swagger:model PostTableRequest
type PostTableRequest struct {

	// end date
	// Format: date-time
	EndDate *strfmt.DateTime `json:"endDate,omitempty"`

	// export
	// Enum: [PDF CSV EXCEL]
	Export *string `json:"export,omitempty"`

	// filters
	Filters []*Filter `json:"filters"`

	// page
	Page *int64 `json:"page,omitempty"`

	// page size
	PageSize *int64 `json:"pageSize,omitempty"`

	// search
	Search *string `json:"search,omitempty"`

	// sort
	Sort map[string]string `json:"sort,omitempty"`

	// start date
	// Format: date-time
	StartDate *strfmt.DateTime `json:"startDate,omitempty"`
}

// Validate validates this post table request
func (m *PostTableRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExport(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFilters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSort(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStartDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostTableRequest) validateEndDate(formats strfmt.Registry) error {
	if swag.IsZero(m.EndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("endDate", "body", "date-time", m.EndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var postTableRequestTypeExportPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PDF","CSV","EXCEL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postTableRequestTypeExportPropEnum = append(postTableRequestTypeExportPropEnum, v)
	}
}

const (

	// PostTableRequestExportPDF captures enum value "PDF"
	PostTableRequestExportPDF string = "PDF"

	// PostTableRequestExportCSV captures enum value "CSV"
	PostTableRequestExportCSV string = "CSV"

	// PostTableRequestExportEXCEL captures enum value "EXCEL"
	PostTableRequestExportEXCEL string = "EXCEL"
)

// prop value enum
func (m *PostTableRequest) validateExportEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postTableRequestTypeExportPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostTableRequest) validateExport(formats strfmt.Registry) error {
	if swag.IsZero(m.Export) { // not required
		return nil
	}

	// value enum
	if err := m.validateExportEnum("export", "body", *m.Export); err != nil {
		return err
	}

	return nil
}

func (m *PostTableRequest) validateFilters(formats strfmt.Registry) error {
	if swag.IsZero(m.Filters) { // not required
		return nil
	}

	for i := 0; i < len(m.Filters); i++ {
		if swag.IsZero(m.Filters[i]) { // not required
			continue
		}

		if m.Filters[i] != nil {
			if err := m.Filters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// additional properties value enum
var postTableRequestSortValueEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["asc","desc"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		postTableRequestSortValueEnum = append(postTableRequestSortValueEnum, v)
	}
}

func (m *PostTableRequest) validateSortValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, postTableRequestSortValueEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *PostTableRequest) validateSort(formats strfmt.Registry) error {
	if swag.IsZero(m.Sort) { // not required
		return nil
	}

	for k := range m.Sort {

		// value enum
		if err := m.validateSortValueEnum("sort"+"."+k, "body", m.Sort[k]); err != nil {
			return err
		}

	}

	return nil
}

func (m *PostTableRequest) validateStartDate(formats strfmt.Registry) error {
	if swag.IsZero(m.StartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("startDate", "body", "date-time", m.StartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this post table request based on the context it is used
func (m *PostTableRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateFilters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PostTableRequest) contextValidateFilters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Filters); i++ {

		if m.Filters[i] != nil {
			if err := m.Filters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("filters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("filters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *PostTableRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PostTableRequest) UnmarshalBinary(b []byte) error {
	var res PostTableRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
