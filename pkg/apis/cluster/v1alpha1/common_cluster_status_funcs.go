package v1alpha1

import (
	"sort"
	"time"
)

func (s CommonClusterStatus) GetCreatedCondition() CommonClusterStatusCondition {
	return getCondition(s.Conditions, ClusterStatusConditionCreated)
}

func (s CommonClusterStatus) GetCreatingCondition() CommonClusterStatusCondition {
	return getCondition(s.Conditions, ClusterStatusConditionCreating)
}

func (s CommonClusterStatus) GetDeletedCondition() CommonClusterStatusCondition {
	return getCondition(s.Conditions, ClusterStatusConditionDeleted)
}

func (s CommonClusterStatus) GetDeletingCondition() CommonClusterStatusCondition {
	return getCondition(s.Conditions, ClusterStatusConditionDeleting)
}

func (s CommonClusterStatus) GetUpdatedCondition() CommonClusterStatusCondition {
	return getCondition(s.Conditions, ClusterStatusConditionUpdated)
}

func (s CommonClusterStatus) GetUpdatingCondition() CommonClusterStatusCondition {
	return getCondition(s.Conditions, ClusterStatusConditionUpdating)
}

func (s CommonClusterStatus) HasCreatedCondition() bool {
	return hasCondition(s.Conditions, ClusterStatusConditionCreated)
}

func (s CommonClusterStatus) HasCreatingCondition() bool {
	return hasCondition(s.Conditions, ClusterStatusConditionCreating)
}

func (s CommonClusterStatus) HasDeletedCondition() bool {
	return hasCondition(s.Conditions, ClusterStatusConditionDeleted)
}

func (s CommonClusterStatus) HasDeletingCondition() bool {
	return hasCondition(s.Conditions, ClusterStatusConditionDeleting)
}

func (s CommonClusterStatus) HasUpdatedCondition() bool {
	return hasCondition(s.Conditions, ClusterStatusConditionUpdated)
}

func (s CommonClusterStatus) HasUpdatingCondition() bool {
	return hasCondition(s.Conditions, ClusterStatusConditionUpdating)
}

func (s CommonClusterStatus) HasVersion(semver string) bool {
	return hasVersion(s.Versions, semver)
}

func (s CommonClusterStatus) LatestVersion() string {
	if len(s.Versions) == 0 {
		return ""
	}

	latest := s.Versions[0]

	for _, v := range s.Versions {
		if latest.LastTransitionTime.Time.Before(v.LastTransitionTime.Time) {
			latest = v
		}
	}

	return latest.Version
}

func (s CommonClusterStatus) WithCreatedCondition() []CommonClusterStatusCondition {
	newCondition := CommonClusterStatusCondition{
		LastTransitionTime: DeepCopyTime{time.Now()},
		Condition:          ClusterStatusConditionCreated,
	}

	return withCondition(s.Conditions, newCondition, ClusterConditionLimit)
}

func (s CommonClusterStatus) WithCreatingCondition() []CommonClusterStatusCondition {
	newCondition := CommonClusterStatusCondition{
		LastTransitionTime: DeepCopyTime{time.Now()},
		Condition:          ClusterStatusConditionCreating,
	}

	return withCondition(s.Conditions, newCondition, ClusterConditionLimit)
}

func (s CommonClusterStatus) WithDeletedCondition() []CommonClusterStatusCondition {
	newCondition := CommonClusterStatusCondition{
		LastTransitionTime: DeepCopyTime{time.Now()},
		Condition:          ClusterStatusConditionDeleted,
	}

	return withCondition(s.Conditions, newCondition, ClusterConditionLimit)
}

func (s CommonClusterStatus) WithDeletingCondition() []CommonClusterStatusCondition {
	newCondition := CommonClusterStatusCondition{
		LastTransitionTime: DeepCopyTime{time.Now()},
		Condition:          ClusterStatusConditionDeleting,
	}

	return withCondition(s.Conditions, newCondition, ClusterConditionLimit)
}

func (s CommonClusterStatus) WithNewVersion(version string) []CommonClusterStatusVersion {
	newVersion := CommonClusterStatusVersion{
		LastTransitionTime: DeepCopyTime{time.Now()},
		Version:            version,
	}

	return withVersion(s.Versions, newVersion, ClusterVersionLimit)
}

func (s CommonClusterStatus) WithUpdatedCondition() []CommonClusterStatusCondition {
	newCondition := CommonClusterStatusCondition{
		LastTransitionTime: DeepCopyTime{time.Now()},
		Condition:          ClusterStatusConditionUpdated,
	}

	return withCondition(s.Conditions, newCondition, ClusterConditionLimit)
}

func (s CommonClusterStatus) WithUpdatingCondition() []CommonClusterStatusCondition {
	newCondition := CommonClusterStatusCondition{
		LastTransitionTime: DeepCopyTime{time.Now()},
		Condition:          ClusterStatusConditionUpdating,
	}

	return withCondition(s.Conditions, newCondition, ClusterConditionLimit)
}

func getCondition(conditions []CommonClusterStatusCondition, condition string) CommonClusterStatusCondition {
	for _, c := range conditions {
		if c.Condition == condition {
			return c
		}
	}

	return CommonClusterStatusCondition{}
}

