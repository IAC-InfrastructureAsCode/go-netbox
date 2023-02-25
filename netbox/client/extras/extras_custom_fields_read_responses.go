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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/v3/netbox/models"
)

// ExtrasCustomFieldsReadReader is a Reader for the ExtrasCustomFieldsRead structure.
type ExtrasCustomFieldsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExtrasCustomFieldsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasCustomFieldsReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasCustomFieldsReadOK creates a ExtrasCustomFieldsReadOK with default headers values
func NewExtrasCustomFieldsReadOK() *ExtrasCustomFieldsReadOK {
	return &ExtrasCustomFieldsReadOK{}
}

/*
ExtrasCustomFieldsReadOK describes a response with status code 200, with default header values.

ExtrasCustomFieldsReadOK extras custom fields read o k
*/
type ExtrasCustomFieldsReadOK struct {
	Payload *models.CustomField
}

// IsSuccess returns true when this extras custom fields read o k response has a 2xx status code
func (o *ExtrasCustomFieldsReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras custom fields read o k response has a 3xx status code
func (o *ExtrasCustomFieldsReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras custom fields read o k response has a 4xx status code
func (o *ExtrasCustomFieldsReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras custom fields read o k response has a 5xx status code
func (o *ExtrasCustomFieldsReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this extras custom fields read o k response a status code equal to that given
func (o *ExtrasCustomFieldsReadOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the extras custom fields read o k response
func (o *ExtrasCustomFieldsReadOK) Code() int {
	return 200
}

func (o *ExtrasCustomFieldsReadOK) Error() string {
	return fmt.Sprintf("[GET /extras/custom-fields/{id}/][%d] extrasCustomFieldsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldsReadOK) String() string {
	return fmt.Sprintf("[GET /extras/custom-fields/{id}/][%d] extrasCustomFieldsReadOK  %+v", 200, o.Payload)
}

func (o *ExtrasCustomFieldsReadOK) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExtrasCustomFieldsReadDefault creates a ExtrasCustomFieldsReadDefault with default headers values
func NewExtrasCustomFieldsReadDefault(code int) *ExtrasCustomFieldsReadDefault {
	return &ExtrasCustomFieldsReadDefault{
		_statusCode: code,
	}
}

/*
ExtrasCustomFieldsReadDefault describes a response with status code -1, with default header values.

ExtrasCustomFieldsReadDefault extras custom fields read default
*/
type ExtrasCustomFieldsReadDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras custom fields read default response has a 2xx status code
func (o *ExtrasCustomFieldsReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras custom fields read default response has a 3xx status code
func (o *ExtrasCustomFieldsReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras custom fields read default response has a 4xx status code
func (o *ExtrasCustomFieldsReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras custom fields read default response has a 5xx status code
func (o *ExtrasCustomFieldsReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras custom fields read default response a status code equal to that given
func (o *ExtrasCustomFieldsReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras custom fields read default response
func (o *ExtrasCustomFieldsReadDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasCustomFieldsReadDefault) Error() string {
	return fmt.Sprintf("[GET /extras/custom-fields/{id}/][%d] extras_custom-fields_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsReadDefault) String() string {
	return fmt.Sprintf("[GET /extras/custom-fields/{id}/][%d] extras_custom-fields_read default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasCustomFieldsReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasCustomFieldsReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
