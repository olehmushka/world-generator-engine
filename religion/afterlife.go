package religion

import (
	"encoding/json"
	"fmt"
	"math"
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

type AfterlifeExist struct {
	Exists               bool         `json:"exists"`
	BaseCoef             BaseCoefType `json:"base_coef"`
	Stats                *Stats       `json:"stats"`
	OmitTypeSlugs        []string     `json:"omit_type_slugs"`
	OmitTraitSlugs       []string     `json:"omit_trait_slugs"`
	OmitStatsKeys        []string     `json:"omit_stats_keys"`
	OmitDominatedGenders []string     `json:"omit_dominated_genders"`
}

func (ale AfterlifeExist) IsZero() bool {
	return ale.Exists == false &&
		ale.BaseCoef == "" &&
		ale.Stats == nil
}

func LoadAllAfterlifeExists(opts ...types.ChangeStringFunc) chan either.Either[[]AfterlifeExist] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/afterlife/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	if strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}
	fn := dirname + "exists.json"
	ch := make(chan either.Either[[]AfterlifeExist], MaxLoadDataConcurrency)
	go func() {
		b, err := os.ReadFile(fn)
		if err != nil {
			ch <- either.Either[[]AfterlifeExist]{Err: we.NewInternalServerError(err, fmt.Sprintf("can not read file by filename (filename=%s)", fn))}
			return
		}
		var ts []AfterlifeExist
		if err := json.Unmarshal(b, &ts); err != nil {
			ch <- either.Either[[]AfterlifeExist]{Err: err}
			return
		}
		for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, ts) {
			ch <- either.Either[[]AfterlifeExist]{Value: chunk}
		}

		close(ch)
	}()

	return ch
}

type Afterlife struct {
	Exists bool `json:"exists"`

	Participants map[AfterlifeParticipantSlug]AfterlifeParticipance `json:"participants"`
}

type AfterlifeParticipance struct {
	Slug         string  `json:"slug"`
	GoodnessCoef float64 `json:"goodness_coef"`
}

func LoadAllAfterlifeParticipances(opts ...types.ChangeStringFunc) chan either.Either[[]AfterlifeParticipance] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/afterlife/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	if strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}
	fn := dirname + "participances.json"
	ch := make(chan either.Either[[]AfterlifeParticipance], MaxLoadDataConcurrency)
	go func() {
		b, err := os.ReadFile(fn)
		if err != nil {
			ch <- either.Either[[]AfterlifeParticipance]{Err: we.NewInternalServerError(err, fmt.Sprintf("can not read file by filename (filename=%s)", fn))}
			return
		}
		var ts []AfterlifeParticipance
		if err := json.Unmarshal(b, &ts); err != nil {
			ch <- either.Either[[]AfterlifeParticipance]{Err: err}
			return
		}
		for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, ts) {
			ch <- either.Either[[]AfterlifeParticipance]{Value: chunk}
		}

		close(ch)
	}()

	return ch
}

type AfterlifeParticipantSlug string

type AfterlifeParticipant struct {
	Slug        AfterlifeParticipantSlug `json:"slug"`
	Description string                   `json:"description"`
	Order       int                      `json:"order"`
}

func LoadAllAfterlifeParticipants(opts ...types.ChangeStringFunc) chan either.Either[[]AfterlifeParticipant] {
	_, filename, _, _ := runtime.Caller(1)
	currDirname := tools.PreparePath(path.Dir(filename), "religion")
	dirname := currDirname + "data/afterlife/"
	for _, fn := range opts {
		dirname = fn(dirname)
	}
	if strings.HasSuffix(dirname, "/") {
		dirname += "/"
	}
	fn := dirname + "participants.json"
	ch := make(chan either.Either[[]AfterlifeParticipant], MaxLoadDataConcurrency)
	go func() {
		b, err := os.ReadFile(fn)
		if err != nil {
			ch <- either.Either[[]AfterlifeParticipant]{Err: we.NewInternalServerError(err, fmt.Sprintf("can not read file by filename (filename=%s)", fn))}
			return
		}
		var ts []AfterlifeParticipant
		if err := json.Unmarshal(b, &ts); err != nil {
			ch <- either.Either[[]AfterlifeParticipant]{Err: err}
			return
		}
		for _, chunk := range sliceTools.Chunk(MaxLoadDataChunkSize, ts) {
			ch <- either.Either[[]AfterlifeParticipant]{Value: chunk}
		}

		close(ch)
	}()

	return ch
}

