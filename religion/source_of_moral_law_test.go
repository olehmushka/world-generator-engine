package religion

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadAllSourcesOfMoralLaw(t *testing.T) {
	var count int

	for chunk := range LoadAllSourcesOfMoralLaw() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Errorf("unexpected length of source_of_moral_law traits")
		}
		for _, c := range chunk.Value {
			if !strings.HasSuffix(c.Slug, RequireSourceOfMoralLawSlugSuffix) {
				t.Errorf("unexpected source_of_moral_law slug suffix (slug=%s)", c.Slug)
			}
		}

		count += len(chunk.Value)
	}

	if expecCount := 6; count != expecCount {
		t.Errorf("unexpected count of source_of_moral_law traits (expected=%d, actual=%d)", expecCount, count)
	}
}
