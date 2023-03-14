// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

//go:build ipam_provider_alibabacloud

package cmd

import (
	"github.com/go-faster/cilium/pkg/ipam/allocator/alibabacloud"
	ipamOption "github.com/go-faster/cilium/pkg/ipam/option"
)

func init() {
	allocatorProviders[ipamOption.IPAMAlibabaCloud] = &alibabacloud.AllocatorAlibabaCloud{}
}
