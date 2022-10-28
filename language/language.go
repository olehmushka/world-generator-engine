package language

type RawLanguage struct {
	Slug          string `json:"slug" bson:"slug"`
	SubfamilySlug string `json:"subfamily_slug" bson:"subfamily_slug"`
	WordbaseSlug  string `json:"wordbase_slug" bson:"wordbase_slug"`
}

type Language struct {
	Slug      string     `json:"slug" bson:"slug"`
	Subfamily *Subfamily `json:"subfamily" bson:"subfamily"`
	Wordbase  *Wordbase  `json:"wordbase" bson:"wordbase"`
}
