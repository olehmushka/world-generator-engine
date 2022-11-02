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

func LoadAllBases() chan either.Either[[]string] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := path.Dir(filename) + "/"
	dirname := currDirname + "/data/bases/"
	ch := make(chan either.Either[[]string], MaxLoadDataConcurrency)

	go func() {
		files, err := ioutil.ReadDir(dirname)
		if err != nil {
			ch <- either.Either[[]string]{Err: wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not read dir by dirname (dirname=%s)", currDirname))}
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
