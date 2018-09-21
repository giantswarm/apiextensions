package v1alpha1

import (
	"encoding/json"
	"testing"
	"time"

	yaml "gopkg.in/yaml.v2"
)

func Test_Core_Release_DeepCopy_YAML(t *testing.T) {
	testCases := []struct {
		Name              string
		Bytes             []byte
		ExpectedAzureHost string
		ExpectedDateMonth time.Month
	}{
		{
			Name: "case 0",
			Bytes: []byte(`
        active: false
        authorities:
        - endpoint: http://azure-operator:8000
          name: azure-operator
          version: 2.0.0
        - endpoint: http://cert-operator:8000
          name: cert-operator
          version: 0.1.0
        - endpoint: http://chart-operator:8000
          name: chart-operator
          version: 0.3.0
        - endpoint: http://cluster-operator:8000
          name: cluster-operator
          provider: azure
          version: 0.7.0
        date: 2018-08-16T18:00:00Z
        version: 2.0.0
      `),
			ExpectedAzureHost: "azure-operator:8000",
			ExpectedDateMonth: time.August,
		},
	}

	for _, tc := range testCases {
		var r Release
		err := yaml.Unmarshal(tc.Bytes, &r.Spec)
		if err != nil {
			t.Fatalf("expected %#v got %#v", nil, err)
		}

		{
			e := tc.ExpectedAzureHost
			m := r.Spec.Authorities[0].Endpoint.Host
			if e != m {
				t.Fatalf("expected %s got %s", e, m)
			}
		}

		{
			e := tc.ExpectedDateMonth
			m := r.Spec.Date.Month()
			if e != m {
				t.Fatalf("expected %d got %d", e, m)
			}
		}
	}
}

func Test_Core_Release_DeepCopy_JSON(t *testing.T) {
	testCases := []struct {
		Name              string
		Bytes             []byte
		ExpectedAzureHost string
		ExpectedDateMonth time.Month
	}{
		{
			Name: "case 0",
			Bytes: []byte(`
        {
          "active": false,
          "authorities": [
            {
              "endpoint": "http://azure-operator:8000",
              "name": "azure-operator",
              "version": "2.0.0"
            },
            {
              "endpoint": "http://cert-operator:8000",
              "name": "cert-operator",
              "version": "0.1.0"
            },
            {
              "endpoint": "http://chart-operator:8000",
              "name": "chart-operator",
              "version": "0.3.0"
            },
            {
              "endpoint": "http://cluster-operator:8000",
              "name": "cluster-operator",
              "provider": "azure",
              "version": "0.7.0"
            }
          ],
          "date": "2018-08-16T18:00:00Z",
          "version": "2.0.0"
        }
      `),
			ExpectedAzureHost: "azure-operator:8000",
			ExpectedDateMonth: time.August,
		},
	}

	for _, tc := range testCases {
		var r Release
		err := json.Unmarshal(tc.Bytes, &r.Spec)
		if err != nil {
			t.Fatalf("expected %#v got %#v", nil, err)
		}

		{
			e := tc.ExpectedAzureHost
			m := r.Spec.Authorities[0].Endpoint.Host
			if e != m {
				t.Fatalf("expected %s got %s", e, m)
			}
		}

		{
			e := tc.ExpectedDateMonth
			m := r.Spec.Date.Month()
			if e != m {
				t.Fatalf("expected %d got %d", e, m)
			}
		}
	}
}
