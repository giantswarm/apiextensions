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
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/watch"
	"k8s.io/client-go/discovery"
	fakediscovery "k8s.io/client-go/discovery/fake"
	"k8s.io/client-go/testing"

	clientset "github.com/giantswarm/apiextensions/pkg/clientset/versioned"
	applicationv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/application/v1alpha1"
	fakeapplicationv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/application/v1alpha1/fake"
	backupv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/backup/v1alpha1"
	fakebackupv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/backup/v1alpha1/fake"
	corev1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/core/v1alpha1"
	fakecorev1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/core/v1alpha1/fake"
	examplev1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/example/v1alpha1"
	fakeexamplev1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/example/v1alpha1/fake"
	infrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/infrastructure/v1alpha2"
	fakeinfrastructurev1alpha2 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/infrastructure/v1alpha2/fake"
	providerv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/provider/v1alpha1"
	fakeproviderv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/provider/v1alpha1/fake"
	releasev1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/release/v1alpha1"
	fakereleasev1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/release/v1alpha1/fake"
	securityv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/security/v1alpha1"
	fakesecurityv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/security/v1alpha1/fake"
	toolingv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/tooling/v1alpha1"
	faketoolingv1alpha1 "github.com/giantswarm/apiextensions/pkg/clientset/versioned/typed/tooling/v1alpha1/fake"
)

// NewSimpleClientset returns a clientset that will respond with the provided objects.
// It's backed by a very simple object tracker that processes creates, updates and deletions as-is,
// without applying any validations and/or defaults. It shouldn't be considered a replacement
// for a real clientset and is mostly useful in simple unit tests.
func NewSimpleClientset(objects ...runtime.Object) *Clientset {
	o := testing.NewObjectTracker(scheme, codecs.UniversalDecoder())
	for _, obj := range objects {
		if err := o.Add(obj); err != nil {
			panic(err)
		}
	}

	cs := &Clientset{tracker: o}
	cs.discovery = &fakediscovery.FakeDiscovery{Fake: &cs.Fake}
	cs.AddReactor("*", "*", testing.ObjectReaction(o))
	cs.AddWatchReactor("*", func(action testing.Action) (handled bool, ret watch.Interface, err error) {
		gvr := action.GetResource()
		ns := action.GetNamespace()
		watch, err := o.Watch(gvr, ns)
		if err != nil {
			return false, nil, err
		}
		return true, watch, nil
	})

	return cs
}

// Clientset implements clientset.Interface. Meant to be embedded into a
// struct to get a default implementation. This makes faking out just the method
// you want to test easier.
type Clientset struct {
	testing.Fake
	discovery *fakediscovery.FakeDiscovery
	tracker   testing.ObjectTracker
}

func (c *Clientset) Discovery() discovery.DiscoveryInterface {
	return c.discovery
}

func (c *Clientset) Tracker() testing.ObjectTracker {
	return c.tracker
}

var _ clientset.Interface = &Clientset{}

// ApplicationV1alpha1 retrieves the ApplicationV1alpha1Client
func (c *Clientset) ApplicationV1alpha1() applicationv1alpha1.ApplicationV1alpha1Interface {
	return &fakeapplicationv1alpha1.FakeApplicationV1alpha1{Fake: &c.Fake}
}

// BackupV1alpha1 retrieves the BackupV1alpha1Client
func (c *Clientset) BackupV1alpha1() backupv1alpha1.BackupV1alpha1Interface {
	return &fakebackupv1alpha1.FakeBackupV1alpha1{Fake: &c.Fake}
}

// CoreV1alpha1 retrieves the CoreV1alpha1Client
func (c *Clientset) CoreV1alpha1() corev1alpha1.CoreV1alpha1Interface {
	return &fakecorev1alpha1.FakeCoreV1alpha1{Fake: &c.Fake}
}

// ExampleV1alpha1 retrieves the ExampleV1alpha1Client
func (c *Clientset) ExampleV1alpha1() examplev1alpha1.ExampleV1alpha1Interface {
	return &fakeexamplev1alpha1.FakeExampleV1alpha1{Fake: &c.Fake}
}

// InfrastructureV1alpha2 retrieves the InfrastructureV1alpha2Client
func (c *Clientset) InfrastructureV1alpha2() infrastructurev1alpha2.InfrastructureV1alpha2Interface {
	return &fakeinfrastructurev1alpha2.FakeInfrastructureV1alpha2{Fake: &c.Fake}
}

// ProviderV1alpha1 retrieves the ProviderV1alpha1Client
func (c *Clientset) ProviderV1alpha1() providerv1alpha1.ProviderV1alpha1Interface {
	return &fakeproviderv1alpha1.FakeProviderV1alpha1{Fake: &c.Fake}
}

// ReleaseV1alpha1 retrieves the ReleaseV1alpha1Client
func (c *Clientset) ReleaseV1alpha1() releasev1alpha1.ReleaseV1alpha1Interface {
	return &fakereleasev1alpha1.FakeReleaseV1alpha1{Fake: &c.Fake}
}

// SecurityV1alpha1 retrieves the SecurityV1alpha1Client
func (c *Clientset) SecurityV1alpha1() securityv1alpha1.SecurityV1alpha1Interface {
	return &fakesecurityv1alpha1.FakeSecurityV1alpha1{Fake: &c.Fake}
}

// ToolingV1alpha1 retrieves the ToolingV1alpha1Client
func (c *Clientset) ToolingV1alpha1() toolingv1alpha1.ToolingV1alpha1Interface {
	return &faketoolingv1alpha1.FakeToolingV1alpha1{Fake: &c.Fake}
}
