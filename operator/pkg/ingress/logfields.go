// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package ingress

import (
	"github.com/go-faster/cilium/pkg/logging"
	"github.com/go-faster/cilium/pkg/logging/logfields"
)

const Subsys = "ingress-controller"

var log = logging.DefaultLogger.WithField(logfields.LogSubsys, Subsys)
