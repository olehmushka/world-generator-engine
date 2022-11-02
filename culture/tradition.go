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
	"github.com/olehmushka/world-generator-engine/gender"
)

type Tradition struct {
	Slug                string        `json:"slug" bson:"slug"`
	Description         string        `json:"description" bson:"description"`
	PreferredEthosSlugs []string      `json:"preferred_ethos_slugs" bson:"preferred_ethos_slugs"`
	Type                TraditionType `json:"type" bson:"type"`
	OmitTraditionSlugs  []string      `json:"omit_tradition_slugs" bson:"omit_tradition_slugs"`
	OmitGenderDominance []gender.Sex  `json:"omit_gender_dominance" bson:"omit_gender_dominance"`
	OmitEthosSLugs      []string      `json:"omit_ethos_slugs" bson:"omit_ethos_slugs"`
}

func LoadAllTraditions() chan either.Either[[]*Tradition] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/traditions/"
	ch := make(chan either.Either[[]*Tradition], MaxLoadDataConcurrency)
	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[[]*Tradition]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))}
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
					ch <- either.Either[[]*Tradition]{Err: err}
					return
				}
				var ts []*Tradition
				if err := json.Unmarshal(b, &ts); err != nil {
					ch <- either.Either[[]*Tradition]{Err: err}
					return
				}
				for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, ts) {
					ch <- either.Either[[]*Tradition]{Value: chunk}
				}
			}(file)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

func SearchTradition(slug string) (*Tradition, error) {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/traditions/"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))
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
		var ts []*Tradition
		if err := json.Unmarshal(b, &ts); err != nil {
			return nil, err
		}
		for _, t := range ts {
			if t.Slug == slug {
				return t, nil
			}
		}
	}

	return nil, nil
}
