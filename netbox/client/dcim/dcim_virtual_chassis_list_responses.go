// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	"github.com/netbox-community/go-netbox/v3/netbox/models"
)

// DcimVirtualChassisListReader is a Reader for the DcimVirtualChassisList structure.
type DcimVirtualChassisListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimVirtualChassisListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimVirtualChassisListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimVirtualChassisListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimVirtualChassisListOK creates a DcimVirtualChassisListOK with default headers values
func NewDcimVirtualChassisListOK() *DcimVirtualChassisListOK {
	return &DcimVirtualChassisListOK{}
}

/*
DcimVirtualChassisListOK describes a response with status code 200, with default header values.

DcimVirtualChassisListOK dcim virtual chassis list o k
*/
type DcimVirtualChassisListOK struct {
	Payload *DcimVirtualChassisListOKBody
}

// IsSuccess returns true when this dcim virtual chassis list o k response has a 2xx status code
func (o *DcimVirtualChassisListOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim virtual chassis list o k response has a 3xx status code
func (o *DcimVirtualChassisListOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim virtual chassis list o k response has a 4xx status code
func (o *DcimVirtualChassisListOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim virtual chassis list o k response has a 5xx status code
func (o *DcimVirtualChassisListOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim virtual chassis list o k response a status code equal to that given
func (o *DcimVirtualChassisListOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the dcim virtual chassis list o k response
func (o *DcimVirtualChassisListOK) Code() int {
	return 200
}

func (o *DcimVirtualChassisListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/virtual-chassis/][%d] dcimVirtualChassisListOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualChassisListOK) String() string {
	return fmt.Sprintf("[GET /dcim/virtual-chassis/][%d] dcimVirtualChassisListOK  %+v", 200, o.Payload)
}

func (o *DcimVirtualChassisListOK) GetPayload() *DcimVirtualChassisListOKBody {
	return o.Payload
}

func (o *DcimVirtualChassisListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(DcimVirtualChassisListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimVirtualChassisListDefault creates a DcimVirtualChassisListDefault with default headers values
func NewDcimVirtualChassisListDefault(code int) *DcimVirtualChassisListDefault {
	return &DcimVirtualChassisListDefault{
		_statusCode: code,
	}
}

/*
DcimVirtualChassisListDefault describes a response with status code -1, with default header values.

DcimVirtualChassisListDefault dcim virtual chassis list default
*/
type DcimVirtualChassisListDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim virtual chassis list default response has a 2xx status code
func (o *DcimVirtualChassisListDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim virtual chassis list default response has a 3xx status code
func (o *DcimVirtualChassisListDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim virtual chassis list default response has a 4xx status code
func (o *DcimVirtualChassisListDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim virtual chassis list default response has a 5xx status code
func (o *DcimVirtualChassisListDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim virtual chassis list default response a status code equal to that given
func (o *DcimVirtualChassisListDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim virtual chassis list default response
func (o *DcimVirtualChassisListDefault) Code() int {
	return o._statusCode
}

func (o *DcimVirtualChassisListDefault) Error() string {
	return fmt.Sprintf("[GET /dcim/virtual-chassis/][%d] dcim_virtual-chassis_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisListDefault) String() string {
	return fmt.Sprintf("[GET /dcim/virtual-chassis/][%d] dcim_virtual-chassis_list default  %+v", o._statusCode, o.Payload)
}

func (o *DcimVirtualChassisListDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimVirtualChassisListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
DcimVirtualChassisListOKBody dcim virtual chassis list o k body
swagger:model DcimVirtualChassisListOKBody
*/
type DcimVirtualChassisListOKBody struct {

	// count
	// Required: true
	Count *int64 `json:"count"`

	// next
	// Format: uri
	Next *strfmt.URI `json:"next,omitempty"`

	// previous
	// Format: uri
	Previous *strfmt.URI `json:"previous,omitempty"`

	// results
	// Required: true
	Results []*models.VirtualChassis `json:"results"`
}

// Validate validates this dcim virtual chassis list o k body
func (o *DcimVirtualChassisListOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrevious(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimVirtualChassisListOKBody) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("dcimVirtualChassisListOK"+"."+"count", "body", o.Count); err != nil {
		return err
	}

	return nil
}

func (o *DcimVirtualChassisListOKBody) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(o.Next) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimVirtualChassisListOK"+"."+"next", "body", "uri", o.Next.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimVirtualChassisListOKBody) validatePrevious(formats strfmt.Registry) error {
	if swag.IsZero(o.Previous) { // not required
		return nil
	}

	if err := validate.FormatOf("dcimVirtualChassisListOK"+"."+"previous", "body", "uri", o.Previous.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *DcimVirtualChassisListOKBody) validateResults(formats strfmt.Registry) error {

	if err := validate.Required("dcimVirtualChassisListOK"+"."+"results", "body", o.Results); err != nil {
		return err
	}

	for i := 0; i < len(o.Results); i++ {
		if swag.IsZero(o.Results[i]) { // not required
			continue
		}

		if o.Results[i] != nil {
			if err := o.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimVirtualChassisListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimVirtualChassisListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this dcim virtual chassis list o k body based on the context it is used
func (o *DcimVirtualChassisListOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResults(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DcimVirtualChassisListOKBody) contextValidateResults(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Results); i++ {

		if o.Results[i] != nil {
			if err := o.Results[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dcimVirtualChassisListOK" + "." + "results" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dcimVirtualChassisListOK" + "." + "results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *DcimVirtualChassisListOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *DcimVirtualChassisListOKBody) UnmarshalBinary(b []byte) error {
	var res DcimVirtualChassisListOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
