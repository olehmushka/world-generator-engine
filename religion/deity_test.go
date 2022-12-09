package religion

import (
	"strings"
	"testing"

	"github.com/olehmushka/world-generator-engine/favour"
	"github.com/stretchr/testify/require"
)

func TestLoadAllDeityFavours(t *testing.T) {
	var count int

	for chunk := range LoadAllDeityFavours() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of divorce traits")
		}
		for _, c := range chunk.Value {
			if !favour.IsValid(c.Favour.String()) {
				t.Errorf("unexpected deity favour (favour=%s)", c.Favour.String())
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 17; count != expecCount {
		t.Errorf("unexpected count of favour traits (expected=%d, actual=%d)", expecCount, count)
	}
}

func TestLoadAllDeityNatureTraits(t *testing.T) {
	var count int

	for chunk := range LoadAllDeityNatureTraits() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of deity_nature traits")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequireDeityNatureSlugSuffix) {
				t.Errorf("unexpected deity_nature slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 7; count != expecCount {
		t.Errorf("unexpected count of deity_nature traits (expected=%d, actual=%d)", expecCount, count)
	}
}
