package religion

const (
	MaxLoadDataChunkSize   = 100
	MaxLoadDataConcurrency = 10

	RequiredCoreTypeSlugSuffix             = "_core_type"
	RequiredTypeSlugSuffix                 = "_type_trait"
	RequiredMarriageKindSlugSuffix         = "_marriage_kind_trait"
	RequireBastardySlugSuffix              = "_bastardy_trait"
	RequireConsanguinitySlugSuffix         = "_consanguinity_trait"
	RequireHighGoalSlugSuffix              = "_high_goal_trait"
	RequireDeityNatureSlugSuffix           = "_deity_nature_trait"
	RequireHumanNatureSlugSuffix           = "_human_nature_trait"
	RequireSocialSlugSuffix                = "_social_trait"
	RequireSourceOfMoralLawSlugSuffix      = "_source_of_moral_law_trait"
	RequireAfterlifeParticipanceSlugSuffix = "_afterlife_participance"
	RequireAfterlifeParticipantSlugSuffix  = "_afterlife_participant"
)

const (
	MarriageTraitBastardyKind      = "marriage_trait_bastardy_kind"
	MarriageTraitConsanguinityKind = "marriage_trait_consanguinity_kind"
	MarriageKindTraitKind          = "marriage_kind_trait_kind"
)
