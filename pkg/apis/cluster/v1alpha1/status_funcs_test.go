package v1alpha1

import (
	"reflect"
	"testing"
	"time"
)

func Test_Provider_Status_LatestVersion(t *testing.T) {
	testCases := []struct {
		Name                string
		CommonClusterStatus CommonClusterStatus
		ExpectedVersion     string
	}{
		{
			Name: "case 0",
			CommonClusterStatus: CommonClusterStatus{
				Versions: []CommonClusterStatusVersion{},
			},
			ExpectedVersion: "",
		},
		{
			Name: "case 1",
			CommonClusterStatus: CommonClusterStatus{
				Versions: []CommonClusterStatusVersion{
					{
						LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
						Version:            "1.0.0",
					},
				},
			},
			ExpectedVersion: "1.0.0",
		},
		{
			Name: "case 2",
			CommonClusterStatus: CommonClusterStatus{
				Versions: []CommonClusterStatusVersion{
					{
						LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
						Version:            "1.0.0",
					},
					{
						LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
						Version:            "2.0.0",
					},
				},
			},
			ExpectedVersion: "2.0.0",
		},
		{
			Name: "case 3",
			CommonClusterStatus: CommonClusterStatus{
				Versions: []CommonClusterStatusVersion{
					{
						LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
						Version:            "1.0.0",
					},
					{
						LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
						Version:            "2.0.0",
					},
					{
						LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
						Version:            "3.0.0",
					},
				},
			},
			ExpectedVersion: "3.0.0",
		},
		{
			Name: "case 4",
			CommonClusterStatus: CommonClusterStatus{
				Versions: []CommonClusterStatusVersion{
					{
						LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
						Version:            "2.0.0",
					},
					{
						LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
						Version:            "3.0.0",
					},
					{
						LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
						Version:            "1.0.0",
					},
				},
			},
			ExpectedVersion: "3.0.0",
		},
		{
			Name: "case 5",
			CommonClusterStatus: CommonClusterStatus{
				Versions: []CommonClusterStatusVersion{
					{
						LastTransitionTime: DeepCopyTime{
							time.Unix(20, 0),
						},
						Version: "2.0.0",
					},
					{
						LastTransitionTime: DeepCopyTime{
							time.Unix(30, 0),
						},
						Version: "3.0.0",
					},
					{
						LastTransitionTime: DeepCopyTime{
							time.Unix(10, 0),
						},
						Version: "1.0.0",
					},
				},
			},
			ExpectedVersion: "3.0.0",
		},
	}

	for _, tc := range testCases {
		semver := tc.CommonClusterStatus.LatestVersion()

		if semver != tc.ExpectedVersion {
			t.Fatalf("expected %#v got %#v", tc.ExpectedVersion, semver)
		}
	}
}

func Test_Provider_Status_withCondition(t *testing.T) {
	testTime := time.Unix(20, 0)

	testCases := []struct {
		Name               string
		Conditions         []CommonClusterStatusCondition
		Search             string
		Replace            string
		ExpectedConditions []CommonClusterStatusCondition
	}{
		{
			Name:       "case 0",
			Conditions: []CommonClusterStatusCondition{},
			Search:     ClusterStatusConditionCreating,
			Replace:    ClusterStatusConditionCreated,
			ExpectedConditions: []CommonClusterStatusCondition{
				{
					LastTransitionTime: DeepCopyTime{testTime},
					Condition:          ClusterStatusConditionCreated,
				},
			},
		},
		{
			Name: "case 1",
			Conditions: []CommonClusterStatusCondition{
				{
					LastTransitionTime: DeepCopyTime{testTime},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			Search:  ClusterStatusConditionCreating,
			Replace: ClusterStatusConditionCreated,
			ExpectedConditions: []CommonClusterStatusCondition{
				{
					LastTransitionTime: DeepCopyTime{testTime},
					Condition:          ClusterStatusConditionCreated,
				},
			},
		},
	}

	for _, tc := range testCases {
		conditions := withCondition(tc.Conditions, tc.Search, tc.Replace, testTime)

		if !reflect.DeepEqual(conditions, tc.ExpectedConditions) {
			t.Fatalf("%s: expected %#v got %#v", tc.Name, tc.ExpectedConditions, conditions)
		}
	}
}

func Test_Provider_Status_withVersion(t *testing.T) {
	testCases := []struct {
		Name             string
		Versions         []CommonClusterStatusVersion
		Version          CommonClusterStatusVersion
		Limit            int
		ExpectedVersions []CommonClusterStatusVersion
	}{
		{
			Name:     "case 0: list with zero items results in a list with one item",
			Versions: []CommonClusterStatusVersion{},
			Version: CommonClusterStatusVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
				Version:            "1.0.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
		},
		{
			Name: "case 1: list with one item results in a list with two items",
			Versions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
			Version: CommonClusterStatusVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
				Version:            "1.1.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
			},
		},
		{
			Name: "case 2: list with two items results in a list with three items",
			Versions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
			},
			Version: CommonClusterStatusVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
				Version:            "1.5.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
					Version:            "1.5.0",
				},
			},
		},
		{
			Name: "case 3: list with three items results in a list with three items due to limit",
			Versions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
					Version:            "1.5.0",
				},
			},
			Version: CommonClusterStatusVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
				Version:            "3.0.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
					Version:            "1.5.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
					Version:            "3.0.0",
				},
			},
		},
		{
			Name: "case 4: list with five items results in a list with three items due to limit",
			Versions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
					Version:            "1.5.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
					Version:            "3.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(50, 0)},
					Version:            "3.2.0",
				},
			},
			Version: CommonClusterStatusVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(60, 0)},
				Version:            "3.3.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
					Version:            "3.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(50, 0)},
					Version:            "3.2.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(60, 0)},
					Version:            "3.3.0",
				},
			},
		},
		{
			Name: "case 5: same as 4 but checking items are ordered by date before cutting off",
			Versions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
					Version:            "3.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(50, 0)},
					Version:            "3.2.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
					Version:            "1.5.0",
				},
			},
			Version: CommonClusterStatusVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(60, 0)},
				Version:            "3.3.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
					Version:            "3.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(50, 0)},
					Version:            "3.2.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(60, 0)},
					Version:            "3.3.0",
				},
			},
		},
		{
			Name: "case 6: list with one item results in a list with one item in case the version already exists",
			Versions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
			Version: CommonClusterStatusVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
				Version:            "1.0.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			versions := withVersion(tc.Versions, tc.Version, tc.Limit)

			if !reflect.DeepEqual(versions, tc.ExpectedVersions) {
				t.Fatalf("expected %#v got %#v", tc.ExpectedVersions, versions)
			}
		})
	}
}
