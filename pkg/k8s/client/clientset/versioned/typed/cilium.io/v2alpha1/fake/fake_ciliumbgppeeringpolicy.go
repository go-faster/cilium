// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v2alpha1 "github.com/go-faster/cilium/pkg/k8s/apis/cilium.io/v2alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeCiliumBGPPeeringPolicies implements CiliumBGPPeeringPolicyInterface
type FakeCiliumBGPPeeringPolicies struct {
	Fake *FakeCiliumV2alpha1
}

var ciliumbgppeeringpoliciesResource = schema.GroupVersionResource{Group: "cilium.io", Version: "v2alpha1", Resource: "ciliumbgppeeringpolicies"}

var ciliumbgppeeringpoliciesKind = schema.GroupVersionKind{Group: "cilium.io", Version: "v2alpha1", Kind: "CiliumBGPPeeringPolicy"}

// Get takes name of the ciliumBGPPeeringPolicy, and returns the corresponding ciliumBGPPeeringPolicy object, and an error if there is any.
func (c *FakeCiliumBGPPeeringPolicies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2alpha1.CiliumBGPPeeringPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(ciliumbgppeeringpoliciesResource, name), &v2alpha1.CiliumBGPPeeringPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPPeeringPolicy), err
}

// List takes label and field selectors, and returns the list of CiliumBGPPeeringPolicies that match those selectors.
func (c *FakeCiliumBGPPeeringPolicies) List(ctx context.Context, opts v1.ListOptions) (result *v2alpha1.CiliumBGPPeeringPolicyList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(ciliumbgppeeringpoliciesResource, ciliumbgppeeringpoliciesKind, opts), &v2alpha1.CiliumBGPPeeringPolicyList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v2alpha1.CiliumBGPPeeringPolicyList{ListMeta: obj.(*v2alpha1.CiliumBGPPeeringPolicyList).ListMeta}
	for _, item := range obj.(*v2alpha1.CiliumBGPPeeringPolicyList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ciliumBGPPeeringPolicies.
func (c *FakeCiliumBGPPeeringPolicies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(ciliumbgppeeringpoliciesResource, opts))
}

// Create takes the representation of a ciliumBGPPeeringPolicy and creates it.  Returns the server's representation of the ciliumBGPPeeringPolicy, and an error, if there is any.
func (c *FakeCiliumBGPPeeringPolicies) Create(ctx context.Context, ciliumBGPPeeringPolicy *v2alpha1.CiliumBGPPeeringPolicy, opts v1.CreateOptions) (result *v2alpha1.CiliumBGPPeeringPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(ciliumbgppeeringpoliciesResource, ciliumBGPPeeringPolicy), &v2alpha1.CiliumBGPPeeringPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPPeeringPolicy), err
}

// Update takes the representation of a ciliumBGPPeeringPolicy and updates it. Returns the server's representation of the ciliumBGPPeeringPolicy, and an error, if there is any.
func (c *FakeCiliumBGPPeeringPolicies) Update(ctx context.Context, ciliumBGPPeeringPolicy *v2alpha1.CiliumBGPPeeringPolicy, opts v1.UpdateOptions) (result *v2alpha1.CiliumBGPPeeringPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(ciliumbgppeeringpoliciesResource, ciliumBGPPeeringPolicy), &v2alpha1.CiliumBGPPeeringPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPPeeringPolicy), err
}

// Delete takes name of the ciliumBGPPeeringPolicy and deletes it. Returns an error if one occurs.
func (c *FakeCiliumBGPPeeringPolicies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(ciliumbgppeeringpoliciesResource, name, opts), &v2alpha1.CiliumBGPPeeringPolicy{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCiliumBGPPeeringPolicies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(ciliumbgppeeringpoliciesResource, listOpts)

	_, err := c.Fake.Invokes(action, &v2alpha1.CiliumBGPPeeringPolicyList{})
	return err
}

// Patch applies the patch and returns the patched ciliumBGPPeeringPolicy.
func (c *FakeCiliumBGPPeeringPolicies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2alpha1.CiliumBGPPeeringPolicy, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(ciliumbgppeeringpoliciesResource, name, pt, data, subresources...), &v2alpha1.CiliumBGPPeeringPolicy{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v2alpha1.CiliumBGPPeeringPolicy), err
}
