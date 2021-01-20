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

package v1alpha2

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha2 "github.com/giantswarm/apiextensions/v3/pkg/apis/infrastructure/v1alpha2"
	scheme "github.com/giantswarm/apiextensions/v3/pkg/clientset/versioned/scheme"
)

// KVMClustersGetter has a method to return a KVMClusterInterface.
// A group's client should implement this interface.
type KVMClustersGetter interface {
	KVMClusters(namespace string) KVMClusterInterface
}

// KVMClusterInterface has methods to work with KVMCluster resources.
type KVMClusterInterface interface {
	Create(ctx context.Context, kVMCluster *v1alpha2.KVMCluster, opts v1.CreateOptions) (*v1alpha2.KVMCluster, error)
	Update(ctx context.Context, kVMCluster *v1alpha2.KVMCluster, opts v1.UpdateOptions) (*v1alpha2.KVMCluster, error)
	UpdateStatus(ctx context.Context, kVMCluster *v1alpha2.KVMCluster, opts v1.UpdateOptions) (*v1alpha2.KVMCluster, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha2.KVMCluster, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha2.KVMClusterList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.KVMCluster, err error)
	KVMClusterExpansion
}

// kVMClusters implements KVMClusterInterface
type kVMClusters struct {
	client rest.Interface
	ns     string
}

// newKVMClusters returns a KVMClusters
func newKVMClusters(c *InfrastructureV1alpha2Client, namespace string) *kVMClusters {
	return &kVMClusters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the kVMCluster, and returns the corresponding kVMCluster object, and an error if there is any.
func (c *kVMClusters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha2.KVMCluster, err error) {
	result = &v1alpha2.KVMCluster{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kvmclusters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of KVMClusters that match those selectors.
func (c *kVMClusters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha2.KVMClusterList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha2.KVMClusterList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("kvmclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested kVMClusters.
func (c *kVMClusters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("kvmclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a kVMCluster and creates it.  Returns the server's representation of the kVMCluster, and an error, if there is any.
func (c *kVMClusters) Create(ctx context.Context, kVMCluster *v1alpha2.KVMCluster, opts v1.CreateOptions) (result *v1alpha2.KVMCluster, err error) {
	result = &v1alpha2.KVMCluster{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("kvmclusters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kVMCluster).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a kVMCluster and updates it. Returns the server's representation of the kVMCluster, and an error, if there is any.
func (c *kVMClusters) Update(ctx context.Context, kVMCluster *v1alpha2.KVMCluster, opts v1.UpdateOptions) (result *v1alpha2.KVMCluster, err error) {
	result = &v1alpha2.KVMCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kvmclusters").
		Name(kVMCluster.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kVMCluster).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *kVMClusters) UpdateStatus(ctx context.Context, kVMCluster *v1alpha2.KVMCluster, opts v1.UpdateOptions) (result *v1alpha2.KVMCluster, err error) {
	result = &v1alpha2.KVMCluster{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("kvmclusters").
		Name(kVMCluster.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(kVMCluster).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the kVMCluster and deletes it. Returns an error if one occurs.
func (c *kVMClusters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kvmclusters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *kVMClusters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("kvmclusters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched kVMCluster.
func (c *kVMClusters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha2.KVMCluster, err error) {
	result = &v1alpha2.KVMCluster{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("kvmclusters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}