package v1alpha1

type sortClusterStatusVersionsByLastTransitionTime []StatusClusterVersion

func (s sortClusterStatusVersionsByLastTransitionTime) Len() int      { return len(s) }
func (s sortClusterStatusVersionsByLastTransitionTime) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s sortClusterStatusVersionsByLastTransitionTime) Less(i, j int) bool {
	return s[i].LastTransitionTime.UnixNano() < s[j].LastTransitionTime.UnixNano()
}
