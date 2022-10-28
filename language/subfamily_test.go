package language

import (
	"testing"
)

func TestLoadAllSubfamilies(t *testing.T) {
	result, err := LoadAllSubfamilies()
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}

	if len(result) == 0 {
		t.Fatalf("unexpected length of subfamilies")
	}
}

func TestSearchSubfamily(t *testing.T) {
	slug := "ruthenian_lang_subfamily"
	result, err := SearchSubfamily(slug)
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
