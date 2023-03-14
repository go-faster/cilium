// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2alpha1

import (
	v2alpha1 "github.com/go-faster/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumBGPPeeringPolicyLister helps list CiliumBGPPeeringPolicies.
// All objects returned here must be treated as read-only.
type CiliumBGPPeeringPolicyLister interface {
	// List lists all CiliumBGPPeeringPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2alpha1.CiliumBGPPeeringPolicy, err error)
	// Get retrieves the CiliumBGPPeeringPolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2alpha1.CiliumBGPPeeringPolicy, error)
	CiliumBGPPeeringPolicyListerExpansion
}

// ciliumBGPPeeringPolicyLister implements the CiliumBGPPeeringPolicyLister interface.
type ciliumBGPPeeringPolicyLister struct {
	indexer cache.Indexer
}

// NewCiliumBGPPeeringPolicyLister returns a new CiliumBGPPeeringPolicyLister.
func NewCiliumBGPPeeringPolicyLister(indexer cache.Indexer) CiliumBGPPeeringPolicyLister {
	return &ciliumBGPPeeringPolicyLister{indexer: indexer}
}

// List lists all CiliumBGPPeeringPolicies in the indexer.
func (s *ciliumBGPPeeringPolicyLister) List(selector labels.Selector) (ret []*v2alpha1.CiliumBGPPeeringPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2alpha1.CiliumBGPPeeringPolicy))
	})
	return ret, err
}

// Get retrieves the CiliumBGPPeeringPolicy from the index for a given name.
func (s *ciliumBGPPeeringPolicyLister) Get(name string) (*v2alpha1.CiliumBGPPeeringPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2alpha1.Resource("ciliumbgppeeringpolicy"), name)
	}
	return obj.(*v2alpha1.CiliumBGPPeeringPolicy), nil
}
