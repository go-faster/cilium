// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by informer-gen. DO NOT EDIT.

package v2

import (
	"context"
	time "time"

	ciliumiov2 "github.com/go-faster/cilium/pkg/k8s/apis/cilium.io/v2"
	versioned "github.com/go-faster/cilium/pkg/k8s/client/clientset/versioned"
	internalinterfaces "github.com/go-faster/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
	v2 "github.com/go-faster/cilium/pkg/k8s/client/listers/cilium.io/v2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CiliumEgressGatewayPolicyInformer provides access to a shared informer and lister for
// CiliumEgressGatewayPolicies.
type CiliumEgressGatewayPolicyInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v2.CiliumEgressGatewayPolicyLister
}

type ciliumEgressGatewayPolicyInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// NewCiliumEgressGatewayPolicyInformer constructs a new informer for CiliumEgressGatewayPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCiliumEgressGatewayPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCiliumEgressGatewayPolicyInformer(client, resyncPeriod, indexers, nil)
}

// NewFilteredCiliumEgressGatewayPolicyInformer constructs a new informer for CiliumEgressGatewayPolicy type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCiliumEgressGatewayPolicyInformer(client versioned.Interface, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumEgressGatewayPolicies().List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CiliumV2().CiliumEgressGatewayPolicies().Watch(context.TODO(), options)
			},
		},
		&ciliumiov2.CiliumEgressGatewayPolicy{},
		resyncPeriod,
		indexers,
	)
}

func (f *ciliumEgressGatewayPolicyInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCiliumEgressGatewayPolicyInformer(client, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *ciliumEgressGatewayPolicyInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&ciliumiov2.CiliumEgressGatewayPolicy{}, f.defaultInformer)
}

func (f *ciliumEgressGatewayPolicyInformer) Lister() v2.CiliumEgressGatewayPolicyLister {
	return v2.NewCiliumEgressGatewayPolicyLister(f.Informer().GetIndexer())
}