type CreateAfterlifeData struct {
	AfterlifeParticipances []AfterlifeParticipance
	AfterlifeParticipants  []AfterlifeParticipant
	AllDeityFavours        []*FavourTrait
	AllAfterlifeExistsOpts []AfterlifeExist
}

func NewAfterlife(cfg StatsConfig, r *Religion, data CreateAfterlifeData) (Afterlife, *Stats, error) {
	stats := r.Stats
	out := Afterlife{}
	existanceOpt, calcStats, err := CalcAfterlifeExistance(cfg, r, data.AllAfterlifeExistsOpts)
	if err != nil {
		return Afterlife{}, nil, err
	}
	stats = calcStats
	out.Exists = existanceOpt.Exists
	if !out.Exists {
		return out, stats, nil
	}

	var (
		highestCoef      = (r.DeityFavour.Favour.Float64() + 1) / 2
		lowestCoef       = 1 - highestCoef
		step             = (highestCoef - lowestCoef) / float64(len(data.AfterlifeParticipants))
		participantCoefs = make([]map[AfterlifeParticipantSlug]float64, len(data.AfterlifeParticipants))
	)
	fmt.Printf("favour:%f\nhigh:%f\nlow:%f\nstep:%f\n\n", r.DeityFavour.Favour.Float64(), highestCoef, lowestCoef, step)
	for _, p := range data.AfterlifeParticipants {
		participantCoefs[p.Order] = map[AfterlifeParticipantSlug]float64{
			p.Slug: highestCoef - float64(p.Order)*step,
		}
	}

	participants := make(map[AfterlifeParticipantSlug]AfterlifeParticipance, len(data.AfterlifeParticipants))
	for _, pc := range participantCoefs {
		var (
			key  AfterlifeParticipantSlug
			coef float64
			p    AfterlifeParticipance
			diff = math.MaxFloat64
		)
		for k, v := range pc {
			key = k
			coef = v
		}
		fmt.Printf("part:%s,coef:%f\n", key, coef)

		for _, d := range data.AfterlifeParticipances {
			if calcDiff := math.Abs(d.GoodnessCoef - coef); diff > calcDiff {
				diff = calcDiff
				p = d
			}
		}

		participants[key] = p
	}
	out.Participants = participants

	return out, stats, nil
}

func CalcAfterlifeExistance(cfg StatsConfig, r *Religion, opts []AfterlifeExist) (AfterlifeExist, *Stats, error) {
	if len(opts) != 2 {
		return AfterlifeExist{}, nil, we.NewInternalServerError(nil, "can not calc afterlife existance if opts len is not 2")
	}

	var out AfterlifeExist
	includedTraitSlugs := ExtractTraitSlugs(r)
	statsKeys := r.Stats.GetActualKeys()
	for count := 0; count < 100; count++ {
		for _, t := range sliceTools.Shuffle(opts) {
			if !out.IsZero() {
				break
			}
			// if iterated trait is not compatible with religion traits it should be skipped
			if sliceTools.Includes(t.OmitDominatedGenders, r.GenderDominance.DominatedSex.String()) {
				continue
			}
			// if iterated trait is not compatible with religion dominated gender it should be skipped
			if cross := sliceTools.GetCrossOfSlices(
				includedTraitSlugs,
				t.OmitTraitSlugs,
				func(left, right string) bool { return left == right },
			); len(cross) > 0 {
				continue
			}
			// if iterated trait is not compatible with religion stats keys it should be skipped
			if cross := sliceTools.GetCrossOfSlices(
				statsKeys,
				t.OmitStatsKeys,
				func(left, right string) bool { return left == right },
			); len(cross) > 0 {
				continue
			}

			ok, err := CalcProbFromReligionStats(GetBaseCoef(cfg, t.BaseCoef), r.Stats, t.Stats, CalcProbOpts{})
			if err != nil {
				return AfterlifeExist{}, nil, we.NewInternalServerError(err, "can not calc probability of trait comatibility")
			}
			if ok {
				out = t
			}
		}
		if !out.IsZero() {
			break
		}
	}

	stats := r.Stats
	merged, err := MergeReligionStats(cfg, stats, out.Stats)
	if err != nil {
		return AfterlifeExist{}, nil, err
	}
	stats = merged

	return out, stats, nil
}
