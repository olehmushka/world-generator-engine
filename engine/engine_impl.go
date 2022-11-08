package engine

import (
	"github.com/olehmushka/golang-toolkit/either"
	wGen "github.com/olehmushka/word-generator"
	"github.com/olehmushka/world-generator-engine/culture"
	"github.com/olehmushka/world-generator-engine/gender"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	"github.com/olehmushka/world-generator-engine/influence"
	"github.com/olehmushka/world-generator-engine/language"
)

type engine struct {
	wordGenerator wGen.Generator
}

func New(wordGenerator wGen.Generator) Engine {
	return &engine{
		wordGenerator: wordGenerator,
	}
}

func (e *engine) LoadLanguageFamilies() chan either.Either[[]string] {
	return language.LoadAllFamilies()
}

func (e *engine) LoadLanguageSubfamilies() chan either.Either[[]*language.Subfamily] {
	return language.LoadAllSubfamilies()
}

func (e *engine) LoadLanguages() chan either.Either[*language.Language] {
	return language.LoadAllLanguages()
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

func (e *engine) LoadCulturesBases() chan either.Either[[]string] {
	return culture.LoadAllBases()
}

func (e *engine) LoadCultureSubbases() chan either.Either[[]*culture.Subbase] {
	return culture.LoadAllSubbases()
}

func (e *engine) LoadAllEthoses() chan either.Either[[]culture.Ethos] {
	return culture.LoadAllEthoses()
}

func (e *engine) LoadAllTraditions() chan either.Either[[]*culture.Tradition] {
	return culture.LoadAllTraditions()
}

func (e *engine) LoadAllParentRawCultures() chan either.Either[[]*culture.RawCulture] {
	return culture.LoadAllRawCultures()
}

func (e *engine) LoadAllParentCultures() chan either.Either[*culture.Culture] {
	return culture.LoadAllCultures()
}

func (e *engine) Generate(opts *culture.CreateCultureOpts, parents ...*culture.Culture) (*culture.Culture, error) {
	return culture.Generate(opts, parents...)
}
