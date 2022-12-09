package culture

import (
	"testing"

	randomtools "github.com/olehmushka/golang-toolkit/random_tools"
	"github.com/olehmushka/world-generator-engine/gender"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	"github.com/olehmushka/world-generator-engine/influence"
	"github.com/olehmushka/world-generator-engine/language"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate(t *testing.T) {
	c, err := Generate(&CreateCultureOpts{
		LanguageSlugs: language.MockLanguageSlugs,
		Subbases:      mockSubbases,
		Ethoses:       mockEthoses,
		Traditions:    mockTraditions,
	}, mockCultures...)
	require.NoError(t, err)
	assert.Zero(t, c.Slug)
	assert.Contains(t, c.BaseSlug, RequiredBaseSlugSuffix)
	assert.Contains(t, c.Subbase.Slug, RequiredSubbaseSlugSuffix)
	assert.Equal(t, len(c.ParentCultures), len(mockCultures))
	assert.Contains(t, c.Ethos.Slug, RequiredEthosSlugSuffix)
	assert.Contains(t, c.LanguageSlug, RequiredLanguageSlugSuffix)
	assert.Contains(t, []string{gender.MaleSex.String(), "", gender.FemaleSex.String()}, c.GenderDominance.DominatedSex.String())
	assert.Contains(t, []string{influence.StrongInfluence.String(), influence.ModerateInfluence.String(), influence.WeakInfluence.String()}, c.GenderDominance.Influence.String())
	assert.Contains(t, []string{genderAcceptance.OnlyMen.String(), genderAcceptance.MenAndWomen.String(), genderAcceptance.OnlyWomen.String()}, c.MartialCustom.String())
	assert.GreaterOrEqual(t, len(c.Traditions), 2)
	assert.LessOrEqual(t, len(c.Traditions), 7)
}

func TestNew(t *testing.T) {
	c, err := New(&CreateCultureOpts{
		LanguageSlugs: language.MockLanguageSlugs,
		Subbases:      mockSubbases,
		Ethoses:       mockEthoses,
		Traditions:    mockTraditions,
	})
	require.NoError(t, err)
	assert.Zero(t, c.Slug)
	assert.Contains(t, c.BaseSlug, RequiredBaseSlugSuffix)
	assert.Contains(t, c.Subbase.Slug, RequiredSubbaseSlugSuffix)
	assert.Empty(t, c.ParentCultures)
	assert.Contains(t, c.Ethos.Slug, RequiredEthosSlugSuffix)
	assert.Contains(t, c.LanguageSlug, RequiredLanguageSlugSuffix)
	assert.Contains(t, []string{gender.MaleSex.String(), "", gender.FemaleSex.String()}, c.GenderDominance.DominatedSex.String())
	assert.Contains(t, []string{influence.StrongInfluence.String(), influence.ModerateInfluence.String(), influence.WeakInfluence.String()}, c.GenderDominance.Influence.String())
	assert.Contains(t, []string{genderAcceptance.OnlyMen.String(), genderAcceptance.MenAndWomen.String(), genderAcceptance.OnlyWomen.String()}, c.MartialCustom.String())
	assert.GreaterOrEqual(t, len(c.Traditions), 2)
	assert.LessOrEqual(t, len(c.Traditions), 7)
}

func TestGenerateSlug(t *testing.T) {
	for i := 0; i < 1000; i++ {
		assert.Contains(t, GenerateSlug(randomtools.UUIDString()), RequiredCultureSlugSuffix)
	}
}

func TestLoadAllRawCultures(t *testing.T) {
	var count int
	for chunk := range LoadAllRawCultures() {
		require.NoError(t, chunk.Err)
		assert.NotEmpty(t, chunk.Value)
		for _, c := range chunk.Value {
			assert.Contains(t, c.Slug, RequiredCultureSlugSuffix)
		}
		count += len(chunk.Value)
	}
	assert.Equal(t, 219, count)
}

func TestSearchCultures(t *testing.T) {
	cultures, err := SearchCultures([]string{"vlach_culture"})
	require.NoError(t, err)
	assert.NotEmpty(t, cultures)
}

func TestLoadAllCultures(t *testing.T) {
	for chunk := range LoadAllCultures() {
		require.NoError(t, chunk.Err)
		assert.NotNil(t, chunk.Value)
	}
}
