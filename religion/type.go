package religion

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/olehmushka/golang-toolkit/either"
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	we "github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/olehmushka/world-generator-engine/types"
)

func LoadAllTypeTraits(opts ...types.ChangeStringFunc) chan either.Either[[]*Trait] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/types/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	if strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}
	fn := dirname + "traits.json"
	ch := make(chan either.Either[[]*Trait], MaxLoadDataConcurrency)
	go func() {
		b, err := os.ReadFile(fn)
		if err != nil {
			ch <- either.Either[[]*Trait]{Err: we.NewInternalServerError(err, fmt.Sprintf("can not read file by filename (filename=%s)", fn))}
			return
		}
		var ts []*Trait
		if err := json.Unmarshal(b, &ts); err != nil {
			ch <- either.Either[[]*Trait]{Err: err}
			return
		}
		for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, ts) {
			ch <- either.Either[[]*Trait]{Value: chunk}
		}

		close(ch)
	}()

	return ch
}

func NewType(list []*Trait, cfg StatsConfig, r *Religion) (*Trait, *Stats, error) {
	t, stats, err := FilterTypeTrait(cfg, r, list)
	if err != nil {
		return nil, nil, err
	}

	return t, stats, nil
}
