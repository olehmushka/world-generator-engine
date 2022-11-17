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

func RandomBase(in []string) (string, error) {
	if len(in) == 0 {
		return "", wrapped_error.NewBadRequestError(nil, "can not get random base of empty slice of bases")
	}
	if len(in) == 1 {
		return in[0], nil
	}

	out, err := sliceTools.RandomValueOfSlice(randomTools.RandFloat64, in)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not get random base")
	}

	return out, nil
}

func ExtractBases(cultures []*Culture) []string {
	out := make([]string, len(cultures))
	for i := range out {
		out[i] = cultures[i].BaseSlug
	}

	return out
}

func SelectBaseByMostRecent(in []string) (string, error) {
	m := make(map[string]int, 3)
	for _, baseSlug := range in {
		if count, ok := m[baseSlug]; ok {
			m[baseSlug] = count + 1
		} else {
			m[baseSlug] = 1
		}
	}
	probs := make(map[string]float64, 3)
	total := float64(len(in))
	for baseSlug, count := range m {
		probs[baseSlug] = float64(count) / total
	}

	baseSlug, err := mapTools.PickOneByProb(probs)
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, "can not select base slug by most recent")
	}

	return baseSlug, nil
}

func LoadAllBases(opts ...types.ChangeStringFunc) chan either.Either[[]string] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "culture")
	dirname := currDirname + "data/bases/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	ch := make(chan either.Either[[]string], MaxLoadDataConcurrency)

	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[[]string]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", dirname))}
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
					ch <- either.Either[[]string]{Err: err}
					return
				}
				var bs []string
				if err := json.Unmarshal(b, &bs); err != nil {
					ch <- either.Either[[]string]{Err: err}
					return
				}
				for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, bs) {
					ch <- either.Either[[]string]{Value: chunk}
				}
				wg.Done()
			}(file)
		}
		wg.Wait()
		close(ch)
	}()

	return ch
}
