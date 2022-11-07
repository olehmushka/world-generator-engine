package acceptance

import (
	"fmt"

	mapTools "github.com/olehmushka/golang-toolkit/map_tools"
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

type Acceptance string

func (a Acceptance) String() string {
	return string(a)
}

const (
	OnlyMen     Acceptance = "only_men"
	MenAndWomen Acceptance = "men_and_women"
	OnlyWomen   Acceptance = "only_women"
)

func GetAcceptanceByProbability(onlyMen, menAndWomen, onlyWomen float64) (Acceptance, error) {
	out, err := mapTools.PickOneByProb(map[string]float64{
		string(OnlyMen):     randomTools.PrepareProbability(onlyMen),
		string(MenAndWomen): randomTools.PrepareProbability(menAndWomen),
		string(OnlyWomen):   randomTools.PrepareProbability(onlyWomen),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not generate acceptance (only_men=%f, men_and_woman=%f, only_women=%f)", onlyMen, menAndWomen, onlyWomen))
	}

	return Acceptance(out), nil
}

func SelectAcceptanceByMostRecent(in []Acceptance) (Acceptance, error) {
	m := make(map[Acceptance]int, 3)
	for _, mc := range in {
		if count, ok := m[mc]; ok {
			m[mc] = count + 1
		} else {
			m[mc] = 1
		}
	}
	probs := make(map[string]float64, 3)
	total := float64(len(in))
	for mc, count := range m {
		probs[mc.String()] = float64(count) / total
	}

	mc, err := mapTools.PickOneByProb(probs)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not select martial custom by most recent")
	}

	return Acceptance(mc), nil
}
