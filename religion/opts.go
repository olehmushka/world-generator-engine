package religion

import genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"

type CreateReligionTraitsOpts struct {
	Slug  string `json:"slug"`
	Stats *Stats `json:"stats"`

	MinHighGoalsNum         int
	MaxHighGoalsNum         int
	MinDeityNatureTraitsNum int
	MaxDeityNatureTraitsNum int
	MinHumanNatureTraitsNum int
	MaxHumanNatureTraitsNum int
	MinSocialTraitsNum      int
	MaxSocialTraitsNum      int
}

type CreateReligionOpts struct {
	GenderDominance genderDominance.Dominance `json:"gender_dominance"`
	Opts            CreateReligionTraitsOpts  `json:"opts"`
}

type CreateReligionByCultureOpts struct {
	Opts CreateReligionTraitsOpts `json:"opts"`
}

type Data struct {
	Types                   []*Trait                `json:"types"`
	HighGoals               []*Trait                `json:"high_goals"`
	SocialTraits            []*Trait                `json:"social_traits"`
	MarriageKinds           []*Trait                `json:"marriage_kinds"`
	BastardyTraditions      []*Trait                `json:"bastardy_traditions"`
	ConsanguinityTraditions []*Trait                `json:"consanguinity_traditions"`
	DivorceTraditions       []*PermissionTrait      `json:"divorce_traditions"`
	AfterlifeParticipances  []AfterlifeParticipance `json:"afterlife_participances"`
	AfterlifeParticipants   []AfterlifeParticipant  `json:"afterlife_participants"`
	AfterlifeExistOpts      []AfterlifeExist        `json:"afterlife_exist_opts"`
	DeityFavourTraits       []*FavourTrait          `json:"deity_favour_traits"`
	DeityNatureTraits       []*Trait                `json:"deity_nature_traits"`
	HumanNatureTraits       []*Trait                `json:"human_nature_traits"`
	SourceOfMoralLawTraits  []*Trait                `json:"source_of_moral_law_traits"`
}
