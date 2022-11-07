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
	"github.com/olehmushka/world-generator-engine/gender"
)

type Tradition struct {
	Slug                string        `json:"slug" bson:"slug"`
	Description         string        `json:"description" bson:"description"`
	PreferredEthosSlugs []string      `json:"preferred_ethos_slugs" bson:"preferred_ethos_slugs"`
	Type                TraditionType `json:"type" bson:"type"`
	OmitTraditionSlugs  []string      `json:"omit_tradition_slugs" bson:"omit_tradition_slugs"`
	OmitGenderDominance []gender.Sex  `json:"omit_gender_dominance" bson:"omit_gender_dominance"`
	OmitEthosSlugs      []string      `json:"omit_ethos_slugs" bson:"omit_ethos_slugs"`
}

func RandomTraditions(ts []*Tradition, min, max int, e Ethos, ds gender.Sex) ([]*Tradition, error) {
	if len(ts) == 0 {
		return nil, wrapped_error.NewBadRequestError(nil, "can not get random traditions of empty slice of traditions")
	}
	if len(ts) < min {
		return nil, wrapped_error.NewBadRequestError(nil, "can not get random traditions from less slice of traditions than min expected traditions number")
	}

	size, err := randomTools.RandIntInRange(min, max)
	if err != nil {
		return nil, wrapped_error.NewInternalServerError(err, "can not generate size of random traditions slise")
	}
	out := make([]*Tradition, 0, size)
	for _, t := range FilterTraditionsByDomitatedSex(FilterTraditionsByEthos(sliceTools.Shuffle(ts), e), ds) {
		if len(out) >= max {
			break
		}
		isOk := true
		for _, traditionToOmit := range t.OmitTraditionSlugs {
			for _, selected := range out {
				if selected.Slug == traditionToOmit {
					isOk = false
					break
				}
			}
		}
		if isOk {
			out = append(out, t)
		}
	}
	if len(out) < min {
		return nil, wrapped_error.NewBadRequestError(nil, fmt.Sprintf("can not generate expected number of traditions (expected=%d, actual=%d)", min, len(out)))
	}

	return out, nil
}

func FilterTraditionsByEthos(in []*Tradition, e Ethos) []*Tradition {
	if e.IsZero() {
		return in
	}

	out := make([]*Tradition, 0, len(in))
	for _, t := range in {
		isOk := true
		for _, ethosToOmit := range t.OmitEthosSlugs {
			if ethosToOmit == e.Slug {
				isOk = false
				break
			}
		}
		if isOk {
			out = append(out, t)
		}
	}

	return out
}

func FilterTraditionsByDomitatedSex(in []*Tradition, ds gender.Sex) []*Tradition {
	out := make([]*Tradition, 0, len(in))
	for _, t := range in {
		isOk := true
		for _, sexToOmit := range t.OmitGenderDominance {
			if sexToOmit == ds {
				isOk = false
				break
			}
		}
		if isOk {
			out = append(out, t)
		}
	}

	return out
}

func ExtractTraditions(cultures []*Culture) []*Tradition {
	out := make([]*Tradition, 0, len(cultures))
	for _, c := range cultures {
		out = append(out, c.Traditions...)
	}

	return out
}

func UniqueTraditions(in []*Tradition) []*Tradition {
	if len(in) <= 1 {
		return in
	}

	preOut := make(map[string]*Tradition)
	for _, e := range in {
		preOut[e.Slug] = e
	}

	out := make([]*Tradition, 0, len(preOut))
	for _, e := range preOut {
		out = append(out, e)
	}

	return out
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
