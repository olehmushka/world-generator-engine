package acceptance

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
)

func TestGetAcceptanceByProbability(t *testing.T) {
	iterationsNumber := 100

	tCases := []struct {
		name                                        string
		onlyMenProb, menAndWomenProb, onlyWomenProb float64
		expectedOutput                              string
	}{
		{
			name:            "should returns only_men for greater strong probability",
			onlyMenProb:     0.35,
			menAndWomenProb: 0.2,
			onlyWomenProb:   0.2,
			expectedOutput:  OnlyMen.String(),
		},
		{
			name:            "should returns men_and_women for greater strong probability",
			onlyMenProb:     0.2,
			menAndWomenProb: 0.35,
			onlyWomenProb:   0.2,
			expectedOutput:  MenAndWomen.String(),
		},
		{
			name:            "should returns only_women for greater strong probability",
			onlyMenProb:     0.2,
			menAndWomenProb: 0.2,
			onlyWomenProb:   0.35,
			expectedOutput:  OnlyWomen.String(),
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := GetAcceptanceByProb(tc.onlyMenProb, tc.menAndWomenProb, tc.onlyWomenProb)
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
			if a := tools.GetKeyWithGreatestValue(m); a != tc.expectedOutput {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedOutput, a)
			}
		})
	}
}

func TestSelectAcceptanceByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := []struct {
		name           string
		input          []Acceptance
		expectedOutput Acceptance
	}{
		{
			name: "should works for 10 the same dominated sexes",
			input: []Acceptance{
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
			},
			expectedOutput: OnlyMen,
		},
		{
			name: "should works for 9 the same dominated sexes and 1 another",
			input: []Acceptance{
				MenAndWomen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
			},
			expectedOutput: OnlyMen,
		},
		{
			name: "should works for 8 the same dominated sexes and 2 other",
			input: []Acceptance{
				MenAndWomen,
				MenAndWomen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
			},
			expectedOutput: OnlyMen,
		},
		{
			name: "should works for 7 the same dominated sexes and 3 other",
			input: []Acceptance{
				OnlyWomen,
				OnlyWomen,
				OnlyWomen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
			},
			expectedOutput: OnlyMen,
		},
		{
			name: "should works for 6 the same dominated sexes and 4 other",
			input: []Acceptance{
				OnlyWomen,
				OnlyWomen,
				OnlyWomen,
				OnlyWomen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
				OnlyMen,
			},
			expectedOutput: OnlyMen,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectAcceptanceByMostRecent(tc.input)
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
			if a := tools.GetKeyWithGreatestValue(m); a != tc.expectedOutput.String() {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedOutput, a)
			}
		})
	}
}
