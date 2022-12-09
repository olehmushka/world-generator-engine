package language

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadAllWordbases(t *testing.T) {
	for chunk := range LoadAllWordbases() {
		if chunk.Err != nil {
			t.Fatalf("unexpected error (err=%+v)", chunk.Err)
			return
		}
		if chunk.Value == nil {
			t.Fatalf("unexpected nil value of loaded wordbase")
		}
	}
}

func TestSearchWordbase(t *testing.T) {
	slug := "ruthenian_wb"
	result, err := SearchWordbase(slug)
	require.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, result.Slug, slug)
}
