// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by lister-gen. DO NOT EDIT.

package v2

import (
	v2 "github.com/go-faster/cilium/pkg/k8s/apis/cilium.io/v2"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CiliumNetworkPolicyLister helps list CiliumNetworkPolicies.
// All objects returned here must be treated as read-only.
type CiliumNetworkPolicyLister interface {
	// List lists all CiliumNetworkPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.CiliumNetworkPolicy, err error)
	// CiliumNetworkPolicies returns an object that can list and get CiliumNetworkPolicies.
	CiliumNetworkPolicies(namespace string) CiliumNetworkPolicyNamespaceLister
	CiliumNetworkPolicyListerExpansion
}

// ciliumNetworkPolicyLister implements the CiliumNetworkPolicyLister interface.
type ciliumNetworkPolicyLister struct {
	indexer cache.Indexer
}

// NewCiliumNetworkPolicyLister returns a new CiliumNetworkPolicyLister.
func NewCiliumNetworkPolicyLister(indexer cache.Indexer) CiliumNetworkPolicyLister {
	return &ciliumNetworkPolicyLister{indexer: indexer}
}

// List lists all CiliumNetworkPolicies in the indexer.
func (s *ciliumNetworkPolicyLister) List(selector labels.Selector) (ret []*v2.CiliumNetworkPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumNetworkPolicy))
	})
	return ret, err
}

// CiliumNetworkPolicies returns an object that can list and get CiliumNetworkPolicies.
func (s *ciliumNetworkPolicyLister) CiliumNetworkPolicies(namespace string) CiliumNetworkPolicyNamespaceLister {
	return ciliumNetworkPolicyNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CiliumNetworkPolicyNamespaceLister helps list and get CiliumNetworkPolicies.
// All objects returned here must be treated as read-only.
type CiliumNetworkPolicyNamespaceLister interface {
	// List lists all CiliumNetworkPolicies in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v2.CiliumNetworkPolicy, err error)
	// Get retrieves the CiliumNetworkPolicy from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v2.CiliumNetworkPolicy, error)
	CiliumNetworkPolicyNamespaceListerExpansion
}

// ciliumNetworkPolicyNamespaceLister implements the CiliumNetworkPolicyNamespaceLister
// interface.
type ciliumNetworkPolicyNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CiliumNetworkPolicies in the indexer for a given namespace.
func (s ciliumNetworkPolicyNamespaceLister) List(selector labels.Selector) (ret []*v2.CiliumNetworkPolicy, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v2.CiliumNetworkPolicy))
	})
	return ret, err
}

// Get retrieves the CiliumNetworkPolicy from the indexer for a given namespace and name.
func (s ciliumNetworkPolicyNamespaceLister) Get(name string) (*v2.CiliumNetworkPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v2.Resource("ciliumnetworkpolicy"), name)
	}
	return obj.(*v2.CiliumNetworkPolicy), nil
}
