package genderdominance

import (
	"github.com/olehmushka/world-generator-engine/gender"
	"github.com/olehmushka/world-generator-engine/influence"
)

type Dominance struct {
	DominatedSex gender.Sex          `json:"dominated_sex" bson:"dominated_sex"`
	Influence    influence.Influence `json:"influence" bson:"influence"`
}
