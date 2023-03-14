// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package fqdn

import (
	"testing"

	. "gopkg.in/check.v1"

	"github.com/go-faster/cilium/pkg/defaults"
	"github.com/go-faster/cilium/pkg/fqdn/re"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) {
	TestingT(t)
}

func (ds *FQDNTestSuite) SetUpSuite(c *C) {
	re.InitRegexCompileLRU(defaults.FQDNRegexCompileLRUSize)
}

type FQDNTestSuite struct{}

var _ = Suite(&FQDNTestSuite{})
