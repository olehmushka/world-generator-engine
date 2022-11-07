package influence

import (
	mapTools "github.com/olehmushka/golang-toolkit/map_tools"
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

type Influence string

func (i Influence) String() string {
	return string(i)
}

const (
	StrongInfluence   Influence = "strong_influence"
	ModerateInfluence Influence = "moderate_influence"
	WeakInfluence     Influence = "weak_influence"
)

func GetRandom() (Influence, error) {
	return GetInfluenceByProbability(0.32, 0.35, 0.32)
}

func GetInfluenceByProbability(strong, moderate, weak float64) (Influence, error) {
	i, err := mapTools.PickOneByProb(map[string]float64{
		string(StrongInfluence):   randomTools.PrepareProbability(strong),
		string(ModerateInfluence): randomTools.PrepareProbability(moderate),
		string(WeakInfluence):     randomTools.PrepareProbability(weak),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate infuence")
	}

	return Influence(i), nil
}
