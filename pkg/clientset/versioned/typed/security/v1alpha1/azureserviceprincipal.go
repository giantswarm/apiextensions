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

package v1alpha1

import (
	"context"
	"time"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"

	v1alpha1 "github.com/giantswarm/apiextensions/v3/pkg/apis/security/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/v3/pkg/clientset/versioned/scheme"
)

// AzureServicePrincipalsGetter has a method to return a AzureServicePrincipalInterface.
// A group's client should implement this interface.
type AzureServicePrincipalsGetter interface {
	AzureServicePrincipals(namespace string) AzureServicePrincipalInterface
}

// AzureServicePrincipalInterface has methods to work with AzureServicePrincipal resources.
type AzureServicePrincipalInterface interface {
	Create(ctx context.Context, azureServicePrincipal *v1alpha1.AzureServicePrincipal, opts v1.CreateOptions) (*v1alpha1.AzureServicePrincipal, error)
	Update(ctx context.Context, azureServicePrincipal *v1alpha1.AzureServicePrincipal, opts v1.UpdateOptions) (*v1alpha1.AzureServicePrincipal, error)
	UpdateStatus(ctx context.Context, azureServicePrincipal *v1alpha1.AzureServicePrincipal, opts v1.UpdateOptions) (*v1alpha1.AzureServicePrincipal, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.AzureServicePrincipal, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.AzureServicePrincipalList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureServicePrincipal, err error)
	AzureServicePrincipalExpansion
}

// azureServicePrincipals implements AzureServicePrincipalInterface
type azureServicePrincipals struct {
	client rest.Interface
	ns     string
}

// newAzureServicePrincipals returns a AzureServicePrincipals
func newAzureServicePrincipals(c *SecurityV1alpha1Client, namespace string) *azureServicePrincipals {
	return &azureServicePrincipals{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the azureServicePrincipal, and returns the corresponding azureServicePrincipal object, and an error if there is any.
func (c *azureServicePrincipals) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.AzureServicePrincipal, err error) {
	result = &v1alpha1.AzureServicePrincipal{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureserviceprincipals").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of AzureServicePrincipals that match those selectors.
func (c *azureServicePrincipals) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.AzureServicePrincipalList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.AzureServicePrincipalList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("azureserviceprincipals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested azureServicePrincipals.
func (c *azureServicePrincipals) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("azureserviceprincipals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a azureServicePrincipal and creates it.  Returns the server's representation of the azureServicePrincipal, and an error, if there is any.
func (c *azureServicePrincipals) Create(ctx context.Context, azureServicePrincipal *v1alpha1.AzureServicePrincipal, opts v1.CreateOptions) (result *v1alpha1.AzureServicePrincipal, err error) {
	result = &v1alpha1.AzureServicePrincipal{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("azureserviceprincipals").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureServicePrincipal).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a azureServicePrincipal and updates it. Returns the server's representation of the azureServicePrincipal, and an error, if there is any.
func (c *azureServicePrincipals) Update(ctx context.Context, azureServicePrincipal *v1alpha1.AzureServicePrincipal, opts v1.UpdateOptions) (result *v1alpha1.AzureServicePrincipal, err error) {
	result = &v1alpha1.AzureServicePrincipal{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureserviceprincipals").
		Name(azureServicePrincipal.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureServicePrincipal).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *azureServicePrincipals) UpdateStatus(ctx context.Context, azureServicePrincipal *v1alpha1.AzureServicePrincipal, opts v1.UpdateOptions) (result *v1alpha1.AzureServicePrincipal, err error) {
	result = &v1alpha1.AzureServicePrincipal{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("azureserviceprincipals").
		Name(azureServicePrincipal.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(azureServicePrincipal).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the azureServicePrincipal and deletes it. Returns an error if one occurs.
func (c *azureServicePrincipals) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureserviceprincipals").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *azureServicePrincipals) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("azureserviceprincipals").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched azureServicePrincipal.
func (c *azureServicePrincipals) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.AzureServicePrincipal, err error) {
	result = &v1alpha1.AzureServicePrincipal{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("azureserviceprincipals").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
