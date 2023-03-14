// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package auth

import (
	"github.com/go-faster/cilium/pkg/logging"
	"github.com/go-faster/cilium/pkg/logging/logfields"
)

var (
	subsystem = "auth"
	log       = logging.DefaultLogger.WithField(logfields.LogSubsys, subsystem)
)
