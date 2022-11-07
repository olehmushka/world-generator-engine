package language

import (
	"testing"

	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/world-generator-engine/tools"
)

func TestGetLanguageKinship(t *testing.T) {
	tCases := []struct {
		name     string
		l1, l2   *Language
		expected int
	}{
		{
			name: "should return for belarusian & ukrainian",
			l1: &Language{
				Slug: "belarusian",
				Subfamily: &Subfamily{
					Slug:       "ruthenian_lang_subfamily",
					FamilySlug: "indo_european_lang_family",
					ExtendedSubfamily: &Subfamily{
						Slug:       "slavic_lang_subfamily",
						FamilySlug: "indo_european_lang_family",
						ExtendedSubfamily: &Subfamily{
							Slug:       "balto_slavic_lang_subfamily",
							FamilySlug: "indo_european_lang_family",
						},
					},
				},
			},
			l2: &Language{
				Slug: "ukrainian",
				Subfamily: &Subfamily{
					Slug:       "ruthenian_lang_subfamily",
					FamilySlug: "indo_european_lang_family",
					ExtendedSubfamily: &Subfamily{
						Slug:       "slavic_lang_subfamily",
						FamilySlug: "indo_european_lang_family",
						ExtendedSubfamily: &Subfamily{
							Slug:       "balto_slavic_lang_subfamily",
							FamilySlug: "indo_european_lang_family",
						},
					},
				},
			},
			expected: 1,
		},
		{
			name: "should return for belarusian & russian",
			l1: &Language{
				Slug: "russian",
				Subfamily: &Subfamily{
					Slug:       "moscovian_lang_subfamily",
					FamilySlug: "indo_european_lang_family",
					ExtendedSubfamily: &Subfamily{
						Slug:       "slavic_lang_subfamily",
						FamilySlug: "indo_european_lang_family",
						ExtendedSubfamily: &Subfamily{
							Slug:       "balto_slavic_lang_subfamily",
							FamilySlug: "indo_european_lang_family",
						},
					},
				},
			},
			l2: &Language{
				Slug: "ukrainian",
				Subfamily: &Subfamily{
					Slug:       "ruthenian_lang_subfamily",
					FamilySlug: "indo_european_lang_family",
					ExtendedSubfamily: &Subfamily{
						Slug:       "slavic_lang_subfamily",
						FamilySlug: "indo_european_lang_family",
						ExtendedSubfamily: &Subfamily{
							Slug:       "balto_slavic_lang_subfamily",
							FamilySlug: "indo_european_lang_family",
						},
					},
				},
			},
			expected: 2,
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			if out := GetLanguageKinship(tc.l1, tc.l2); out != tc.expected {
				tt.Errorf("expected = %d, actual = %d", tc.expected, out)
			}
		})
	}
}

func TestRandomLanguageSlug(t *testing.T) {
	result, err := RandomLanguageSlug(MockLanguageSlugs)
	if err != nil {
		t.Fatalf("unexpected err (err=%+v)", err)
	}
	if result == "" {
		t.Errorf("result should not be empty string")
		return
	}
	if !sliceTools.Includes(MockLanguageSlugs, result) {
		t.Errorf("result language_slug should be picked from input slice")
	}
}

func TestFindLanguageBySlug(t *testing.T) {
	slug := "lithuanian_lang"
	lang := FindLanguageBySlug(MockLanguages, slug)
	if lang == nil {
		t.Error("language should not be nil")
		return
	}
	if lang.Slug != slug {
		t.Errorf("unexpected found language (expected=%s, actual=%s)", slug, lang.Slug)
	}
}

func TestSelectLanguageSlugByMostRecent(t *testing.T) {
	iterationsNumber := 100
	tCases := []struct {
		name           string
		input          []string
		expectedOutput string
	}{
		{
			name:           "should works for 10 the same langs",
			input:          []string{"frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang"},
			expectedOutput: "frankish_lang",
		},
		{
			name:           "should works for 9 the same langs and 1 another",
			input:          []string{"anglic_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang"},
			expectedOutput: "frankish_lang",
		},
		{
			name:           "should works for 8 the same langs and 2 another",
			input:          []string{"anglic_lang", "anglic_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang"},
			expectedOutput: "frankish_lang",
		},
		{
			name:           "should works for 7 the same langs and 3 another",
			input:          []string{"anglic_lang", "anglic_lang", "anglic_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang"},
			expectedOutput: "frankish_lang",
		},
		{
			name:           "should works for 6 the same langs and 4 another",
			input:          []string{"anglic_lang", "anglic_lang", "anglic_lang", "anglic_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang", "frankish_lang"},
			expectedOutput: "frankish_lang",
		},
	}

	for _, tc := range tCases {
		t.Run(tc.name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := SelectLanguageSlugByMostRecent(tc.input)
				if err != nil {
					t.Fatalf("unexpected error (err=%v)", err)
					return
				}
				if count, ok := m[out]; ok {
					m[out] = count + 1
				} else {
					m[out] = 1
				}
			}
			if slug := tools.GetKeyWithGreatestValue(m); slug != tc.expectedOutput {
				t.Errorf("unexpected result (expected=%s, actual=%s)", tc.expectedOutput, slug)
			}
		})
	}
}

func TestLoadAllLanguages(t *testing.T) {
	var count int
	for chunk := range LoadAllLanguages() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("loaded language can not be nil")
		}
		count++
	}
	t.Logf("counted langs: %d\n", count)
}
