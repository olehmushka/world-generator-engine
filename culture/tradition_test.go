package culture

import (
	"strings"
	"testing"
)

func TestRandomTraditions(t *testing.T) {
	t.Error("test is not written yet")
}

func TestFilterTraditionsByEthos(t *testing.T) {
	t.Error("test is not written yet")
}

func TestFilterTraditionsByDomitatedSex(t *testing.T) {
	t.Error("test is not written yet")
}

func TestExtractTraditions(t *testing.T) {
	t.Error("test is not written yet")
}

func TestUniqueTraditions(t *testing.T) {
	t.Error("test is not written yet")
}

func TestLoadAllTraditions(t *testing.T) {
	var count int

	for chunk := range LoadAllTraditions() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of traditions")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredTraditionSlugSuffix) {
				t.Errorf("unexpected tradition slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 147; count != expecCount {
		t.Errorf("unexpected count of traditions (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestSearchTradition(t *testing.T) {
	slug := "astute_diplomats_tradition"
	result, err := SearchTradition(slug)
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}
	if result == nil {
		t.Fatal("result should not be nil")
	}
	if result.Slug != slug {
		t.Fatalf("unexpected result (expected slug=%s, actual slug=%s)", slug, result.Slug)
	}
}
