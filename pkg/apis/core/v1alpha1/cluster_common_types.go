package v1alpha1

type CommonGuestConfig struct {
	API            CommonGuestConfigAPI             `json:"api" yaml:"api"`
	ID             string                           `json:"id" yaml:"id"`
	Name           string                           `json:"name,omitempty" yaml:"name,omitempty"`
	Owner          string                           `json:"owner,omitempty" yaml:"owner,omitempty"`
	VersionBundles []CommonGuestConfigVersionBundle `json:"versionBundles,omitempty" yaml:"versionBundles,omitempty"`
}

type CommonGuestConfigAPI struct {
	Endpoint string `json:"endpoint" yaml:"endpoint"`
}

type CommonGuestConfigVersionBundle struct {
	Name    string `json:"name" yaml:"name"`
	Version string `json:"version" yaml:"version"`
	WIP     bool   `json:"wip" yaml:"wip"`
}
