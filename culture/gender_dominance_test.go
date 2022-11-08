package culture

import "testing"

func TestExtractGenderDominances(t *testing.T) {
	gds := ExtractGenderDominances(mockCultures)
	if len(gds) != len(mockCultures) {
		t.Errorf("unexpected extracted lang_slug length (expected=%d, actual=%d)", len(mockCultures), len(gds))
	}
}
