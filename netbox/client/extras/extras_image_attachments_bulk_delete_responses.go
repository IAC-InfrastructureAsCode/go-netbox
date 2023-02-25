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
)

// ExtrasImageAttachmentsBulkDeleteReader is a Reader for the ExtrasImageAttachmentsBulkDelete structure.
type ExtrasImageAttachmentsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasImageAttachmentsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewExtrasImageAttachmentsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExtrasImageAttachmentsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExtrasImageAttachmentsBulkDeleteNoContent creates a ExtrasImageAttachmentsBulkDeleteNoContent with default headers values
func NewExtrasImageAttachmentsBulkDeleteNoContent() *ExtrasImageAttachmentsBulkDeleteNoContent {
	return &ExtrasImageAttachmentsBulkDeleteNoContent{}
}

/*
ExtrasImageAttachmentsBulkDeleteNoContent describes a response with status code 204, with default header values.

ExtrasImageAttachmentsBulkDeleteNoContent extras image attachments bulk delete no content
*/
type ExtrasImageAttachmentsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this extras image attachments bulk delete no content response has a 2xx status code
func (o *ExtrasImageAttachmentsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this extras image attachments bulk delete no content response has a 3xx status code
func (o *ExtrasImageAttachmentsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this extras image attachments bulk delete no content response has a 4xx status code
func (o *ExtrasImageAttachmentsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this extras image attachments bulk delete no content response has a 5xx status code
func (o *ExtrasImageAttachmentsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this extras image attachments bulk delete no content response a status code equal to that given
func (o *ExtrasImageAttachmentsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the extras image attachments bulk delete no content response
func (o *ExtrasImageAttachmentsBulkDeleteNoContent) Code() int {
	return 204
}

func (o *ExtrasImageAttachmentsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/][%d] extrasImageAttachmentsBulkDeleteNoContent ", 204)
}

func (o *ExtrasImageAttachmentsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/][%d] extrasImageAttachmentsBulkDeleteNoContent ", 204)
}

func (o *ExtrasImageAttachmentsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewExtrasImageAttachmentsBulkDeleteDefault creates a ExtrasImageAttachmentsBulkDeleteDefault with default headers values
func NewExtrasImageAttachmentsBulkDeleteDefault(code int) *ExtrasImageAttachmentsBulkDeleteDefault {
	return &ExtrasImageAttachmentsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
ExtrasImageAttachmentsBulkDeleteDefault describes a response with status code -1, with default header values.

ExtrasImageAttachmentsBulkDeleteDefault extras image attachments bulk delete default
*/
type ExtrasImageAttachmentsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// IsSuccess returns true when this extras image attachments bulk delete default response has a 2xx status code
func (o *ExtrasImageAttachmentsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this extras image attachments bulk delete default response has a 3xx status code
func (o *ExtrasImageAttachmentsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this extras image attachments bulk delete default response has a 4xx status code
func (o *ExtrasImageAttachmentsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this extras image attachments bulk delete default response has a 5xx status code
func (o *ExtrasImageAttachmentsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this extras image attachments bulk delete default response a status code equal to that given
func (o *ExtrasImageAttachmentsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the extras image attachments bulk delete default response
func (o *ExtrasImageAttachmentsBulkDeleteDefault) Code() int {
	return o._statusCode
}

func (o *ExtrasImageAttachmentsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/][%d] extras_image-attachments_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /extras/image-attachments/][%d] extras_image-attachments_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *ExtrasImageAttachmentsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *ExtrasImageAttachmentsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
