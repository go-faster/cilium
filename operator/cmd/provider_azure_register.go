// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build ipam_provider_azure

package cmd

import (
	// These dependencies should be included only when this file is included in the build.
	allocatorAzure "github.com/go-faster/cilium/pkg/ipam/allocator/azure" // Azure allocator task.
	ipamOption "github.com/go-faster/cilium/pkg/ipam/option"
)

func init() {
	allocatorProviders[ipamOption.IPAMAzure] = &allocatorAzure.AllocatorAzure{}
}
