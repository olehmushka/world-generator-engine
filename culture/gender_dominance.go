package culture

import (
	genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"
)

func ExtractGenderDominances(cultures []*Culture) []genderDominance.Dominance {
	out := make([]genderDominance.Dominance, len(cultures))
	for i := range out {
		out[i] = cultures[i].GenderDominance
	}

	return out
}
