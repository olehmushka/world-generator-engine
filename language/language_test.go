package language

import (
	"testing"
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
