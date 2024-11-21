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

// FakeCatalogTemplateVersions implements CatalogTemplateVersionInterface
type FakeCatalogTemplateVersions struct {
	Fake *FakeManagementV3
	ns   string
}

var catalogtemplateversionsResource = v3.SchemeGroupVersion.WithResource("catalogtemplateversions")

var catalogtemplateversionsKind = v3.SchemeGroupVersion.WithKind("CatalogTemplateVersion")

// Get takes name of the catalogTemplateVersion, and returns the corresponding catalogTemplateVersion object, and an error if there is any.
func (c *FakeCatalogTemplateVersions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.CatalogTemplateVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(catalogtemplateversionsResource, c.ns, name), &v3.CatalogTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CatalogTemplateVersion), err
}

// List takes label and field selectors, and returns the list of CatalogTemplateVersions that match those selectors.
func (c *FakeCatalogTemplateVersions) List(ctx context.Context, opts v1.ListOptions) (result *v3.CatalogTemplateVersionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(catalogtemplateversionsResource, catalogtemplateversionsKind, c.ns, opts), &v3.CatalogTemplateVersionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.CatalogTemplateVersionList{ListMeta: obj.(*v3.CatalogTemplateVersionList).ListMeta}
	for _, item := range obj.(*v3.CatalogTemplateVersionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested catalogTemplateVersions.
func (c *FakeCatalogTemplateVersions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(catalogtemplateversionsResource, c.ns, opts))

}

// Create takes the representation of a catalogTemplateVersion and creates it.  Returns the server's representation of the catalogTemplateVersion, and an error, if there is any.
func (c *FakeCatalogTemplateVersions) Create(ctx context.Context, catalogTemplateVersion *v3.CatalogTemplateVersion, opts v1.CreateOptions) (result *v3.CatalogTemplateVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(catalogtemplateversionsResource, c.ns, catalogTemplateVersion), &v3.CatalogTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CatalogTemplateVersion), err
}

// Update takes the representation of a catalogTemplateVersion and updates it. Returns the server's representation of the catalogTemplateVersion, and an error, if there is any.
func (c *FakeCatalogTemplateVersions) Update(ctx context.Context, catalogTemplateVersion *v3.CatalogTemplateVersion, opts v1.UpdateOptions) (result *v3.CatalogTemplateVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(catalogtemplateversionsResource, c.ns, catalogTemplateVersion), &v3.CatalogTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CatalogTemplateVersion), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCatalogTemplateVersions) UpdateStatus(ctx context.Context, catalogTemplateVersion *v3.CatalogTemplateVersion, opts v1.UpdateOptions) (*v3.CatalogTemplateVersion, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(catalogtemplateversionsResource, "status", c.ns, catalogTemplateVersion), &v3.CatalogTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CatalogTemplateVersion), err
}

// Delete takes name of the catalogTemplateVersion and deletes it. Returns an error if one occurs.
func (c *FakeCatalogTemplateVersions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(catalogtemplateversionsResource, c.ns, name, opts), &v3.CatalogTemplateVersion{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCatalogTemplateVersions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(catalogtemplateversionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.CatalogTemplateVersionList{})
	return err
}

// Patch applies the patch and returns the patched catalogTemplateVersion.
func (c *FakeCatalogTemplateVersions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.CatalogTemplateVersion, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(catalogtemplateversionsResource, c.ns, name, pt, data, subresources...), &v3.CatalogTemplateVersion{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.CatalogTemplateVersion), err
}
