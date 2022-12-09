package gender

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetRandomSex(t *testing.T) {
	result, err := GetRandomSex()
	require.NoError(t, err)
	assert.Contains(t, []string{MaleSex.String(), FemaleSex.String()}, result.String())
}
