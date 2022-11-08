package genderdominance

import (
	"testing"

	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/world-generator-engine/gender"
	"github.com/olehmushka/world-generator-engine/influence"
	"github.com/olehmushka/world-generator-engine/tools"
)

func TestGetRandom(t *testing.T) {
	result, err := GetRandom()
	if err != nil {
		t.Fatalf("unexpected err (err=%+v)", err)
	}
	if !sliceTools.Includes([]string{
		gender.MaleSex.String(),
		"",
		gender.FemaleSex.String(),
	}, result.DominatedSex.String()) {
		t.Errorf("result dominated sex should be picked from available sexes or empty")
	}

	if !sliceTools.Includes([]string{
		influence.StrongInfluence.String(),
		influence.ModerateInfluence.String(),
		influence.WeakInfluence.String(),
	}, result.Influence.String()) {
		t.Errorf("result influence should be picked from available influences")
	}
}

func TestSelectGenderDominanceByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := []struct {
		name           string
		input          []Dominance
		expectedOutput gender.Sex
	}{
		{
			name: "should works for 10 the same dominated sexes",
			input: []Dominance{
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
			},
			expectedOutput: gender.MaleSex,
		},
		{
			name: "should works for 9 the same dominated sexes and 1 another",
			input: []Dominance{
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
			},
			expectedOutput: gender.MaleSex,
		},
		{
			name: "should works for 8 the same dominated sexes and 2 other",
			input: []Dominance{
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
			},
			expectedOutput: gender.MaleSex,
		},
		{
			name: "should works for 7 the same dominated sexes and 3 other",
			input: []Dominance{
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
			},
			expectedOutput: gender.MaleSex,
		},
		{
			name: "should works for 6 the same dominated sexes and 4 other",
			input: []Dominance{
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.FemaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
				{DominatedSex: gender.MaleSex},
			},
			expectedOutput: gender.MaleSex,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectGenderDominanceByMostRecent(tc.input)
				if err != nil {
					t.Fatalf("unexpected error (err=%v)", err)
					return
				}
				if count, ok := m[out.DominatedSex.String()]; ok {
					m[out.DominatedSex.String()] = count + 1
				} else {
					m[out.DominatedSex.String()] = 1
				}
			}
			if ds := tools.GetKeyWithGreatestValue(m); ds != tc.expectedOutput.String() {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedOutput, ds)
			}
		})
	}
}
