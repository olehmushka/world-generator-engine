package language

type Subfamily struct {
	Name              string     `json:"name" bson:"name"`
	FamilyName        string     `json:"family_name" bson:"family_name"`
	ExtendedSubfamily *Subfamily `json:"extended_subfamily" bson:"extended_subfamily"`
}
