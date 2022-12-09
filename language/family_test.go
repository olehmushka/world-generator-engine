package language

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadAllFamilies(t *testing.T) {
	for chunk := range LoadAllFamilies() {
		require.NoError(t, chunk.Err)
		assert.Greater(t, len(chunk.Value), 0)
	}
}
