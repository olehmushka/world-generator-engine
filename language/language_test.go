package language

import "testing"

func TestLoadAllLanguages(t *testing.T) {
	for chunk := range LoadAllLanguages() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("loaded language can not be nil")
		}
	}
}
