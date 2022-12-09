package favour

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFloat64(t *testing.T) {
	tCases := map[string]struct {
		v        float64
		expected Favour
	}{
		"should parse loved correctly": {
			v:        Loved.Float64(),
			expected: Loved,
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out := ParseFloat64(tc.v)
			assert.Equal(tt, out, tc.expected)
		})
	}
}

func TestGenerateFavour(t *testing.T) {
	tCases := map[string]struct {
		positivity float64
		dev        float64
		expected   Favour
	}{
		"should work for positive": {
			positivity: 1,
			dev:        2,
			expected:   Loved,
		},
		"should work for negative": {
			positivity: -1,
			dev:        2,
			expected:   Damned,
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out, err := GenerateFavour(tc.positivity, tc.dev)
			require.NoError(tt, err)
			assert.Equal(tt, out.Positive(), tc.expected.Positive())
			assert.Equal(tt, out.Zero(), tc.expected.Zero())
			assert.Equal(tt, out.Negative(), tc.expected.Negative())
		})
	}
}
