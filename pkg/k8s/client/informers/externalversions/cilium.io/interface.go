// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by informer-gen. DO NOT EDIT.

package cilium

import (
	v2 "github.com/go-faster/cilium/pkg/k8s/client/informers/externalversions/cilium.io/v2"
	v2alpha1 "github.com/go-faster/cilium/pkg/k8s/client/informers/externalversions/cilium.io/v2alpha1"
	internalinterfaces "github.com/go-faster/cilium/pkg/k8s/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to each of this group's versions.
type Interface interface {
	// V2 provides access to shared informers for resources in V2.
	V2() v2.Interface
	// V2alpha1 provides access to shared informers for resources in V2alpha1.
	V2alpha1() v2alpha1.Interface
}

type group struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &group{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// V2 returns a new v2.Interface.
func (g *group) V2() v2.Interface {
	return v2.New(g.factory, g.namespace, g.tweakListOptions)
}

// V2alpha1 returns a new v2alpha1.Interface.
func (g *group) V2alpha1() v2alpha1.Interface {
	return v2alpha1.New(g.factory, g.namespace, g.tweakListOptions)
}
