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

package versioned

import (
	"fmt"

	discovery "k8s.io/client-go/discovery"
	rest "k8s.io/client-go/rest"
	flowcontrol "k8s.io/client-go/util/flowcontrol"

	applicationv1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/typed/application/v1alpha1"
	backupv1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/typed/backup/v1alpha1"
	corev1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/typed/core/v1alpha1"
	examplev1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/typed/example/v1alpha1"
	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/typed/infrastructure/v1alpha2"
	providerv1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/typed/provider/v1alpha1"
	releasev1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/typed/release/v1alpha1"
	securityv1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/typed/security/v1alpha1"
	toolingv1alpha1 "github.com/giantswarm/apiextensions/v2/pkg/clientset/versioned/typed/tooling/v1alpha1"
)

type Interface interface {
	Discovery() discovery.DiscoveryInterface
	ApplicationV1alpha1() applicationv1alpha1.ApplicationV1alpha1Interface
	BackupV1alpha1() backupv1alpha1.BackupV1alpha1Interface
	CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface
	ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1Interface
	InfrastructureV1alpha2() infrastructurev1alpha2.InfrastructureV1alpha2Interface
	ProviderV1alpha1() providerv1alpha1.ProviderV1alpha1Interface
	ReleaseV1alpha1() releasev1alpha1.ReleaseV1alpha1Interface
	SecurityV1alpha1() securityv1alpha1.SecurityV1alpha1Interface
	ToolingV1alpha1() toolingv1alpha1.ToolingV1alpha1Interface
}

// Clientset contains the clients for groups. Each group has exactly one
// version included in a Clientset.
type Clientset struct {
	*discovery.DiscoveryClient
	applicationV1alpha1    *applicationv1alpha1.ApplicationV1alpha1Client
	backupV1alpha1         *backupv1alpha1.BackupV1alpha1Client
	coreV1alpha1           *corev1alpha1.CoreV1alpha1Client
	exampleV1alpha1        *examplev1alpha1.ExampleV1alpha1Client
	infrastructureV1alpha2 *infrastructurev1alpha2.InfrastructureV1alpha2Client
	providerV1alpha1       *providerv1alpha1.ProviderV1alpha1Client
	releaseV1alpha1        *releasev1alpha1.ReleaseV1alpha1Client
	securityV1alpha1       *securityv1alpha1.SecurityV1alpha1Client
	toolingV1alpha1        *toolingv1alpha1.ToolingV1alpha1Client
}

// ApplicationV1alpha1 retrieves the ApplicationV1alpha1Client
func (c *Clientset) ApplicationV1alpha1() applicationv1alpha1.ApplicationV1alpha1Interface {
	return c.applicationV1alpha1
}

// BackupV1alpha1 retrieves the BackupV1alpha1Client
func (c *Clientset) BackupV1alpha1() backupv1alpha1.BackupV1alpha1Interface {
	return c.backupV1alpha1
}

// CoreV1alpha1 retrieves the CoreV1alpha1Client
func (c *Clientset) CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface {
	return c.coreV1alpha1
}

// ExampleV1alpha1 retrieves the ExampleV1alpha1Client
func (c *Clientset) ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1Interface {
	return c.exampleV1alpha1
}

// InfrastructureV1alpha2 retrieves the InfrastructureV1alpha2Client
func (c *Clientset) InfrastructureV1alpha2() infrastructurev1alpha2.InfrastructureV1alpha2Interface {
	return c.infrastructureV1alpha2
}

// ProviderV1alpha1 retrieves the ProviderV1alpha1Client
func (c *Clientset) ProviderV1alpha1() providerv1alpha1.ProviderV1alpha1Interface {
	return c.providerV1alpha1
}

// ReleaseV1alpha1 retrieves the ReleaseV1alpha1Client
func (c *Clientset) ReleaseV1alpha1() releasev1alpha1.ReleaseV1alpha1Interface {
	return c.releaseV1alpha1
}

// SecurityV1alpha1 retrieves the SecurityV1alpha1Client
func (c *Clientset) SecurityV1alpha1() securityv1alpha1.SecurityV1alpha1Interface {
	return c.securityV1alpha1
}

// ToolingV1alpha1 retrieves the ToolingV1alpha1Client
func (c *Clientset) ToolingV1alpha1() toolingv1alpha1.ToolingV1alpha1Interface {
	return c.toolingV1alpha1
}

// Discovery retrieves the DiscoveryClient
func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	if c == nil {
		return nil
	}
	return c.DiscoveryClient
}

// NewForConfig creates a new Clientset for the given config.
// If config's RateLimiter is not set and QPS and Burst are acceptable,
// NewForConfig will generate a rate-limiter in configShallowCopy.
func NewForConfig(c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		if configShallowCopy.Burst <= 0 {
			return nil, fmt.Errorf("burst is required to be greater than 0 when RateLimiter is not set and QPS is set to greater than 0")
		}
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error
	cs.applicationV1alpha1, err = applicationv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.backupV1alpha1, err = backupv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.coreV1alpha1, err = corev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.exampleV1alpha1, err = examplev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.infrastructureV1alpha2, err = infrastructurev1alpha2.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.providerV1alpha1, err = providerv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.releaseV1alpha1, err = releasev1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.securityV1alpha1, err = securityv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	cs.toolingV1alpha1, err = toolingv1alpha1.NewForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}

	cs.DiscoveryClient, err = discovery.NewDiscoveryClientForConfig(&configShallowCopy)
	if err != nil {
		return nil, err
	}
	return &cs, nil
}

// NewForConfigOrDie creates a new Clientset for the given config and
// panics if there is an error in the config.
func NewForConfigOrDie(c *rest.Config) *Clientset {
	var cs Clientset
	cs.applicationV1alpha1 = applicationv1alpha1.NewForConfigOrDie(c)
	cs.backupV1alpha1 = backupv1alpha1.NewForConfigOrDie(c)
	cs.coreV1alpha1 = corev1alpha1.NewForConfigOrDie(c)
	cs.exampleV1alpha1 = examplev1alpha1.NewForConfigOrDie(c)
	cs.infrastructureV1alpha2 = infrastructurev1alpha2.NewForConfigOrDie(c)
	cs.providerV1alpha1 = providerv1alpha1.NewForConfigOrDie(c)
	cs.releaseV1alpha1 = releasev1alpha1.NewForConfigOrDie(c)
	cs.securityV1alpha1 = securityv1alpha1.NewForConfigOrDie(c)
	cs.toolingV1alpha1 = toolingv1alpha1.NewForConfigOrDie(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClientForConfigOrDie(c)
	return &cs
}

// New creates a new Clientset for the given RESTClient.
func New(c rest.Interface) *Clientset {
	var cs Clientset
	cs.applicationV1alpha1 = applicationv1alpha1.New(c)
	cs.backupV1alpha1 = backupv1alpha1.New(c)
	cs.coreV1alpha1 = corev1alpha1.New(c)
	cs.exampleV1alpha1 = examplev1alpha1.New(c)
	cs.infrastructureV1alpha2 = infrastructurev1alpha2.New(c)
	cs.providerV1alpha1 = providerv1alpha1.New(c)
	cs.releaseV1alpha1 = releasev1alpha1.New(c)
	cs.securityV1alpha1 = securityv1alpha1.New(c)
	cs.toolingV1alpha1 = toolingv1alpha1.New(c)

	cs.DiscoveryClient = discovery.NewDiscoveryClient(c)
	return &cs
}
