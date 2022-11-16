package worldgeneratorengine

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/language"
)

func TestLoadAllFamilies(t *testing.T) {
	for chunk := range language.LoadAllFamilies() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error: %+v", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of families")
		}
	}
}

func TestLoadAllSubfamilies(t *testing.T) {
	for chunk := range language.LoadAllSubfamilies() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error: %+v", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of subfamilies")
		}
	}
}

func TestLoadAllLanguages(t *testing.T) {
	var count int
	for chunk := range language.LoadAllLanguages() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("loaded language can not be nil")
		}
		count++
	}
	t.Logf("counted langs: %d\n", count)
}
