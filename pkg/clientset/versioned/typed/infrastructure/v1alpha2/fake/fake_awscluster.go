/*
Copyright 2019 Giant Swarm GmbH.

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
	v1alpha2 "github.com/giantswarm/apiextensions/pkg/apis/infrastructure/v1alpha2"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeAWSClusters implements AWSClusterInterface
type FakeAWSClusters struct {
	Fake *FakeInfrastructureV1alpha2
}

var awsclustersResource = schema.GroupVersionResource{Group: "infrastructure.giantswarm.io", Version: "v1alpha2", Resource: "awsclusters"}

var awsclustersKind = schema.GroupVersionKind{Group: "infrastructure.giantswarm.io", Version: "v1alpha2", Kind: "AWSCluster"}

// Get takes name of the aWSCluster, and returns the corresponding aWSCluster object, and an error if there is any.
func (c *FakeAWSClusters) Get(name string, options v1.GetOptions) (result *v1alpha2.AWSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(awsclustersResource, name), &v1alpha2.AWSCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}

// List takes label and field selectors, and returns the list of AWSClusters that match those selectors.
func (c *FakeAWSClusters) List(opts v1.ListOptions) (result *v1alpha2.AWSClusterList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(awsclustersResource, awsclustersKind, opts), &v1alpha2.AWSClusterList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.AWSClusterList{ListMeta: obj.(*v1alpha2.AWSClusterList).ListMeta}
	for _, item := range obj.(*v1alpha2.AWSClusterList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSClusters.
func (c *FakeAWSClusters) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(awsclustersResource, opts))
}

// Create takes the representation of a aWSCluster and creates it.  Returns the server's representation of the aWSCluster, and an error, if there is any.
func (c *FakeAWSClusters) Create(aWSCluster *v1alpha2.AWSCluster) (result *v1alpha2.AWSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(awsclustersResource, aWSCluster), &v1alpha2.AWSCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}

// Update takes the representation of a aWSCluster and updates it. Returns the server's representation of the aWSCluster, and an error, if there is any.
func (c *FakeAWSClusters) Update(aWSCluster *v1alpha2.AWSCluster) (result *v1alpha2.AWSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(awsclustersResource, aWSCluster), &v1alpha2.AWSCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSClusters) UpdateStatus(aWSCluster *v1alpha2.AWSCluster) (*v1alpha2.AWSCluster, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(awsclustersResource, "status", aWSCluster), &v1alpha2.AWSCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}

// Delete takes name of the aWSCluster and deletes it. Returns an error if one occurs.
func (c *FakeAWSClusters) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(awsclustersResource, name), &v1alpha2.AWSCluster{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSClusters) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(awsclustersResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.AWSClusterList{})
	return err
}

// Patch applies the patch and returns the patched aWSCluster.
func (c *FakeAWSClusters) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AWSCluster, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(awsclustersResource, name, pt, data, subresources...), &v1alpha2.AWSCluster{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSCluster), err
}
