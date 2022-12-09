package worldgeneratorengine

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/language"
	"github.com/olehmushka/world-generator-engine/religion"
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

func TestLoadAllMarriageKinds(t *testing.T) {
	var count int
	for chunk := range religion.LoadAllMarriageKinds() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("loaded marriage_kinds can not be nil")
		}
		count++
	}
	t.Logf("counted marriage_kinds: %d\n", count)
}

func TestLoadAllBastardies(t *testing.T) {
	var count int
	for chunk := range religion.LoadAllBastardies() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("loaded bastardies can not be nil")
		}
		count++
	}
	t.Logf("counted bastardies: %d\n", count)
}

func TestLoadAllConsanguinities(t *testing.T) {
	var count int
	for chunk := range religion.LoadAllConsanguinities() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("loaded consanguinities can not be nil")
		}
		count++
	}
	t.Logf("counted bastardies: %d\n", count)
}

func TestLoadAllDivorceOpts(t *testing.T) {
	var count int
	for chunk := range religion.LoadAllDivorceOpts() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("loaded divorce opts can not be nil")
		}
		count++
	}
	t.Logf("counted divorce opts: %d\n", count)
}
