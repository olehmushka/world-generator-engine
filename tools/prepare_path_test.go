package tools

import "testing"

func TestPreparePath(t *testing.T) {
	tCases := []struct {
		name             string
		inPath, inSuffix string
		expectedOut      string
	}{
		{
			name:        "should return xxx/xxx/ for xxx/xxx & xxx suffix",
			inPath:      "xxx/xxx",
			inSuffix:    "xxx",
			expectedOut: "xxx/xxx/",
		},
		{
			name:        "should return xxx/xxx/ for xxx/xxx/ & xxx suffix",
			inPath:      "xxx/xxx/",
			inSuffix:    "xxx",
			expectedOut: "xxx/xxx/",
		},
		{
			name:        "should return xxx/xxx/ for xxx/xxx// & xxx suffix",
			inPath:      "xxx/xxx//",
			inSuffix:    "xxx",
			expectedOut: "xxx/xxx/",
		},
		{
			name:        "should return xxx/xxx/yyy/xxx/ for xxx/xxx/yyy/ & xxx suffix",
			inPath:      "xxx/xxx/yyy/",
			inSuffix:    "xxx",
			expectedOut: "xxx/xxx/yyy/xxx/",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := PreparePath(tc.inPath, tc.inSuffix); out != tc.expectedOut {
				tt.Errorf("unnexpected result (expected=%s, actual=%s)", tc.expectedOut, out)
			}
		})
	}
}

func TestRemoveLastSlash(t *testing.T) {
	tCases := []struct {
		name        string
		in          string
		expectedOut string
	}{
		{
			name:        "should work for path without last slash",
			in:          "xxx/xxx",
			expectedOut: "xxx/xxx",
		},
		{
			name:        "should work for path with one last slash",
			in:          "xxx/xxx/",
			expectedOut: "xxx/xxx",
		},
		{
			name:        "should work for path with double last slash",
			in:          "xxx/xxx//",
			expectedOut: "xxx/xxx",
		},
		{
			name:        "should work for path with triple last slash",
			in:          "xxx/xxx///",
			expectedOut: "xxx/xxx",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := removeLastSlash(tc.in); out != tc.expectedOut {
				tt.Errorf("unnexpected result (expected=%s, actual=%s)", tc.expectedOut, out)
			}
		})
	}
}
