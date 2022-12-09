package favour

import (
	"math"

	randomtools "github.com/olehmushka/golang-toolkit/random_tools"
	slicetools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

type Favour string

func (f Favour) String() string {
	return string(f)
}

func (f Favour) Float64() float64 {
	switch f {
	case Loved:
		return 1
	case Cherished:
		return 0.9
	case Honored:
		return 0.8
	case Praised:
		return 0.7
	case Favored:
		return 0.6
	case Respected:
		return 0.5
	case Liked:
		return 0.35
	case Tolerated:
		return 0.2
	case Ignored:
		return 0
	case Shunned:
		return -0.2
	case Disliked:
		return -0.35
	case Dishonored:
		return -0.5
	case Disowned:
		return -0.6
	case Abandoned:
		return -0.7
	case Despited:
		return -0.8
	case Hated:
		return -0.9
	case Damned:
		return -1
	}
	return InvalidFloat64Value
}

func (f Favour) Positive() bool {
	return f.Float64() > 0
}

func (f Favour) Zero() bool {
	return f.Float64() == 0
}

func (f Favour) Negative() bool {
	return f.Float64() < 0
}

const (
	Loved      Favour = "loved_favour"
	Cherished  Favour = "cherished_favour"
	Honored    Favour = "honored_favour"
	Praised    Favour = "praised_favour"
	Favored    Favour = "favored_favour"
	Respected  Favour = "respected_favour"
	Liked      Favour = "liked_favour"
	Tolerated  Favour = "tolerated_favour"
	Ignored    Favour = "ignored_favour"
	Shunned    Favour = "shunned_favour"
	Disliked   Favour = "disliked_favour"
	Dishonored Favour = "dishonored_favour"
	Disowned   Favour = "disowned_favour"
	Abandoned  Favour = "abandoned_favour"
	Despited   Favour = "despited_favour"
	Hated      Favour = "hated_favour"
	Damned     Favour = "damned_favour"
)

var AllFavours = []Favour{
	Loved,
	Cherished,
	Honored,
	Praised,
	Favored,
	Respected,
	Liked,
	Tolerated,
	Ignored,
	Shunned,
	Disliked,
	Dishonored,
	Disowned,
	Abandoned,
	Despited,
	Hated,
	Damned,
}

func ParseFloat64(v float64) Favour {
	return ParseFloat64Of(AllFavours, v)
}

func ParseFloat64Of(list []Favour, v float64) Favour {
	if v > Max.Float64() {
		return Max
	}
	if v < Min.Float64() {
		return Min
	}

	var (
		minDiff = math.MaxFloat64
		f       Favour
	)
	for _, el := range slicetools.Shuffle(list) {
		if diff := math.Abs(el.Float64() - v); minDiff > diff {
			minDiff = diff
			f = el
		}
	}

	return f
}

func IsValid(v string) bool {
	for _, f := range AllFavours {
		if f.String() == v {
			return true
		}
	}

	return false
}

func GenerateFavourOf(list []Favour, positivity, dev float64) (Favour, error) {
	v, err := randomtools.RandFloat64NormInRange(Min.Float64(), Max.Float64(), dev, positivity)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate favour")
	}

	return ParseFloat64Of(list, v), nil
}

func GenerateFavour(positivity, dev float64) (Favour, error) {
	v, err := randomtools.RandFloat64NormInRange(Min.Float64(), Max.Float64(), dev, positivity)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not generate favour")
	}

	return ParseFloat64(v), nil
}
