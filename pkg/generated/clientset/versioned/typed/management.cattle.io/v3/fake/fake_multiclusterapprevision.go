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

// FakeMultiClusterAppRevisions implements MultiClusterAppRevisionInterface
type FakeMultiClusterAppRevisions struct {
	Fake *FakeManagementV3
	ns   string
}

var multiclusterapprevisionsResource = v3.SchemeGroupVersion.WithResource("multiclusterapprevisions")

var multiclusterapprevisionsKind = v3.SchemeGroupVersion.WithKind("MultiClusterAppRevision")

// Get takes name of the multiClusterAppRevision, and returns the corresponding multiClusterAppRevision object, and an error if there is any.
func (c *FakeMultiClusterAppRevisions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v3.MultiClusterAppRevision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(multiclusterapprevisionsResource, c.ns, name), &v3.MultiClusterAppRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.MultiClusterAppRevision), err
}

// List takes label and field selectors, and returns the list of MultiClusterAppRevisions that match those selectors.
func (c *FakeMultiClusterAppRevisions) List(ctx context.Context, opts v1.ListOptions) (result *v3.MultiClusterAppRevisionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(multiclusterapprevisionsResource, multiclusterapprevisionsKind, c.ns, opts), &v3.MultiClusterAppRevisionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v3.MultiClusterAppRevisionList{ListMeta: obj.(*v3.MultiClusterAppRevisionList).ListMeta}
	for _, item := range obj.(*v3.MultiClusterAppRevisionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested multiClusterAppRevisions.
func (c *FakeMultiClusterAppRevisions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(multiclusterapprevisionsResource, c.ns, opts))

}

// Create takes the representation of a multiClusterAppRevision and creates it.  Returns the server's representation of the multiClusterAppRevision, and an error, if there is any.
func (c *FakeMultiClusterAppRevisions) Create(ctx context.Context, multiClusterAppRevision *v3.MultiClusterAppRevision, opts v1.CreateOptions) (result *v3.MultiClusterAppRevision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(multiclusterapprevisionsResource, c.ns, multiClusterAppRevision), &v3.MultiClusterAppRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.MultiClusterAppRevision), err
}

// Update takes the representation of a multiClusterAppRevision and updates it. Returns the server's representation of the multiClusterAppRevision, and an error, if there is any.
func (c *FakeMultiClusterAppRevisions) Update(ctx context.Context, multiClusterAppRevision *v3.MultiClusterAppRevision, opts v1.UpdateOptions) (result *v3.MultiClusterAppRevision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(multiclusterapprevisionsResource, c.ns, multiClusterAppRevision), &v3.MultiClusterAppRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.MultiClusterAppRevision), err
}

// Delete takes name of the multiClusterAppRevision and deletes it. Returns an error if one occurs.
func (c *FakeMultiClusterAppRevisions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(multiclusterapprevisionsResource, c.ns, name, opts), &v3.MultiClusterAppRevision{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeMultiClusterAppRevisions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(multiclusterapprevisionsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v3.MultiClusterAppRevisionList{})
	return err
}

// Patch applies the patch and returns the patched multiClusterAppRevision.
func (c *FakeMultiClusterAppRevisions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v3.MultiClusterAppRevision, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(multiclusterapprevisionsResource, c.ns, name, pt, data, subresources...), &v3.MultiClusterAppRevision{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v3.MultiClusterAppRevision), err
}