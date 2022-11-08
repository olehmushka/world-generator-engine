package culture

import (
	"strings"
	"testing"

	randomtools "github.com/olehmushka/golang-toolkit/random_tools"
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/world-generator-engine/gender"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	"github.com/olehmushka/world-generator-engine/influence"
	"github.com/olehmushka/world-generator-engine/language"
)

func TestGenerate(t *testing.T) {
	c, err := Generate(&CreateCultureOpts{
		LanguageSlugs: language.MockLanguageSlugs,
		Subbases:      mockSubbases,
		Ethoses:       mockEthoses,
		Traditions:    mockTraditions,
	}, mockCultures...)
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}
	if c.Slug != "" {
		t.Errorf("unexpected empty slug name (slug=%s)", c.Slug)
	}
	if !strings.HasSuffix(c.BaseSlug, RequiredBaseSlugSuffix) {
		t.Errorf("unexpected base slug suffix (slug=%s)", c.BaseSlug)
	}
	if !strings.HasSuffix(c.Subbase.Slug, RequiredSubbaseSlugSuffix) {
		t.Errorf("unexpected subbase slug suffix (slug=%s)", c.Subbase.Slug)
	}
	if len(c.ParentCultures) != len(mockCultures) {
		t.Errorf("parent cultures should have %d count slice (legnth=%d)", len(mockCultures), len(c.ParentCultures))
	}
	if !strings.HasSuffix(c.Ethos.Slug, RequiredEthosSlugSuffix) {
		t.Errorf("unexpected ethos slug suffix (slug=%s)", c.Ethos.Slug)
	}
	if !strings.HasSuffix(c.LanguageSlug, RequiredLanguageSlugSuffix) {
		t.Errorf("unexpected lang slug suffix (slug=%s)", c.LanguageSlug)
	}
	if !sliceTools.Includes([]string{
		gender.MaleSex.String(),
		"",
		gender.FemaleSex.String(),
	}, c.GenderDominance.DominatedSex.String()) {
		t.Errorf("result dominated sex should be picked from available sexes or empty")
	}

	if !sliceTools.Includes([]string{
		influence.StrongInfluence.String(),
		influence.ModerateInfluence.String(),
		influence.WeakInfluence.String(),
	}, c.GenderDominance.Influence.String()) {
		t.Errorf("result influence should be picked from available influences")
	}
	if !sliceTools.Includes([]string{
		genderAcceptance.OnlyMen.String(),
		genderAcceptance.MenAndWomen.String(),
		genderAcceptance.OnlyWomen.String(),
	}, c.MartialCustom.String()) {
		t.Errorf("result martial custom should be picked from available acceptances")
	}
	if len(c.Traditions) < 2 || len(c.Traditions) > 7 {
		t.Errorf("traditions number is out of range (length=%d)", len(c.Traditions))
	}
}

func TestNew(t *testing.T) {
	c, err := New(&CreateCultureOpts{
		LanguageSlugs: language.MockLanguageSlugs,
		Subbases:      mockSubbases,
		Ethoses:       mockEthoses,
		Traditions:    mockTraditions,
	})
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}
	if c.Slug != "" {
		t.Errorf("unexpected empty slug name (slug=%s)", c.Slug)
	}
	if !strings.HasSuffix(c.BaseSlug, RequiredBaseSlugSuffix) {
		t.Errorf("unexpected base slug suffix (slug=%s)", c.BaseSlug)
	}
	if !strings.HasSuffix(c.Subbase.Slug, RequiredSubbaseSlugSuffix) {
		t.Errorf("unexpected subbase slug suffix (slug=%s)", c.Subbase.Slug)
	}
	if len(c.ParentCultures) != 0 {
		t.Errorf("parent cultures should be empty slice (legnth=%d)", len(c.ParentCultures))
	}
	if !strings.HasSuffix(c.Ethos.Slug, RequiredEthosSlugSuffix) {
		t.Errorf("unexpected ethos slug suffix (slug=%s)", c.Ethos.Slug)
	}
	if !strings.HasSuffix(c.LanguageSlug, RequiredLanguageSlugSuffix) {
		t.Errorf("unexpected lang slug suffix (slug=%s)", c.LanguageSlug)
	}
	if !sliceTools.Includes([]string{
		gender.MaleSex.String(),
		"",
		gender.FemaleSex.String(),
	}, c.GenderDominance.DominatedSex.String()) {
		t.Errorf("result dominated sex should be picked from available sexes or empty")
	}

	if !sliceTools.Includes([]string{
		influence.StrongInfluence.String(),
		influence.ModerateInfluence.String(),
		influence.WeakInfluence.String(),
	}, c.GenderDominance.Influence.String()) {
		t.Errorf("result influence should be picked from available influences")
	}
	if !sliceTools.Includes([]string{
		genderAcceptance.OnlyMen.String(),
		genderAcceptance.MenAndWomen.String(),
		genderAcceptance.OnlyWomen.String(),
	}, c.MartialCustom.String()) {
		t.Errorf("result martial custom should be picked from available acceptances")
	}
	if len(c.Traditions) < 2 || len(c.Traditions) > 7 {
		t.Errorf("traditions number is out of range (length=%d)", len(c.Traditions))
	}
}

func TestGenerateSlug(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if slug := GenerateSlug(randomtools.UUIDString()); !strings.HasSuffix(slug, RequiredCultureSlugSuffix) {
			t.Errorf("unexpected generated lang_slug suffix (slug=%s)", slug)
		}
	}
}

func TestLoadAllRawCultures(t *testing.T) {
	var count int

	for chunk := range LoadAllRawCultures() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of raw_cultures")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredCultureSlugSuffix) {
				t.Errorf("unexpected culture slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}
	if expecCount := 219; count != expecCount {
		t.Errorf("unexpected count of cultures (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestSearchCultures(t *testing.T) {
	cultures, err := SearchCultures([]string{"vlach_culture"})
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}
	if len(cultures) == 0 {
		t.Errorf("unexpected number or found cultures")
	}
}

func TestLoadAllCultures(t *testing.T) {
	var count int
	for chunk := range LoadAllCultures() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("loaded culture can not be nil")
		}
		count++
		if count%25 == 0 {
			t.Logf("counted cultures: %d\n", count)
		}
	}
	t.Logf("counted cultures: %d\n", count)
}
