package v1alpha2

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

func Test_GenerateAWSMachineDeploymentYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_awsmachinedeployment.yaml", group, version),
			resource: newAWSMachineDeploymentExampleCR(),
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

func newAWSMachineDeploymentExampleCR() *AWSMachineDeployment {
	cr := NewAWSMachineDeploymentCR()

	cr.Name = "general-purpose-node-pool"
	cr.Spec = AWSMachineDeploymentSpec{
		NodePool: AWSMachineDeploymentSpecNodePool{
			Description: "General purpose worker nodes",
			Machine: AWSMachineDeploymentSpecNodePoolMachine{
				DockerVolumeSizeGB:  30,
				KubeletVolumeSizeGB: 100,
			},
			Scaling: AWSMachineDeploymentSpecNodePoolScaling{
				Min: 2,
				Max: 50,
			},
		},
		Provider: AWSMachineDeploymentSpecProvider{
			AvailabilityZones: []string{"eu-central-1b", "eu-central-1c"},
			InstanceDistribution: AWSMachineDeploymentSpecInstanceDistribution{
				OnDemandBaseCapacity:                0,
				OnDemandPercentageAboveBaseCapacity: 0,
			},
			Worker: AWSMachineDeploymentSpecProviderWorker{
				InstanceType:          "m5.4xlarge",
				UseAlikeInstanceTypes: true,
			},
		},
	}
	cr.Status = AWSMachineDeploymentStatus{
		Provider: AWSMachineDeploymentStatusProvider{
			Worker: AWSMachineDeploymentStatusProviderWorker{
				InstanceTypes: []string{
					"m4.xlarge",
					"m5.xlarge",
				},
				SpotInstances: 39,
			},
		},
	}

	return cr
}
