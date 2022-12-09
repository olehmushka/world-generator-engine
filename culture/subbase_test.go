package culture

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFilterSubbasesByBaseSlug(t *testing.T) {
	baseSlug := "europe_continental_base"
	result := FilterSubbasesByBaseSlug(mockSubbases, baseSlug)
	var count int
	for i, sb := range result {
		assert.Equal(t, sb.BaseSlug, baseSlug)
		count = i + 1
	}
	assert.Equal(t, 24, count)
}

func TestRandomSubbase(t *testing.T) {
	result, err := RandomSubbase(mockSubbases)
	require.NoError(t, err)
	assert.False(t, result.IsZero())
	var isSubbasesIncludeResult bool
	for _, sb := range mockSubbases {
		if sb.Slug == result.Slug {
			isSubbasesIncludeResult = true
		}
	}
	assert.True(t, isSubbasesIncludeResult)
}

func TestExtractSubbases(t *testing.T) {
	subbases := ExtractSubbases(mockCultures)
	assert.Equal(t, len(subbases), len(mockCultures))
	for _, subbase := range subbases {
		assert.Contains(t, subbase.Slug, RequiredSubbaseSlugSuffix)
	}
}

func TestSelectSubbaseByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := map[string]struct {
		input          []Subbase
		expectedOutput Subbase
	}{
		"should works for 10 the same subbases": {
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
		"should works for 9 the same subbases and 1 another": {
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
		"should works for 8 the same subbases and 2 other": {
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
		"should works for 7 the same subbases and 3 other": {
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
		"should works for 6 the same subbases and 4 other": {
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

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectSubbaseByMostRecent(tc.input)
				require.NoError(t, err)
				if count, ok := m[out.Slug]; ok {
					m[out.Slug] = count + 1
				} else {
					m[out.Slug] = 1
				}
			}
			assert.Equal(tt, tools.GetKeyWithGreatestValue(m), tc.expectedOutput.Slug)
		})
	}
}

func TestLoadAllSubbases(t *testing.T) {
	var count int

	for chunk := range LoadAllSubbases() {
		require.NoError(t, chunk.Err)
		assert.NotEmpty(t, chunk.Value)
		for _, c := range chunk.Value {
			assert.Contains(t, c.Slug, RequiredSubbaseSlugSuffix)
		}

		count += len(chunk.Value)
	}
	assert.Equal(t, 70, count)
}
