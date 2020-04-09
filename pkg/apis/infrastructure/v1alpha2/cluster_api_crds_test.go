package v1alpha2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	apiv1alpha2 "sigs.k8s.io/cluster-api/api/v1alpha2"
	"sigs.k8s.io/yaml"
)

func Test_NewClusterCRD(t *testing.T) {
	crd := NewClusterCRD()
	if crd == nil {
		t.Error("Cluster CRD was nil.")
	}
	if crd.Name == "" {
		t.Error("Cluster CRD name was empty.")
	}
}

func Test_GenerateClusterYAML(t *testing.T) {
	crd := NewClusterCRD()

	crdGroup := crd.Spec.Group
	crdKindLower := strings.ToLower(crd.Spec.Names.Kind)

	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "crd",
			name:     fmt.Sprintf("%s_%s.yaml", crdGroup, crdKindLower),
			resource: clusterCRD,
		},
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_%s.yaml", crdGroup, version, crdKindLower),
			resource: newClusterExampleCR(),
		},
	}

	docs := filepath.Join(root, "..", "..", "..", "..", "docs")
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d: generates %s successfully", i, tc.name), func(t *testing.T) {
			rendered, err := yaml.Marshal(tc.resource)
			if err != nil {
				t.Fatal(err)
			}
			directory := filepath.Join(docs, tc.category)
			path := filepath.Join(directory, tc.name)

			// We don't want a status in the docs YAML for the CR and CRD so that they work with `kubectl create -f <file>.yaml`.
			// This just strips off the top level `status:` and everything following.
			statusRegex := regexp.MustCompile(`(?ms)^status:.*$`)
			rendered = statusRegex.ReplaceAll(rendered, []byte(""))

			if *update {
				err := ioutil.WriteFile(path, rendered, 0644)
				if err != nil {
					t.Fatal(err)
				}
			}
			goldenFile, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(rendered, goldenFile) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(goldenFile), string(rendered)))
			}
		})
	}
}

func newClusterExampleCR() *apiv1alpha2.Cluster {
	cr := NewClusterCR()

	cr.Name = "ca1p0"
	cr.Spec = apiv1alpha2.ClusterSpec{
		// ClusterNetwork does not occur in our practice, so leaving it empty.
		//ClusterNetwork:    &apiv1alpha2.ClusterNetwork{},
		InfrastructureRef: &corev1.ObjectReference{
			APIVersion:      "infrastructure.giantswarm.io/v1alpha2",
			Kind:            "AWSCluster",
			Name:            "ca1p0",
			Namespace:       "default",
			ResourceVersion: "57975957",
			UID:             "2dc05fcd-ba76-4135-b9ea-76955e3a7966",
		},
	}

	return cr
}

func Test_NewMachineDeploymentCRD(t *testing.T) {
	crd := NewMachineDeploymentCRD()
	if crd == nil {
		t.Error("MachineDeployment CRD was nil.")
	}
	if crd.Name == "" {
		t.Error("MachineDeployment CRD name was empty.")
	}
}

func Test_GenerateMachineDeploymentYAML(t *testing.T) {
	crd := NewMachineDeploymentCRD()

	crdGroup := crd.Spec.Group
	crdKindLower := strings.ToLower(crd.Spec.Names.Kind)

	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "crd",
			name:     fmt.Sprintf("%s_%s.yaml", crdGroup, crdKindLower),
			resource: clusterCRD,
		},
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_%s.yaml", crdGroup, version, crdKindLower),
			resource: newClusterExampleCR(),
		},
	}

	docs := filepath.Join(root, "..", "..", "..", "..", "docs")
	for i, tc := range testCases {
		t.Run(fmt.Sprintf("case %d: generates %s successfully", i, tc.name), func(t *testing.T) {
			rendered, err := yaml.Marshal(tc.resource)
			if err != nil {
				t.Fatal(err)
			}
			directory := filepath.Join(docs, tc.category)
			path := filepath.Join(directory, tc.name)

			// We don't want a status in the docs YAML for the CR and CRD so that they work with `kubectl create -f <file>.yaml`.
			// This just strips off the top level `status:` and everything following.
			statusRegex := regexp.MustCompile(`(?ms)^status:.*$`)
			rendered = statusRegex.ReplaceAll(rendered, []byte(""))

			if *update {
				err := ioutil.WriteFile(path, rendered, 0644)
				if err != nil {
					t.Fatal(err)
				}
			}
			goldenFile, err := ioutil.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}

			if !bytes.Equal(rendered, goldenFile) {
				t.Fatalf("\n\n%s\n", cmp.Diff(string(goldenFile), string(rendered)))
			}
		})
	}
}

func newMachineDeploymentExampleCR() *apiv1alpha2.MachineDeployment {
	cr := NewMachineDeploymentCR()

	ten := int32(10)

	cr.Name = "e3z8q"
	cr.Spec = apiv1alpha2.MachineDeploymentSpec{
		Replicas: &ten,
		Template: apiv1alpha2.MachineTemplateSpec{
			Spec: apiv1alpha2.MachineSpec{
				InfrastructureRef: corev1.ObjectReference{
					APIVersion:      "infrastructure.giantswarm.io/v1alpha2",
					Kind:            "AWSMachineDeployment",
					Name:            "e3z8q",
					Namespace:       "default",
					ResourceVersion: "58776588",
					UID:             "b5fdf153-6f10-462b-a49f-0e821873d16b",
				},
			},
		},
	}
	cr.Status = apiv1alpha2.MachineDeploymentStatus{
		ReadyReplicas: 8,
		Replicas:      10,
	}

	return cr
}
