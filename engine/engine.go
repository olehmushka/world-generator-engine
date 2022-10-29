package engine

import (
	"github.com/olehmushka/golang-toolkit/either"
	"github.com/olehmushka/world-generator-engine/language"
)

type Engine interface {
	LoadLanguageFamilies() chan either.Either[[]string]
	LoadLanguageSubfamilies() chan either.Either[[]*language.Subfamily]
	LoadLanguages() chan either.Either[*language.Language]
	GenerateWord(lang *language.Language) (string, error)
}
