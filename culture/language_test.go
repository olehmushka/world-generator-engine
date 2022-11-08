package culture

import (
	"strings"
	"testing"
)

func TestExtractLanguageSlugs(t *testing.T) {
	langSlugs := ExtractLanguageSlugs(mockCultures)
	if len(langSlugs) != len(mockCultures) {
		t.Errorf("unexpected extracted lang_slug length (expected=%d, actual=%d)", len(mockCultures), len(langSlugs))
	}
	for _, langSlug := range langSlugs {
		if !strings.HasSuffix(langSlug, RequiredLanguageSlugSuffix) {
			t.Errorf("unexpected subbase lang_slug suffix (slug=%s)", langSlug)
		}
	}
}
