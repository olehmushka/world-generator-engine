package religion

import (
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	we "github.com/olehmushka/golang-toolkit/wrapped_error"
	genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"
)

type Religion struct {
	Stats             *Stats                    `json:"stats"`
	Slug              string                    `json:"slug"`
	Type              *Trait                    `json:"type"`
	GenderDominance   genderDominance.Dominance `json:"gender_dominance"`
	Traits            []*Trait                  `json:"traits"`
	MarriageTradition MarriageTradition         `json:"marriage_tradition"`
	Afterlife         Afterlife                 `json:"afterlife"`
	DeityFavour       *FavourTrait              `json:"deity_favour"`
	SourceOfMoralLaw  *Trait                    `json:"source_of_moral_law"`
}

func New(opts CreateReligionOpts, data Data) (*Religion, error) {
	out := &Religion{
		Slug:            opts.Opts.Slug,
		GenderDominance: opts.GenderDominance,
		Stats:           opts.Opts.Stats,
	}
	if out.GenderDominance.IsZero() {
		gd, err := genderDominance.GetRandom()
		if err != nil {
			return nil, we.NewInternalServerError(err, "can not generate gender dominance for generated religion")
		}
		out.GenderDominance = gd
	}
	if out.Stats == nil {
		out.Stats = &Stats{}
	}

	// cfg generation
	cfg, err := NewStatsConfig()
	if err != nil {
		return nil, we.NewInternalServerError(err, "can not create stats config for generated religion")
	}
	// religion type generation
	t, stats, err := NewType(data.Types, cfg, out)
	if err != nil {
		return nil, we.NewInternalServerError(err, "can not generate type for generated religion")
	}
	out.Stats = stats
	out.Type = t

	for _, brick := range sliceTools.Shuffle(ReligionConstructBricks) {
		var brickErr error
		out, brickErr = brick(cfg, opts, data, out)
		if brickErr != nil {
			return nil, brickErr
		}
	}

	al, alStats, err := NewAfterlife(cfg, out, CreateAfterlifeData{
		AfterlifeParticipances: data.AfterlifeParticipances,
		AfterlifeParticipants:  data.AfterlifeParticipants,
		AllDeityFavours:        data.DeityFavourTraits,
		AllAfterlifeExistsOpts: data.AfterlifeExistOpts,
	})
	if err != nil {
		return nil, we.NewInternalServerError(err, "can not generate afterlife for generated religion")
	}
	out.Stats = alStats
	out.Afterlife = al

	return out, nil
}

type ReligionConstructBrick func(cfg StatsConfig, opts CreateReligionOpts, data Data, r *Religion) (*Religion, error)

var ReligionConstructBricks = []ReligionConstructBrick{
	func(cfg StatsConfig, opts CreateReligionOpts, data Data, r *Religion) (*Religion, error) {
		traits, stats, err := NewHighGoals(cfg, r, data.HighGoals, opts.Opts.MinHighGoalsNum, opts.Opts.MaxHighGoalsNum)
		if err != nil {
			return nil, we.NewInternalServerError(err, "can not generate high goals for generated religion")
		}
		r.Stats = stats
		r.Traits = append(r.Traits, traits...)

		return r, nil
	},
	func(cfg StatsConfig, opts CreateReligionOpts, data Data, r *Religion) (*Religion, error) {
		traits, stats, err := NewSocialTraits(cfg, r, data.SocialTraits, opts.Opts.MinSocialTraitsNum, opts.Opts.MaxSocialTraitsNum)
		if err != nil {
			return nil, we.NewInternalServerError(err, "can not generate social traits for generated religion")
		}
		r.Stats = stats
		r.Traits = append(r.Traits, traits...)

		return r, nil
	},
	func(cfg StatsConfig, opts CreateReligionOpts, data Data, r *Religion) (*Religion, error) {
		deityFavour, stats, err := NewDeityFavour(data.DeityFavourTraits, cfg, r)
		if err != nil {
			return nil, we.NewInternalServerError(err, "can not generate deity favour for generated religion")
		}
		r.Stats = stats
		r.DeityFavour = deityFavour

		return r, nil
	},
	func(cfg StatsConfig, opts CreateReligionOpts, data Data, r *Religion) (*Religion, error) {
		traits, stats, err := NewDeityNatureTraits(cfg, r, data.DeityNatureTraits, opts.Opts.MinDeityNatureTraitsNum, opts.Opts.MaxDeityNatureTraitsNum)
		if err != nil {
			return nil, we.NewInternalServerError(err, "can not generate deity nature traits for generated religion")
		}
		r.Stats = stats
		r.Traits = append(r.Traits, traits...)

		return r, nil
	},
	func(cfg StatsConfig, opts CreateReligionOpts, data Data, r *Religion) (*Religion, error) {
		traits, stats, err := NewHumanNatureTraits(cfg, r, data.DeityNatureTraits, opts.Opts.MinHumanNatureTraitsNum, opts.Opts.MaxHumanNatureTraitsNum)
		if err != nil {
			return nil, we.NewInternalServerError(err, "can not generate human nature traits for generated religion")
		}
		r.Stats = stats
		r.Traits = append(r.Traits, traits...)

		return r, nil
	},
	func(cfg StatsConfig, opts CreateReligionOpts, data Data, r *Religion) (*Religion, error) {
		mt, stats, err := NewMarriageTradition(cfg, r, CreateMarriageTraditionData{
			MarriageKinds:           data.MarriageKinds,
			BastardyTraditions:      data.BastardyTraditions,
			ConsanguinityTraditions: data.ConsanguinityTraditions,
			DivorceTraditions:       data.DivorceTraditions,
		})
		if err != nil {
			return nil, we.NewInternalServerError(err, "can not generate marriage tradition for generated religion")
		}
		r.Stats = stats
		r.MarriageTradition = mt

		return r, nil
	},
	func(cfg StatsConfig, opts CreateReligionOpts, data Data, r *Religion) (*Religion, error) {
		soml, stats, err := NewSourceOfMoralLaw(data.SourceOfMoralLawTraits, cfg, r)
		if err != nil {
			return nil, we.NewInternalServerError(err, "can not generate deity nature traits for generated religion")
		}
		r.Stats = stats
		r.SourceOfMoralLaw = soml

		return r, nil
	},
}

