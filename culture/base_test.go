package culture

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomBase(t *testing.T) {
	result, err := RandomBase(mockBases)
	require.NoError(t, err)
	assert.NotZero(t, result)
	assert.Contains(t, mockBases, result)
}

func TestExtractBases(t *testing.T) {
	bases := ExtractBases(mockCultures)
	assert.Equal(t, len(bases), len(mockCultures))
	for _, base := range bases {
		assert.Contains(t, base, RequiredBaseSlugSuffix)
	}
}

func TestSelectBaseByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := map[string]struct {
		input          []string
		expectedOutput string
	}{
		"should works for 10 the same langs": {
			input:          []string{"europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
		"should works for 9 the same langs and 1 another": {
			input:          []string{"mediterranean_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
		"should works for 8 the same langs and 2 another": {
			input:          []string{"mediterranean_base", "mediterranean_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
		"should works for 7 the same langs and 3 another": {
			input:          []string{"mediterranean_base", "mediterranean_base", "mediterranean_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
		"should works for 6 the same langs and 4 another": {
			input:          []string{"mediterranean_base", "mediterranean_base", "mediterranean_base", "mediterranean_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base", "europe_continental_base"},
			expectedOutput: "europe_continental_base",
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectBaseByMostRecent(tc.input)
				require.NoError(t, err)
				if count, ok := m[out]; ok {
					m[out] = count + 1
				} else {
					m[out] = 1
				}
			}
			assert.Equal(tt, tools.GetKeyWithGreatestValue(m), tc.expectedOutput)
		})
	}
}

func TestLoadAllBases(t *testing.T) {
	var count int

	for chunk := range LoadAllBases() {
		require.NoError(t, chunk.Err)
		assert.Greater(t, len(chunk.Value), 0)

		for _, base := range chunk.Value {
			assert.Contains(t, base, RequiredBaseSlugSuffix)
		}

		count += len(chunk.Value)
	}

	assert.Equal(t, 28, count)
}
