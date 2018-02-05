/*
Copyright 2018 Giant Swarm GmbH.

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

// ClusterConfigsGetter has a method to return a ClusterConfigInterface.
// A group's client should implement this interface.
type ClusterConfigsGetter interface {
	ClusterConfigs(namespace string) ClusterConfigInterface
}

// ClusterConfigInterface has methods to work with ClusterConfig resources.
type ClusterConfigInterface interface {
	Create(*v1alpha1.ClusterConfig) (*v1alpha1.ClusterConfig, error)
	Update(*v1alpha1.ClusterConfig) (*v1alpha1.ClusterConfig, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ClusterConfig, error)
	List(opts v1.ListOptions) (*v1alpha1.ClusterConfigList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterConfig, err error)
	ClusterConfigExpansion
}

// clusterConfigs implements ClusterConfigInterface
type clusterConfigs struct {
	client rest.Interface
	ns     string
}

// newClusterConfigs returns a ClusterConfigs
func newClusterConfigs(c *CoreV1alpha1Client, namespace string) *clusterConfigs {
	return &clusterConfigs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the clusterConfig, and returns the corresponding clusterConfig object, and an error if there is any.
func (c *clusterConfigs) Get(name string, options v1.GetOptions) (result *v1alpha1.ClusterConfig, err error) {
	result = &v1alpha1.ClusterConfig{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterconfigs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ClusterConfigs that match those selectors.
func (c *clusterConfigs) List(opts v1.ListOptions) (result *v1alpha1.ClusterConfigList, err error) {
	result = &v1alpha1.ClusterConfigList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clusterconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested clusterConfigs.
func (c *clusterConfigs) Watch(opts v1.ListOptions) (watch.Interface, error) {
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clusterconfigs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Watch()
}

// Create takes the representation of a clusterConfig and creates it.  Returns the server's representation of the clusterConfig, and an error, if there is any.
func (c *clusterConfigs) Create(clusterConfig *v1alpha1.ClusterConfig) (result *v1alpha1.ClusterConfig, err error) {
	result = &v1alpha1.ClusterConfig{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clusterconfigs").
		Body(clusterConfig).
		Do().
		Into(result)
	return
}

// Update takes the representation of a clusterConfig and updates it. Returns the server's representation of the clusterConfig, and an error, if there is any.
func (c *clusterConfigs) Update(clusterConfig *v1alpha1.ClusterConfig) (result *v1alpha1.ClusterConfig, err error) {
	result = &v1alpha1.ClusterConfig{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clusterconfigs").
		Name(clusterConfig.Name).
		Body(clusterConfig).
		Do().
		Into(result)
	return
}

// Delete takes name of the clusterConfig and deletes it. Returns an error if one occurs.
func (c *clusterConfigs) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterconfigs").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *clusterConfigs) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clusterconfigs").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched clusterConfig.
func (c *clusterConfigs) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ClusterConfig, err error) {
	result = &v1alpha1.ClusterConfig{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clusterconfigs").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
