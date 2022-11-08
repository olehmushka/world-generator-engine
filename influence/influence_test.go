package influence

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
)

func TestGetRandom(t *testing.T) {
	iterationsNumber := 10000
	tCases := []struct {
		name           string
		expectedOutput string
	}{
		{
			name:           "should returns moderate_influence for most cases",
			expectedOutput: ModerateInfluence.String(),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := GetRandom()
				if err != nil {
					t.Fatalf("unexpected error (err=%v)", err)
					return
				}
				if count, ok := m[out.String()]; ok {
					m[out.String()] = count + 1
				} else {
					m[out.String()] = 1
				}
			}
			if slug := tools.GetKeyWithGreatestValue(m); slug != tc.expectedOutput {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedOutput, slug)
			}
		})
	}
}

func TestGetInfluenceByProbability(t *testing.T) {
	iterationsNumber := 1000
	tCases := []struct {
		name                               string
		strongProb, moderateProb, weakProb float64
		expectedOutput                     string
	}{
		{
			name:           "should returns strong_influence for greater strong probability",
			strongProb:     0.35,
			moderateProb:   0.2,
			weakProb:       0.2,
			expectedOutput: StrongInfluence.String(),
		},
		{
			name:           "should returns moderate_influence for greater moderate probability",
			strongProb:     0.2,
			moderateProb:   0.35,
			weakProb:       0.2,
			expectedOutput: ModerateInfluence.String(),
		},
		{
			name:           "should returns weak_influence for greater weak probability",
			strongProb:     0.2,
			moderateProb:   0.2,
			weakProb:       0.35,
			expectedOutput: WeakInfluence.String(),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := GetInfluenceByProbability(tc.strongProb, tc.moderateProb, tc.weakProb)
				if err != nil {
					t.Fatalf("unexpected error (err=%v)", err)
					return
				}
				if count, ok := m[out.String()]; ok {
					m[out.String()] = count + 1
				} else {
					m[out.String()] = 1
				}
			}
			if slug := tools.GetKeyWithGreatestValue(m); slug != tc.expectedOutput {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedOutput, slug)
			}
		})
	}
}
