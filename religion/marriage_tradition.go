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

type MarriageTradition struct {
	Kind          *Trait           `json:"kind"`
	Bastardry     *Trait           `json:"bastardy"`
	Consanguinity *Trait           `json:"consanguinity"`
	Divorce       *PermissionTrait `json:"divorce"`
}

func (mt MarriageTradition) IsZero() bool {
	return mt == MarriageTradition{}
}

type CreateMarriageTraditionData struct {
	MarriageKinds           []*Trait           `json:"marriage_kinds"`
	BastardyTraditions      []*Trait           `json:"bastardy_traditions"`
	ConsanguinityTraditions []*Trait           `json:"consanguinity_traditions"`
	DivorceTraditions       []*PermissionTrait `json:"divorce_traditions"`
}

func NewMarriageTradition(cfg StatsConfig, r *Religion, data CreateMarriageTraditionData) (MarriageTradition, *Stats, error) {
	stats := r.Stats
	out := MarriageTradition{}
	// creation marriage tradition kind
	kind, kindStats, err := FilterTraits(cfg, r, data.MarriageKinds, 1, 2)
	if err != nil {
		return MarriageTradition{}, nil, err
	}
	if len(kind) < 1 {
		return MarriageTradition{}, nil, we.NewInternalServerError(nil, fmt.Sprintf("unexpected kinds number (len=%d)", len(kind)))
	}
	out.Kind = kind[0]
	stats = kindStats
	// creation marriage tradition bastardy
	bastardy, bastardyStats, err := FilterTraits(cfg, r, data.BastardyTraditions, 1, 2)
	if err != nil {
		return MarriageTradition{}, nil, err
	}
	if len(bastardy) < 1 {
		return MarriageTradition{}, nil, we.NewInternalServerError(nil, fmt.Sprintf("unexpected bastardies number (len=%d)", len(bastardy)))
	}
	out.Bastardry = bastardy[0]
	stats = bastardyStats
	// creation marriage tradition consanguinity
	consanguinity, consanguinityStats, err := FilterTraits(cfg, r, data.ConsanguinityTraditions, 1, 2)
	if err != nil {
		return MarriageTradition{}, nil, err
	}
	if len(consanguinity) < 1 {
		return MarriageTradition{}, nil, we.NewInternalServerError(nil, fmt.Sprintf("unexpected consanguinities number (len=%d)", len(consanguinity)))
	}
	out.Consanguinity = consanguinity[0]
	stats = consanguinityStats
	// creation divorce tradition
	divorce, divorceStats, err := FilterPermissionTraits(cfg, r, data.DivorceTraditions, 1, 2)
	if err != nil {
		return MarriageTradition{}, nil, err
	}
	if len(divorce) < 1 {
		return MarriageTradition{}, nil, we.NewInternalServerError(nil, fmt.Sprintf("unexpected divorce traditions number (len=%d)", len(divorce)))
	}
	out.Divorce = divorce[0]
	stats = divorceStats

	return out, stats, nil
}

func LoadAllMarriageKinds(opts ...types.ChangeStringFunc) chan either.Either[[]*Trait] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/marriage_traits/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	if strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}
	fn := dirname + "kinds.json"
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

func LoadAllBastardies(opts ...types.ChangeStringFunc) chan either.Either[[]*Trait] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/marriage_traits/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	if strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}
	fn := dirname + "bastardies.json"
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

func LoadAllConsanguinities(opts ...types.ChangeStringFunc) chan either.Either[[]*Trait] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/marriage_traits/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	if strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}
	fn := dirname + "consanguinities.json"
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

func LoadAllDivorceOpts(opts ...types.ChangeStringFunc) chan either.Either[[]*PermissionTrait] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/marriage_traits/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	if strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}
	fn := dirname + "divorce_traits.json"
	ch := make(chan either.Either[[]*PermissionTrait], MaxLoadDataConcurrency)
	go func() {
		b, err := os.ReadFile(fn)
		if err != nil {
			ch <- either.Either[[]*PermissionTrait]{Err: we.NewInternalServerError(err, fmt.Sprintf("can not read file by filename (filename=%s)", fn))}
			return
		}
		var ts []*PermissionTrait
		if err := json.Unmarshal(b, &ts); err != nil {
			ch <- either.Either[[]*PermissionTrait]{Err: err}
			return
		}
		for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, ts) {
			ch <- either.Either[[]*PermissionTrait]{Value: chunk}
		}

		close(ch)
	}()

	return ch
}

func (mt MarriageTradition) ExtractTraitSlugs() []string {
	if mt.IsZero() {
		return []string{}
	}

	out := make([]string, 0, 3)
	if mt.Bastardry != nil {
		out = append(out, mt.Bastardry.Slug)
	}
	if mt.Kind != nil {
		out = append(out, mt.Kind.Slug)
	}
	if mt.Consanguinity != nil {
		out = append(out, mt.Consanguinity.Slug)
	}

	return out
}

type PureMarriageTradition struct {
	Kind          *PureTrait `json:"kind"`
	Bastardry     *PureTrait `json:"bastardy"`
	Consanguinity *PureTrait `json:"consanguinity"`
	Divorce       *PureTrait `json:"divorce"`
}

func PurifyMarriageTradition(in MarriageTradition) PureMarriageTradition {
	return PureMarriageTradition{
		Kind:          PurifyTrait(in.Kind),
		Bastardry:     PurifyTrait(in.Bastardry),
		Consanguinity: PurifyTrait(in.Consanguinity),
		Divorce:       PurifyPermissionTrait(in.Divorce),
	}
}
