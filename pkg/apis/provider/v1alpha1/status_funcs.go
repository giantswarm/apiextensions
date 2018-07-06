package v1alpha1

import (
	"sort"
	"time"
)

func (s StatusCluster) HasCreatedCondition() bool {
	return hasCondition(s.Conditions, StatusClusterStatusTrue, StatusClusterTypeCreated)
}

func (s StatusCluster) HasCreatingCondition() bool {
	return hasCondition(s.Conditions, StatusClusterStatusTrue, StatusClusterTypeCreating)
}

func (s StatusCluster) HasUpdatedCondition() bool {
	return hasCondition(s.Conditions, StatusClusterStatusTrue, StatusClusterTypeUpdated)
}

func (s StatusCluster) HasUpdatingCondition() bool {
	return hasCondition(s.Conditions, StatusClusterStatusTrue, StatusClusterTypeUpdating)
}

func (s StatusCluster) HasVersion(semver string) bool {
	return hasVersion(s.Versions, semver)
}

func (s StatusCluster) LatestVersion() string {
	if len(s.Versions) == 0 {
		return ""
	}

	latest := s.Versions[0]

	for _, v := range s.Versions {
		if latest.Date.Before(v.Date) {
			latest = v
		}
	}

	return latest.Semver
}

func (s StatusCluster) WithCreatedCondition() []StatusClusterCondition {
	return withCondition(s.Conditions, StatusClusterTypeCreating, StatusClusterTypeCreated, StatusClusterStatusTrue)
}

func (s StatusCluster) WithCreatingCondition() []StatusClusterCondition {
	return withCondition(s.Conditions, StatusClusterTypeCreated, StatusClusterTypeCreating, StatusClusterStatusTrue)
}

func (s StatusCluster) WithNewVersion(version string) []StatusClusterVersion {
	return withVersion(s.Versions, time.Now(), version, ClusterVersionLimit)
}

func (s StatusCluster) WithUpdatedCondition() []StatusClusterCondition {
	return withCondition(s.Conditions, StatusClusterTypeUpdating, StatusClusterTypeUpdated, StatusClusterStatusTrue)
}

func (s StatusCluster) WithUpdatingCondition() []StatusClusterCondition {
	return withCondition(s.Conditions, StatusClusterTypeUpdated, StatusClusterTypeUpdating, StatusClusterStatusTrue)
}

func hasCondition(conditions []StatusClusterCondition, search string, status string) bool {
	for _, c := range conditions {
		if c.Status == search && c.Type == status {
			return true
		}
	}

	return false
}

func hasVersion(versions []StatusClusterVersion, search string) bool {
	for _, v := range versions {
		if v.Semver == search {
			return true
		}
	}

	return false
}

func withCondition(conditions []StatusClusterCondition, search string, replace string, status string) []StatusClusterCondition {
	newConditions := []StatusClusterCondition{
		{
			Status: status,
			Type:   replace,
		},
	}

	for _, c := range conditions {
		if c.Type == search {
			continue
		}

		newConditions = append(newConditions, c)
	}

	return newConditions
}

func withVersion(versions []StatusClusterVersion, date time.Time, version string, limit int) []StatusClusterVersion {
	var newVersions []StatusClusterVersion

	start := 0
	if len(versions) >= limit {
		start = len(versions) - limit + 1
	}

	sort.Sort(sortClusterStatusVersionsByDate(versions))

	for i := start; i < len(versions); i++ {
		newVersions = append(newVersions, versions[i])
	}

	newVersions = append(newVersions, StatusClusterVersion{
		Date:   date,
		Semver: version,
	})

	return newVersions
}
