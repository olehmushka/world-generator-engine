package language

import (
	"testing"
)

func TestLoadAllSubfamilies(t *testing.T) {
	result, err := LoadAllSubfamilies()
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}

	if len(result) == 0 {
		t.Fatalf("unexpected length of subfamilies")
	}
}
