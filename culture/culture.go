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
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	genderAcceptance "github.com/olehmushka/world-generator-engine/gender_acceptance"
	genderDominance "github.com/olehmushka/world-generator-engine/gender_dominance"
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
