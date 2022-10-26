package wordbase

type WordBase struct {
	FemaleOwnNames []string `json:"female_own_names" bson:"female_own_names"`
	MaleOwnNames   []string `json:"male_own_names" bson:"male_own_names"`
	Words          []string `json:"words" bson:"words"`
	Name           string   `json:"name" bson:"name"`
	Min            int      `json:"min" bson:"min"`
	Max            int      `json:"max" bson:"max"`
	Dupl           string   `json:"dupl" bson:"dupl"`
	M              float64  `json:"m" bson:"m"`
}
