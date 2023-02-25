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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DcimModuleBayTemplatesDeleteReader is a Reader for the DcimModuleBayTemplatesDelete structure.
type DcimModuleBayTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimModuleBayTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimModuleBayTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimModuleBayTemplatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimModuleBayTemplatesDeleteNoContent creates a DcimModuleBayTemplatesDeleteNoContent with default headers values
func NewDcimModuleBayTemplatesDeleteNoContent() *DcimModuleBayTemplatesDeleteNoContent {
	return &DcimModuleBayTemplatesDeleteNoContent{}
}

/*
DcimModuleBayTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimModuleBayTemplatesDeleteNoContent dcim module bay templates delete no content
*/
type DcimModuleBayTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim module bay templates delete no content response has a 2xx status code
func (o *DcimModuleBayTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim module bay templates delete no content response has a 3xx status code
func (o *DcimModuleBayTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim module bay templates delete no content response has a 4xx status code
func (o *DcimModuleBayTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim module bay templates delete no content response has a 5xx status code
func (o *DcimModuleBayTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim module bay templates delete no content response a status code equal to that given
func (o *DcimModuleBayTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim module bay templates delete no content response
func (o *DcimModuleBayTemplatesDeleteNoContent) Code() int {
	return 204
}

func (o *DcimModuleBayTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesDeleteNoContent ", 204)
}

func (o *DcimModuleBayTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcimModuleBayTemplatesDeleteNoContent ", 204)
}

func (o *DcimModuleBayTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimModuleBayTemplatesDeleteDefault creates a DcimModuleBayTemplatesDeleteDefault with default headers values
func NewDcimModuleBayTemplatesDeleteDefault(code int) *DcimModuleBayTemplatesDeleteDefault {
	return &DcimModuleBayTemplatesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimModuleBayTemplatesDeleteDefault describes a response with status code -1, with default header values.

DcimModuleBayTemplatesDeleteDefault dcim module bay templates delete default
*/
type DcimModuleBayTemplatesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim module bay templates delete default response has a 2xx status code
func (o *DcimModuleBayTemplatesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim module bay templates delete default response has a 3xx status code
func (o *DcimModuleBayTemplatesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim module bay templates delete default response has a 4xx status code
func (o *DcimModuleBayTemplatesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim module bay templates delete default response has a 5xx status code
func (o *DcimModuleBayTemplatesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim module bay templates delete default response a status code equal to that given
func (o *DcimModuleBayTemplatesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim module bay templates delete default response
func (o *DcimModuleBayTemplatesDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimModuleBayTemplatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcim_module-bay-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/module-bay-templates/{id}/][%d] dcim_module-bay-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimModuleBayTemplatesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimModuleBayTemplatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
