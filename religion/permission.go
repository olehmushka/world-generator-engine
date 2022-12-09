package religion

import (
	"fmt"

	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	we "github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/world-generator-engine/permission"
)

type PermissionTrait struct {
	Permission           permission.Permission `json:"permission"`
	BaseCoef             BaseCoefType          `json:"base_coef"`
	Stats                *Stats                `json:"stats"`
	OmitTypeSlugs        []string              `json:"omit_type_slugs"`
	OmitTraitSlugs       []string              `json:"omit_trait_slugs"`
	OmitStatsKeys        []string              `json:"omit_stats_keys"`
	OmitDominatedGenders []string              `json:"omit_dominated_genders"`
}

func FilterPermissionTraits(cfg StatsConfig, r *Religion, in []*PermissionTrait, min, max int) ([]*PermissionTrait, *Stats, error) {
	if min < 0 {
		return nil, nil, we.NewInternalServerError(nil, "min expected filtered permission traits can not be less than 0")
	}
	if max > len(in) {
		return nil, nil, we.NewInternalServerError(nil, fmt.Sprintf("max expected filtered can not be greater than all permission traits length (len=%d)", len(in)))
	}

	out := make([]*PermissionTrait, 0, len(in))
	includedTraitSlugs := ExtractTraitSlugs(r)
	statsKeys := r.Stats.GetActualKeys()
	for count := 0; count < 100; count++ {
		for _, t := range in {
			if len(out) == max {
				break
			}
			// if iterated permission trait is not compatible with religion type it should be skipped
			if sliceTools.Includes(t.OmitTypeSlugs, r.Type.Slug) {
				continue
			}
			// if iterated permission trait is not compatible with religion traits it should be skipped
			if sliceTools.Includes(t.OmitDominatedGenders, r.GenderDominance.DominatedSex.String()) {
				continue
			}
			// if iterated permission trait is not compatible with religion dominated gender it should be skipped
			if cross := sliceTools.GetCrossOfSlices(
				includedTraitSlugs,
				t.OmitTraitSlugs,
				func(left, right string) bool { return left == right },
			); len(cross) > 0 {
				continue
			}
			// if iterated permission trait is not compatible with religion stats keys it should be skipped
			if cross := sliceTools.GetCrossOfSlices(
				statsKeys,
				t.OmitStatsKeys,
				func(left, right string) bool { return left == right },
			); len(cross) > 0 {
				continue
			}

			ok, err := CalcProbFromReligionStats(GetBaseCoef(cfg, t.BaseCoef), r.Stats, t.Stats, CalcProbOpts{})
			if err != nil {
				return nil, nil, we.NewInternalServerError(err, "can not calc probability of permission trait comatibility")
			}
			if ok {
				out = append(out, t)
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

func PurifyPermissionTrait(in *PermissionTrait) *PureTrait {
	if in == nil {
		return nil
	}

	return &PureTrait{
		Slug:        in.Permission.String(),
		Description: "",
	}
}
