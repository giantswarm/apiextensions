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

package v1alpha2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/apis/infrastructure/v1alpha2"
	scheme "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/scheme"
)

// AWSClustersGetter has a method to return a AWSClusterInterface.
// A group's client should implement this interface.
type AWSClustersGetter interface {
	AWSClusters(namespace string) AWSClusterInterface
}

// AWSClusterInterface has methods to work with AWSCluster resources.
type AWSClusterInterface interface {
	Create(ctx context.Context, aWSCluster *v1alpha2.AWSCluster, opts v1.CreateOptions) (*v1alpha2.AWSCluster, error)
	Update(ctx context.Context, aWSCluster *v1alpha2.AWSCluster, opts v1.UpdateOptions) (*v1alpha2.AWSCluster, error)
	UpdateStatus(ctx context.Context, aWSCluster *v1alpha2.AWSCluster, opts v1.UpdateOptions) (*v1alpha2.AWSCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.AWSCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.AWSClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.AWSCluster, err error)
	AWSClusterExpansion
}

// aWSClusters implements AWSClusterInterface
type aWSClusters struct {
	client rest.Interface
	ns     string
}

// newAWSClusters returns a AWSClusters
func newAWSClusters(c *InfrastructureV1alpha2Client, namespace string) *aWSClusters {
	return &aWSClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the aWSCluster, and returns the corresponding aWSCluster object, and an error if there is any.
func (c *aWSClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.AWSCluster, err error) {
	result = &v1alpha2.AWSCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AWSClusters that match those selectors.
func (c *aWSClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.AWSClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.AWSClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("awsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested aWSClusters.
func (c *aWSClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("awsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a aWSCluster and creates it.  Returns the server's representation of the aWSCluster, and an error, if there is any.
func (c *aWSClusters) Create(ctx context.Context, aWSCluster *v1alpha2.AWSCluster, opts v1.CreateOptions) (result *v1alpha2.AWSCluster, err error) {
	result = &v1alpha2.AWSCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("awsclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a aWSCluster and updates it. Returns the server's representation of the aWSCluster, and an error, if there is any.
func (c *aWSClusters) Update(ctx context.Context, aWSCluster *v1alpha2.AWSCluster, opts v1.UpdateOptions) (result *v1alpha2.AWSCluster, err error) {
	result = &v1alpha2.AWSCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsclusters").
		Name(aWSCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *aWSClusters) UpdateStatus(ctx context.Context, aWSCluster *v1alpha2.AWSCluster, opts v1.UpdateOptions) (result *v1alpha2.AWSCluster, err error) {
	result = &v1alpha2.AWSCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("awsclusters").
		Name(aWSCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(aWSCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the aWSCluster and deletes it. Returns an error if one occurs.
func (c *aWSClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *aWSClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("awsclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched aWSCluster.
func (c *aWSClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.AWSCluster, err error) {
	result = &v1alpha2.AWSCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("awsclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
