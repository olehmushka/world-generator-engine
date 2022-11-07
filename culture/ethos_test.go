package culture

import (
	"strings"
	"testing"
)

func TestRandomEthos(t *testing.T) {
	t.Error("test is not written yet")
}

func TestExtractEthoses(t *testing.T) {
	t.Error("test is not written yet")
}

func TestSelectEthosByMostRecent(t *testing.T) {
	t.Error("test is not written yet")
}

func TestUniqueEthoses(t *testing.T) {
	t.Error("test is not written yet")
}

func TestLoadAllEthoses(t *testing.T) {
	var count int

	for chunk := range LoadAllEthoses() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of ethoses")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredEthosSlugSuffix) {
				t.Errorf("unexpected ethos slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 7; count != expecCount {
		t.Errorf("unexpected count of ethoses (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestSearchEthos(t *testing.T) {
	slug := "courtly_ethos"
	result, err := SearchEthos(slug)
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}
	if result.IsZero() {
		t.Fatal("result should not be nil")
	}
	if result.Slug != slug {
		t.Fatalf("unexpected result (expected slug=%s, actual slug=%s)", slug, result.Slug)
	}
}
