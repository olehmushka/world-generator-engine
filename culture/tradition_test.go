package culture

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/gender"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRandomTraditions(t *testing.T) {
	tCases := map[string]struct {
		inputTraditions []*Tradition
		min, max        int
		ethos           Ethos
		ds              gender.Sex
		expectedOutput  []*Tradition
	}{
		"should generate random tradition": {
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

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out, err := RandomTraditions(tc.inputTraditions, tc.min, tc.max, tc.ethos, tc.ds)
			require.NoError(t, err)
			assert.Equal(tt, len(out), len(tc.expectedOutput))
		})
	}
}

func TestFilterTraditionsByEthos(t *testing.T) {
	tCases := map[string]struct {
		name           string
		input          []*Tradition
		inputEthos     Ethos
		expectedOutput []*Tradition
	}{
		"should work for bellicose_ethos": {
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

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out := FilterTraditionsByEthos(tc.input, tc.inputEthos)
			assert.Equal(tt, len(out), len(tc.expectedOutput))
		})
	}
}

func TestFilterTraditionsByDomitatedSex(t *testing.T) {
	tCases := map[string]struct {
		input          []*Tradition
		inputSex       gender.Sex
		expectedOutput []*Tradition
	}{
		"should work for male sex": {
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

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out := FilterTraditionsByDomitatedSex(tc.input, tc.inputSex)
			assert.Equal(tt, len(out), len(tc.expectedOutput))
		})
	}
}

func TestExtractTraditions(t *testing.T) {
	traditions := ExtractTraditions(mockCultures)
	var expectedCount int
	for _, c := range mockCultures {
		expectedCount += len(c.Traditions)
	}
	assert.Equal(t, len(traditions), expectedCount)
	for _, tradition := range traditions {
		assert.Contains(t, tradition.Slug, RequiredTraditionSlugSuffix)
	}
}

func TestUniqueTraditions(t *testing.T) {
	tCases := map[string]struct {
		name           string
		input          []*Tradition
		expectedOutput []*Tradition
	}{
		"should work for shuffeled": {
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

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			out := UniqueTraditions(tc.input)
			assert.Equal(tt, len(out), len(tc.expectedOutput))
		})
	}
}

func TestLoadAllTraditions(t *testing.T) {
	var count int

	for chunk := range LoadAllTraditions() {
		require.NoError(t, chunk.Err)
		assert.NotEmpty(t, chunk.Value)
		for _, c := range chunk.Value {
			assert.Contains(t, c.Slug, RequiredTraditionSlugSuffix)
		}

		count += len(chunk.Value)
	}

	assert.Equal(t, 147, count)
}

func TestSearchTradition(t *testing.T) {
	slug := "astute_diplomats_tradition"
	result, err := SearchTradition(slug)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.Slug, slug)
}
