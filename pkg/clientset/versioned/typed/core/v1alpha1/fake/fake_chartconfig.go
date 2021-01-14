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

	v1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/core/v1alpha1"
)

// FakeChartConfigs implements ChartConfigInterface
type FakeChartConfigs struct {
	Fake *FakeCoreV1alpha1
	ns   string
}

var chartconfigsResource = schema.GroupVersionResource{Group: "core.giantswarm.io", Version: "v1alpha1", Resource: "chartconfigs"}

var chartconfigsKind = schema.GroupVersionKind{Group: "core.giantswarm.io", Version: "v1alpha1", Kind: "ChartConfig"}

// Get takes name of the chartConfig, and returns the corresponding chartConfig object, and an error if there is any.
func (c *FakeChartConfigs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ChartConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(chartconfigsResource, c.ns, name), &v1alpha1.ChartConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartConfig), err
}

// List takes label and field selectors, and returns the list of ChartConfigs that match those selectors.
func (c *FakeChartConfigs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ChartConfigList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(chartconfigsResource, chartconfigsKind, c.ns, opts), &v1alpha1.ChartConfigList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ChartConfigList{ListMeta: obj.(*v1alpha1.ChartConfigList).ListMeta}
	for _, item := range obj.(*v1alpha1.ChartConfigList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested chartConfigs.
func (c *FakeChartConfigs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(chartconfigsResource, c.ns, opts))

}

// Create takes the representation of a chartConfig and creates it.  Returns the server's representation of the chartConfig, and an error, if there is any.
func (c *FakeChartConfigs) Create(ctx context.Context, chartConfig *v1alpha1.ChartConfig, opts v1.CreateOptions) (result *v1alpha1.ChartConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(chartconfigsResource, c.ns, chartConfig), &v1alpha1.ChartConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartConfig), err
}

// Update takes the representation of a chartConfig and updates it. Returns the server's representation of the chartConfig, and an error, if there is any.
func (c *FakeChartConfigs) Update(ctx context.Context, chartConfig *v1alpha1.ChartConfig, opts v1.UpdateOptions) (result *v1alpha1.ChartConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(chartconfigsResource, c.ns, chartConfig), &v1alpha1.ChartConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartConfig), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeChartConfigs) UpdateStatus(ctx context.Context, chartConfig *v1alpha1.ChartConfig, opts v1.UpdateOptions) (*v1alpha1.ChartConfig, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(chartconfigsResource, "status", c.ns, chartConfig), &v1alpha1.ChartConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartConfig), err
}

// Delete takes name of the chartConfig and deletes it. Returns an error if one occurs.
func (c *FakeChartConfigs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(chartconfigsResource, c.ns, name), &v1alpha1.ChartConfig{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeChartConfigs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(chartconfigsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.ChartConfigList{})
	return err
}

// Patch applies the patch and returns the patched chartConfig.
func (c *FakeChartConfigs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ChartConfig, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(chartconfigsResource, c.ns, name, pt, data, subresources...), &v1alpha1.ChartConfig{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ChartConfig), err
}
