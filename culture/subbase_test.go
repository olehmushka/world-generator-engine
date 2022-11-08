package culture

import (
	"strings"
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
)

func TestFilterSubbasesByBaseSlug(t *testing.T) {
	baseSlug := "europe_continental_base"
	result := FilterSubbasesByBaseSlug(mockSubbases, baseSlug)
	var count int
	for i, sb := range result {
		if sb.BaseSlug != baseSlug {
			t.Errorf("unexpected base_slug (expected=%s, actual=%s)", baseSlug, sb.BaseSlug)
		}
		count = i + 1
	}

	if expecCount := 24; count != expecCount {
		t.Errorf("unexpected count of filtered subbases (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestRandomSubbase(t *testing.T) {
	result, err := RandomSubbase(mockSubbases)
	if err != nil {
		t.Fatalf("unexpected err (err=%+v)", err)
	}
	if result.IsZero() {
		t.Errorf("result should not be empty string")
		return
	}
	var isSubbasesIncludeResult bool
	for _, sb := range mockSubbases {
		if sb.Slug == result.Slug {
			isSubbasesIncludeResult = true
		}
	}
	if !isSubbasesIncludeResult {
		t.Errorf("result subbase should be picked from input slice")
	}
}

func TestExtractSubbases(t *testing.T) {
	subbases := ExtractSubbases(mockCultures)
	if len(subbases) != len(mockCultures) {
		t.Errorf("unexpected extracted subbase length (expected=%d, actual=%d)", len(mockCultures), len(subbases))
	}
	for _, subbase := range subbases {
		if !strings.HasSuffix(subbase.Slug, RequiredSubbaseSlugSuffix) {
			t.Errorf("unexpected subbase slug suffix (slug=%s)", subbase.Slug)
		}
	}
}

func TestSelectSubbaseByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := []struct {
		name           string
		input          []Subbase
		expectedOutput Subbase
	}{
		{
			name: "should works for 10 the same subbases",
			input: []Subbase{
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
			},
			expectedOutput: Subbase{Slug: "ancient_levantine_subbase"},
		},
		{
			name: "should works for 9 the same subbases and 1 another",
			input: []Subbase{
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
			},
			expectedOutput: Subbase{Slug: "ancient_levantine_subbase"},
		},
		{
			name: "should works for 8 the same subbases and 2 other",
			input: []Subbase{
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
			},
			expectedOutput: Subbase{Slug: "ancient_levantine_subbase"},
		},
		{
			name: "should works for 7 the same subbases and 3 other",
			input: []Subbase{
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
			},
			expectedOutput: Subbase{Slug: "ancient_levantine_subbase"},
		},
		{
			name: "should works for 6 the same subbases and 4 other",
			input: []Subbase{
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_egyptian_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
				{Slug: "ancient_levantine_subbase"},
			},
			expectedOutput: Subbase{Slug: "ancient_levantine_subbase"},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectSubbaseByMostRecent(tc.input)
				if err != nil {
					t.Fatalf("unexpected error (err=%v)", err)
					return
				}
				if count, ok := m[out.Slug]; ok {
					m[out.Slug] = count + 1
				} else {
					m[out.Slug] = 1
				}
			}
			if slug := tools.GetKeyWithGreatestValue(m); slug != tc.expectedOutput.Slug {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedOutput.Slug, slug)
			}
		})
	}
}

func TestLoadAllSubbases(t *testing.T) {
	var count int

	for chunk := range LoadAllSubbases() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of subbases")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredSubbaseSlugSuffix) {
				t.Errorf("unexpected subbase slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 70; count != expecCount {
		t.Errorf("unexpected count of subbases (expected=%d, actual=%d)", expecCount, count)
	}
}
