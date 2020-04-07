/*
Copyright 2020 Giant Swarm GmbH.

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

// Code generated by apiextensions/generator. DO NOT EDIT.

package fake

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIgnitions implements IgnitionInterface
type FakeIgnitions struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var ignitionsResource = schema.GroupVersionResource{Group: "core.giantswarm.io", Version: "v1alpha1", Resource: "ignitions"}

var ignitionsKind = schema.GroupVersionKind{Group: "core.giantswarm.io", Version: "v1alpha1", Kind: "Ignition"}

// Get takes name of the ignition, and returns the corresponding ignition object, and an error if there is any.
func (c *FakeIgnitions) Get(name string, options v1.GetOptions) (result *v1alpha1.Ignition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(ignitionsResource, c.ns, name), &v1alpha1.Ignition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ignition), err
}

// List takes label and field selectors, and returns the list of Ignitions that match those selectors.
func (c *FakeIgnitions) List(opts v1.ListOptions) (result *v1alpha1.IgnitionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(ignitionsResource, ignitionsKind, c.ns, opts), &v1alpha1.IgnitionList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.IgnitionList{ListMeta: obj.(*v1alpha1.IgnitionList).ListMeta}
	for _, item := range obj.(*v1alpha1.IgnitionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested ignitions.
func (c *FakeIgnitions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(ignitionsResource, c.ns, opts))

}

// Create takes the representation of a ignition and creates it.  Returns the server's representation of the ignition, and an error, if there is any.
func (c *FakeIgnitions) Create(ignition *v1alpha1.Ignition) (result *v1alpha1.Ignition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(ignitionsResource, c.ns, ignition), &v1alpha1.Ignition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ignition), err
}

// Update takes the representation of a ignition and updates it. Returns the server's representation of the ignition, and an error, if there is any.
func (c *FakeIgnitions) Update(ignition *v1alpha1.Ignition) (result *v1alpha1.Ignition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(ignitionsResource, c.ns, ignition), &v1alpha1.Ignition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ignition), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIgnitions) UpdateStatus(ignition *v1alpha1.Ignition) (*v1alpha1.Ignition, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(ignitionsResource, "status", c.ns, ignition), &v1alpha1.Ignition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ignition), err
}

// Delete takes name of the ignition and deletes it. Returns an error if one occurs.
func (c *FakeIgnitions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(ignitionsResource, c.ns, name), &v1alpha1.Ignition{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIgnitions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(ignitionsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.IgnitionList{})
	return err
}

// Patch applies the patch and returns the patched ignition.
func (c *FakeIgnitions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Ignition, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(ignitionsResource, c.ns, name, pt, data, subresources...), &v1alpha1.Ignition{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.Ignition), err
}
