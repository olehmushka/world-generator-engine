package culture

import genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"

var mockCultures = []*Culture{
	{
		Slug:     "slavic_ilmenian_culture",
		BaseSlug: "europe_continental_base",
		Subbase: Subbase{
			Slug:     "medieval_east_slavic_subbase",
			BaseSlug: "europe_continental_base",
		},
		Ethos: Ethos{Slug: "communal_ethos"},
		Traditions: []*Tradition{
			{Slug: "agrarian_tradition"},
			{Slug: "druzhina_tradition"},
			{Slug: "forest_folk_tradition"},
			{Slug: "hit_and_run_tacticians_tradition"},
		},
		LanguageSlug: "ruthenian_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "slavic_ruthenian_culture",
		BaseSlug: "europe_continental_base",
		Subbase: Subbase{
			Slug:     "medieval_east_slavic_subbase",
			BaseSlug: "europe_continental_base",
		},
		Ethos: Ethos{Slug: "courtly_ethos"},
		Traditions: []*Tradition{
			{Slug: "agrarian_tradition"},
			{Slug: "druzhina_tradition"},
			{Slug: "charismatic_tradition"},
			{Slug: "city_keepers_tradition"},
		},
		LanguageSlug: "ruthenian_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "slavic_severian_culture",
		BaseSlug: "europe_continental_base",
		Subbase: Subbase{
			Slug:     "medieval_east_slavic_subbase",
			BaseSlug: "europe_continental_base",
		},
		Ethos: Ethos{Slug: "stoic_ethos"},
		Traditions: []*Tradition{
			{Slug: "druzhina_tradition"},
			{Slug: "forest_folk_tradition"},
			{Slug: "sacred_groves_tradition"},
		},
		LanguageSlug: "ruthenian_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "slavic_volhynian_culture",
		BaseSlug: "europe_continental_base",
		Subbase: Subbase{
			Slug:     "medieval_east_slavic_subbase",
			BaseSlug: "europe_continental_base",
		},
		Ethos: Ethos{Slug: "courtly_ethos"},
		Traditions: []*Tradition{
			{Slug: "agrarian_tradition"},
			{Slug: "druzhina_tradition"},
			{Slug: "astute_diplomats_tradition"},
			{Slug: "city_keepers_tradition"},
		},
		LanguageSlug: "ruthenian_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "latin_roman_culture",
		BaseSlug: "mediterranean_base",
		Subbase:  Subbase{Slug: "medieval_latin_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "bellicose_ethos"},
		Traditions: []*Tradition{
			{Slug: "formation_fighting_experts_tradition"},
			{Slug: "hereditary_hierarchy_tradition"},
			{Slug: "legalistic_tradition"},
			{Slug: "refined_poetry_tradition"},
		},
		LanguageSlug: "latin_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "latin_italian_culture",
		BaseSlug: "mediterranean_base",
		Subbase:  Subbase{Slug: "medieval_latin_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "spiritual_ethos"},
		Traditions: []*Tradition{
			{Slug: "formation_fighting_experts_tradition"},
			{Slug: "refined_poetry_tradition"},
			{Slug: "republican_legacy_tradition"},
		},
		LanguageSlug: "italian_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "latin_sardinian_culture",
		BaseSlug: "mediterranean_base",
		Subbase:  Subbase{Slug: "medieval_latin_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "communal_ethos"},
		Traditions: []*Tradition{
			{Slug: "isolationist_tradition"},
			{Slug: "stalwart_defenders_tradition"},
		},
		LanguageSlug: "sardinian_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "latin_lombard_culture",
		BaseSlug: "mediterranean_base",
		Subbase:  Subbase{Slug: "medieval_latin_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "stoic_ethos"},
		Traditions: []*Tradition{
			{Slug: "isolationist_tradition"},
			{Slug: "martial_admiration_tradition"},
			{Slug: "republican_legacy_tradition"},
			{Slug: "stand_and_fight_tradition"},
		},
		LanguageSlug: "lombard_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "latin_cisalpine_culture",
		BaseSlug: "mediterranean_base",
		Subbase:  Subbase{Slug: "medieval_latin_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "communal_ethos"},
		Traditions: []*Tradition{
			{Slug: "martial_admiration_tradition"},
			{Slug: "republican_legacy_tradition"},
			{Slug: "mountain_homes_tradition"},
		},
		LanguageSlug: "italian_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "latin_sicilian_culture",
		BaseSlug: "mediterranean_base",
		Subbase:  Subbase{Slug: "medieval_latin_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "courtly_ethos"},
		Traditions: []*Tradition{
			{Slug: "republican_legacy_tradition"},
			{Slug: "refined_poetry_tradition"},
			{Slug: "ruling_caste_tradition"},
			{Slug: "swords_for_hire_tradition"},
			{Slug: "xenophilic_tradition"},
		},
		LanguageSlug: "sicilian_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "frankish_frankish_culture",
		BaseSlug: "europe_continental_base",
		Subbase:  Subbase{Slug: "medieval_frankish_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "bellicose_ethos"},
		Traditions: []*Tradition{
			{Slug: "hereditary_hierarchy_tradition"},
			{Slug: "stand_and_fight_tradition"},
			{Slug: "warrior_culture_tradition"},
		},
		LanguageSlug: "frankish_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "frankish_french_culture",
		BaseSlug: "europe_continental_base",
		Subbase:  Subbase{Slug: "medieval_frankish_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "courtly_ethos"},
		Traditions: []*Tradition{
			{Slug: "hereditary_hierarchy_tradition"},
			{Slug: "chanson_de_geste_tradition"},
			{Slug: "chivalry_tradition"},
		},
		LanguageSlug: "french_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "frankish_norman_culture",
		BaseSlug: "europe_continental_base",
		Subbase:  Subbase{Slug: "medieval_frankish_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "bellicose_ethos"},
		Traditions: []*Tradition{
			{Slug: "hereditary_hierarchy_tradition"},
			{Slug: "chanson_de_geste_tradition"},
			{Slug: "stand_and_fight_tradition"},
		},
		LanguageSlug: "french_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "male_sex",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
	{
		Slug:     "frankish_occitan_culture",
		BaseSlug: "europe_continental_base",
		Subbase:  Subbase{Slug: "medieval_frankish_subbase", BaseSlug: "europe_continental_base"},
		Ethos:    Ethos{Slug: "egalitarian_ethos"},
		Traditions: []*Tradition{
			{Slug: "chanson_de_geste_tradition"},
			{Slug: "chivalry_tradition"},
			{Slug: "equal_inheritance_tradition"},
			{Slug: "parochialism_tradition"},
		},
		LanguageSlug: "occitan_lang",
		GenderDominance: genderDominance.Dominance{
			DominatedSex: "",
			Influence:    "moderate_influence",
		},
		MartialCustom: "only_men",
	},
}
