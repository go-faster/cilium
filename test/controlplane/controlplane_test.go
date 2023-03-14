// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package controlplane_test

import (
	"testing"

	_ "github.com/go-faster/cilium/test/controlplane/ciliumnetworkpolicies"
	_ "github.com/go-faster/cilium/test/controlplane/node"
	_ "github.com/go-faster/cilium/test/controlplane/node/ciliumnodes"
	_ "github.com/go-faster/cilium/test/controlplane/services/dualstack"
	_ "github.com/go-faster/cilium/test/controlplane/services/graceful-termination"
	_ "github.com/go-faster/cilium/test/controlplane/services/nodeport"
	"github.com/go-faster/cilium/test/controlplane/suite"
)

func TestControlPlane(t *testing.T) {
	suite.RunSuite(t)
}
