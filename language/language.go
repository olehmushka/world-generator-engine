package language

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"io/ioutil"
	"path"
	"runtime"
	"sync"

	"github.com/olehmushka/golang-toolkit/either"
	"github.com/olehmushka/golang-toolkit/list"
	mapTools "github.com/olehmushka/golang-toolkit/map_tools"
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/world-generator-engine/gender"
)

type RawLanguage struct {
	Slug          string `json:"slug" bson:"slug"`
	SubfamilySlug string `json:"subfamily_slug" bson:"subfamily_slug"`
	WordbaseSlug  string `json:"wordbase_slug" bson:"wordbase_slug"`
}

type Language struct {
	Slug      string     `json:"slug" bson:"slug"`
	Subfamily *Subfamily `json:"subfamily" bson:"subfamily"`
	Wordbase  *Wordbase  `json:"wordbase" bson:"wordbase"`
}

func (l *Language) IsZero() bool {
	return l == nil
}

func GetOwnName(lang *Language, sex gender.Sex) (string, error) {
	switch sex {
	case gender.MaleSex:
		return sliceTools.RandomValueOfSlice(randomTools.RandFloat64, lang.Wordbase.MaleOwnNames)
	case gender.FemaleSex:
		return sliceTools.RandomValueOfSlice(randomTools.RandFloat64, lang.Wordbase.FemaleOwnNames)
	default:
		return "", nil
	}
}

func GetFullLanguageChains(lang *Language) []string {
	if lang == nil {
		return []string{}
	}
	out := []string{lang.Slug}
	out = GetLanguageSubfamilyChains(out, lang.Subfamily)

	return append(out, string(lang.Subfamily.FamilySlug))
}

func GetLanguageKinship(l1, l2 *Language) int {
	var (
		l1Chains = GetFullLanguageChains(l1)
		l2Chains = GetFullLanguageChains(l2)
	)
	if len(l1Chains) == 0 || len(l2Chains) == 0 {
		return -1
	}
	if l1Chains[0] == l2Chains[0] {
		return 0
	}
	if l1Chains[len(l1Chains)-1] != l2Chains[len(l2Chains)-1] {
		return -1
	}

	maxIter := len(l1Chains)
	if maxIter < len(l2Chains) {
		maxIter = len(l2Chains)
	}
	for i := 0; i < maxIter; i++ {
		var (
			l1El = l1Chains[len(l1Chains)-1]
			l2El = l2Chains[len(l2Chains)-1]
		)
		if len(l1Chains) > i {
			l1El = l1Chains[i]
		}
		if len(l2Chains) > i {
			l2El = l2Chains[i]
		}
		if l1El == l2El {
			return i
		}
	}

	return -1
}

func GetLanguageSimilarityCoef(l1, l2 *Language) (float64, error) {
	if l1.IsZero() && l2.IsZero() {
		return 1, nil
	}
	if l1.IsZero() || l2.IsZero() {
		return 0, wrapped_error.NewInternalServerError(nil, "can not compare languages if one of it is <nil>")
	}

	switch kinship := GetLanguageKinship(l1, l2); kinship {
	case -1:
		return 0, nil
	case 0:
		return 1, nil
	case 1:
		return 0.75, nil
	default:
		return 1 / float64(kinship), nil
	}
}

func RandomLanguageSlug(in []string) (string, error) {
	if len(in) == 0 {
		return "", wrapped_error.NewBadRequestError(nil, "can not get random language of empty slice of languages")
	}
	if len(in) == 1 {
		return in[0], nil
	}

	out, err := sliceTools.RandomValueOfSlice(randomTools.RandFloat64, in)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not get random language")
	}

	return out, nil
}

func FindLanguageBySlug(langs []*Language, slug string) *Language {
	for _, lang := range langs {
		if lang.Slug == slug {
			return lang
		}
	}

	return nil
}

func SelectLanguageSlugByMostRecent(in []string) (string, error) {
	m := make(map[string]int, 3)
	for _, slug := range in {
		if count, ok := m[slug]; ok {
			m[slug] = count + 1
		} else {
			m[slug] = 1
		}
	}
	probs := make(map[string]float64, 3)
	total := float64(len(in))
	for slug, count := range m {
		probs[slug] = float64(count) / total
	}

	slug, err := mapTools.PickOneByProb(probs)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not select language slug by most recent")
	}

	return slug, nil
}

func LoadAllLanguages() chan either.Either[*Language] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/languages/"
	rawLangCh := make(chan []*RawLanguage, MaxLoadDataConcurrency)
	ch := make(chan either.Either[*Language], MaxLoadDataConcurrency)

	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[*Language]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))}
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
					ch <- either.Either[*Language]{Err: err}
					return
				}
				var ls []*RawLanguage
				if err := json.Unmarshal(b, &ls); err != nil {
					ch <- either.Either[*Language]{Err: err}
					return
				}
				if len(ls) == 0 {
					return
				}
				for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, ls) {
					rawLangCh <- chunk
				}

			}(file)
		}
		wg.Wait()
		close(rawLangCh)
	}()

	go func() {
		subfamilies := list.NewFIFOUniqueList(100, func(sf1, sf2 *Subfamily) bool {
			return sf1.Slug == sf2.Slug
		})
		wordbases := list.NewFIFOUniqueList(100, func(w1, w2 *Wordbase) bool {
			return w1.Slug == w2.Slug
		})

		for rawLangs := range rawLangCh {
			for _, rLang := range rawLangs {
				// get subfamily
				sf, isSFFound := subfamilies.FindOne(func(_, curr, _ *Subfamily) bool { return curr.Slug == rLang.SubfamilySlug })
				if !isSFFound {
					found, err := SearchSubfamily(rLang.SubfamilySlug)
					if err != nil {
						ch <- either.Either[*Language]{Err: err}
						return
					}
					if found == nil {
						ch <- either.Either[*Language]{Err: wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not found subfamily by slug (slug=%s)", rLang.SubfamilySlug))}
						return
					}
					sf = found
				}
				subfamilies.Push(sf)

				// get wordbase
				wb, isWBFound := wordbases.FindOne(func(_, curr, _ *Wordbase) bool { return curr.Slug == rLang.WordbaseSlug })
				if !isWBFound {
					found, err := SearchWordbase(rLang.WordbaseSlug)
					if err != nil {
						ch <- either.Either[*Language]{Err: err}
						return
					}
					if found == nil {
						ch <- either.Either[*Language]{Err: wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not found wordbase by slug (slug=%s)", rLang.WordbaseSlug))}
						return
					}
					wb = found
				}
				wordbases.Push(wb)

				ch <- either.Either[*Language]{Value: &Language{
					Slug:      rLang.Slug,
					Subfamily: sf,
					Wordbase:  wb,
				}}
			}
		}
		close(ch)
	}()

	return ch
}
