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

func (a Acceptance) IsAccepted() bool {
	return a == Accepted
}

func (a Acceptance) IsShunned() bool {
	return a == Shunned
}

func (a Acceptance) IsCriminal() bool {
	return a == Criminal
}

const (
	Accepted Acceptance = "accepted"
	Shunned  Acceptance = "shunned"
	Criminal Acceptance = "criminal"
)

func GetAcceptanceByProb(accepted, shunned, criminal float64) (Acceptance, error) {
	out, err := mapTools.PickOneByProb(map[string]float64{
		string(Accepted): randomTools.PrepareProbability(accepted),
		string(Shunned):  randomTools.PrepareProbability(shunned),
		string(Criminal): randomTools.PrepareProbability(criminal),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not generate acceptance (accepted=%f, shunned=%f, criminal=%f)", accepted, shunned, criminal))
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
		return "", wrapped_error.NewInternalServerError(err, "can not select acceptance by most recent")
	}

	return Acceptance(mc), nil
}
