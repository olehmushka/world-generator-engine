package religion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAllTypeTraits(t *testing.T) {
	var count int

	for chunk := range LoadAllTypeTraits() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of type traits")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequiredTypeSlugSuffix) {
				t.Errorf("unexpected type slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 8; count != expecCount {
		t.Errorf("unexpected count of type traits (expected=%d, actual=%d)", expecCount, count)
	}
}
