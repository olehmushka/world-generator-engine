package religion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAllHighGoals(t *testing.T) {
	var count int

	for chunk := range LoadAllHighGoals() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of high_goal")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequireHighGoalSlugSuffix) {
				t.Errorf("unexpected high_goal slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 16; count != expecCount {
		t.Errorf("unexpected count of marriage_kinds (expected=%d, actual=%d)", expecCount, count)
	}
}
