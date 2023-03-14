// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package cmd

import (
	"github.com/go-openapi/runtime/middleware"

	. "github.com/go-faster/cilium/api/v1/server/restapi/daemon"
	"github.com/go-faster/cilium/pkg/datapath"
)

type getNodeIDHandler struct {
	nodeHandler datapath.NodeHandler
}

func NewGetNodeIDsHandler(h datapath.NodeHandler) GetNodeIdsHandler {
	return &getNodeIDHandler{nodeHandler: h}
}

func (h *getNodeIDHandler) Handle(_ GetNodeIdsParams) middleware.Responder {
	dump := h.nodeHandler.DumpNodeIDs()
	return NewGetNodeIdsOK().WithPayload(dump)
}
