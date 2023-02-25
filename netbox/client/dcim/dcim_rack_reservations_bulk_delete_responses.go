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

// DcimRackReservationsBulkDeleteReader is a Reader for the DcimRackReservationsBulkDelete structure.
type DcimRackReservationsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimRackReservationsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimRackReservationsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimRackReservationsBulkDeleteNoContent creates a DcimRackReservationsBulkDeleteNoContent with default headers values
func NewDcimRackReservationsBulkDeleteNoContent() *DcimRackReservationsBulkDeleteNoContent {
	return &DcimRackReservationsBulkDeleteNoContent{}
}

/*
DcimRackReservationsBulkDeleteNoContent describes a response with status code 204, with default header values.

DcimRackReservationsBulkDeleteNoContent dcim rack reservations bulk delete no content
*/
type DcimRackReservationsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this dcim rack reservations bulk delete no content response has a 2xx status code
func (o *DcimRackReservationsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim rack reservations bulk delete no content response has a 3xx status code
func (o *DcimRackReservationsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim rack reservations bulk delete no content response has a 4xx status code
func (o *DcimRackReservationsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim rack reservations bulk delete no content response has a 5xx status code
func (o *DcimRackReservationsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim rack reservations bulk delete no content response a status code equal to that given
func (o *DcimRackReservationsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the dcim rack reservations bulk delete no content response
func (o *DcimRackReservationsBulkDeleteNoContent) Code() int {
	return 204
}

func (o *DcimRackReservationsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-reservations/][%d] dcimRackReservationsBulkDeleteNoContent ", 204)
}

func (o *DcimRackReservationsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/rack-reservations/][%d] dcimRackReservationsBulkDeleteNoContent ", 204)
}

func (o *DcimRackReservationsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimRackReservationsBulkDeleteDefault creates a DcimRackReservationsBulkDeleteDefault with default headers values
func NewDcimRackReservationsBulkDeleteDefault(code int) *DcimRackReservationsBulkDeleteDefault {
	return &DcimRackReservationsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimRackReservationsBulkDeleteDefault describes a response with status code -1, with default header values.

DcimRackReservationsBulkDeleteDefault dcim rack reservations bulk delete default
*/
type DcimRackReservationsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this dcim rack reservations bulk delete default response has a 2xx status code
func (o *DcimRackReservationsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim rack reservations bulk delete default response has a 3xx status code
func (o *DcimRackReservationsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim rack reservations bulk delete default response has a 4xx status code
func (o *DcimRackReservationsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim rack reservations bulk delete default response has a 5xx status code
func (o *DcimRackReservationsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim rack reservations bulk delete default response a status code equal to that given
func (o *DcimRackReservationsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the dcim rack reservations bulk delete default response
func (o *DcimRackReservationsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *DcimRackReservationsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/rack-reservations/][%d] dcim_rack-reservations_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackReservationsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/rack-reservations/][%d] dcim_rack-reservations_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimRackReservationsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimRackReservationsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
