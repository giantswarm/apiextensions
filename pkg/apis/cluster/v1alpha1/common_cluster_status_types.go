package v1alpha1

// CommonStatusCluster is shared type to contain provider independent cluster status
// information.
type CommonStatusCluster struct {
	Conditions []CommonStatusClusterCondition `json:"conditions" yaml:"conditions"`
	ID         string                         `json:"id" yaml:"id"`
	Versions   []CommonStatusClusterVersion   `json:"versions" yaml:"versions"`
}

type CommonStatusClusterCondition struct {
	LastTransitionTime DeepCopyTime `json:"lastTransitionTime" yaml:"lastTransitionTime"`
	Type               string       `json:"type" yaml:"type"`
}

type CommonStatusClusterVersion struct {
	LastTransitionTime DeepCopyTime `json:"lastTransitionTime" yaml:"lastTransitionTime"`
	Version            string       `json:"version" yaml:"version"`
}
