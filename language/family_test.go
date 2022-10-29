package language

import (
	"testing"
)

func TestLoadAllFamilies(t *testing.T) {
	for chunk := range LoadAllFamilies() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Fatalf("unexpected length of families")
		}
	}
}
