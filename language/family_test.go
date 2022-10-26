package language

import "testing"

func TestLoadAllFamilies(t *testing.T) {
	result, err := loadAllFamilies("")
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}

	if len(result) == 0 {
		t.Fatalf("unexpected length of families")
	}
}
