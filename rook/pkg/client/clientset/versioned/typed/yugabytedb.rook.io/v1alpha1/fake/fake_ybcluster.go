/*
Copyright The Kubernetes Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/rook/rook/pkg/apis/yugabytedb.rook.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeYBClusters implements YBClusterInterface
type FakeYBClusters struct {
	Fake *FakeYugabytedbV1alpha1
	ns   string
}

var ybclustersResource = schema.GroupVersionResource{Group: "yugabytedb.rook.io", Version: "v1alpha1", Resource: "ybclusters"}

var ybclustersKind = schema.GroupVersionKind{Group: "yugabytedb.rook.io", Version: "v1alpha1", Kind: "YBCluster"}

// Get takes name of the yBCluster, and returns the corresponding yBCluster object, and an error if there is any.
func (c *FakeYBClusters) Get(name string, options v1.GetOptions) (result *v1alpha1.YBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ybclustersResource, c.ns, name), &v1alpha1.YBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.YBCluster), err
}

// List takes label and field selectors, and returns the list of YBClusters that match those selectors.
func (c *FakeYBClusters) List(opts v1.ListOptions) (result *v1alpha1.YBClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ybclustersResource, ybclustersKind, c.ns, opts), &v1alpha1.YBClusterList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.YBClusterList{ListMeta: obj.(*v1alpha1.YBClusterList).ListMeta}
	for _, item := range obj.(*v1alpha1.YBClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested yBClusters.
func (c *FakeYBClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ybclustersResource, c.ns, opts))

}

// Create takes the representation of a yBCluster and creates it.  Returns the server's representation of the yBCluster, and an error, if there is any.
func (c *FakeYBClusters) Create(yBCluster *v1alpha1.YBCluster) (result *v1alpha1.YBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ybclustersResource, c.ns, yBCluster), &v1alpha1.YBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.YBCluster), err
}

// Update takes the representation of a yBCluster and updates it. Returns the server's representation of the yBCluster, and an error, if there is any.
func (c *FakeYBClusters) Update(yBCluster *v1alpha1.YBCluster) (result *v1alpha1.YBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ybclustersResource, c.ns, yBCluster), &v1alpha1.YBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.YBCluster), err
}

// Delete takes name of the yBCluster and deletes it. Returns an error if one occurs.
func (c *FakeYBClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ybclustersResource, c.ns, name), &v1alpha1.YBCluster{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeYBClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ybclustersResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.YBClusterList{})
	return err
}

// Patch applies the patch and returns the patched yBCluster.
func (c *FakeYBClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.YBCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ybclustersResource, c.ns, name, pt, data, subresources...), &v1alpha1.YBCluster{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.YBCluster), err
}
