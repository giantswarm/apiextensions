package v1alpha2

import (
	"reflect"
	"strconv"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_Provider_Status_LatestVersion(t *testing.T) {
	testCases := []struct {
		Name                       string
		CommonClusterStatusCluster CommonClusterStatusCluster
		ExpectedVersion            string
	}{
		{
			Name: "case 0",
			CommonClusterStatusCluster: CommonClusterStatusCluster{
				Versions: []CommonClusterStatusClusterVersion{},
			},
			ExpectedVersion: "",
		},
		{
			Name: "case 1",
			CommonClusterStatusCluster: CommonClusterStatusCluster{
				Versions: []CommonClusterStatusClusterVersion{
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
			CommonClusterStatusCluster: CommonClusterStatusCluster{
				Versions: []CommonClusterStatusClusterVersion{
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
			CommonClusterStatusCluster: CommonClusterStatusCluster{
				Versions: []CommonClusterStatusClusterVersion{
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
			CommonClusterStatusCluster: CommonClusterStatusCluster{
				Versions: []CommonClusterStatusClusterVersion{
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
			CommonClusterStatusCluster: CommonClusterStatusCluster{
				Versions: []CommonClusterStatusClusterVersion{
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
		semver := tc.CommonClusterStatusCluster.LatestVersion()

		if semver != tc.ExpectedVersion {
			t.Fatalf("expected %#v got %#v", tc.ExpectedVersion, semver)
		}
	}
}

func Test_Provider_Status_withCondition(t *testing.T) {
	testCases := []struct {
		name               string
		conditions         []CommonClusterStatusClusterCondition
		condition          CommonClusterStatusClusterCondition
		limit              int
		expectedConditions []CommonClusterStatusClusterCondition
	}{
		{
			name:       "case 0: the creation of the tenant cluster starts",
			conditions: []CommonClusterStatusClusterCondition{},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
				Condition:          ClusterStatusConditionCreating,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
		{
			name: "case 1: the creation of the tenant cluster ends",
			conditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
				Condition:          ClusterStatusConditionCreated,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
		{
			name: "case 2: the first update of the tenant cluster starts",
			conditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
				Condition:          ClusterStatusConditionUpdating,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
		{
			name: "case 3: the first update of the tenant cluster ends",
			conditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
				Condition:          ClusterStatusConditionUpdated,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
		{
			name: "case 4: the second update of the tenant cluster starts",
			conditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(60, 0)},
				Condition:          ClusterStatusConditionUpdating,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(60, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
		{
			name: "case 5: the second update of the tenant cluster ends",
			conditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(60, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(70, 0)},
				Condition:          ClusterStatusConditionUpdated,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(70, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(60, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
		{
			name: "case 6: the third update of the tenant cluster starts",
			conditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(70, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(60, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(80, 0)},
				Condition:          ClusterStatusConditionUpdating,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(80, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(70, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(60, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
		{
			name: "case 7: the third update of the tenant cluster ends",
			conditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(80, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(70, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(60, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(90, 0)},
				Condition:          ClusterStatusConditionUpdated,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(90, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(80, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(70, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(60, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
		{
			name: "case 8: the second update of the tenant cluster starts before the first ended",
			conditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
				Condition:          ClusterStatusConditionUpdating,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				// This Updated condition is added automatically when adding the
				// Updating condition twice. That way any failure tracking the
				// conditions properly would be fixed on reconciliation to keep the
				// integrity of the condition list. Only unfortunate effect is that the
				// tracked timestamp for the automatically added condition is off and
				// does not reflect the truth.
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, -1)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
		{
			name: "case 9: the fourth update of the tenant cluster starts before the thrird ended",
			conditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(80, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(70, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(60, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(50, 0)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(40, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
			condition: CommonClusterStatusClusterCondition{
				LastTransitionTime: DeepCopyTime{time.Unix(90, 0)},
				Condition:          ClusterStatusConditionUpdating,
			},
			limit: 2,
			expectedConditions: []CommonClusterStatusClusterCondition{
				{
					LastTransitionTime: DeepCopyTime{time.Unix(90, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				// This Updated condition is added automatically when adding the
				// Updating condition twice. That way any failure tracking the
				// conditions properly would be fixed on reconciliation to keep the
				// integrity of the condition list. Only unfortunate effect is that the
				// tracked timestamp for the automatically added condition is off and
				// does not reflect the truth.
				{
					LastTransitionTime: DeepCopyTime{time.Unix(90, -1)},
					Condition:          ClusterStatusConditionUpdated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(80, 0)},
					Condition:          ClusterStatusConditionUpdating,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(30, 0)},
					Condition:          ClusterStatusConditionCreated,
				},
				{
					LastTransitionTime: DeepCopyTime{time.Unix(20, 0)},
					Condition:          ClusterStatusConditionCreating,
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			conditions := withCondition(tc.conditions, tc.condition, tc.limit)

			if !reflect.DeepEqual(conditions, tc.expectedConditions) {
				t.Fatalf("\n\n%s\n", cmp.Diff(conditions, tc.expectedConditions))
			}
		})
	}
}

func Test_Provider_Status_withVersion(t *testing.T) {
	testCases := []struct {
		Name             string
		Versions         []CommonClusterStatusClusterVersion
		Version          CommonClusterStatusClusterVersion
		Limit            int
		ExpectedVersions []CommonClusterStatusClusterVersion
	}{
		{
			Name:     "case 0: list with zero items results in a list with one item",
			Versions: []CommonClusterStatusClusterVersion{},
			Version: CommonClusterStatusClusterVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
				Version:            "1.0.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
		},
		{
			Name: "case 1: list with one item results in a list with two items",
			Versions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
			Version: CommonClusterStatusClusterVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
				Version:            "1.1.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
		},
		{
			Name: "case 2: list with two items results in a list with three items",
			Versions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
			},
			Version: CommonClusterStatusClusterVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
				Version:            "1.5.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
					Version:            "1.5.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
		},
		{
			Name: "case 3: list with three items results in a list with three items due to limit",
			Versions: []CommonClusterStatusClusterVersion{
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
			Version: CommonClusterStatusClusterVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
				Version:            "3.0.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
					Version:            "3.0.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(30, 0)},
					Version:            "1.5.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
					Version:            "1.1.0",
				},
			},
		},
		{
			Name: "case 4: list with five items results in a list with three items due to limit",
			Versions: []CommonClusterStatusClusterVersion{
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
			Version: CommonClusterStatusClusterVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(60, 0)},
				Version:            "3.3.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(60, 0)},
					Version:            "3.3.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(50, 0)},
					Version:            "3.2.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
					Version:            "3.0.0",
				},
			},
		},
		{
			Name: "case 5: same as 4 but checking items are ordered by date before cutting off",
			Versions: []CommonClusterStatusClusterVersion{
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
			Version: CommonClusterStatusClusterVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(60, 0)},
				Version:            "3.3.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(60, 0)},
					Version:            "3.3.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(50, 0)},
					Version:            "3.2.0",
				},
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(40, 0)},
					Version:            "3.0.0",
				},
			},
		},
		{
			Name: "case 6: list with one item results in a list with one item in case the version already exists",
			Versions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
			Version: CommonClusterStatusClusterVersion{
				LastTransitionTime: DeepCopyTime{Time: time.Unix(20, 0)},
				Version:            "1.0.0",
			},
			Limit: 3,
			ExpectedVersions: []CommonClusterStatusClusterVersion{
				{
					LastTransitionTime: DeepCopyTime{Time: time.Unix(10, 0)},
					Version:            "1.0.0",
				},
			},
		},
	}

	for i, tc := range testCases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			versions := withVersion(tc.Versions, tc.Version, tc.Limit)

			if !reflect.DeepEqual(versions, tc.ExpectedVersions) {
				t.Fatalf("\n\n%s\n", cmp.Diff(versions, tc.ExpectedVersions))
			}
		})
	}
}
