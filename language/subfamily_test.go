package language

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadAllSubfamilies(t *testing.T) {
	for chunk := range LoadAllSubfamilies() {
		require.NoError(t, chunk.Err)
		if len(chunk.Value) == 0 {
			t.Fatalf("unexpected length of subfamilies")
		}
	}
}

func TestSearchSubfamily(t *testing.T) {
	slug := "ruthenian_lang_subfamily"
	result, err := SearchSubfamily(slug)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.Slug, slug)
}
