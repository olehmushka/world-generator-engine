package culture

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/gender"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"
	"github.com/olehmushka/world-generator-engine/tools"
)

func TestRandomMartialCustom(t *testing.T) {
	iterationsNumber := 100
	tCases := []struct {
		name                       string
		ds                         gender.Sex
		expectedMajorMartialCustom genderAcceptance.Acceptance
	}{
		{
			name:                       "should work for male dominated sex",
			ds:                         gender.MaleSex,
			expectedMajorMartialCustom: genderAcceptance.OnlyMen,
		},
		{
			name:                       "should work for empty dominated sex",
			ds:                         "",
			expectedMajorMartialCustom: genderAcceptance.MenAndWomen,
		},
		{
			name:                       "should work for female dominated sex",
			ds:                         gender.FemaleSex,
			expectedMajorMartialCustom: genderAcceptance.OnlyWomen,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := RandomMartialCustom(genderDominance.Dominance{DominatedSex: tc.ds})
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
			if mc := tools.GetKeyWithGreatestValue(m); mc != tc.expectedMajorMartialCustom.String() {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedMajorMartialCustom, mc)
			}
		})
	}
}

func TestExtractMartialCusoms(t *testing.T) {
	mcs := ExtractMartialCusoms(mockCultures)
	if len(mcs) != len(mockCultures) {
		t.Errorf("unexpected extracted ethos length (expected=%d, actual=%d)", len(mockCultures), len(mcs))
	}
}
