package culture

import (
	"strings"
	"testing"

	"github.com/olehmushka/world-generator-engine/gender"
)

func TestRandomTraditions(t *testing.T) {
	tCases := []struct {
		name            string
		inputTraditions []*Tradition
		min, max        int
		ethos           Ethos
		ds              gender.Sex
		expectedOutput  []*Tradition
	}{
		{
			name: "",
			inputTraditions: []*Tradition{
				{Slug: "equal_inheritance_tradition", OmitGenderDominance: []gender.Sex{"male_sex", "female_sex"}},
				{Slug: "equal_inheritance_tradition", OmitGenderDominance: []gender.Sex{"male_sex", ""}},
				{Slug: "city_keepers_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "forest_folk_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "hit_and_run_tacticians_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "agrarian_tradition", OmitGenderDominance: []gender.Sex{}, OmitEthosSlugs: []string{"bellicose_ethos"}},
			},
			min:   2,
			max:   5,
			ethos: Ethos{Slug: "bellicose_ethos"},
			ds:    gender.MaleSex,
			expectedOutput: []*Tradition{
				{Slug: "city_keepers_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "forest_folk_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "hit_and_run_tacticians_tradition", OmitGenderDominance: []gender.Sex{}},
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			out, err := RandomTraditions(tc.inputTraditions, tc.min, tc.max, tc.ethos, tc.ds)
			if err != nil {
				t.Fatalf("unexpected err (err=%+v)", err)
			}
			if len(out) != len(tc.expectedOutput) {
				t.Errorf("unexpected output (expected_length=%+v, actual_length=%+v)", len(tc.expectedOutput), len(out))
			}
		})
	}
}

func TestFilterTraditionsByEthos(t *testing.T) {
	tCases := []struct {
		name           string
		input          []*Tradition
		inputEthos     Ethos
		expectedOutput []*Tradition
	}{
		{
			name: "should work for bellicose_ethos",
			input: []*Tradition{
				{Slug: "equal_inheritance_tradition", OmitEthosSlugs: []string{"bellicose_ethos"}},
				{Slug: "city_keepers_tradition", OmitEthosSlugs: []string{"bellicose_ethos"}},
				{Slug: "forest_folk_tradition"},
				{Slug: "hit_and_run_tacticians_tradition"},
				{Slug: "agrarian_tradition", OmitEthosSlugs: []string{"bellicose_ethos"}},
			},
			inputEthos: Ethos{Slug: "bellicose_ethos"},
			expectedOutput: []*Tradition{
				{Slug: "forest_folk_tradition"},
				{Slug: "hit_and_run_tacticians_tradition"},
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := FilterTraditionsByEthos(tc.input, tc.inputEthos); len(out) != len(tc.expectedOutput) {
				t.Errorf("unexpected output (expected_length=%+v, actual_length=%+v)", len(tc.expectedOutput), len(out))
			}
		})
	}
}

func TestFilterTraditionsByDomitatedSex(t *testing.T) {
	tCases := []struct {
		name           string
		input          []*Tradition
		inputSex       gender.Sex
		expectedOutput []*Tradition
	}{
		{
			name: "should work for male sex",
			input: []*Tradition{
				{Slug: "equal_inheritance_tradition", OmitGenderDominance: []gender.Sex{"male_sex", "female_sex"}},
				{Slug: "equal_inheritance_tradition", OmitGenderDominance: []gender.Sex{"male_sex", ""}},
				{Slug: "city_keepers_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "forest_folk_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "hit_and_run_tacticians_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "agrarian_tradition", OmitGenderDominance: []gender.Sex{}},
			},
			inputSex: gender.MaleSex,
			expectedOutput: []*Tradition{
				{Slug: "city_keepers_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "forest_folk_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "hit_and_run_tacticians_tradition", OmitGenderDominance: []gender.Sex{}},
				{Slug: "agrarian_tradition", OmitGenderDominance: []gender.Sex{}},
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := FilterTraditionsByDomitatedSex(tc.input, tc.inputSex); len(out) != len(tc.expectedOutput) {
				t.Errorf("unexpected output (expected_length=%+v, actual_length=%+v)", len(tc.expectedOutput), len(out))
			}
		})
	}
}

func TestExtractTraditions(t *testing.T) {
	traditions := ExtractTraditions(mockCultures)
	var expectedCount int
	for _, c := range mockCultures {
		expectedCount += len(c.Traditions)
	}
	if len(traditions) != expectedCount {
		t.Errorf("unexpected extracted tradition length (expected=%d, actual=%d)", expectedCount, len(traditions))
	}
	for _, tradition := range traditions {
		if !strings.HasSuffix(tradition.Slug, RequiredTraditionSlugSuffix) {
			t.Errorf("unexpected tradition slug suffix (slug=%s)", tradition.Slug)
		}
	}
}

func TestUniqueTraditions(t *testing.T) {
	tCases := []struct {
		name           string
		input          []*Tradition
		expectedOutput []*Tradition
	}{
		{
			name: "should work for shuffeled",
			input: []*Tradition{
				{Slug: "agrarian_tradition"},
				{Slug: "druzhina_tradition"},
				{Slug: "charismatic_tradition"},
				{Slug: "city_keepers_tradition"},
				{Slug: "agrarian_tradition"},
				{Slug: "druzhina_tradition"},
				{Slug: "forest_folk_tradition"},
				{Slug: "hit_and_run_tacticians_tradition"},
			},
			expectedOutput: []*Tradition{
				{Slug: "druzhina_tradition"},
				{Slug: "charismatic_tradition"},
				{Slug: "city_keepers_tradition"},
				{Slug: "forest_folk_tradition"},
				{Slug: "hit_and_run_tacticians_tradition"},
				{Slug: "agrarian_tradition"},
			},
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := UniqueTraditions(tc.input); len(out) != len(tc.expectedOutput) {
				t.Errorf("unexpected output (expected_length=%+v, actual_length=%+v)", len(tc.expectedOutput), len(out))
			}
		})
	}
}

func TestLoadAllTraditions(t *testing.T) {
	var count int

	for chunk := range LoadAllTraditions() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of traditions")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredTraditionSlugSuffix) {
				t.Errorf("unexpected tradition slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 147; count != expecCount {
		t.Errorf("unexpected count of traditions (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestSearchTradition(t *testing.T) {
	slug := "astute_diplomats_tradition"
	result, err := SearchTradition(slug)
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.Slug != slug {
		t.Fatalf("unexpected result (expected slug=%s, actual slug=%s)", slug, result.Slug)
	}
}
