// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package daemon

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/go-faster/cilium/api/v1/models"
)

// GetNodeIdsOKCode is the HTTP code returned for type GetNodeIdsOK
const GetNodeIdsOKCode int = 200

/*
GetNodeIdsOK Success

swagger:response getNodeIdsOK
*/
type GetNodeIdsOK struct {

	/*
	  In: Body
	*/
	Payload []*models.NodeID `json:"body,omitempty"`
}

// NewGetNodeIdsOK creates GetNodeIdsOK with default headers values
func NewGetNodeIdsOK() *GetNodeIdsOK {

	return &GetNodeIdsOK{}
}

// WithPayload adds the payload to the get node ids o k response
func (o *GetNodeIdsOK) WithPayload(payload []*models.NodeID) *GetNodeIdsOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get node ids o k response
func (o *GetNodeIdsOK) SetPayload(payload []*models.NodeID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetNodeIdsOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		// return empty array
		payload = make([]*models.NodeID, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}
