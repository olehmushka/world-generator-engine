package permission

import (
	"testing"

	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetPermissionByProb(t *testing.T) {
	iterationsNumber := 100

	tCases := map[string]struct {
		alwaysAllowedProb, mustBeApprovedProb, disallowedProb float64
		expectedOutput                                        string
	}{
		"should returns always_allowed for greater strong probability": {
			alwaysAllowedProb:  0.35,
			mustBeApprovedProb: 0.2,
			disallowedProb:     0.2,
			expectedOutput:     AlwaysAllowed.String(),
		},
		"should returns must_be_approved for greater strong probability": {
			alwaysAllowedProb:  0.2,
			mustBeApprovedProb: 0.35,
			disallowedProb:     0.2,
			expectedOutput:     MustBeApproved.String(),
		},
		"should returns disallowed for greater strong probability": {
			alwaysAllowedProb:  0.2,
			mustBeApprovedProb: 0.2,
			disallowedProb:     0.35,
			expectedOutput:     Disallowed.String(),
		},
	}

	for name, tc := range tCases {
		t.Run(name, func(tt *testing.T) {
			m := make(map[string]int)
			for i := 0; i < iterationsNumber; i++ {
				out, err := GetPermissionByProb(tc.alwaysAllowedProb, tc.mustBeApprovedProb, tc.disallowedProb)
				require.NoError(t, err)
				if count, ok := m[out.String()]; ok {
					m[out.String()] = count + 1
				} else {
					m[out.String()] = 1
				}
			}
			assert.Equal(tt, tools.GetKeyWithGreatestValue(m), tc.expectedOutput)
		})
	}
}
