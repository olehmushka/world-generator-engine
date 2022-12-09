package engine

import (
	"github.com/olehmushka/golang-toolkit/either"
	"github.com/olehmushka/world-generator-engine/culture"
	"github.com/olehmushka/world-generator-engine/gender"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	"github.com/olehmushka/world-generator-engine/influence"
	"github.com/olehmushka/world-generator-engine/language"
	"github.com/olehmushka/world-generator-engine/religion"
	"github.com/olehmushka/world-generator-engine/types"
)

type Engine interface {
	LoadLanguageFamilies(opts ...types.ChangeStringFunc) chan either.Either[[]string]
	LoadLanguageSubfamilies(opts ...types.ChangeStringFunc) chan either.Either[[]*language.Subfamily]
	LoadLanguages(opts ...types.ChangeStringFunc) chan either.Either[*language.Language]
	GenerateWord(lang *language.Language) (string, error)
	LoadGenders() []gender.Sex
	LoadGenderAcceptances() []genderAcceptance.Acceptance
	LoadInfluences() []influence.Influence
	LoadCulturesBases(opts ...types.ChangeStringFunc) chan either.Either[[]string]
	LoadCultureSubbases(opts ...types.ChangeStringFunc) chan either.Either[[]*culture.Subbase]
	LoadAllEthoses(opts ...types.ChangeStringFunc) chan either.Either[[]culture.Ethos]
	LoadAllTraditions(opts ...types.ChangeStringFunc) chan either.Either[[]*culture.Tradition]
	LoadAllParentRawCultures(opts ...types.ChangeStringFunc) chan either.Either[[]*culture.RawCulture]
	LoadAllParentCultures(opts ...types.ChangeStringFunc) chan either.Either[*culture.Culture]
	GenerateCulture(opts *culture.CreateCultureOpts, parents ...*culture.Culture) (*culture.Culture, error)
	LoadAllReligionTypes(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait]
	LoadAllReligionHighGoals(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait]
	LoadAllReligionMarriageKinds(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait]
	LoadAllReligionBastardies(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait]
	LoadAllReligionConsanguinities(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.Trait]
	LoadAllReligionDivorceOpts(opts ...types.ChangeStringFunc) chan either.Either[[]*religion.PermissionTrait]
	GenerateReligion(opts religion.CreateReligionOpts, data religion.Data) (*religion.Religion, error)
	GenerateReligionByCulture(c *culture.Culture, opts religion.CreateReligionByCultureOpts, data religion.Data) (*religion.Religion, error)
}
