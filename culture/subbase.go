package culture

type Subbase struct {
	Slug     string `json:"slug" bson:"slug"`
	BaseSlug string `json:"base_slug" bson:"base_slug"`
}
