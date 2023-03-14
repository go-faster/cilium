// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/go-faster/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/core/v1"
	rest "k8s.io/client-go/rest"
	testing "k8s.io/client-go/testing"
)

type FakeCoreV1 struct {
	*testing.Fake
}

func (c *FakeCoreV1) Endpoints(namespace string) v1.EndpointsInterface {
	return &FakeEndpoints{c, namespace}
}

func (c *FakeCoreV1) Namespaces() v1.NamespaceInterface {
	return &FakeNamespaces{c}
}

func (c *FakeCoreV1) Nodes() v1.NodeInterface {
	return &FakeNodes{c}
}

func (c *FakeCoreV1) Pods(namespace string) v1.PodInterface {
	return &FakePods{c, namespace}
}

func (c *FakeCoreV1) Secrets(namespace string) v1.SecretInterface {
	return &FakeSecrets{c, namespace}
}

func (c *FakeCoreV1) Services(namespace string) v1.ServiceInterface {
	return &FakeServices{c, namespace}
}

// RESTClient returns a RESTClient that is used to communicate
// with API server by this client implementation.
func (c *FakeCoreV1) RESTClient() rest.Interface {
	var ret *rest.RESTClient
	return ret
}
