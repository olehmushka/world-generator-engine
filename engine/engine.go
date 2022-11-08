package engine

import (
	"github.com/olehmushka/golang-toolkit/either"
	"github.com/olehmushka/world-generator-engine/culture"
	"github.com/olehmushka/world-generator-engine/gender"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	"github.com/olehmushka/world-generator-engine/influence"
	"github.com/olehmushka/world-generator-engine/language"
)

type Engine interface {
	LoadLanguageFamilies() chan either.Either[[]string]
	LoadLanguageSubfamilies() chan either.Either[[]*language.Subfamily]
	LoadLanguages() chan either.Either[*language.Language]
	GenerateWord(lang *language.Language) (string, error)
	LoadGenders() []gender.Sex
	LoadGenderAcceptances() []genderAcceptance.Acceptance
	LoadInfluences() []influence.Influence
	LoadCulturesBases() chan either.Either[[]string]
	LoadCultureSubbases() chan either.Either[[]*culture.Subbase]
	LoadAllEthoses() chan either.Either[[]culture.Ethos]
	LoadAllTraditions() chan either.Either[[]*culture.Tradition]
	LoadAllParentRawCultures() chan either.Either[[]*culture.RawCulture]
	Generate(opts *culture.CreateCultureOpts, parents ...*culture.Culture) (*culture.Culture, error)
}
