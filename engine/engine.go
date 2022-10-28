package engine

import "github.com/olehmushka/world-generator-engine/language"

type Engine interface {
	LoadLanguageFamilies() ([]string, error)
	LoadLanguageSubfamilies() ([]*language.Subfamily, error)
	GenerateWord(lang *language.Language) (string, error)
}
