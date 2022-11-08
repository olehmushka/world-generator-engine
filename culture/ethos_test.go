package culture

import (
	"strings"
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
)

func TestRandomEthos(t *testing.T) {
	result, err := RandomEthos(mockEthoses)
	if err != nil {
		t.Fatalf("unexpected err (err=%+v)", err)
	}
	if result.IsZero() {
		t.Errorf("result should not be empty string")
		return
	}
	var isEthosesIncludeResult bool
	for _, e := range mockEthoses {
		if e.Slug == result.Slug {
			isEthosesIncludeResult = true
		}
	}
	if !isEthosesIncludeResult {
		t.Errorf("result ethos should be picked from input slice")
	}
}

func TestExtractEthoses(t *testing.T) {
	ethoses := ExtractEthoses(mockCultures)
	if len(ethoses) != len(mockCultures) {
		t.Errorf("unexpected extracted ethos length (expected=%d, actual=%d)", len(mockCultures), len(ethoses))
	}
	for _, ethos := range ethoses {
		if !strings.HasSuffix(ethos.Slug, RequiredEthosSlugSuffix) {
			t.Errorf("unexpected ethos slug suffix (slug=%s)", ethos.Slug)
		}
	}
}

func TestSelectEthosByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := []struct {
		name           string
		input          []Ethos
		expectedOutput Ethos
	}{
		{
			name: "should works for 10 the same ethoses",
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
		{
			name: "should works for 9 the same ethoses and 1 another",
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
		{
			name: "should works for 8 the same ethoses and 2 other",
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
		{
			name: "should works for 7 the same ethoses and 3 other",
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
		{
			name: "should works for 6 the same ethoses and 4 other",
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

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectEthosByMostRecent(tc.input)
				if err != nil {
					t.Fatalf("unexpected error (err=%v)", err)
					return
				}
				if count, ok := m[out.Slug]; ok {
					m[out.Slug] = count + 1
				} else {
					m[out.Slug] = 1
				}
			}
			if slug := tools.GetKeyWithGreatestValue(m); slug != tc.expectedOutput.Slug {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedOutput.Slug, slug)
			}
		})
	}
}

func TestUniqueEthoses(t *testing.T) {
	tCases := []struct {
		name           string
		input          []Ethos
		expectedOutput []Ethos
	}{
		{
			name: "should work for shuffeled",
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

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := UniqueEthoses(tc.input); len(out) != len(tc.expectedOutput) {
				t.Errorf("unexpected output (expected_length=%+v, actual_length=%+v)", len(tc.expectedOutput), len(out))
			}
		})
	}
}

func TestLoadAllEthoses(t *testing.T) {
	var count int

	for chunk := range LoadAllEthoses() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of ethoses")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredEthosSlugSuffix) {
				t.Errorf("unexpected ethos slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 7; count != expecCount {
		t.Errorf("unexpected count of ethoses (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestSearchEthos(t *testing.T) {
	slug := "communal_ethos"
	result, err := SearchEthos(slug)
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}
	if result.IsZero() {
		t.Fatal("result should not be nil")
	}
	if result.Slug != slug {
		t.Fatalf("unexpected result (expected slug=%s, actual slug=%s)", slug, result.Slug)
	}
}
