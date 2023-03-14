// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package node

import (
	"github.com/go-faster/cilium/pkg/logging"
	"github.com/go-faster/cilium/pkg/logging/logfields"
)

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, "node")
