package v1alpha1

import (
	"encoding/json"
	"reflect"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func Test_DeepCopyDate_UnmarshalJSON(t *testing.T) {
	testCases := []struct {
		name          string
		inputJSON     string
		expectedDay   int
		expectedMonth time.Month
		expectedYear  int
		errorMatcher  func(err error) bool
	}{
		{
			name:          "case 0: valid date",
			inputJSON:     `{"testDate": "2019-02-08"}`,
			expectedDay:   8,
			expectedMonth: 2,
			expectedYear:  2019,
			errorMatcher:  nil,
		},
		{
			name:         "case 1: malformed date",
			inputJSON:    `{"testDate": "2019-02-08T12:04:00"}`,
			errorMatcher: func(err error) bool { return err != nil },
		},
		{
			name:          "case 2: null",
			inputJSON:     `{"testDate": null}`,
			expectedDay:   1,
			expectedMonth: 1,
			expectedYear:  1,
			errorMatcher:  nil,
		},
		{
			name:         "case 3: wrong type",
			inputJSON:    `{"testDate": 5}`,
			errorMatcher: func(err error) bool { return err != nil },
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			type JSONObject struct {
				TestDate DeepCopyDate `json:"testDate"`
			}

			var jsonObject JSONObject
			err := json.Unmarshal([]byte(tc.inputJSON), &jsonObject)

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

			if !reflect.DeepEqual(jsonObject.TestDate.Day(), tc.expectedDay) {
				t.Errorf("\n\n%s\n", cmp.Diff(jsonObject.TestDate.Day(), tc.expectedDay))
			}
			if !reflect.DeepEqual(jsonObject.TestDate.Month(), tc.expectedMonth) {
				t.Errorf("\n\n%s\n", cmp.Diff(jsonObject.TestDate.Month(), tc.expectedMonth))
			}
			if !reflect.DeepEqual(jsonObject.TestDate.Year(), tc.expectedYear) {
				t.Errorf("\n\n%s\n", cmp.Diff(jsonObject.TestDate.Year(), tc.expectedYear))
			}
		})
	}
}
