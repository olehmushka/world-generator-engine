package culture

import (
	"strings"
	"testing"

	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/world-generator-engine/tools"
)

func TestRandomBase(t *testing.T) {
	result, err := RandomBase(mockBases)
	if err != nil {
		t.Fatalf("unexpected err (err=%+v)", err)
	}
	if result == "" {
		t.Errorf("result should not be empty string")
		return
	}
	if !sliceTools.Includes(mockBases, result) {
		t.Errorf("result base should be picked from input slice")
	}
}

func TestExtractBases(t *testing.T) {
	bases := ExtractBases(mockCultures)
	if len(bases) != len(mockCultures) {
		t.Errorf("unexpected extracted base length (expected=%d, actual=%d)", len(mockCultures), len(bases))
	}
	for _, base := range bases {
		if !strings.HasSuffix(base, RequiredBaseSlugSuffix) {
			t.Errorf("unexpected base slug suffix (slug=%s)", base)
		}
	}
}

func TestSelectBaseByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := []struct {
		name           string
		input          []string
		expectedOutput string
	}{
		{
			name:           "should works for 10 the same langs",
			input:          []string{"europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
		{
			name:           "should works for 9 the same langs and 1 another",
			input:          []string{"mediterranean_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
		{
			name:           "should works for 8 the same langs and 2 another",
			input:          []string{"mediterranean_base", "mediterranean_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
		{
			name:           "should works for 7 the same langs and 3 another",
			input:          []string{"mediterranean_base", "mediterranean_base", "mediterranean_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
		{
			name:           "should works for 6 the same langs and 4 another",
			input:          []string{"mediterranean_base", "mediterranean_base", "mediterranean_base", "mediterranean_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectBaseByMostRecent(tc.input)
				if err != nil {
					t.Fatalf("unexpected error (err=%v)", err)
					return
				}
				if count, ok := m[out]; ok {
					m[out] = count + 1
				} else {
					m[out] = 1
				}
			}
			if slug := tools.GetKeyWithGreatestValue(m); slug != tc.expectedOutput {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedOutput, slug)
			}
		})
	}
}

func TestLoadAllBases(t *testing.T) {
	var count int

	for chunk := range LoadAllBases() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of bases")
		}

		for _, base := range chunk.Value {
			if !strings.HasSuffix(base, RequiredBaseSlugSuffix) {
				t.Errorf("unexpected base slug suffix (slug=%s)", base)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 28; count != expecCount {
		t.Errorf("unexpected count of bases (expected=%d, actual=%d)", expecCount, count)
	}
}
