package religion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAllSocialTraits(t *testing.T) {
	var count int

	for chunk := range LoadAllSocialTraits() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of high_goal")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequireSocialSlugSuffix) {
				t.Errorf("unexpected social slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 14; count != expecCount {
		t.Errorf("unexpected count of social (expected=%d, actual=%d)", expecCount, count)
	}
}
