package v1alpha1

import (
	"testing"
	"time"
)

func Test_Provider_Status_LatestVersion(t *testing.T) {
	testCases := []struct {
		Name           string
		StatusCluster  StatusCluster
		ExpectedSemver string
	}{
		{
			Name: "case 0",
			StatusCluster: StatusCluster{
				Versions: []StatusClusterVersion{},
			},
			ExpectedSemver: "",
		},
		{
			Name: "case 1",
			StatusCluster: StatusCluster{
				Versions: []StatusClusterVersion{
					{
						Date:   time.Unix(10, 0),
						Semver: "1.0.0",
					},
				},
			},
			ExpectedSemver: "1.0.0",
		},
		{
			Name: "case 2",
			StatusCluster: StatusCluster{
				Versions: []StatusClusterVersion{
					{
						Date:   time.Unix(10, 0),
						Semver: "1.0.0",
					},
					{
						Date:   time.Unix(20, 0),
						Semver: "2.0.0",
					},
				},
			},
			ExpectedSemver: "2.0.0",
		},
		{
			Name: "case 3",
			StatusCluster: StatusCluster{
				Versions: []StatusClusterVersion{
					{
						Date:   time.Unix(10, 0),
						Semver: "1.0.0",
					},
					{
						Date:   time.Unix(20, 0),
						Semver: "2.0.0",
					},
					{
						Date:   time.Unix(30, 0),
						Semver: "3.0.0",
					},
				},
			},
			ExpectedSemver: "3.0.0",
		},
		{
			Name: "case 4",
			StatusCluster: StatusCluster{
				Versions: []StatusClusterVersion{
					{
						Date:   time.Unix(20, 0),
						Semver: "2.0.0",
					},
					{
						Date:   time.Unix(30, 0),
						Semver: "3.0.0",
					},
					{
						Date:   time.Unix(10, 0),
						Semver: "1.0.0",
					},
				},
			},
			ExpectedSemver: "3.0.0",
		},
	}

	for _, tc := range testCases {
		semver := tc.StatusCluster.LatestVersion()

		if semver != tc.ExpectedSemver {
			t.Fatalf("expected %#v got %#v", tc.ExpectedSemver, semver)
		}
	}
}
