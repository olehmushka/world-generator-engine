package culture

import (
	"testing"
)

func TestLoadAllBases(t *testing.T) {
	for chunk := range LoadAllBases() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Fatalf("unexpected length of bases")
		}
	}
}
