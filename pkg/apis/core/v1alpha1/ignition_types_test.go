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

	"github.com/giantswarm/apiextensions/pkg/apis/core"
)

var (
	_, b, _, _ = goruntime.Caller(0)
	root       = filepath.Dir(b)
	update     = flag.Bool("update", false, "update generated YAMLs")
)

func Test_GenerateIgnitionYAML(t *testing.T) {
	testCases := []struct {
		name     string
		category string
		filename string
		resource runtime.Object
	}{
		{
			name:     fmt.Sprintf("case 1: %s_%s_ignition.yaml is generated successfully", core.Group, version),
			category: "cr",
			filename: fmt.Sprintf("%s_%s_ignition.yaml", core.Group, version),
			resource: NewIgnitionCR("abc12-master", IgnitionSpec{
				ClusterID:               "abc12",
				DisableEncryptionAtRest: false,
				Docker: IgnitionSpecDocker{
					Daemon: IgnitionSpecDockerDaemon{
						CIDR: "172.100.0.1/16",
					},
					NetworkSetup: IgnitionSpecDockerNetworkSetup{
						Image: "quay.io/giantswarm/k8s-setup-network-environment",
					},
				},
				Etcd: IgnitionSpecEtcd{
					Domain: "https://etcd.abc12.k8s.example.eu-west-1.aws.gigantic.io",
					Port:   2379,
					Prefix: "",
				},
				Extension: IgnitionSpecExtension{
					Files: nil,
					Units: nil,
					Users: nil,
				},
				Ingress: IgnitionSpecIngress{
					Disable: false,
				},
				IsMaster: true,
				Kubernetes: IgnitionSpecKubernetes{
					API: IgnitionSpecKubernetesAPI{
						Domain:     "https://abc12.k8s.example.eu-west-1.aws.gigantic.io",
						SecurePort: 443,
					},
					CloudProvider: "aws",
					DNS: IgnitionSpecKubernetesDNS{
						IP: "10.1.2.3/32",
					},
					Domain: "https://abc12.k8s.example.eu-west-1.aws.gigantic.io",
					Kubelet: IgnitionSpecKubernetesKubelet{
						Domain: "https://abc12.k8s.example.eu-west-1.aws.gigantic.io",
					},
					IPRange: "10.2.3.4/24",
					OIDC: IgnitionSpecOIDC{
						Enabled:        true,
						ClientID:       "abc12",
						IssuerURL:      "https://giantswarm.io",
						UsernameClaim:  "",
						UsernamePrefix: "gs",
						GroupsClaim:    "",
						GroupsPrefix:   "gs",
					},
				},
				Provider: "aws",
				Registry: IgnitionSpecRegistry{
					Domain:               "quay.io",
					PullProgressDeadline: "10s",
				},
				SSO: IgnitionSpecSSO{
					PublicKey: "ssh-rsa 1234567890",
				},
			}),
		},
	}

	docs := filepath.Join(root, "..", "..", "..", "..", "docs")
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			rendered, err := yaml.Marshal(tc.resource)
			if err != nil {
				t.Fatal(err)
			}

			// We don't want a status in the docs YAML for the CR and CRD so that they work with `kubectl create -f <file>.yaml`.
			// This just strips off the top level `status:` and everything following.
			re := regexp.MustCompile(`(?ms)^status:.*$`)
			rendered = re.ReplaceAll(rendered, []byte(""))

			path := filepath.Join(docs, tc.category, tc.filename)
			if *update {
				err := ioutil.WriteFile(path, rendered, 0644) // nolint
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
