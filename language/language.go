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
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
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
