package religion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAllHumanNatureTraits(t *testing.T) {
	var count int

	for chunk := range LoadAllHumanNatureTraits() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of human_nature traits")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequireHumanNatureSlugSuffix) {
				t.Errorf("unexpected human_nature slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 2; count != expecCount {
		t.Errorf("unexpected count of human_nature traits (expected=%d, actual=%d)", expecCount, count)
	}
}