func hasCondition(conditions []CommonClusterStatusCondition, condition string) bool {
	for _, c := range conditions {
		if c.Condition == condition {
			return true
		}
	}

	return false
}

func hasVersion(versions []CommonClusterStatusVersion, search string) bool {
	for _, v := range versions {
		if v.Version == search {
			return true
		}
	}

	return false
}

func isConditionPair(a CommonClusterStatusCondition, b CommonClusterStatusCondition) bool {
	conditionPairs := [][]string{
		[]string{
			ClusterStatusConditionCreated,
			ClusterStatusConditionCreating,
		},
		[]string{
			ClusterStatusConditionDeleted,
			ClusterStatusConditionDeleting,
		},
		[]string{
			ClusterStatusConditionUpdated,
			ClusterStatusConditionUpdating,
		},
	}

	for _, p := range conditionPairs {
		if p[0] == a.Condition && p[1] == b.Condition {
			return true
		}
		if p[1] == a.Condition && p[0] == b.Condition {
			return true
		}
	}

	return false
}

func withCondition(conditions []CommonClusterStatusCondition, condition CommonClusterStatusCondition, limit int) []CommonClusterStatusCondition {
	// We create a new list which acts like a copy so the input parameters are not
	// manipulated.
	var newConditions []CommonClusterStatusCondition
	{
		for _, c := range conditions {
			newConditions = append(newConditions, c)
		}
		newConditions = append([]CommonClusterStatusCondition{condition}, newConditions...)
	}

	// The new list is sorted to have the first item being the oldest. This is to
	// have an easier grouping mechanism below. When the first item of a new pair
	// is added, it would throw of the grouping when the order would be kept as
	// given.
	sort.Sort(sortClusterStatusConditionsByDate(newConditions))

	// The conditions are grouped into their corresponding pairs of transitioning
	// states. Associated Creating/Created, Updating/Updated and Deleting/Deleted
	// conditions are put together.
	var conditionGroups [][]CommonClusterStatusCondition
	for len(newConditions) > 0 {
		var g []CommonClusterStatusCondition

		for _, c := range newConditions {
			// If the list only contains one item anymore, we process it separately
			// here and be done. Otherwhise the pruning of the list below panics due
			// to the range calculations.
			if len(newConditions) == 1 {
				g = append(g, c)
				newConditions = []CommonClusterStatusCondition{}
				break
			}

			// Put the first item from the top of the list into the group and drop
			// the grouped item from the list.
			if len(g) == 0 {
				g = append(g, c)
				newConditions = newConditions[1:len(newConditions)]
				continue
			}

			// When we find the second item of the pair we are done for this group.
			if len(g) == 1 {
				if isConditionPair(g[0], c) {
					g = append(g, c)
					newConditions = newConditions[1:len(newConditions)]
				}
				break
			}
		}

		conditionGroups = append(conditionGroups, g)
	}

	// The pairs are now grouped. When there are only three group kinds for
	// create/update/delete, conditionPairs has a length of 3. Each of the groups
	// has then as many pairs as grouped together. Below these groups are limited.
	var conditionPairs [][]CommonClusterStatusCondition
	for len(conditionGroups) > 0 {
		var p []CommonClusterStatusCondition

		for _, g := range conditionGroups {
			if len(p) == 0 {
				p = append(p, g...)
				conditionGroups = conditionGroups[1:len(conditionGroups)]
				continue
			}

			if len(g) >= 1 {
				if isConditionPair(p[0], g[0]) || isConditionPair(p[1], g[0]) {
					p = append(p, g...)
					conditionGroups = conditionGroups[1:len(conditionGroups)]
				}
			}
		}

		conditionPairs = append(conditionPairs, p)
	}

	// Here the list is finally flattened again and its pairs are limitted to the
	// input parameter.
	var limittedList []CommonClusterStatusCondition
	for _, p := range conditionPairs {
		// We cmpute the pair limit here for the total number of items. This is why
		// we multiply by 2. When the limit is 5, we want to track for instance 5
		// Updating/Updated pairs. Additionally when there is an item of a new pair
		// and the list must be capped, the additional odd of the new item has to be
		// considered when computing the limit. This results in an additional pair
		// being dropped. Test case 6 demonstrates that.
		l := (limit * 2) - (len(p) % 2)
		if len(p) < l {
			l = len(p)
		}

		limittedList = append(limittedList, p[len(p)-l:len(p)]...)
	}

	// We reverse the list order to have the item with the highest timestamp at
	// the top again.
	sort.Sort(sort.Reverse(sortClusterStatusConditionsByDate(limittedList)))

	return limittedList
}

// withVersion computes a list of version history using the given list and new
// version structure to append. withVersion also limits total amount of elements
// in the list by cutting off the tail with respect to the limit parameter.
func withVersion(versions []CommonClusterStatusVersion, version CommonClusterStatusVersion, limit int) []CommonClusterStatusVersion {
	if hasVersion(versions, version.Version) {
		return versions
	}

	var newVersions []CommonClusterStatusVersion

	start := 0
	if len(versions) >= limit {
		start = len(versions) - limit + 1
	}

	sort.Sort(sortClusterStatusVersionsByDate(versions))

	for i := start; i < len(versions); i++ {
		newVersions = append(newVersions, versions[i])
	}

	newVersions = append(newVersions, version)

	return newVersions
}
