package culture

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractLanguageSlugs(t *testing.T) {
	langSlugs := ExtractLanguageSlugs(mockCultures)
	assert.Equal(t, len(langSlugs), len(mockCultures))
	for _, langSlug := range langSlugs {
		assert.Contains(t, langSlug, RequiredLanguageSlugSuffix)
	}
}
