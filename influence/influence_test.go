package influence

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRandom(t *testing.T) {
	iterationsNumber := 10000
	tCases := map[string]struct {
		expectedOutput string
	}{
		"should returns moderate_influence for most cases": {
			expectedOutput: ModerateInfluence.String(),
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := GetRandom()
				require.NoError(tt, err)
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

func TestGetInfluenceByProbability(t *testing.T) {
	iterationsNumber := 1000
	tCases := map[string]struct {
		strongProb, moderateProb, weakProb float64
		expectedOutput                     string
	}{
		"should returns strong_influence for greater strong probability": {
			strongProb:     0.35,
			moderateProb:   0.2,
			weakProb:       0.2,
			expectedOutput: StrongInfluence.String(),
		},
		"should returns moderate_influence for greater moderate probability": {
			strongProb:     0.2,
			moderateProb:   0.35,
			weakProb:       0.2,
			expectedOutput: ModerateInfluence.String(),
		},
		"should returns weak_influence for greater weak probability": {
			strongProb:     0.2,
			moderateProb:   0.2,
			weakProb:       0.35,
			expectedOutput: WeakInfluence.String(),
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := GetInfluenceByProbability(tc.strongProb, tc.moderateProb, tc.weakProb)
				require.NoError(tt, err)
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
