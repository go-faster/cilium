// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package versioned

import (
	"fmt"
	"net/http"

	corev1 "github.com/go-faster/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/core/v1"
	discoveryv1 "github.com/go-faster/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/discovery/v1"
	discoveryv1beta1 "github.com/go-faster/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/discovery/v1beta1"
	networkingv1 "github.com/go-faster/cilium/pkg/k8s/slim/k8s/client/clientset/versioned/typed/networking/v1"
	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	CoreV1() corev1.CoreV1Interface
	DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface
	DiscoveryV1() discoveryv1.DiscoveryV1Interface
	NetworkingV1() networkingv1.NetworkingV1Interface
}

// Clientset contains the clients for groups.
type Clientset struct {
	*discovery.DiscoveryClient
	coreV1           *corev1.CoreV1Client
	discoveryV1beta1 *discoveryv1beta1.DiscoveryV1beta1Client
	discoveryV1      *discoveryv1.DiscoveryV1Client
	networkingV1     *networkingv1.NetworkingV1Client
}

// CoreV1 retrieves the CoreV1Client
func (c *Clientset) CoreV1() corev1.CoreV1Interface {
	return c.coreV1
}

// DiscoveryV1beta1 retrieves the DiscoveryV1beta1Client
func (c *Clientset) DiscoveryV1beta1() discoveryv1beta1.DiscoveryV1beta1Interface {
	return c.discoveryV1beta1
}

// DiscoveryV1 retrieves the DiscoveryV1Client
func (c *Clientset) DiscoveryV1() discoveryv1.DiscoveryV1Interface {
	return c.discoveryV1
}

// NetworkingV1 retrieves the NetworkingV1Client
func (c *Clientset) NetworkingV1() networkingv1.NetworkingV1Interface {
	return c.networkingV1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
// NewForConfig is equivalent to NewForConfigAndClient(c, httpClient),
// where httpClient was generated with rest.HTTPClientFor(c).
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c

	if configShallowCopy.UserAgent == "" {
		configShallowCopy.UserAgent = rest.DefaultKubernetesUserAgent()
	}

	// share the transport between all clients
	httpClient, err := rest.HTTPClientFor(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	return NewForConfigAndClient(&configShallowCopy, httpClient)
}

// NewForConfigAndClient creates a new Clientset for the given config and http client.
// Note the http client provided takes precedence over the configured transport values.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfigAndClient will generate a rate-limiter in configShallowCopy.
func NewForConfigAndClient(c *rest.Config, httpClient *http.Client) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}

	var cs Clientset
	var err error
	cs.coreV1, err = corev1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.discoveryV1beta1, err = discoveryv1beta1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.discoveryV1, err = discoveryv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	cs.networkingV1, err = networkingv1.NewForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfigAndClient(&configShallowCopy, httpClient)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	cs, err := NewForConfig(c)
	if err != nil {
		panic(err)
	}
	return cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.coreV1 = corev1.New(c)
	cs.discoveryV1beta1 = discoveryv1beta1.New(c)
	cs.discoveryV1 = discoveryv1.New(c)
	cs.networkingV1 = networkingv1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
