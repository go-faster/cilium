// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build ipam_provider_operator

package cmd

import (
	// These dependencies should be included only when this file is included in the build.
	"github.com/go-faster/cilium/pkg/ipam/allocator/clusterpool"
	ipamOption "github.com/go-faster/cilium/pkg/ipam/option"
)

func init() {
	allocatorProviders[ipamOption.IPAMClusterPool] = &clusterpool.AllocatorOperator{}
	allocatorProviders[ipamOption.IPAMClusterPoolV2] = &clusterpool.AllocatorOperator{}
}
