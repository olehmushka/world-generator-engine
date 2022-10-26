package language

type Language struct {
	ID           string     `json:"id" bson:"id"`
	Name         string     `json:"name" bson:"name"`
	Subfamily    *Subfamily `json:"subfamily" bson:"subfamily,omitempty"`
	WordBaseName string     `json:"word_base_name" bson:"word_base_name"`
}
