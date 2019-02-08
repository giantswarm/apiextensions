package v1alpha1

import (
	"encoding/json"
	"testing"
)

func Test_DeepCopyDate_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name         string
		json         string
		errorMatcher func(err error) bool
	}{
		{
			name:         "case 0: valid date",
			json:         `{"testDate": "2019-02-08"}`,
			errorMatcher: nil,
		},
		{
			name:         "case 1: malformed date",
			json:         `{"testDate": "2019-02-08T12:04:00"}`,
			errorMatcher: func(err error) bool { return err != nil },
		},
		{
			name:         "case 2: null",
			json:         `{"testDate": null}`,
			errorMatcher: nil,
		},
		{
			name:         "case 3: wrong type",
			json:         `{"testDate": 5}`,
			errorMatcher: func(err error) bool { return err != nil },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			type JSONObject struct {
				TestDate DeepCopyDate `json:"testDate"`
			}

			var jsonObject JSONObject
			err := json.Unmarshal([]byte(tc.json), &jsonObject)

			switch {
			case err == nil && tc.errorMatcher == nil:
				// correct; carry on
			case err != nil && tc.errorMatcher == nil:
				t.Fatalf("error == %v, want nil", err)
			case err == nil && tc.errorMatcher != nil:
				t.Fatalf("error == nil, want non-nil")
			case !tc.errorMatcher(err):
				t.Fatalf("error == %v, want matching", err)
			}

			if tc.errorMatcher != nil {
				return
			}
		})
	}
}
