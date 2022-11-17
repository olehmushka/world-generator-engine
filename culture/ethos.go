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
	mapTools "github.com/olehmushka/golang-toolkit/map_tools"
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/olehmushka/world-generator-engine/types"
)

type Ethos struct {
	Slug             string `json:"slug" bson:"slug"`
	Description      string `json:"description" bson:"description"`
	IsDiplomatic     bool   `json:"is_diplomatic" bson:"is_diplomatic"`
	IsWarlike        bool   `json:"is_warlike" bson:"is_warlike"`
	IsAdministrative bool   `json:"is_administrative" bson:"is_administrative"`
	IsIntrigue       bool   `json:"is_intrigue" bson:"is_intrigue"`
	IsScholarly      bool   `json:"is_scholarly" bson:"is_scholarly"`
}

func (e Ethos) IsZero() bool {
	return e == Ethos{}
}

func RandomEthos(in []Ethos) (Ethos, error) {
	if len(in) == 0 {
		return Ethos{}, wrapped_error.NewBadRequestError(nil, "can not get random ethos of empty slice of ethoses")
	}
	if len(in) == 1 {
		return in[0], nil
	}

	out, err := sliceTools.RandomValueOfSlice(randomTools.RandFloat64, in)
	if err != nil {
		return Ethos{}, wrapped_error.NewInternalServerError(err, "can not get random ethos")
	}

	return out, nil
}

func ExtractEthoses(cultures []*Culture) []Ethos {
	if len(cultures) == 0 {
		return []Ethos{}
	}

	out := make([]Ethos, 0, len(cultures))
	for _, c := range cultures {
		if c == nil {
			continue
		}
		out = append(out, c.Ethos)
	}

	return out
}

func SelectEthosByMostRecent(in []Ethos) (Ethos, error) {
	m := make(map[string]int, 3)
	for _, e := range in {
		if count, ok := m[e.Slug]; ok {
			m[e.Slug] = count + 1
		} else {
			m[e.Slug] = 1
		}
	}
	probs := make(map[string]float64, 3)
	total := float64(len(in))
	for slug, count := range m {
		probs[slug] = float64(count) / total
	}

	slug, err := mapTools.PickOneByProb(probs)
	if err != nil {
		return Ethos{}, wrapped_error.NewInternalServerError(err, "can not select language slug by most recent")
	}
	for _, e := range in {
		if e.Slug == slug {
			return e, nil
		}
	}

	return Ethos{}, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not select language by slug (slug=%s)", slug))
}

func UniqueEthoses(ethoses []Ethos) []Ethos {
	if len(ethoses) <= 1 {
		return ethoses
	}

	preOut := make(map[string]Ethos)
	for _, e := range ethoses {
		preOut[e.Slug] = e
	}

	out := make([]Ethos, 0, len(preOut))
	for _, e := range preOut {
		out = append(out, e)
	}

	return out
}

func GetEthosSimilarityCoef(e1, e2 Ethos) float64 {
	if e1.IsZero() && e2.IsZero() {
		return 1
	}
	if e1.IsZero() || e2.IsZero() {
		return 0
	}

	if e1.Slug == e2.Slug {
		return 1
	}

	similarityTraits := []struct {
		enable bool
		coef   float64
	}{
		{
			enable: e1.IsDiplomatic == e2.IsDiplomatic,
			coef:   0.2,
		},
		{
			enable: e1.IsWarlike == e2.IsWarlike,
			coef:   0.2,
		},
		{
			enable: e1.IsAdministrative == e2.IsAdministrative,
			coef:   0.2,
		},
		{
			enable: e1.IsIntrigue == e2.IsIntrigue,
			coef:   0.2,
		},
		{
			enable: e1.IsScholarly == e2.IsScholarly,
			coef:   0.2,
		},
	}

	var out float64
	for _, t := range similarityTraits {
		if t.enable {
			out += t.coef
		}
	}

	return out
}

func LoadAllEthoses(opts ...types.ChangeStringFunc) chan either.Either[[]Ethos] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "culture")
	dirname := currDirname + "data/ethoses/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	ch := make(chan either.Either[[]Ethos], MaxLoadDataConcurrency)
	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[[]Ethos]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", dirname))}
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
					ch <- either.Either[[]Ethos]{Err: err}
					return
				}
				var sfs []Ethos
				if err := json.Unmarshal(b, &sfs); err != nil {
					ch <- either.Either[[]Ethos]{Err: err}
					return
				}
				for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, sfs) {
					ch <- either.Either[[]Ethos]{Value: chunk}
				}
			}(file)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

func SearchEthos(slug string, opts ...types.ChangeStringFunc) (*Ethos, error) {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "culture")
	dirname := currDirname + "data/ethoses/"
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
		var es []*Ethos
		if err := json.Unmarshal(b, &es); err != nil {
			return nil, err
		}
		for _, e := range es {
			if e.Slug == slug {
				return e, nil
			}
		}
	}

	return nil, nil
}
