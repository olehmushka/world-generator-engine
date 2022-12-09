package religion

import (
	"strings"
	"testing"

	"github.com/olehmushka/world-generator-engine/permission"
	"github.com/stretchr/testify/require"
)

func TestLoadAllMarriageKinds(t *testing.T) {
	var count int

	for chunk := range LoadAllMarriageKinds() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of marriage_kinds")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredMarriageKindSlugSuffix) {
				t.Errorf("unexpected marriage_kind slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 3; count != expecCount {
		t.Errorf("unexpected count of marriage_kinds (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestLoadAllBastardies(t *testing.T) {
	var count int

	for chunk := range LoadAllBastardies() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of bastardies")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequireBastardySlugSuffix) {
				t.Errorf("unexpected bastard slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 3; count != expecCount {
		t.Errorf("unexpected count of bastardies (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestLoadAllConsanguinities(t *testing.T) {
	var count int

	for chunk := range LoadAllConsanguinities() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of consanguinities")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequireConsanguinitySlugSuffix) {
				t.Errorf("unexpected consanguinity slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 4; count != expecCount {
		t.Errorf("unexpected count of consanguinities (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestLoadAllDivorceOpts(t *testing.T) {
	var count int

	for chunk := range LoadAllDivorceOpts() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of divorce traits")
		}
		for _, c := range chunk.Value {
			if !permission.IsValid(c.Permission.String()) {
				t.Errorf("unexpected divorce (permission=%s)", c.Permission.String())
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 3; count != expecCount {
		t.Errorf("unexpected count of divorce traits (expected=%d, actual=%d)", expecCount, count)
	}
}
