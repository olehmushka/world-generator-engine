package culture

import (
	"strings"
	"testing"
)

func TestRandomBase(t *testing.T) {
	t.Error("test is not written yet")
}

func TestExtractBases(t *testing.T) {
	t.Error("test is not written yet")
}

func TestSelectBaseByMostRecent(t *testing.T) {
	t.Error("test is not written yet")
}

func TestLoadAllBases(t *testing.T) {
	var count int

	for chunk := range LoadAllBases() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of bases")
		}

		for _, base := range chunk.Value {
			if !strings.HasSuffix(base, RequiredBaseSlugSuffix) {
				t.Errorf("unexpected base slug suffix (slug=%s)", base)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 28; count != expecCount {
		t.Errorf("unexpected count of bases (expected=%d, actual=%d)", expecCount, count)
	}
}
