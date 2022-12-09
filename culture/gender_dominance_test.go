package culture

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExtractGenderDominances(t *testing.T) {
	gds := ExtractGenderDominances(mockCultures)
	assert.Equal(t, len(gds), len(mockCultures))
}
