package engine

import (
	wGen "github.com/olehmushka/word-generator"
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

func (e *engine) LoadLanguageFamilies() ([]string, error) {
	return language.LoadAllFamilies()
}

func (e *engine) LoadLanguageSubfamilies() ([]*language.Subfamily, error) {
	return language.LoadAllSubfamilies()
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
