package religion

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/olehmushka/golang-toolkit/either"
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	we "github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/world-generator-engine/favour"
	"github.com/olehmushka/world-generator-engine/tools"
	"github.com/olehmushka/world-generator-engine/types"
)

func LoadAllDeityFavours(opts ...types.ChangeStringFunc) chan either.Either[[]*FavourTrait] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/deity/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	if strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}
	fn := dirname + "favour.json"
	ch := make(chan either.Either[[]*FavourTrait], MaxLoadDataConcurrency)
	go func() {
		b, err := os.ReadFile(fn)
		if err != nil {
			ch <- either.Either[[]*FavourTrait]{Err: we.NewInternalServerError(err, fmt.Sprintf("can not read file by filename (filename=%s)", fn))}
			return
		}
		var ts []*FavourTrait
		if err := json.Unmarshal(b, &ts); err != nil {
			ch <- either.Either[[]*FavourTrait]{Err: err}
			return
		}
		for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, ts) {
			ch <- either.Either[[]*FavourTrait]{Value: chunk}
		}

		close(ch)
	}()

	return ch
}

func NewDeityFavour(list []*FavourTrait, cfg StatsConfig, r *Religion) (*FavourTrait, *Stats, error) {
	filtered, err := FilterFavourTraits(cfg, r, list)
	if err != nil {
		return nil, nil, we.NewInternalServerError(err, "can not generate deity favour")
	}
	favours := ExtractFavoursFromFavourTraits(filtered)
	positivity, err := randomTools.RandFloat64InRange(0, 1)
	if err != nil {
		return nil, nil, we.NewInternalServerError(err, "can not generate positivity for deity favour generation")
	}
	dev, err := randomTools.RandFloat64InRange(1, 5)
	if err != nil {
		return nil, nil, we.NewInternalServerError(err, "can not generate deviation for deity favour generation")
	}
	f, err := favour.GenerateFavourOf(favours, positivity, dev)
	if err != nil {
		return nil, nil, we.NewInternalServerError(err, "can not pick favour for deity favour generation")
	}
	out := FindFavourTrait(list, f)
	if out == nil {
		return nil, nil, we.NewInternalServerError(nil, fmt.Sprintf("can not find favour (favour=%s) trait for deity favour generation", f.String()))
	}

	stats := r.Stats
	merged, err := MergeReligionStats(cfg, stats, out.Stats)
	if err != nil {
		return nil, nil, err
	}
	stats = merged

	return out, stats, nil
}

func LoadAllDeityNatureTraits(opts ...types.ChangeStringFunc) chan either.Either[[]*Trait] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/deity/"
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

func NewDeityNatureTraits(cfg StatsConfig, r *Religion, in []*Trait, min, max int) ([]*Trait, *Stats, error) {
	return FilterTraits(cfg, r, in, min, max)
}
