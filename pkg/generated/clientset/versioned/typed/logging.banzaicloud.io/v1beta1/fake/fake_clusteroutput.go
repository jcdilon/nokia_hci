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

	v1beta1 "github.com/kube-logging/logging-operator/pkg/sdk/logging/api/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeClusterOutputs implements ClusterOutputInterface
type FakeClusterOutputs struct {
	Fake *FakeLoggingV1beta1
}

var clusteroutputsResource = v1beta1.GroupVersion.WithResource("clusteroutputs")

var clusteroutputsKind = v1beta1.GroupVersion.WithKind("ClusterOutput")

// Get takes name of the clusterOutput, and returns the corresponding clusterOutput object, and an error if there is any.
func (c *FakeClusterOutputs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ClusterOutput, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(clusteroutputsResource, name), &v1beta1.ClusterOutput{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterOutput), err
}

// List takes label and field selectors, and returns the list of ClusterOutputs that match those selectors.
func (c *FakeClusterOutputs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ClusterOutputList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(clusteroutputsResource, clusteroutputsKind, opts), &v1beta1.ClusterOutputList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ClusterOutputList{ListMeta: obj.(*v1beta1.ClusterOutputList).ListMeta}
	for _, item := range obj.(*v1beta1.ClusterOutputList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested clusterOutputs.
func (c *FakeClusterOutputs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(clusteroutputsResource, opts))
}

// Create takes the representation of a clusterOutput and creates it.  Returns the server's representation of the clusterOutput, and an error, if there is any.
func (c *FakeClusterOutputs) Create(ctx context.Context, clusterOutput *v1beta1.ClusterOutput, opts v1.CreateOptions) (result *v1beta1.ClusterOutput, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(clusteroutputsResource, clusterOutput), &v1beta1.ClusterOutput{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterOutput), err
}

// Update takes the representation of a clusterOutput and updates it. Returns the server's representation of the clusterOutput, and an error, if there is any.
func (c *FakeClusterOutputs) Update(ctx context.Context, clusterOutput *v1beta1.ClusterOutput, opts v1.UpdateOptions) (result *v1beta1.ClusterOutput, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(clusteroutputsResource, clusterOutput), &v1beta1.ClusterOutput{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterOutput), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeClusterOutputs) UpdateStatus(ctx context.Context, clusterOutput *v1beta1.ClusterOutput, opts v1.UpdateOptions) (*v1beta1.ClusterOutput, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(clusteroutputsResource, "status", clusterOutput), &v1beta1.ClusterOutput{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterOutput), err
}

// Delete takes name of the clusterOutput and deletes it. Returns an error if one occurs.
func (c *FakeClusterOutputs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteActionWithOptions(clusteroutputsResource, name, opts), &v1beta1.ClusterOutput{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeClusterOutputs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(clusteroutputsResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ClusterOutputList{})
	return err
}

// Patch applies the patch and returns the patched clusterOutput.
func (c *FakeClusterOutputs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ClusterOutput, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(clusteroutputsResource, name, pt, data, subresources...), &v1beta1.ClusterOutput{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ClusterOutput), err
}