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
)

type Subbase struct {
	Slug     string `json:"slug" bson:"slug"`
	BaseSlug string `json:"base_slug" bson:"base_slug"`
}

func (sb Subbase) IsZero() bool {
	return sb == Subbase{}
}

func FilterSubbasesByBaseSlug(subbases []Subbase, baseSlug string) []Subbase {
	out := make([]Subbase, 0, (len(subbases)/2)+1)
	for _, sb := range subbases {
		if sb.BaseSlug == baseSlug {
			out = append(out, sb)
		}
	}

	return out
}

func RandomSubbase(in []Subbase) (Subbase, error) {
	if len(in) == 0 {
		return Subbase{}, wrapped_error.NewBadRequestError(nil, "can not get random subbase of empty slice of subbases")
	}
	if len(in) == 1 {
		return in[0], nil
	}

	out, err := sliceTools.RandomValueOfSlice(randomTools.RandFloat64, in)
	if err != nil {
		return Subbase{}, wrapped_error.NewInternalServerError(err, "can not get random subbase")
	}

	return out, nil
}

func ExtractSubbases(cultures []*Culture) []Subbase {
	out := make([]Subbase, len(cultures))
	for i := range out {
		out[i] = cultures[i].Subbase
	}

	return out
}

func SelectSubbaseByMostRecent(in []Subbase) (Subbase, error) {
	m := make(map[string]int, 3)
	for _, sb := range in {
		if count, ok := m[sb.Slug]; ok {
			m[sb.Slug] = count + 1
		} else {
			m[sb.Slug] = 1
		}
	}
	probs := make(map[string]float64, 3)
	total := float64(len(in))
	for slug, count := range m {
		probs[slug] = float64(count) / total
	}

	slug, err := mapTools.PickOneByProb(probs)
	if err != nil {
		return Subbase{}, wrapped_error.NewInternalServerError(err, "can not select subbase by most recent")
	}
	for _, sb := range in {
		if sb.Slug == slug {
			return sb, nil
		}
	}

	return Subbase{}, wrapped_error.NewInternalServerError(nil, fmt.Sprintf("can not select subbase by slug (slug=%s)", slug))
}

func LoadAllSubbases(opts ...PathChangeLoadOpts) chan either.Either[[]*Subbase] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "culture")
	dirname := currDirname + "data/subbases/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	ch := make(chan either.Either[[]*Subbase], MaxLoadDataConcurrency)
	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[[]*Subbase]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", dirname))}
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
					ch <- either.Either[[]*Subbase]{Err: err}
					return
				}
				var sbs []*Subbase
				if err := json.Unmarshal(b, &sbs); err != nil {
					ch <- either.Either[[]*Subbase]{Err: err}
					return
				}
				for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, sbs) {
					ch <- either.Either[[]*Subbase]{Value: chunk}
				}
			}(file)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}

func SearchSubbase(slug string, opts ...PathChangeLoadOpts) (*Subbase, error) {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "culture")
	dirname := currDirname + "data/subbases/"
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
		var sbs []*Subbase
		if err := json.Unmarshal(b, &sbs); err != nil {
			return nil, err
		}
		for _, sb := range sbs {
			if sb.Slug == slug {
				return sb, nil
			}
		}
	}

	return nil, nil
}
