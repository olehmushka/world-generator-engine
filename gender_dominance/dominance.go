package genderdominance

import (
	"fmt"

	mapTools "github.com/olehmushka/golang-toolkit/map_tools"
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/world-generator-engine/gender"
	"github.com/olehmushka/world-generator-engine/influence"
)

type Dominance struct {
	DominatedSex gender.Sex          `json:"dominated_sex" bson:"dominated_sex"`
	Influence    influence.Influence `json:"influence" bson:"influence"`
}

func GetRandom() (Dominance, error) {
	ds, err := mapTools.PickOneByProb(map[string]float64{
		gender.MaleSex.String():   randomTools.PrepareProbability(0.33),
		"":                        randomTools.PrepareProbability(0.33),
		gender.FemaleSex.String(): randomTools.PrepareProbability(0.33),
	})
	if err != nil {
		return Dominance{}, wrapped_error.NewInternalServerError(err, "can not generate random geneder dominance")
	}

	i, err := influence.GetRandom()
	if err != nil {
		return Dominance{}, wrapped_error.NewInternalServerError(err, "can not generate infulence for gender_dominance")
	}

	return Dominance{
		DominatedSex: gender.Sex(ds),
		Influence:    i,
	}, nil
}

func SelectGenderDominanceByMostRecent(in []Dominance) (Dominance, error) {
	m := make(map[string]int, 3)
	for _, gd := range in {
		if count, ok := m[gd.DominatedSex.String()]; ok {
			m[gd.DominatedSex.String()] = count + 1
		} else {
			m[gd.DominatedSex.String()] = 1
		}
	}
	probs := make(map[string]float64, 3)
	total := float64(len(in))
	for ds, count := range m {
		probs[ds] = float64(count) / total
	}

	ds, err := mapTools.PickOneByProb(probs)
	if err != nil {
		return Dominance{}, wrapped_error.NewInternalServerError(err, "can not select dominated sex by most recent")
	}
	for _, gd := range in {
		if gd.DominatedSex.String() == ds {
			return gd, nil
		}
	}

	return Dominance{}, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not select gender dominance by dominanted_sex (dominated_sex=%s)", ds))
}
