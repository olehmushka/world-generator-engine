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
	"github.com/olehmushka/golang-toolkit/list"
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

func FilterRawCulturesBySlugs(rCultures []*RawCulture, slugs []string) []*RawCulture {
	out := make([]*RawCulture, 0, len(rCultures))
	for _, rc := range rCultures {
		if sliceTools.Includes(slugs, rc.Slug) {
			out = append(out, rc)
		}
	}

	return out
}

func SearchCultures(slugs []string) ([]*Culture, error) {
	if len(slugs) == 0 {
		return []*Culture{}, nil
	}

	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/cultures/"
	rawCultureCh := make(chan []*RawCulture, MaxLoadDataConcurrency)
	ch := make(chan either.Either[*Culture], MaxLoadDataConcurrency)

	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[*Culture]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))}
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
					ch <- either.Either[*Culture]{Err: err}
					return
				}
				var rcs []*RawCulture
				if err := json.Unmarshal(b, &rcs); err != nil {
					ch <- either.Either[*Culture]{Err: err}
					return
				}
				rcs = FilterRawCulturesBySlugs(rcs, slugs)
				if len(rcs) == 0 {
					return
				}
				for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, rcs) {
					rawCultureCh <- chunk
				}
			}(file)
		}
		wg.Wait()
		close(rawCultureCh)
	}()

	go func() {
		subbases := list.NewFIFOUniqueList(100, func(sb1, sb2 *Subbase) bool {
			return sb1.Slug == sb2.Slug
		})
		ethoses := list.NewFIFOUniqueList(100, func(e1, e2 *Ethos) bool {
			return e1.Slug == e2.Slug
		})
		traditions := list.NewFIFOUniqueList(100, func(t1, t2 *Tradition) bool {
			return t1.Slug == t2.Slug
		})

		for rawCultures := range rawCultureCh {
			for _, rCulture := range rawCultures {
				// get subbase
				sb, isSBFound := subbases.FindOne(func(_, curr, _ *Subbase) bool { return curr.Slug == rCulture.SubbaseSlug })
				if !isSBFound {
					found, err := SearchSubbase(rCulture.SubbaseSlug)
					if err != nil {
						ch <- either.Either[*Culture]{Err: err}
						return
					}
					if found == nil {
						ch <- either.Either[*Culture]{Err: wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not found subbase by slug (slug=%s)", rCulture.SubbaseSlug))}
						return
					}
					sb = found
				}
				subbases.Push(sb)

				// get ethos
				e, isEFound := ethoses.FindOne(func(_, curr, _ *Ethos) bool { return curr.Slug == rCulture.EthosSlug })
				if !isEFound {
					found, err := SearchEthos(rCulture.EthosSlug)
					if err != nil {
						ch <- either.Either[*Culture]{Err: err}
						return
					}
					if found == nil {
						ch <- either.Either[*Culture]{Err: wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not found ethos by slug (slug=%s)", rCulture.EthosSlug))}
						return
					}
					e = found
				}
				ethoses.Push(e)

				ts := traditions.FindMany(func(_, curr, _ *Tradition) bool {
					return sliceTools.Includes(rCulture.TraditionSlugs, curr.Slug)
				})
				if len(ts) != len(rCulture.TraditionSlugs) {
					_, noFound := SepareteTraditionsByPresent(ts, rCulture.TraditionSlugs)
					found, err := SearchTraditions(noFound)
					if err != nil {
						ch <- either.Either[*Culture]{Err: err}
						return
					}
					if len(found) != len(noFound) {
						ch <- either.Either[*Culture]{Err: wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not found all traditions (found.length=%d, all_required.length=%d)", len(found), len(noFound)))}
						return
					}
					ts = sliceTools.Merge(ts, found)
				}
				for _, t := range ts {
					traditions.Push(t)
				}
				parentCultures, err := SearchCultures(rCulture.ParentCultureSlugs)
				if err != nil {
					ch <- either.Either[*Culture]{Err: err}
					return
				}

				ch <- either.Either[*Culture]{Value: &Culture{
					Slug:            rCulture.Slug,
					BaseSlug:        rCulture.BaseSlug,
					Subbase:         *sb,
					ParentCultures:  parentCultures,
					Ethos:           *e,
					Traditions:      ts,
					LanguageSlug:    rCulture.LanguageSlug,
					GenderDominance: rCulture.GenderDominance,
					MartialCustom:   rCulture.MartialCustom,
				}}
			}
		}
		close(ch)
	}()
	out := make([]*Culture, 0, len(slugs))
	for chunk := range ch {
		if chunk.Err != nil {
			close(ch)
			return nil, chunk.Err
		}
		out = append(out, chunk.Value)
	}

	return out, nil
}

func SepareteCulturesByPresent(present []*Culture, ts []string) ([]string, []string) {
	if len(present) == 0 {
		return []string{}, []string{}
	}
	included := make([]string, 0, len(present)/2)
	notIncluded := make([]string, 0, len(present)/2)
	for _, t := range ts {
		var isFound bool
		for _, pr := range present {
			if pr.Slug == t {
				isFound = true
				break
			}
		}
		if isFound {
			included = append(included, t)
			continue
		}
		notIncluded = append(notIncluded, t)
	}

	return included, notIncluded
}

func LoadAllCultures() chan either.Either[*Culture] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/cultures/"
	rawCultureCh := make(chan []*RawCulture, MaxLoadDataConcurrency)
	ch := make(chan either.Either[*Culture], MaxLoadDataConcurrency)

	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[*Culture]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))}
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
					ch <- either.Either[*Culture]{Err: err}
					return
				}
				var sfs []*RawCulture
				if err := json.Unmarshal(b, &sfs); err != nil {
					ch <- either.Either[*Culture]{Err: err}
					return
				}
				if len(sfs) == 0 {
					return
				}
				for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, sfs) {
					rawCultureCh <- chunk
				}
			}(file)
		}
		wg.Wait()
		close(rawCultureCh)
	}()

	go func() {
		subbases := list.NewFIFOUniqueList(100, func(sb1, sb2 *Subbase) bool {
			return sb1.Slug == sb2.Slug
		})
		ethoses := list.NewFIFOUniqueList(100, func(e1, e2 *Ethos) bool {
			return e1.Slug == e2.Slug
		})
		traditions := list.NewFIFOUniqueList(100, func(t1, t2 *Tradition) bool {
			return t1.Slug == t2.Slug
		})
		cultures := list.NewFIFOUniqueList(100, func(t1, t2 *Culture) bool {
			return t1.Slug == t2.Slug
		})

		for rawCultures := range rawCultureCh {
			for _, rCulture := range rawCultures {
				// get subbase
				sb, isSBFound := subbases.FindOne(func(_, curr, _ *Subbase) bool { return curr.Slug == rCulture.SubbaseSlug })
				if !isSBFound {
					found, err := SearchSubbase(rCulture.SubbaseSlug)
					if err != nil {
						ch <- either.Either[*Culture]{Err: err}
						return
					}
					if found == nil {
						ch <- either.Either[*Culture]{Err: wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not found subbase by slug (slug=%s)", rCulture.SubbaseSlug))}
						return
					}
					sb = found
				}
				subbases.Push(sb)

				// get ethos
				e, isEFound := ethoses.FindOne(func(_, curr, _ *Ethos) bool { return curr.Slug == rCulture.EthosSlug })
				if !isEFound {
					found, err := SearchEthos(rCulture.EthosSlug)
					if err != nil {
						ch <- either.Either[*Culture]{Err: err}
						return
					}
					if found == nil {
						ch <- either.Either[*Culture]{Err: wrapped_error.NewNotFoundError(nil, fmt.Sprintf("can not found ethos by slug (slug=%s)", rCulture.EthosSlug))}
						return
					}
					e = found
				}
				ethoses.Push(e)

				ts := traditions.FindMany(func(_, curr, _ *Tradition) bool {
					return sliceTools.Includes(rCulture.TraditionSlugs, curr.Slug)
				})
				if len(ts) != len(rCulture.TraditionSlugs) {
					_, noFound := SepareteTraditionsByPresent(ts, rCulture.TraditionSlugs)
					found, err := SearchTraditions(noFound)
					if err != nil {
						ch <- either.Either[*Culture]{Err: err}
						return
					}
					if len(found) != len(noFound) {
						ch <- either.Either[*Culture]{Err: wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not found all traditions (found.length=%d, all_required.length=%d)", len(found), len(noFound)))}
						return
					}
					ts = sliceTools.Merge(ts, found)
				}
				for _, t := range ts {
					traditions.Push(t)
				}

				cs := cultures.FindMany(func(_, curr, _ *Culture) bool {
					return sliceTools.Includes(rCulture.ParentCultureSlugs, curr.Slug)
				})
				if len(cs) != len(rCulture.ParentCultureSlugs) {
					_, noFound := SepareteCulturesByPresent(cs, rCulture.ParentCultureSlugs)
					found, err := SearchCultures(noFound)
					if err != nil {
						ch <- either.Either[*Culture]{Err: err}
						return
					}
					if len(found) != len(noFound) {
						ch <- either.Either[*Culture]{Err: wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not found all traditions (found.length=%d, all_required.length=%d)", len(found), len(noFound)))}
						return
					}
					cs = sliceTools.Merge(cs, found)
				}
				for _, c := range cs {
					cultures.Push(c)
				}

				ch <- either.Either[*Culture]{Value: &Culture{
					Slug:            rCulture.Slug,
					BaseSlug:        rCulture.BaseSlug,
					Subbase:         *sb,
					ParentCultures:  cs,
					Ethos:           *e,
					Traditions:      ts,
					LanguageSlug:    rCulture.LanguageSlug,
					GenderDominance: rCulture.GenderDominance,
					MartialCustom:   rCulture.MartialCustom,
				}}
			}
		}
		close(ch)
	}()

	return ch
}
