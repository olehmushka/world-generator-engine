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
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/olehmushka/world-generator-engine/types"
)

type Subfamily struct {
	Slug              string     `json:"slug" bson:"slug"`
	FamilySlug        string     `json:"family_slug" bson:"family_slug"`
	ExtendedSubfamily *Subfamily `json:"extended_subfamily" bson:"extended_subfamily"`
}

func LoadAllSubfamilies(opts ...types.ChangeStringFunc) chan either.Either[[]*Subfamily] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "language")
	dirname := currDirname + "data/subfamilies/"
	ch := make(chan either.Either[[]*Subfamily], MaxLoadDataConcurrency)
	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[[]*Subfamily]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", dirname))}
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
					ch <- either.Either[[]*Subfamily]{Err: err}
					return
				}
				var sfs []*Subfamily
				if err := json.Unmarshal(b, &sfs); err != nil {
					ch <- either.Either[[]*Subfamily]{Err: err}
					return
				}
				for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, sfs) {
					ch <- either.Either[[]*Subfamily]{Value: chunk}
				}
			}(file)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

func SearchSubfamily(slug string, opts ...types.ChangeStringFunc) (*Subfamily, error) {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "language")
	dirname := currDirname + "data/subfamilies/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", dirname))
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := dirname + file.Name()
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		var sfs []*Subfamily
		if err := json.Unmarshal(b, &sfs); err != nil {
			return nil, err
		}
		for _, sf := range sfs {
			if sf.Slug == slug {
				return sf, nil
			}
		}
	}

	return nil, nil
}

func GetLanguageSubfamilyChains(accum []string, sf *Subfamily) []string {
	if sf == nil {
		return accum
	}
	if sf.ExtendedSubfamily == nil {
		return accum
	}

	return GetLanguageSubfamilyChains(
		append(accum, sf.Slug),
		sf.ExtendedSubfamily,
	)
}
