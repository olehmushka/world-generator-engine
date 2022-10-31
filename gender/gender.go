package gender

import (
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

type Sex string

const (
	MaleSex   Sex = "male_sex"
	FemaleSex Sex = "female_sex"
)

func (s Sex) String() string {
	return string(s)
}

func (s Sex) IsMale() bool {
	return s == MaleSex
}

func (s Sex) IsFemale() bool {
	return s == FemaleSex
}

func GetRandomSex() (Sex, error) {
	isMale, err := randomTools.GetRandomBool(0.5)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate bool for generation random sex")
	}
	if isMale {
		return MaleSex, nil
	}

	return FemaleSex, nil
}
