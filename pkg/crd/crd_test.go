package crd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	goruntime "runtime"
	"strings"
	"testing"

	"github.com/giantswarm/microerror"
	"github.com/markbates/pkger"
)

var (
	_, b, _, _ = goruntime.Caller(0)
	testDirectory       = filepath.Dir(b)
)

func Test_PkgerUpToDate(t *testing.T) {
	root := filepath.Join(testDirectory, "..", "..")
	err := pkger.Walk("/config/crd/bases", func(fullPath string, info os.FileInfo, err error) error {
		// An unknown error, stop walking
		if err != nil {
			return microerror.Mask(err)
		}
		// Skip directories and any other files after a match has been found
		if info.IsDir() {
			return nil
		}

		// pkger files have a path like github.com/giantswarm/apiextensions:/config/crd/bases/release.giantswarm.io_releases.yaml
		split := strings.Split(fullPath, ":")
		path := split[1]
		extension := filepath.Ext(path)
		// Skip non-yaml files
		if extension != ".yaml" {
			return nil
		}

		virtualFile, err := pkger.Open(path)
		if err != nil {
			return microerror.Mask(err)
		}
		virtualYaml, err := ioutil.ReadAll(virtualFile)
		if err != nil {
			return microerror.Mask(err)
		}

		localPath := filepath.Join(root, strings.TrimPrefix(path, "/"))
		localYaml, err := ioutil.ReadFile(localPath)
		if err != nil {
			return microerror.Mask(err)
		}

		if string(virtualYaml) != string(localYaml) {
			t.Errorf("local file doesn't match virtual file: %s", path)
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
}

func Test_LoadAll(t *testing.T) {
	groupKinds := map[string][]string{
		"application.giantswarm.io": {
			"AppCatalog",
			"App",
			"Chart",
		},
		"backup.giantswarm.io": {
			"ETCDBackup",
		},
		"core.giantswarm.io": {
			"AWSClusterConfig",
			"AzureClusterConfig",
			"CertConfig",
			"ChartConfig",
			"DrainerConfig",
			"DraughtsmanConfig",
			"FlannelConfig",
			"Ignition",
			"IngressConfig",
			"KVMClusterConfig",
			"NodeConfig",
			"StorageConfig",
		},
		"example.giantswarm.io": {
			"MemcachedConfig",
		},
		"infrastructure.giantswarm.io": {
			"AWSCluster",
			"AWSControlPlane",
			"AWSMachineDeployment",
			"G8sControlPlane",
		},
		"provider.giantswarm.io": {
			"AWSConfig",
			"AzureConfig",
			"KVMConfig",
		},
		"release.giantswarm.io": {
			"Release",
			"ReleaseCycle",
		},
	}
	groupCRDVersions := map[string]string{
		"application.giantswarm.io":    "v1beta1",
		"backup.giantswarm.io":         "v1beta1",
		"core.giantswarm.io":           "v1beta1",
		"example.giantswarm.io":        "v1beta1",
		"infrastructure.giantswarm.io": "v1",
		"provider.giantswarm.io":       "v1beta1",
		"release.giantswarm.io":        "v1beta1",
	}

	count := 0
	for group, kinds := range groupKinds {
		crdVersion := groupCRDVersions[group]
		for _, kind := range kinds {
			name := fmt.Sprintf("case %d: %s in %s as %s CRD", count, kind, group, crdVersion)
			count++
			t.Run(name, func(t *testing.T) {
				defer func() {
					err := recover()
					if err != nil {
						t.Errorf("unexpected panic: %#v", err)
					}
				}()
				var crd interface{}
				switch crdVersion {
				case "v1beta1":
					crd = LoadV1Beta1(group, kind)
				case "v1":
					crd = LoadV1(group, kind)
				}
				if crd == nil {
					t.Errorf("nil crd")
				}
			})
		}
	}
}

func Test_Load(t *testing.T) {
	testCases := []struct {
		name            string
		inputGroup      string
		inputKind       string
		inputCRDVersion string
		panicMatcher    func(ret interface{}) bool
	}{
		{
			name:            "case 0: v1beta1 CRD loads normally",
			inputGroup:      "application.giantswarm.io",
			inputKind:       "App",
			inputCRDVersion: "v1beta1",
			panicMatcher: func(ret interface{}) bool {
				return ret == nil
			},
		},
		{
			name:            "case 1: non-existent CRD panics with notFoundError",
			inputGroup:      "application.giantswarm.io",
			inputKind:       "Bad",
			inputCRDVersion: "v1beta1",
			panicMatcher: func(ret interface{}) bool {
				err, ok := ret.(error)
				return ok && IsNotFound(err)
			},
		},
		{
			name:            "case 2: incorrect CRD version panics with conversionFailedError",
			inputGroup:      "application.giantswarm.io",
			inputKind:       "App",
			inputCRDVersion: "v1",
			panicMatcher: func(ret interface{}) bool {
				err, ok := ret.(error)
				return ok && IsConversionFailed(err)
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			defer func() {
				err := recover()
				if !tc.panicMatcher(err) {
					t.Errorf("unexpected panic: %#v", err)
				}
			}()
			var crd interface{}
			switch tc.inputCRDVersion {
			case "v1beta1":
				crd = LoadV1Beta1(tc.inputGroup, tc.inputKind)
			case "v1":
				crd = LoadV1(tc.inputGroup, tc.inputKind)
			}
			if crd == nil {
				t.Errorf("nil crd")
			}
		})
	}
}
