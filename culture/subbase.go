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
)

type Subbase struct {
	Slug     string `json:"slug" bson:"slug"`
	BaseSlug string `json:"base_slug" bson:"base_slug"`
}

func LoadAllSubbases() chan either.Either[[]*Subbase] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/subbases/"
	ch := make(chan either.Either[[]*Subbase], MaxLoadDataConcurrency)
	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[[]*Subbase]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))}
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
