package culture

import (
	"strings"
	"testing"
)

func TestFilterSubbasesByBaseSlug(t *testing.T) {
	baseSlug := "europe_continental_base"
	result := FilterSubbasesByBaseSlug(mockSubbases, baseSlug)
	var count int
	for i, sb := range result {
		if sb.BaseSlug != baseSlug {
			t.Errorf("unexpected base_slug (expected=%s, actual=%s)", baseSlug, sb.BaseSlug)
		}
		count = i + 1
	}

	if expecCount := 24; count != expecCount {
		t.Errorf("unexpected count of filtered subbases (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestRandomSubbase(t *testing.T) {
	result, err := RandomSubbase(mockSubbases)
	if err != nil {
		t.Fatalf("unexpected err (err=%+v)", err)
	}
	if result.IsZero() {
		t.Errorf("result should not be empty string")
		return
	}
	var isSubbasesIncludeResult bool
	for _, sb := range mockSubbases {
		if sb.Slug == result.Slug {
			isSubbasesIncludeResult = true
		}
	}
	if !isSubbasesIncludeResult {
		t.Errorf("result subbase should be picked from input slice")
	}
}

func TestExtractSubbases(t *testing.T) {
	subbases := ExtractSubbases(mockCultures)
	if len(subbases) != len(mockCultures) {
		t.Errorf("unexpected extracted subbase length (expected=%d, actual=%d)", len(mockCultures), len(subbases))
	}
	for _, subbase := range subbases {
		if !strings.HasSuffix(subbase.Slug, RequiredSubbaseSlugSuffix) {
			t.Errorf("unexpected subbase slug suffix (slug=%s)", subbase.Slug)
		}
	}
}

func TestSelectSubbaseByMostRecent(t *testing.T) {
	t.Error("test is not written yet")
}

func TestLoadAllSubbases(t *testing.T) {
	var count int

	for chunk := range LoadAllSubbases() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of subbases")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredSubbaseSlugSuffix) {
				t.Errorf("unexpected subbase slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 70; count != expecCount {
		t.Errorf("unexpected count of subbases (expected=%d, actual=%d)", expecCount, count)
	}
}
