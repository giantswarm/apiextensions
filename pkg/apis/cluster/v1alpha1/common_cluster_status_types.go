package v1alpha1

// StatusCluster is shared type to contain provider independent cluster status
// information.
type StatusCluster struct {
	Conditions []StatusClusterCondition `json:"conditions" yaml:"conditions"`
	ID         string                   `json:"id" yaml:"id"`
	Versions   []StatusClusterVersion   `json:"versions" yaml:"versions"`
}

type StatusClusterCondition struct {
	LastTransitionTime DeepCopyTime `json:"lastTransitionTime" yaml:"lastTransitionTime"`
	Type               string       `json:"type" yaml:"type"`
}

type StatusClusterVersion struct {
	LastTransitionTime DeepCopyTime `json:"lastTransitionTime" yaml:"lastTransitionTime"`
	Version            string       `json:"version" yaml:"version"`
}
