/*
Copyright 2020 The Kubernetes Authors.

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

// Package v1alpha3 contains API Schema definitions for the exp v1alpha3 API group
// +kubebuilder:object:generate=true
// +groupName=exp.infrastructure.cluster.x-k8s.io
package v1alpha3

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// knownTypes is the full list of objects to register with the scheme. It
// should contain all zero values of custom objects and custom object lists
// in the group version.
var knownTypes = []runtime.Object{
	&AzureMachinePool{},
	&AzureMachinePoolList{},
}

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: "exp.infrastructure.cluster.x-k8s.io", Version: "v1alpha3"}

	schemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme is used by the generated client.
	AddToScheme = schemeBuilder.AddToScheme
)

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{
	Group:   GroupVersion.Group,
	Version: GroupVersion.Version,
}

// Adds the list of known types to api.Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion, knownTypes...)
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
