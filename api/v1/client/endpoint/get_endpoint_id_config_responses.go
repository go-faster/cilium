// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package endpoint

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/go-faster/cilium/api/v1/models"
)

// GetEndpointIDConfigReader is a Reader for the GetEndpointIDConfig structure.
type GetEndpointIDConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetEndpointIDConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetEndpointIDConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewGetEndpointIDConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewGetEndpointIDConfigTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetEndpointIDConfigOK creates a GetEndpointIDConfigOK with default headers values
func NewGetEndpointIDConfigOK() *GetEndpointIDConfigOK {
	return &GetEndpointIDConfigOK{}
}

/*
GetEndpointIDConfigOK describes a response with status code 200, with default header values.

Success
*/
type GetEndpointIDConfigOK struct {
	Payload *models.EndpointConfigurationStatus
}

// IsSuccess returns true when this get endpoint Id config o k response has a 2xx status code
func (o *GetEndpointIDConfigOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get endpoint Id config o k response has a 3xx status code
func (o *GetEndpointIDConfigOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint Id config o k response has a 4xx status code
func (o *GetEndpointIDConfigOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get endpoint Id config o k response has a 5xx status code
func (o *GetEndpointIDConfigOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint Id config o k response a status code equal to that given
func (o *GetEndpointIDConfigOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetEndpointIDConfigOK) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/config][%d] getEndpointIdConfigOK  %+v", 200, o.Payload)
}

func (o *GetEndpointIDConfigOK) String() string {
	return fmt.Sprintf("[GET /endpoint/{id}/config][%d] getEndpointIdConfigOK  %+v", 200, o.Payload)
}

func (o *GetEndpointIDConfigOK) GetPayload() *models.EndpointConfigurationStatus {
	return o.Payload
}

func (o *GetEndpointIDConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EndpointConfigurationStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetEndpointIDConfigNotFound creates a GetEndpointIDConfigNotFound with default headers values
func NewGetEndpointIDConfigNotFound() *GetEndpointIDConfigNotFound {
	return &GetEndpointIDConfigNotFound{}
}

/*
GetEndpointIDConfigNotFound describes a response with status code 404, with default header values.

Endpoint not found
*/
type GetEndpointIDConfigNotFound struct {
}

// IsSuccess returns true when this get endpoint Id config not found response has a 2xx status code
func (o *GetEndpointIDConfigNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint Id config not found response has a 3xx status code
func (o *GetEndpointIDConfigNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint Id config not found response has a 4xx status code
func (o *GetEndpointIDConfigNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint Id config not found response has a 5xx status code
func (o *GetEndpointIDConfigNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint Id config not found response a status code equal to that given
func (o *GetEndpointIDConfigNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *GetEndpointIDConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/config][%d] getEndpointIdConfigNotFound ", 404)
}

func (o *GetEndpointIDConfigNotFound) String() string {
	return fmt.Sprintf("[GET /endpoint/{id}/config][%d] getEndpointIdConfigNotFound ", 404)
}

func (o *GetEndpointIDConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetEndpointIDConfigTooManyRequests creates a GetEndpointIDConfigTooManyRequests with default headers values
func NewGetEndpointIDConfigTooManyRequests() *GetEndpointIDConfigTooManyRequests {
	return &GetEndpointIDConfigTooManyRequests{}
}

/*
GetEndpointIDConfigTooManyRequests describes a response with status code 429, with default header values.

Rate-limiting too many requests in the given time frame
*/
type GetEndpointIDConfigTooManyRequests struct {
}

// IsSuccess returns true when this get endpoint Id config too many requests response has a 2xx status code
func (o *GetEndpointIDConfigTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get endpoint Id config too many requests response has a 3xx status code
func (o *GetEndpointIDConfigTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get endpoint Id config too many requests response has a 4xx status code
func (o *GetEndpointIDConfigTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this get endpoint Id config too many requests response has a 5xx status code
func (o *GetEndpointIDConfigTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this get endpoint Id config too many requests response a status code equal to that given
func (o *GetEndpointIDConfigTooManyRequests) IsCode(code int) bool {
	return code == 429
}

func (o *GetEndpointIDConfigTooManyRequests) Error() string {
	return fmt.Sprintf("[GET /endpoint/{id}/config][%d] getEndpointIdConfigTooManyRequests ", 429)
}

func (o *GetEndpointIDConfigTooManyRequests) String() string {
	return fmt.Sprintf("[GET /endpoint/{id}/config][%d] getEndpointIdConfigTooManyRequests ", 429)
}

func (o *GetEndpointIDConfigTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
