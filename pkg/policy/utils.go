// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package policy

import "github.com/go-faster/cilium/pkg/labels"

// JoinPath returns a joined path from a and b.
func JoinPath(a, b string) string {
	return a + labels.PathDelimiter + b
}
