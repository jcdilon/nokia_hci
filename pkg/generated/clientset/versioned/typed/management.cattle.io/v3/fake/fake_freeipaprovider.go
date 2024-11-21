/*
Copyright 2024 Rancher Labs, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by main. DO NOT EDIT.

package fake

import (
	"context"

	v3 "github.com/rancher/rancher/pkg/apis/management.cattle.io/v3"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeFreeIpaProviders implements FreeIpaProviderInterface
type FakeFreeIpaProviders struct {
	Fake *FakeManagementV3
}

var freeipaprovidersResource = v3.SchemeGroupVersion.WithResource("freeipaproviders")

var freeipaprovidersKind = v3.SchemeGroupVersion.WithKind("FreeIpaProvider")

// Get takes name of the freeIpaProvider, and returns the corresponding freeIpaProvider object, and an error if there is any.
func (c *FakeFreeIpaProviders) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.FreeIpaProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(freeipaprovidersResource, name), &v3.FreeIpaProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.FreeIpaProvider), err
}

// List takes label and field selectors, and returns the list of FreeIpaProviders that match those selectors.
func (c *FakeFreeIpaProviders) List(ctx context.Context, opts v1.ListOptions) (result *v3.FreeIpaProviderList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(freeipaprovidersResource, freeipaprovidersKind, opts), &v3.FreeIpaProviderList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.FreeIpaProviderList{ListMeta: obj.(*v3.FreeIpaProviderList).ListMeta}
	for _, item := range obj.(*v3.FreeIpaProviderList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested freeIpaProviders.
func (c *FakeFreeIpaProviders) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(freeipaprovidersResource, opts))
}

// Create takes the representation of a freeIpaProvider and creates it.  Returns the server's representation of the freeIpaProvider, and an error, if there is any.
func (c *FakeFreeIpaProviders) Create(ctx context.Context, freeIpaProvider *v3.FreeIpaProvider, opts v1.CreateOptions) (result *v3.FreeIpaProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(freeipaprovidersResource, freeIpaProvider), &v3.FreeIpaProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.FreeIpaProvider), err
}

// Update takes the representation of a freeIpaProvider and updates it. Returns the server's representation of the freeIpaProvider, and an error, if there is any.
func (c *FakeFreeIpaProviders) Update(ctx context.Context, freeIpaProvider *v3.FreeIpaProvider, opts v1.UpdateOptions) (result *v3.FreeIpaProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(freeipaprovidersResource, freeIpaProvider), &v3.FreeIpaProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.FreeIpaProvider), err
}

// Delete takes name of the freeIpaProvider and deletes it. Returns an error if one occurs.
func (c *FakeFreeIpaProviders) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(freeipaprovidersResource, name, opts), &v3.FreeIpaProvider{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeFreeIpaProviders) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(freeipaprovidersResource, listOpts)

	_, err := c.Fake.Invokes(action, &v3.FreeIpaProviderList{})
	return err
}

// Patch applies the patch and returns the patched freeIpaProvider.
func (c *FakeFreeIpaProviders) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.FreeIpaProvider, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(freeipaprovidersResource, name, pt, data, subresources...), &v3.FreeIpaProvider{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v3.FreeIpaProvider), err
}
