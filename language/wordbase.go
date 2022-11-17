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
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/olehmushka/world-generator-engine/types"
)

type Wordbase struct {
	FemaleOwnNames []string `json:"female_own_names" bson:"female_own_names"`
	MaleOwnNames   []string `json:"male_own_names" bson:"male_own_names"`
	Words          []string `json:"words" bson:"words"`
	Slug           string   `json:"slug" bson:"slug"`
	Min            int      `json:"min" bson:"min"`
	Max            int      `json:"max" bson:"max"`
	Dupl           string   `json:"dupl" bson:"dupl"`
	M              float64  `json:"m" bson:"m"`
}

func LoadAllWordbases(opts ...types.ChangeStringFunc) chan either.Either[*Wordbase] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "language")
	dirname := currDirname + "data/wordbases/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	ch := make(chan either.Either[*Wordbase], MaxLoadDataConcurrency)
	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[*Wordbase]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", dirname))}
			return
		}

		var wg sync.WaitGroup
		wg.Add(len(files))
		for _, file := range files {
			go func(file fs.FileInfo) {
				if file.IsDir() {
					return
				}
				filename := dirname + file.Name()
				b, err := ioutil.ReadFile(filename)
				if err != nil {
					ch <- either.Either[*Wordbase]{Err: err}
					return
				}
				var wb *Wordbase
				if err := json.Unmarshal(b, &wb); err != nil {
					ch <- either.Either[*Wordbase]{Err: err}
					return
				}
				ch <- either.Either[*Wordbase]{Value: wb}
				wg.Done()
			}(file)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

func SearchWordbase(slug string, opts ...types.ChangeStringFunc) (*Wordbase, error) {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "language")
	dirname := currDirname + "data/wordbases/"
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
		var wb Wordbase
		if err := json.Unmarshal(b, &wb); err != nil {
			return nil, err
		}
		if wb.Slug == slug {
			return &wb, nil
		}
	}

	return nil, nil
}
