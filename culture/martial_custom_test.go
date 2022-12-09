package culture

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/gender"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"
	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomMartialCustom(t *testing.T) {
	iterationsNumber := 100
	tCases := map[string]struct {
		ds                         gender.Sex
		expectedMajorMartialCustom genderAcceptance.Acceptance
	}{
		"should work for male dominated sex": {
			ds:                         gender.MaleSex,
			expectedMajorMartialCustom: genderAcceptance.OnlyMen,
		},
		"should work for empty dominated sex": {
			ds:                         "",
			expectedMajorMartialCustom: genderAcceptance.MenAndWomen,
		},
		"should work for female dominated sex": {
			ds:                         gender.FemaleSex,
			expectedMajorMartialCustom: genderAcceptance.OnlyWomen,
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := RandomMartialCustom(genderDominance.Dominance{DominatedSex: tc.ds})
				require.NoError(t, err)
				if count, ok := m[out.String()]; ok {
					m[out.String()] = count + 1
				} else {
					m[out.String()] = 1
				}
			}
			assert.Equal(tt, tools.GetKeyWithGreatestValue(m), tc.expectedMajorMartialCustom.String())
		})
	}
}

func TestExtractMartialCusoms(t *testing.T) {
	assert.Equal(t, len(ExtractMartialCusoms(mockCultures)), len(mockCultures))
}