/*
Doctrine
	Afterlife        *Afterlife       `json:"afterlife" bson:"afterlife"`
		IsExists     bool                   `json:"is_exists" bson:"is_exists"`
    Participants *AfterlifeParticipants `json:"participants" bson:"participants"`
				ForTopBelievers    AfterlifeOption `json:"for_top_belivers" bson:"for_top_belivers"`
				ForBelievers       AfterlifeOption `json:"for_belivers" bson:"for_belivers"`
				ForUntrueBelievers AfterlifeOption `json:"for_untrue_belivers" bson:"for_untrue_belivers"`
				ForAtheists        AfterlifeOption `json:"for_atheists" bson:"for_atheists"`
    Traits       []*trait
Attributes
	Traits        []*trait       `json:"traits" bson:"traits"`
	Clerics       *Clerics       `json:"clerics" bson:"clerics"`
		HasClerics  bool                `json:"has_clerics" bson:"has_clerics"`
		Appointment *ClericsAppointment `json:"appointment" bson:"appointment"`
			IsCivil     bool `json:"is_civil" bson:"is_civil"`
			IsRevocable bool `json:"is_revocable" bson:"is_revocable"`
		Limitations *ClericsLimitations `json:"limitations" bson:"limitations"`
			AcceptableGender g.Acceptance `json:"acceptable" bson:"acceptable"`
			Marriage         Permission   `json:"marriage" bson:"marriage"`
		Traits      []*trait            `json:"traits" bson:"traits"`
		Functions   []*trait            `json:"functions" bson:"functions"`
	HolyScripture *HolyScripture `json:"holy_scripture" bson:"holy_scripture"`
		HasHolyScripture bool     `json:"has_holy_scripture" bson:"has_holy_scripture"`
		Traits           []*trait `json:"traits" bson:"traits"`
	Temples       *Temples       `json:"temples" bson:"temples"`
		HasTemples      bool     `json:"has_temples"  bson:"has_temples"`
		HasSacredPlaces bool     `json:"has_sacred_places"  bson:"has_sacred_places"`
		Traits          []*trait `json:"traits"  bson:"traits"`
Theology
	Traits            []*trait           `json:"traits" bson:"traits"`
	Cults             []*trait           `json:"cults" bson:"cults"`
	Rules             *Rules             `json:"rules" bson:"rules"`
		Rules []*trait `json:"rules" bson:"rules"`
	Taboos            *Taboos            `json:"taboos" bson:"taboos"`
		Taboos []*Taboo `json:"taboos" bson:"taboos"`
	Rituals           *Rituals           `json:"rituals" bson:"rituals"`
		Initiation []*trait `json:"initiations" bson:"initiations"`
		Funeral    []*trait `json:"funeral" bson:"funeral"`
		Sacrifice  []*trait `json:"sacrifice" bson:"sacrifice"`
		Holyday    []*trait `json:"holyday" bson:"holyday"`
	Holydays          *Holydays          `json:"holydays" bson:"holydays"`
		Holydays []*trait `json:"holydays" bson:"holydays"`
	Conversion        *Conversion        `json:"conversion" bson:"conversion"`
		Traits []*trait `json:"traits" bson:"traits"`
*/

func ExtractTraitSlugs(r *Religion) []string {
	if r == nil {
		return []string{}
	}

	out := make([]string, 0, len(r.Traits))
	for _, t := range r.Traits {
		out = append(out, t.Slug)
	}
	if traits := r.MarriageTradition.ExtractTraitSlugs(); len(traits) != 0 {
		out = append(out, traits...)
	}

	return out
}

type PureReligion struct {
	Slug              string                    `json:"slug"`
	Type              *PureTrait                `json:"type"`
	GenderDominance   genderDominance.Dominance `json:"gender_dominance"`
	Traits            []*PureTrait              `json:"traits"`
	MarriageTradition PureMarriageTradition     `json:"marriage_tradition"`
	Afterlife         Afterlife                 `json:"afterlife"`
	DeityFavour       *PureTrait                `json:"deity_favour"`
	SourceOfMoralLaw  *PureTrait                `json:"source_of_moral_law"`
}

func PurifyReligion(in *Religion) *PureReligion {
	if in == nil {
		return nil
	}

	return &PureReligion{
		Slug:              in.Slug,
		Type:              PurifyTrait(in.Type),
		GenderDominance:   genderDominance.Dominance{},
		Traits:            PurifyTraits(in.Traits),
		MarriageTradition: PurifyMarriageTradition(in.MarriageTradition),
		Afterlife:         in.Afterlife,
		DeityFavour:       PurifyFavourTrait(in.DeityFavour),
		SourceOfMoralLaw:  PurifyTrait(in.SourceOfMoralLaw),
	}
}
