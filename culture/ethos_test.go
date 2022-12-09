package culture

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomEthos(t *testing.T) {
	result, err := RandomEthos(mockEthoses)
	require.NoError(t, err)
	assert.False(t, result.IsZero())
	var isEthosesIncludeResult bool
	for _, e := range mockEthoses {
		if e.Slug == result.Slug {
			isEthosesIncludeResult = true
		}
	}
	assert.True(t, isEthosesIncludeResult)
}

func TestExtractEthoses(t *testing.T) {
	ethoses := ExtractEthoses(mockCultures)
	assert.Equal(t, len(ethoses), len(mockCultures))
	for _, ethos := range ethoses {
		assert.Contains(t, ethos.Slug, RequiredEthosSlugSuffix)
	}
}

func TestSelectEthosByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := map[string]struct {
		input          []Ethos
		expectedOutput Ethos
	}{
		"should works for 10 the same ethoses": {
			input: []Ethos{
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
			},
			expectedOutput: Ethos{Slug: "bellicose_ethos"},
		},
		"should works for 9 the same ethoses and 1 another": {
			input: []Ethos{
				{Slug: "bureaucratic_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
			},
			expectedOutput: Ethos{Slug: "bellicose_ethos"},
		},
		"should works for 8 the same ethoses and 2 other": {
			input: []Ethos{
				{Slug: "bureaucratic_ethos"},
				{Slug: "bureaucratic_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
			},
			expectedOutput: Ethos{Slug: "bellicose_ethos"},
		},
		"should works for 7 the same ethoses and 3 other": {
			input: []Ethos{
				{Slug: "bureaucratic_ethos"},
				{Slug: "bureaucratic_ethos"},
				{Slug: "bureaucratic_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
			},
			expectedOutput: Ethos{Slug: "bellicose_ethos"},
		},
		"should works for 6 the same ethoses and 4 other": {
			input: []Ethos{
				{Slug: "bureaucratic_ethos"},
				{Slug: "bureaucratic_ethos"},
				{Slug: "bureaucratic_ethos"},
				{Slug: "bureaucratic_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
			},
			expectedOutput: Ethos{Slug: "bellicose_ethos"},
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectEthosByMostRecent(tc.input)
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

func TestUniqueEthoses(t *testing.T) {
	tCases := map[string]struct {
		input          []Ethos
		expectedOutput []Ethos
	}{
		"should work for shuffeled": {
			input: []Ethos{
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bellicose_ethos"},
				{Slug: "bureaucratic_ethos"},
				{Slug: "communal_ethos"},
				{Slug: "courtly_ethos"},
				{Slug: "courtly_ethos"},
				{Slug: "courtly_ethos"},
				{Slug: "egalitarian_ethos"},
				{Slug: "spiritual_ethos"},
				{Slug: "stoic_ethos"},
				{Slug: "stoic_ethos"},
			},
			expectedOutput: []Ethos{
				{Slug: "bellicose_ethos"},
				{Slug: "bureaucratic_ethos"},
				{Slug: "communal_ethos"},
				{Slug: "courtly_ethos"},
				{Slug: "egalitarian_ethos"},
				{Slug: "spiritual_ethos"},
				{Slug: "stoic_ethos"},
			},
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			assert.Equal(tt, len(UniqueEthoses(tc.input)), len(tc.expectedOutput))
		})
	}
}

func TestLoadAllEthoses(t *testing.T) {
	var count int

	for chunk := range LoadAllEthoses() {
		require.NoError(t, chunk.Err)
		assert.NotEmpty(t, chunk.Value)
		for _, c := range chunk.Value {
			assert.Contains(t, c.Slug, RequiredEthosSlugSuffix)
		}
		count += len(chunk.Value)
	}

	assert.Equal(t, 7, count)
}

func TestSearchEthos(t *testing.T) {
	slug := "communal_ethos"
	result, err := SearchEthos(slug)
	require.NoError(t, err)
	assert.False(t, result.IsZero())
	assert.Equal(t, result.Slug, slug)
}
