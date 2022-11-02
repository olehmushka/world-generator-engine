package culture

import (
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"
)

type RawCulture struct {
	Slug               string                      `json:"slug" bson:"slug"`
	BaseSlug           string                      `json:"base_slug" bson:"base_slug"`
	SubbaseSlug        string                      `json:"subbase_slug" bson:"subbase_slug"`
	ParentCultureSlugs []string                    `json:"parent_culture_slugs" bson:"parent_culture_slugs"`
	EthosSlug          string                      `json:"ethos_slug" bson:"ethos_slug"`
	TraditionSlugs     []string                    `json:"tradition_slugs" bson:"tradition_slugs"`
	LanguageSlug       string                      `json:"language_slug" bson:"language_slug"`
	GenderDominance    genderDominance.Dominance   `json:"gender_dominance" bson:"gender_dominance"`
	MartialCustom      genderAcceptance.Acceptance `json:"martial_custom" bson:"martial_custom"`
}

type Culture struct {
	Slug            string                      `json:"slug" bson:"slug"`
	BaseSlug        string                      `json:"base_slug" bson:"base_slug"`
	Subbase         Subbase                     `json:"subbase" bson:"subbase"`
	ParentCultures  []*Culture                  `json:"parent_cultures" bson:"parent_cultures"`
	Ethos           Ethos                       `json:"ethos" bson:"ethos"`
	Traditions      []*Tradition                `json:"traditions" bson:"traditions"`
	LanguageSlug    string                      `json:"language_slug" bson:"language_slug"`
	GenderDominance genderDominance.Dominance   `json:"gender_dominance" bson:"gender_dominance"`
	MartialCustom   genderAcceptance.Acceptance `json:"martial_custom" bson:"martial_custom"`
}
