package culture

import (
	"strings"
	"testing"
)

func TestRandomEthos(t *testing.T) {
	result, err := RandomEthos(mockEthoses)
	if err != nil {
		t.Fatalf("unexpected err (err=%+v)", err)
	}
	if result.IsZero() {
		t.Errorf("result should not be empty string")
		return
	}
	var isEthosesIncludeResult bool
	for _, e := range mockEthoses {
		if e.Slug == result.Slug {
			isEthosesIncludeResult = true
		}
	}
	if !isEthosesIncludeResult {
		t.Errorf("result ethos should be picked from input slice")
	}
}

func TestExtractEthoses(t *testing.T) {
	ethoses := ExtractEthoses(mockCultures)
	if len(ethoses) != len(mockCultures) {
		t.Errorf("unexpected extracted ethos length (expected=%d, actual=%d)", len(mockCultures), len(ethoses))
	}
	for _, ethos := range ethoses {
		if !strings.HasSuffix(ethos.Slug, RequiredEthosSlugSuffix) {
			t.Errorf("unexpected ethos slug suffix (slug=%s)", ethos.Slug)
		}
	}
}

func TestSelectEthosByMostRecent(t *testing.T) {
	t.Error("test is not written yet")
}

func TestUniqueEthoses(t *testing.T) {
	t.Error("test is not written yet")
}

func TestLoadAllEthoses(t *testing.T) {
	var count int

	for chunk := range LoadAllEthoses() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of ethoses")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredEthosSlugSuffix) {
				t.Errorf("unexpected ethos slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 7; count != expecCount {
		t.Errorf("unexpected count of ethoses (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestSearchEthos(t *testing.T) {
	slug := "courtly_ethos"
	result, err := SearchEthos(slug)
	if err != nil {
		t.Fatalf("unexpected error (err=%+v)", err)
		return
	}
	if result.IsZero() {
		t.Fatal("result should not be nil")
	}
	if result.Slug != slug {
		t.Fatalf("unexpected result (expected slug=%s, actual slug=%s)", slug, result.Slug)
	}
}
