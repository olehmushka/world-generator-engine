package religion

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadAllAfterlifeExists(t *testing.T) {
	var count int

	for chunk := range LoadAllAfterlifeExists() {
		require.NoError(t, chunk.Err)
		assert.Greater(t, len(chunk.Value), 0)
		count += len(chunk.Value)
	}

	assert.Equal(t, 2, count)
}

func TestLoadAllAfterlifeParticipances(t *testing.T) {
	var count int

	for chunk := range LoadAllAfterlifeParticipances() {
		require.NoError(t, chunk.Err)
		assert.Greater(t, len(chunk.Value), 0)
		for _, c := range chunk.Value {
			assert.Contains(t, c.Slug, RequireAfterlifeParticipanceSlugSuffix)
		}
		count += len(chunk.Value)
	}

	assert.Equal(t, 3, count)
}

func TestLoadAllAfterlifeParticipants(t *testing.T) {
	var count int

	for chunk := range LoadAllAfterlifeParticipants() {
		require.NoError(t, chunk.Err)
		assert.Greater(t, len(chunk.Value), 0)
		for _, c := range chunk.Value {
			assert.Contains(t, c.Slug, RequireAfterlifeParticipantSlugSuffix)
		}
		count += len(chunk.Value)
	}

	assert.Equal(t, 4, count)
}
