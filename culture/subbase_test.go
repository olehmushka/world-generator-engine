package culture

import (
	"strings"
	"testing"
)

func TestFilterSubbasesByBaseSlug(t *testing.T) {
	t.Error("test is not written yet")
}

func TestRandomSubbase(t *testing.T) {
	t.Error("test is not written yet")
}

func TestExtractSubbases(t *testing.T) {
	t.Error("test is not written yet")
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
