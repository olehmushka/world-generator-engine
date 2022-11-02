package culture

import (
	"testing"
)

func TestLoadAllTraditions(t *testing.T) {
	for chunk := range LoadAllTraditions() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Fatalf("unexpected length of traditions")
		}
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
