package main

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/language"
)

func TestXxx(t *testing.T) {
	result, err := language.LoadAllFamilies()
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}

	if len(result) == 0 {
		t.Fatalf("unexpected length of families")
	}
}
