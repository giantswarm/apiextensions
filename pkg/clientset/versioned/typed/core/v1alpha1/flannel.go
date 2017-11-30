/*
Copyright 2017 The Kubernetes Authors.

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

package v1alpha1

import (
	v1alpha1 "github.com/giantswarm/apiextensions/pkg/apis/core/v1alpha1"
	scheme "github.com/giantswarm/apiextensions/pkg/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// FlannelsGetter has a method to return a FlannelInterface.
// A group's client should implement this interface.
type FlannelsGetter interface {
	Flannels(namespace string) FlannelInterface
}

// FlannelInterface has methods to work with Flannel resources.
type FlannelInterface interface {
	Create(*v1alpha1.Flannel) (*v1alpha1.Flannel, error)
	Update(*v1alpha1.Flannel) (*v1alpha1.Flannel, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.Flannel, error)
	List(opts v1.ListOptions) (*v1alpha1.FlannelList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Flannel, err error)
	FlannelExpansion
}

// flannels implements FlannelInterface
type flannels struct {
	client rest.Interface
	ns     string
}

// newFlannels returns a Flannels
func newFlannels(c *CoreV1alpha1Client, namespace string) *flannels {
	return &flannels{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the flannel, and returns the corresponding flannel object, and an error if there is any.
func (c *flannels) Get(name string, options v1.GetOptions) (result *v1alpha1.Flannel, err error) {
	result = &v1alpha1.Flannel{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flannels").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of Flannels that match those selectors.
func (c *flannels) List(opts v1.ListOptions) (result *v1alpha1.FlannelList, err error) {
	result = &v1alpha1.FlannelList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("flannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested flannels.
func (c *flannels) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("flannels").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a flannel and creates it.  Returns the server's representation of the flannel, and an error, if there is any.
func (c *flannels) Create(flannel *v1alpha1.Flannel) (result *v1alpha1.Flannel, err error) {
	result = &v1alpha1.Flannel{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("flannels").
		Body(flannel).
		Do().
		Into(result)
	return
}

// Update takes the representation of a flannel and updates it. Returns the server's representation of the flannel, and an error, if there is any.
func (c *flannels) Update(flannel *v1alpha1.Flannel) (result *v1alpha1.Flannel, err error) {
	result = &v1alpha1.Flannel{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("flannels").
		Name(flannel.Name).
		Body(flannel).
		Do().
		Into(result)
	return
}

// Delete takes name of the flannel and deletes it. Returns an error if one occurs.
func (c *flannels) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flannels").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *flannels) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("flannels").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched flannel.
func (c *flannels) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.Flannel, err error) {
	result = &v1alpha1.Flannel{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("flannels").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
