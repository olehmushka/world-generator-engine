package culture

import (
	mapTools "github.com/olehmushka/golang-toolkit/map_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"
)

func RandomMartialCustom(d genderDominance.Dominance) (genderAcceptance.Acceptance, error) {
	var (
		onlyMen     = 0.35
		menAndWomen = 0.2
		onlyWomen   = 0.05
	)
	switch {
	case d.DominatedSex.IsMale():
		onlyMen += 0.2
	case d.DominatedSex.IsZero():
		menAndWomen += 0.2
	case d.DominatedSex.IsFemale():
		onlyWomen += 0.2
	}

	mc, err := mapTools.PickOneByProb(map[string]float64{
		genderAcceptance.OnlyMen.String():     onlyMen,
		genderAcceptance.MenAndWomen.String(): menAndWomen,
		genderAcceptance.OnlyWomen.String():   onlyWomen,
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate random martial custom")
	}

	return genderAcceptance.Acceptance(mc), nil
}

func ExtractMartialCusoms(cultures []*Culture) []genderAcceptance.Acceptance {
	out := make([]genderAcceptance.Acceptance, len(cultures))
	for i := range out {
		out[i] = cultures[i].MartialCustom
	}

	return out
}
