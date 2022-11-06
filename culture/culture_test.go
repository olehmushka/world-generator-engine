package culture

import "testing"

func TestLoadAllRawCultures(t *testing.T) {
	var count int
	for chunk := range LoadAllRawCultures() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Fatalf("unexpected length of raw_cultures")
		}
		count += len(chunk.Value)
	}
	t.Logf("count:%d\n", count)
}
