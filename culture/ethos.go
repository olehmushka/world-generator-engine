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
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
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

func getRandomEthosFromCultures(proto []*Culture) (Ethos, error) {
	ethoses := UniqueEthoses(ExtractEthoses(proto))
	m := make(map[string]int)
	for _, c := range proto {
		if _, ok := m[c.Ethos.Slug]; ok {
			m[c.Ethos.Slug]++
			continue
		}
		m[c.Ethos.Slug] = 1
	}

	mWithProb := make(map[string]float64, len(m))
	for name, amount := range m {
		mWithProb[name] = float64(amount) / float64(len(proto))
	}

	chosenEthos, err := mapTools.PickOneByProb(mWithProb)
	if err != nil {
		return Ethos{}, wrapped_error.NewInternalServerError(err, "can not generate ethos")
	}
	for _, e := range ethoses {
		if e.Slug == chosenEthos {
			return e, nil
		}
	}

	return Ethos{}, nil
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

func LoadAllEthoses() chan either.Either[[]Ethos] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/ethoses/"
	ch := make(chan either.Either[[]Ethos], MaxLoadDataConcurrency)
	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[[]Ethos]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))}
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

func SearchEthos(slug string) (Ethos, error) {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/ethoses/"
	files, err := ioutil.ReadDir(dirname)
	if err != nil {
		return Ethos{}, wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := dirname + file.Name()
		b, err := ioutil.ReadFile(filename)
		if err != nil {
			return Ethos{}, err
		}
		var sfs []Ethos
		if err := json.Unmarshal(b, &sfs); err != nil {
			return Ethos{}, err
		}
		for _, sf := range sfs {
			if sf.Slug == slug {
				return sf, nil
			}
		}
	}

	return Ethos{}, nil
}
