package acceptance

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetAcceptanceByProbability(t *testing.T) {
	iterationsNumber := 100

	tCases := map[string]struct {
		acceptedProb, shunnedProb, criminalProb float64
		expectedOutput                          string
	}{
		"should returns only_men for greater strong probability": {
			acceptedProb:   0.35,
			shunnedProb:    0.2,
			criminalProb:   0.2,
			expectedOutput: Accepted.String(),
		},
		"should returns men_and_women for greater strong probability": {
			acceptedProb:   0.2,
			shunnedProb:    0.35,
			criminalProb:   0.2,
			expectedOutput: Shunned.String(),
		},
		"should returns only_women for greater strong probability": {
			acceptedProb:   0.2,
			shunnedProb:    0.2,
			criminalProb:   0.35,
			expectedOutput: Criminal.String(),
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := GetAcceptanceByProb(tc.acceptedProb, tc.shunnedProb, tc.criminalProb)
				require.NoError(t, err)
				if count, ok := m[out.String()]; ok {
					m[out.String()] = count + 1
				} else {
					m[out.String()] = 1
				}
			}
			assert.Equal(tt, tools.GetKeyWithGreatestValue(m), tc.expectedOutput)
		})
	}
}

func TestSelectAcceptanceByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := map[string]struct {
		input          []Acceptance
		expectedOutput Acceptance
	}{
		"should works for 10 the same dominated sexes": {
			input: []Acceptance{
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
			},
			expectedOutput: Accepted,
		},
		"should works for 9 the same dominated sexes and 1 another": {
			input: []Acceptance{
				Shunned,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
			},
			expectedOutput: Accepted,
		},
		"should works for 8 the same dominated sexes and 2 other": {
			input: []Acceptance{
				Shunned,
				Shunned,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
			},
			expectedOutput: Accepted,
		},
		"should works for 7 the same dominated sexes and 3 other": {
			input: []Acceptance{
				Criminal,
				Criminal,
				Criminal,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
			},
			expectedOutput: Accepted,
		},
		"should works for 6 the same dominated sexes and 4 other": {
			input: []Acceptance{
				Criminal,
				Criminal,
				Criminal,
				Criminal,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
				Accepted,
			},
			expectedOutput: Accepted,
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectAcceptanceByMostRecent(tc.input)
				require.NoError(t, err)
				if count, ok := m[out.String()]; ok {
					m[out.String()] = count + 1
				} else {
					m[out.String()] = 1
				}
			}
			assert.Equal(tt, tools.GetKeyWithGreatestValue(m), tc.expectedOutput.String())
		})
	}
}
