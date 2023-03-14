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

// PutPolicyOKCode is the HTTP code returned for type PutPolicyOK
const PutPolicyOKCode int = 200

/*
PutPolicyOK Success

swagger:response putPolicyOK
*/
type PutPolicyOK struct {

	/*
	  In: Body
	*/
	Payload *models.Policy `json:"body,omitempty"`
}

// NewPutPolicyOK creates PutPolicyOK with default headers values
func NewPutPolicyOK() *PutPolicyOK {

	return &PutPolicyOK{}
}

// WithPayload adds the payload to the put policy o k response
func (o *PutPolicyOK) WithPayload(payload *models.Policy) *PutPolicyOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put policy o k response
func (o *PutPolicyOK) SetPayload(payload *models.Policy) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPolicyOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// PutPolicyInvalidPolicyCode is the HTTP code returned for type PutPolicyInvalidPolicy
const PutPolicyInvalidPolicyCode int = 400

/*
PutPolicyInvalidPolicy Invalid policy

swagger:response putPolicyInvalidPolicy
*/
type PutPolicyInvalidPolicy struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutPolicyInvalidPolicy creates PutPolicyInvalidPolicy with default headers values
func NewPutPolicyInvalidPolicy() *PutPolicyInvalidPolicy {

	return &PutPolicyInvalidPolicy{}
}

// WithPayload adds the payload to the put policy invalid policy response
func (o *PutPolicyInvalidPolicy) WithPayload(payload models.Error) *PutPolicyInvalidPolicy {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put policy invalid policy response
func (o *PutPolicyInvalidPolicy) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPolicyInvalidPolicy) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutPolicyInvalidPathCode is the HTTP code returned for type PutPolicyInvalidPath
const PutPolicyInvalidPathCode int = 460

/*
PutPolicyInvalidPath Invalid path

swagger:response putPolicyInvalidPath
*/
type PutPolicyInvalidPath struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutPolicyInvalidPath creates PutPolicyInvalidPath with default headers values
func NewPutPolicyInvalidPath() *PutPolicyInvalidPath {

	return &PutPolicyInvalidPath{}
}

// WithPayload adds the payload to the put policy invalid path response
func (o *PutPolicyInvalidPath) WithPayload(payload models.Error) *PutPolicyInvalidPath {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put policy invalid path response
func (o *PutPolicyInvalidPath) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPolicyInvalidPath) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(460)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

// PutPolicyFailureCode is the HTTP code returned for type PutPolicyFailure
const PutPolicyFailureCode int = 500

/*
PutPolicyFailure Policy import failed

swagger:response putPolicyFailure
*/
type PutPolicyFailure struct {

	/*
	  In: Body
	*/
	Payload models.Error `json:"body,omitempty"`
}

// NewPutPolicyFailure creates PutPolicyFailure with default headers values
func NewPutPolicyFailure() *PutPolicyFailure {

	return &PutPolicyFailure{}
}

// WithPayload adds the payload to the put policy failure response
func (o *PutPolicyFailure) WithPayload(payload models.Error) *PutPolicyFailure {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put policy failure response
func (o *PutPolicyFailure) SetPayload(payload models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutPolicyFailure) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
