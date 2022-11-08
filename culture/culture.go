package culture

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path"
	"runtime"
	"sync"

	"github.com/olehmushka/golang-toolkit/either"
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"
	"github.com/olehmushka/world-generator-engine/language"
)

type RawCulture struct {
	Slug               string                      `json:"slug" bson:"slug"`
	BaseSlug           string                      `json:"base_slug" bson:"base_slug"`
	SubbaseSlug        string                      `json:"subbase_slug" bson:"subbase_slug"`
	ParentCultureSlugs []string                    `json:"parent_culture_slugs" bson:"parent_culture_slugs"`
	EthosSlug          string                      `json:"ethos_slug" bson:"ethos_slug"`
	TraditionSlugs     []string                    `json:"tradition_slugs" bson:"tradition_slugs"`
	LanguageSlug       string                      `json:"language_slug" bson:"language_slug"`
	GenderDominance    genderDominance.Dominance   `json:"gender_dominance" bson:"gender_dominance"`
	MartialCustom      genderAcceptance.Acceptance `json:"martial_custom" bson:"martial_custom"`
}

type Culture struct {
	Slug            string                      `json:"slug" bson:"slug"`
	BaseSlug        string                      `json:"base_slug" bson:"base_slug"`
	Subbase         Subbase                     `json:"subbase" bson:"subbase"`
	ParentCultures  []*Culture                  `json:"parent_cultures" bson:"parent_cultures"`
	Ethos           Ethos                       `json:"ethos" bson:"ethos"`
	Traditions      []*Tradition                `json:"traditions" bson:"traditions"`
	LanguageSlug    string                      `json:"language_slug" bson:"language_slug"`
	GenderDominance genderDominance.Dominance   `json:"gender_dominance" bson:"gender_dominance"`
	MartialCustom   genderAcceptance.Acceptance `json:"martial_custom" bson:"martial_custom"`
}

type CreateCultureOpts struct {
	LanguageSlugs []string
	Subbases      []Subbase
	Ethoses       []Ethos
	Traditions    []*Tradition
}

// Generate This method generates culture from already generated cultures
// Or it can generate from data from opts passing via arguments
// The method does not generate slug for the culture. It should be done by word_generator (Please, use GenrateSlug func for preparing generated word for using it as culture's slug)
func Generate(opts *CreateCultureOpts, parents ...*Culture) (*Culture, error) {
	if len(parents) == 0 {
		c, err := New(opts)
		if err != nil {
			return nil, wrapped_error.NewInternalServerError(err, "can not generate random culture")
		}
		if c == nil {
			return nil, wrapped_error.NewBadRequestError(nil, "can not genererate culture with options")
		}

		return c, nil
	}
	c := &Culture{
		ParentCultures: parents,
	}
	gd, err := genderDominance.SelectGenderDominanceByMostRecent(ExtractGenderDominances(parents))
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not select dominated gender for generated culture")
	}
	c.GenderDominance = gd

	mc, err := genderAcceptance.SelectAcceptanceByMostRecent(ExtractMartialCusoms(parents))
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not select martial custom for generated culture")
	}
	c.MartialCustom = mc

	e, err := SelectEthosByMostRecent(ExtractEthoses(parents))
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not select ethos for generated culture")
	}
	c.Ethos = e

	ts, err := RandomTraditions(UniqueTraditions(ExtractTraditions(parents)), 2, 7, e, gd.DominatedSex)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not select traditions for generated culture")
	}
	c.Traditions = ts

	subbase, err := SelectSubbaseByMostRecent(ExtractSubbases(parents))
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not select subbase for generated culture")
	}
	c.Subbase = subbase
	c.BaseSlug = subbase.BaseSlug

	langSlug, err := language.SelectLanguageSlugByMostRecent(ExtractLanguageSlugs(parents))
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate language for random culture")
	}
	c.LanguageSlug = langSlug

	return c, nil
}

func New(opts *CreateCultureOpts) (*Culture, error) {
	if opts == nil {
		return nil, nil
	}

	out := &Culture{}
	gd, err := genderDominance.GetRandom()
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate dominated gender for random culture")
	}
	out.GenderDominance = gd

	mc, err := RandomMartialCustom(gd)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate martial custom for random culture")
	}
	out.MartialCustom = mc

	e, err := RandomEthos(opts.Ethoses)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate ethos for random culture")
	}
	out.Ethos = e

	ts, err := RandomTraditions(opts.Traditions, 2, 7, e, gd.DominatedSex)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate traditions for random culture")
	}
	out.Traditions = ts

	subbase, err := RandomSubbase(opts.Subbases)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate subbase for random culture")
	}
	out.Subbase = subbase
	out.BaseSlug = subbase.BaseSlug

	langSlug, err := language.RandomLanguageSlug(opts.LanguageSlugs)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate language for random culture")
	}
	out.LanguageSlug = langSlug

	return out, nil
}

func GenerateSlug(word string) string {
	return fmt.Sprintf("%sian_%s%s", word, randomTools.UUIDString(), RequiredCultureSlugSuffix)
}

func LoadAllRawCultures() chan either.Either[[]*RawCulture] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/cultures/"
	ch := make(chan either.Either[[]*RawCulture], MaxLoadDataConcurrency)
	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[[]*RawCulture]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))}
			return
		}

		var wg sync.WaitGroup
		wg.Add(len(files))
		for _, file := range files {
			go func(file fs.FileInfo) {
				defer wg.Done()
				if file.IsDir() {
					return
				}
				filename := dirname + file.Name()
				b, err := ioutil.ReadFile(filename)
				if err != nil {
					ch <- either.Either[[]*RawCulture]{Err: err}
					return
				}
				var sfs []*RawCulture
				if err := json.Unmarshal(b, &sfs); err != nil {
					ch <- either.Either[[]*RawCulture]{Err: err}
					return
				}
				if len(sfs) == 0 {
					return
				}
				for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, sfs) {
					ch <- either.Either[[]*RawCulture]{Value: chunk}
				}
			}(file)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}
