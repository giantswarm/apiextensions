package v1alpha1

import (
	"fmt"
	"strings"
)

func (a ReleaseSpecAuthority) ApprChannelName() string {
	return fmt.Sprintf("%s-%s", a.Name, strings.Replace(a.Version, ".", "-", -1))
}

func (a ReleaseSpecAuthority) HelmChartName() string {
	return fmt.Sprintf("%s-chart", a.Name)
}
