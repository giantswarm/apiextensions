package v1alpha1

import (
	"encoding/json"
	"net/url"
	"time"

	"github.com/giantswarm/microerror"
)

// DeepCopyTime implements the deep copy logic for time.Time which the k8s
// codegen is not able to generate out of the box.
type DeepCopyTime struct {
	time.Time
}

func (in *DeepCopyTime) DeepCopyInto(out *DeepCopyTime) {
	*out = *in
}

// DeepCopyURL implements the deep copy logic for url.URL which the k8s codegen
// is not able to generate out of the box.
type DeepCopyURL struct {
	*url.URL
}

func (in *DeepCopyURL) DeepCopyInto(out *DeepCopyURL) {
	*out = *in
}

func (in *DeepCopyURL) MarshalJSON() ([]byte, error) {
	return []byte(in.String()), nil
}

func (in *DeepCopyURL) MarshalYAML() (interface{}, error) {
	return in.String(), nil
}

func (in *DeepCopyURL) UnmarshalJSON(b []byte) error {
	var s string
	err := json.Unmarshal(b, &s)
	if err != nil {
		return microerror.Mask(err)
	}

	u, err := url.Parse(s)
	if err != nil {
		return microerror.Mask(err)
	}

	*in = DeepCopyURL{u}

	return nil
}

func (in *DeepCopyURL) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var s string
	err := unmarshal(&s)
	if err != nil {
		return microerror.Mask(err)
	}

	u, err := url.Parse(s)
	if err != nil {
		return microerror.Mask(err)
	}

	*in = DeepCopyURL{u}

	return nil
}
