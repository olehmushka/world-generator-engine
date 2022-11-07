package culture

import (
	"strings"
	"testing"
)

func TestGenerate(t *testing.T) {
	t.Error("test is not written yet")
}

func TestNew(t *testing.T) {
	t.Error("test is not written yet")
}

func TestGenrateSlug(t *testing.T) {
	t.Error("test is not written yet")
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
