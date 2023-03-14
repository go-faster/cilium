// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package policy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/go-faster/cilium/api/v1/models"
)

// GetPolicyOKCode is the HTTP code returned for type GetPolicyOK
const GetPolicyOKCode int = 200

/*
GetPolicyOK Success

swagger:response getPolicyOK
*/
type GetPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.Policy `json:"body,omitempty"`
}

// NewGetPolicyOK creates GetPolicyOK with default headers values
func NewGetPolicyOK() *GetPolicyOK {

	return &GetPolicyOK{}
}

// WithPayload adds the payload to the get policy o k response
func (o *GetPolicyOK) WithPayload(payload *models.Policy) *GetPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get policy o k response
func (o *GetPolicyOK) SetPayload(payload *models.Policy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetPolicyNotFoundCode is the HTTP code returned for type GetPolicyNotFound
const GetPolicyNotFoundCode int = 404

/*
GetPolicyNotFound No policy rules found

swagger:response getPolicyNotFound
*/
type GetPolicyNotFound struct {
}

// NewGetPolicyNotFound creates GetPolicyNotFound with default headers values
func NewGetPolicyNotFound() *GetPolicyNotFound {

	return &GetPolicyNotFound{}
}

// WriteResponse to the client
func (o *GetPolicyNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(404)
}
