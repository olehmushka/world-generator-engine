package culture

import (
	"testing"
)

func TestLoadAllEthoses(t *testing.T) {
	for chunk := range LoadAllEthoses() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Fatalf("unexpected length of subfamilies")
		}
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
