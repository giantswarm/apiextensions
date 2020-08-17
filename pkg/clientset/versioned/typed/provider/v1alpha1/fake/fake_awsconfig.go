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

	v1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/apis/provider/v1alpha1"
)

// FakeAWSConfigs implements AWSConfigInterface
type FakeAWSConfigs struct {
	Fake *FakeProviderV1alpha1
	ns   string
}

var awsconfigsResource = schema.GroupVersionResource{Group: "provider.giantswarm.io", Version: "v1alpha1", Resource: "awsconfigs"}

var awsconfigsKind = schema.GroupVersionKind{Group: "provider.giantswarm.io", Version: "v1alpha1", Kind: "AWSConfig"}

// Get takes name of the aWSConfig, and returns the corresponding aWSConfig object, and an error if there is any.
func (c *FakeAWSConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AWSConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(awsconfigsResource, c.ns, name), &v1alpha1.AWSConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSConfig), err
}

// List takes label and field selectors, and returns the list of AWSConfigs that match those selectors.
func (c *FakeAWSConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AWSConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(awsconfigsResource, awsconfigsKind, c.ns, opts), &v1alpha1.AWSConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.AWSConfigList{ListMeta: obj.(*v1alpha1.AWSConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.AWSConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested aWSConfigs.
func (c *FakeAWSConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(awsconfigsResource, c.ns, opts))

}

// Create takes the representation of a aWSConfig and creates it.  Returns the server's representation of the aWSConfig, and an error, if there is any.
func (c *FakeAWSConfigs) Create(ctx context.Context, aWSConfig *v1alpha1.AWSConfig, opts v1.CreateOptions) (result *v1alpha1.AWSConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(awsconfigsResource, c.ns, aWSConfig), &v1alpha1.AWSConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSConfig), err
}

// Update takes the representation of a aWSConfig and updates it. Returns the server's representation of the aWSConfig, and an error, if there is any.
func (c *FakeAWSConfigs) Update(ctx context.Context, aWSConfig *v1alpha1.AWSConfig, opts v1.UpdateOptions) (result *v1alpha1.AWSConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(awsconfigsResource, c.ns, aWSConfig), &v1alpha1.AWSConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeAWSConfigs) UpdateStatus(ctx context.Context, aWSConfig *v1alpha1.AWSConfig, opts v1.UpdateOptions) (*v1alpha1.AWSConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(awsconfigsResource, "status", c.ns, aWSConfig), &v1alpha1.AWSConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSConfig), err
}

// Delete takes name of the aWSConfig and deletes it. Returns an error if one occurs.
func (c *FakeAWSConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(awsconfigsResource, c.ns, name), &v1alpha1.AWSConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeAWSConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(awsconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.AWSConfigList{})
	return err
}

// Patch applies the patch and returns the patched aWSConfig.
func (c *FakeAWSConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AWSConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(awsconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.AWSConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.AWSConfig), err
}
