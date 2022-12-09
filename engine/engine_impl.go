package engine

import (
	"github.com/olehmushka/golang-toolkit/either"
	wGen "github.com/olehmushka/word-generator"
	"github.com/olehmushka/world-generator-engine/culture"
	"github.com/olehmushka/world-generator-engine/gender"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	"github.com/olehmushka/world-generator-engine/influence"
	"github.com/olehmushka/world-generator-engine/language"
	"github.com/olehmushka/world-generator-engine/religion"
	"github.com/olehmushka/world-generator-engine/types"
)

type engine struct {
	wordGenerator wGen.Generator
}

func New(wordGenerator wGen.Generator) Engine {
	return &engine{
		wordGenerator: wordGenerator,
	}
}

func (e *engine) LoadLanguageFamilies(opts ...types.ChangeStringFunc) chan either.Either[[]string] {
	return language.LoadAllFamilies(opts...)
}

func (e *engine) LoadLanguageSubfamilies(opts ...types.ChangeStringFunc) chan either.Either[[]*language.Subfamily] {
	return language.LoadAllSubfamilies(opts...)
}

func (e *engine) LoadLanguages(opts ...types.ChangeStringFunc) chan either.Either[*language.Language] {
	return language.LoadAllLanguages(opts...)
}

func (e *engine) GenerateWord(lang *language.Language) (string, error) {
	return e.wordGenerator.Generate(wGen.GenerateOpts{
		BaseName:  lang.Wordbase.Slug,
		BaseWords: lang.Wordbase.Words,
		Min:       lang.Wordbase.Min,
		Max:       lang.Wordbase.Max,
		Dupl:      lang.Wordbase.Dupl,
	})
}

func (e *engine) LoadGenders() []gender.Sex {
	return []gender.Sex{
		gender.MaleSex,
		gender.FemaleSex,
	}
}

func (e *engine) LoadGenderAcceptances() []genderAcceptance.Acceptance {
	return []genderAcceptance.Acceptance{
		genderAcceptance.OnlyMen,
		genderAcceptance.MenAndWomen,
		genderAcceptance.OnlyWomen,
	}
}

func (e *engine) LoadInfluences() []influence.Influence {
	return []influence.Influence{
		influence.StrongInfluence,
		influence.ModerateInfluence,
		influence.WeakInfluence,
	}
}

func (e *engine) LoadCulturesBases(opts ...types.ChangeStringFunc) chan either.Either[[]string] {
	return culture.LoadAllBases(opts...)
}

func (e *engine) LoadCultureSubbases(opts ...types.ChangeStringFunc) chan either.Either[[]*culture.Subbase] {
	return culture.LoadAllSubbases(opts...)
}

func (e *engine) LoadAllEthoses(opts ...types.ChangeStringFunc) chan either.Either[[]culture.Ethos] {
	return culture.LoadAllEthoses(opts...)
}

func (e *engine) LoadAllTraditions(opts ...types.ChangeStringFunc) chan either.Either[[]*culture.Tradition] {
	return culture.LoadAllTraditions(opts...)
}

func (e *engine) LoadAllParentRawCultures(opts ...types.ChangeStringFunc) chan either.Either[[]*culture.RawCulture] {
	return culture.LoadAllRawCultures(opts...)
}

func (e *engine) LoadAllParentCultures(opts ...types.ChangeStringFunc) chan either.Either[*culture.Culture] {
	return culture.LoadAllCultures(opts...)
}

func (e *engine) GenerateCulture(data *culture.CreateCultureOpts, parents ...*culture.Culture) (*culture.Culture, error) {
	return culture.Generate(data, parents...)
}

func (e *engine) LoadAllReligionTypes(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait] {
	return religion.LoadAllTypeTraits(opts...)
}

func (e *engine) LoadAllReligionHighGoals(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait] {
	return religion.LoadAllHighGoals(opts...)
}

func (e *engine) LoadAllReligionMarriageKinds(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait] {
	return religion.LoadAllMarriageKinds(opts...)
}

func (e *engine) LoadAllReligionBastardies(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait] {
	return religion.LoadAllBastardies(opts...)
}

func (e *engine) LoadAllReligionConsanguinities(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait] {
	return religion.LoadAllConsanguinities(opts...)
}

func (e *engine) LoadAllReligionDivorceOpts(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.PermissionTrait] {
	return religion.LoadAllDivorceOpts(opts...)
}

func (e *engine) GenerateReligion(opts religion.CreateReligionOpts, data religion.Data) (*religion.Religion, error) {
	return religion.New(opts, data)
}

func (e *engine) GenerateReligionByCulture(c *culture.Culture, opts religion.CreateReligionByCultureOpts, data religion.Data) (*religion.Religion, error) {
	generateOpts := religion.CreateReligionOpts{
		Opts: religion.CreateReligionTraitsOpts{
			Slug:            opts.Opts.Slug,
			Stats:           opts.Opts.Stats,
			MinHighGoalsNum: opts.Opts.MinHighGoalsNum,
			MaxHighGoalsNum: opts.Opts.MaxHighGoalsNum,
		},
		GenderDominance: c.GenderDominance,
	}
	if generateOpts.Opts.Stats == nil {
		generateOpts.Opts.Stats = &religion.Stats{}
	}

	return religion.New(generateOpts, data)
}
