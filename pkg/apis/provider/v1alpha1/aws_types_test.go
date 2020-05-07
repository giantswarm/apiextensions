package v1alpha1

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	goruntime "runtime"
	"testing"

	"github.com/google/go-cmp/cmp"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/yaml"
)

var (
	_, b, _, _ = goruntime.Caller(0)
	root       = filepath.Dir(b)

	// This flag allows to call the tests like
	//
	//   go test -v ./pkg/apis/provider/v1alpha1 -update
	//
	// to create/overwrite the YAML files in /docs/crd and /docs/cr.
	update = flag.Bool("update", false, "update generated YAMLs")
)

func Test_GenerateAWSConfigYAML(t *testing.T) {
	testCases := []struct {
		category string
		name     string
		resource runtime.Object
	}{
		{
			category: "cr",
			name:     fmt.Sprintf("%s_%s_awsconfig.yaml", group, version),
			resource: newAWSConfigExampleCR(),
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

func newAWSConfigExampleCR() *AWSConfig {
	cr := NewAWSConfigCR()

	cr.Name = "l8zrw"
	cr.Spec = AWSConfigSpec{
		AWS: AWSConfigSpecAWS{
			API: AWSConfigSpecAWSAPI{
				ELB: AWSConfigSpecAWSAPIELB{
					IdleTimeoutSeconds: 0,
				},
				HostedZones: "",
			},
			AvailabilityZones: 1,
			AZ:                "eu-central-1a",
			CredentialSecret: CredentialSecret{
				Name:      "credential-default",
				Namespace: "giantswarm",
			},
			Etcd: AWSConfigSpecAWSEtcd{
				ELB: AWSConfigSpecAWSEtcdELB{
					IdleTimeoutSeconds: 0,
				},
				HostedZones: "",
			},
			HostedZones: AWSConfigSpecAWSHostedZones{
				API: AWSConfigSpecAWSHostedZonesZone{
					Name: "gauss.eu-central-1.aws.gigantic.io",
				},
				Etcd: AWSConfigSpecAWSHostedZonesZone{
					Name: "gauss.eu-central-1.aws.gigantic.io",
				},
				Ingress: AWSConfigSpecAWSHostedZonesZone{
					Name: "gauss.eu-central-1.aws.gigantic.io",
				},
			},
			Ingress: AWSConfigSpecAWSIngress{
				ELB: AWSConfigSpecAWSIngressELB{
					IdleTimeoutSeconds: 0,
				},
				HostedZones: "",
			},
			Masters: []AWSConfigSpecAWSNode{
				AWSConfigSpecAWSNode{
					DockerVolumeSizeGB: 0,
					ImageID:            "ami-90c152ff",
					InstanceType:       "m4.xlarge",
				},
			},
			Region: "eu-central-1",
			VPC: AWSConfigSpecAWSVPC{
				CIDR:              "",
				PeerID:            "vpc-02030541ba719061c",
				PrivateSubnetCIDR: "",
				PublicSubnetCIDR:  "",
				RouteTableNames: []string{
					"gauss_private_0",
					"gauss_private_1",
					"gauss_private_2",
				},
			},
			Workers: []AWSConfigSpecAWSNode{
				AWSConfigSpecAWSNode{
					DockerVolumeSizeGB: 0,
					ImageID:            "ami-90c152ff",
					InstanceType:       "m4.xlarge",
				},
				AWSConfigSpecAWSNode{
					DockerVolumeSizeGB: 0,
					ImageID:            "ami-90c152ff",
					InstanceType:       "m4.xlarge",
				},
				AWSConfigSpecAWSNode{
					DockerVolumeSizeGB: 0,
					ImageID:            "ami-90c152ff",
					InstanceType:       "m4.xlarge",
				},
			},
		},
		Cluster: Cluster{
			Calico: ClusterCalico{
				CIDR:   16,
				MTU:    1430,
				Subnet: "10.2.0.0",
			},
			Customer: ClusterCustomer{
				ID: "acme",
			},
			Docker: ClusterDocker{
				Daemon: ClusterDockerDaemon{
					CIDR: "172.17.0.1/16",
				},
			},
			Etcd: ClusterEtcd{
				AltNames: "",
				Domain:   "etcd.l8zrw.k8s.gauss.eu-central-1.aws.gigantic.io",
				Port:     2379,
				Prefix:   "giantswarm.io",
			},
			ID: cr.Name,
			Kubernetes: ClusterKubernetes{
				API: ClusterKubernetesAPI{
					ClusterIPRange: "172.31.0.0/16",
					Domain:         "api.l8zrw.k8s.gauss.eu-central-1.aws.gigantic.io",
					SecurePort:     443,
				},
				CloudProvider: "aws",
				DNS: ClusterKubernetesDNS{
					IP: "172.31.0.10",
				},
				Domain: "cluster.local",
				IngressController: ClusterKubernetesIngressController{
					Docker: ClusterKubernetesIngressControllerDocker{
						Image: "quay.io/giantswarm/nginx-ingress-controller:0.9.0-beta.11",
					},
					Domain:                "ingress.l8zrw.k8s.gauss.eu-central-1.aws.gigantic.io",
					InsecurePort:          30010,
					SecurePort:            30011,
					LoadBalancerType:      LoadBalancerTypePublic,
					ExternalTrafficPolicy: ExternalTrafficPolicyLocal,
				},
				Kubelet: ClusterKubernetesKubelet{
					AltNames: "kubernetes,kubernetes.default,kubernetes.default.svc,kubernetes.default.svc.cluster.local",
					Domain:   "worker.l8zrw.k8s.gauss.eu-central-1.aws.gigantic.io",
					Labels:   "aws-operator.giantswarm.io/version=5.5.1-dev,giantswarm.io/provider=aws",
					Port:     10250,
				},
				NetworkSetup: ClusterKubernetesNetworkSetup{
					Docker: ClusterKubernetesNetworkSetupDocker{
						Image: "quay.io/giantswarm/k8s-setup-network-environment:1f4ffc52095ac368847ce3428ea99b257003d9b9",
					},
					KubeProxy: ClusterKubernetesNetworkSetupKubeProxy{
						ConntrackMaxPerCore: 1000,
						// TODO: find out how to add these fields which occur in practice:
						// - conntrackMin
						// - tcpCloseWaitTimeout
						// - tcpEstablishedTimeout
					},
				},
				SSH: ClusterKubernetesSSH{
					UserList: []ClusterKubernetesSSHUser{
						ClusterKubernetesSSHUser{
							Name:      "joe",
							PublicKey: "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAACAQCuJvxy3FKGrfJ4XB5exEdKXiqqteXEPFzPtex6dC0lHyigtO7l+NXXbs9Lga2+Ifs0Tza92MRhg/FJ+6za3oULFo7+gDyt86DIkZkMFdnSv9+YxYe+g4zqakSV+bLVf2KP6krUGJb7t4Nb+gGH62AiUx+58Onxn5rvYC0/AXOYhkAiH8PydXTDJDPhSA/qWSWEeCQistpZEDFnaVi0e7uq/k3hWJ+v9Gz0q---SHORTENED---G7iIV0Y6o9w5gIHJxf6+8X70DCuVDx9OLHmjjMyGnd+1c3yTFMUdugtvmeiGW== joe",
						},
					},
				},
			},
			Masters: []ClusterNode{
				ClusterNode{
					ID: "6t04n",
				},
			},
			Scaling: ClusterScaling{
				Min: 3,
				Max: 3,
			},
			Workers: []ClusterNode{
				ClusterNode{
					ID: "by3fd",
				},
				ClusterNode{
					ID: "z4yi6",
				},
				ClusterNode{
					ID: "mkpv8",
				},
			},
		},
		VersionBundle: AWSConfigSpecVersionBundle{
			Version: "5.5.1-dev",
		},
	}

	return cr
}
