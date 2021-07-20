/*
Copyright 2021 Giant Swarm GmbH.

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
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"

	v1alpha3 "github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha3"
)

// FakeAWSControlPlanes implements AWSControlPlaneInterface
type FakeAWSControlPlanes struct {
	Fake *FakeInfrastructureV1alpha3
	ns   string
}

var awscontrolplanesResource = schema.GroupVersionResource{Group: "infrastructure.giantswarm.io", Version: "v1alpha3", Resource: "awscontrolplanes"}

var awscontrolplanesKind = schema.GroupVersionKind{Group: "infrastructure.giantswarm.io", Version: "v1alpha3", Kind: "AWSControlPlane"}

// Get takes name of the aWSControlPlane, and returns the corresponding aWSControlPlane object, and an error if there is any.
func (c *FakeAWSControlPlanes) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha3.AWSControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awscontrolplanesResource, c.ns, name), &v1alpha3.AWSControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.AWSControlPlane), err
}

// List takes label and field selectors, and returns the list of AWSControlPlanes that match those selectors.
func (c *FakeAWSControlPlanes) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha3.AWSControlPlaneList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awscontrolplanesResource, awscontrolplanesKind, c.ns, opts), &v1alpha3.AWSControlPlaneList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha3.AWSControlPlaneList{ListMeta: obj.(*v1alpha3.AWSControlPlaneList).ListMeta}
	for _, item := range obj.(*v1alpha3.AWSControlPlaneList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSControlPlanes.
func (c *FakeAWSControlPlanes) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awscontrolplanesResource, c.ns, opts))

}

// Create takes the representation of a aWSControlPlane and creates it.  Returns the server's representation of the aWSControlPlane, and an error, if there is any.
func (c *FakeAWSControlPlanes) Create(ctx context.Context, aWSControlPlane *v1alpha3.AWSControlPlane, opts v1.CreateOptions) (result *v1alpha3.AWSControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awscontrolplanesResource, c.ns, aWSControlPlane), &v1alpha3.AWSControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.AWSControlPlane), err
}

// Update takes the representation of a aWSControlPlane and updates it. Returns the server's representation of the aWSControlPlane, and an error, if there is any.
func (c *FakeAWSControlPlanes) Update(ctx context.Context, aWSControlPlane *v1alpha3.AWSControlPlane, opts v1.UpdateOptions) (result *v1alpha3.AWSControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awscontrolplanesResource, c.ns, aWSControlPlane), &v1alpha3.AWSControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.AWSControlPlane), err
}

// Delete takes name of the aWSControlPlane and deletes it. Returns an error if one occurs.
func (c *FakeAWSControlPlanes) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awscontrolplanesResource, c.ns, name), &v1alpha3.AWSControlPlane{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSControlPlanes) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awscontrolplanesResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha3.AWSControlPlaneList{})
	return err
}

// Patch applies the patch and returns the patched aWSControlPlane.
func (c *FakeAWSControlPlanes) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha3.AWSControlPlane, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awscontrolplanesResource, c.ns, name, pt, data, subresources...), &v1alpha3.AWSControlPlane{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha3.AWSControlPlane), err
}