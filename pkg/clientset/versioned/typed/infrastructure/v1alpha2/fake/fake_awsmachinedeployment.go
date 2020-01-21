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

// FakeAWSMachineDeployments implements AWSMachineDeploymentInterface
type FakeAWSMachineDeployments struct {
	Fake *FakeInfrastructureV1alpha2
	ns   string
}

var awsmachinedeploymentsResource = schema.GroupVersionResource{Group: "infrastructure.giantswarm.io", Version: "v1alpha2", Resource: "awsmachinedeployments"}

var awsmachinedeploymentsKind = schema.GroupVersionKind{Group: "infrastructure.giantswarm.io", Version: "v1alpha2", Kind: "AWSMachineDeployment"}

// Get takes name of the aWSMachineDeployment, and returns the corresponding aWSMachineDeployment object, and an error if there is any.
func (c *FakeAWSMachineDeployments) Get(name string, options v1.GetOptions) (result *v1alpha2.AWSMachineDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsmachinedeploymentsResource, c.ns, name), &v1alpha2.AWSMachineDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSMachineDeployment), err
}

// List takes label and field selectors, and returns the list of AWSMachineDeployments that match those selectors.
func (c *FakeAWSMachineDeployments) List(opts v1.ListOptions) (result *v1alpha2.AWSMachineDeploymentList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsmachinedeploymentsResource, awsmachinedeploymentsKind, c.ns, opts), &v1alpha2.AWSMachineDeploymentList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha2.AWSMachineDeploymentList{ListMeta: obj.(*v1alpha2.AWSMachineDeploymentList).ListMeta}
	for _, item := range obj.(*v1alpha2.AWSMachineDeploymentList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSMachineDeployments.
func (c *FakeAWSMachineDeployments) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsmachinedeploymentsResource, c.ns, opts))

}

// Create takes the representation of a aWSMachineDeployment and creates it.  Returns the server's representation of the aWSMachineDeployment, and an error, if there is any.
func (c *FakeAWSMachineDeployments) Create(aWSMachineDeployment *v1alpha2.AWSMachineDeployment) (result *v1alpha2.AWSMachineDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsmachinedeploymentsResource, c.ns, aWSMachineDeployment), &v1alpha2.AWSMachineDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSMachineDeployment), err
}

// Update takes the representation of a aWSMachineDeployment and updates it. Returns the server's representation of the aWSMachineDeployment, and an error, if there is any.
func (c *FakeAWSMachineDeployments) Update(aWSMachineDeployment *v1alpha2.AWSMachineDeployment) (result *v1alpha2.AWSMachineDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsmachinedeploymentsResource, c.ns, aWSMachineDeployment), &v1alpha2.AWSMachineDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSMachineDeployment), err
}

// Delete takes name of the aWSMachineDeployment and deletes it. Returns an error if one occurs.
func (c *FakeAWSMachineDeployments) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsmachinedeploymentsResource, c.ns, name), &v1alpha2.AWSMachineDeployment{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSMachineDeployments) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsmachinedeploymentsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha2.AWSMachineDeploymentList{})
	return err
}

// Patch applies the patch and returns the patched aWSMachineDeployment.
func (c *FakeAWSMachineDeployments) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha2.AWSMachineDeployment, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsmachinedeploymentsResource, c.ns, name, pt, data, subresources...), &v1alpha2.AWSMachineDeployment{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha2.AWSMachineDeployment), err
}
