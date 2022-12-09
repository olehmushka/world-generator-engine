package religion

import (
	"fmt"

	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	slicetools "github.com/olehmushka/golang-toolkit/slice_tools"
	we "github.com/olehmushka/golang-toolkit/wrapped_error"
)

type Trait struct {
	Slug                 string       `json:"slug"`
	Description          string       `json:"description"`
	Kind                 string       `json:"kind"`
	BaseCoef             BaseCoefType `json:"base_coef"`
	Stats                *Stats       `json:"stats"`
	OmitTypeSlugs        []string     `json:"omit_type_slugs"`
	OmitTraitSlugs       []string     `json:"omit_trait_slugs"`
	OmitStatsKeys        []string     `json:"omit_stats_keys"`
	OmitDominatedGenders []string     `json:"omit_dominated_genders"`
}

func FilterTraits(cfg StatsConfig, r *Religion, in []*Trait, min, max int) ([]*Trait, *Stats, error) {
	if min < 0 {
		return nil, nil, we.NewInternalServerError(nil, "min expected filtered traits can not be less than 0")
	}
	if max > len(in) {
		return nil, nil, we.NewInternalServerError(nil, fmt.Sprintf("max expected filtered can not be greater than traitsToSelect length (len=%d)", len(in)))
	}

	out := make([]*Trait, 0, len(in))
	includedTraitSlugs := ExtractTraitSlugs(r)
	statsKeys := r.Stats.GetActualKeys()
	for count := 0; count < 20; count++ {
		for _, t := range sliceTools.Shuffle(in) {
			if len(out) == max {
				break
			}
			// if iterated trait is not compatible with religion type it should be skipped
			if sliceTools.Includes(t.OmitTypeSlugs, r.Type.Slug) {
				continue
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
				return nil, nil, we.NewInternalServerError(err, "can not calc probability of trait comatibility")
			}
			if ok {
				out = UniqueTraits(append(out, t))
				includedTraitSlugs = slicetools.Unique(append(includedTraitSlugs, t.Slug))
			}
		}
		if len(out) == max {
			break
		}
		if len(out) >= min {
			break
		}
	}

	stats := r.Stats
	for _, t := range out {
		merged, err := MergeReligionStats(cfg, stats, t.Stats)
		if err != nil {
			return nil, nil, err
		}
		stats = merged
	}

	return out, stats, nil
}

func FilterTypeTrait(cfg StatsConfig, r *Religion, in []*Trait) (*Trait, *Stats, error) {
	if len(in) < 1 {
		return nil, nil, we.NewInternalServerError(nil, "can not pick type trait from zero len list")
	}

	var out *Trait
	includedTraitSlugs := ExtractTraitSlugs(r)
	statsKeys := r.Stats.GetActualKeys()
	for count := 0; count < 20; count++ {
		for _, t := range sliceTools.Shuffle(in) {
			if out != nil {
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
				return nil, nil, we.NewInternalServerError(err, "can not calc probability of trait comatibility")
			}
			if ok {
				out = t
			}
		}
		if out != nil {
			break
		}
	}

	stats := r.Stats
	merged, err := MergeReligionStats(cfg, stats, out.Stats)
	if err != nil {
		return nil, nil, err
	}
	stats = merged

	return out, stats, nil
}

func UniqueTraits(in []*Trait) []*Trait {
	if len(in) <= 1 {
		return in
	}

	preOut := make(map[string]*Trait)
	for _, c := range in {
		preOut[c.Slug] = c
	}

	out := make([]*Trait, 0, len(preOut))
	for _, c := range preOut {
		out = append(out, c)
	}

	return out
}

type PureTrait struct {
	Slug        string `json:"slug"`
	Description string `json:"description"`
}

func PurifyTrait(in *Trait) *PureTrait {
	if in == nil {
		return nil
	}

	return &PureTrait{
		Slug:        in.Slug,
		Description: in.Description,
	}
}

func PurifyTraits(in []*Trait) []*PureTrait {
	if len(in) == 0 {
		return []*PureTrait{}
	}

	out := make([]*PureTrait, len(in))
	for i := range out {
		out[i] = PurifyTrait(in[i])
	}

	return out
}
