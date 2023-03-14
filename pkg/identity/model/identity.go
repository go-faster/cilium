// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package model

import (
	"github.com/go-faster/cilium/api/v1/models"
	"github.com/go-faster/cilium/pkg/identity"
	"github.com/go-faster/cilium/pkg/labels"
)

func NewIdentityFromModel(base *models.Identity) *identity.Identity {
	if base == nil {
		return nil
	}

	id := &identity.Identity{
		ID:     identity.NumericIdentity(base.ID),
		Labels: make(labels.Labels, len(base.Labels)),
	}
	for _, v := range base.Labels {
		lbl := labels.ParseLabel(v)
		id.Labels[lbl.Key] = lbl
	}
	id.Sanitize()

	return id
}

func CreateModel(id *identity.Identity) *models.Identity {
	if id == nil {
		return nil
	}

	ret := &models.Identity{
		ID:     int64(id.ID),
		Labels: make([]string, 0, len(id.Labels)),
	}

	for _, v := range id.Labels {
		ret.Labels = append(ret.Labels, v.String())
	}
	return ret
}
