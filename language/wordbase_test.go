package language

import (
	"testing"
)

func TestLoadAllWordbases(t *testing.T) {
	for chunk := range LoadAllWordbases() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("unexpected nil value of loaded wordbase")
		}
	}
}

func TestSearchWordbase(t *testing.T) {
	slug := "ruthenian_wb"
	result, err := SearchWordbase(slug)
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
