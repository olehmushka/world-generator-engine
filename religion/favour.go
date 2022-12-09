package religion

import (
	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
	we "github.com/olehmushka/golang-toolkit/wrapped_error"
	"github.com/olehmushka/world-generator-engine/favour"
)

type FavourTrait struct {
	Favour               favour.Favour `json:"favour"`
	BaseCoef             BaseCoefType  `json:"base_coef"`
	Stats                *Stats        `json:"stats"`
	OmitTypeSlugs        []string      `json:"omit_type_slugs"`
	OmitTraitSlugs       []string      `json:"omit_trait_slugs"`
	OmitStatsKeys        []string      `json:"omit_stats_keys"`
	OmitDominatedGenders []string      `json:"omit_dominated_genders"`
}

func FilterFavourTraits(cfg StatsConfig, r *Religion, in []*FavourTrait) ([]*FavourTrait, error) {
	if len(in) == 0 {
		return nil, we.NewInternalServerError(nil, "can not filter zero len favours list")
	}

	out := make([]*FavourTrait, 0, len(in))
	includedTraitSlugs := ExtractTraitSlugs(r)
	statsKeys := r.Stats.GetActualKeys()
	for count := 0; count < 100; count++ {
		for _, t := range in {
			// if iterated favour trait is not compatible with religion type it should be skipped
			if sliceTools.Includes(t.OmitTypeSlugs, r.Type.Slug) {
				continue
			}
			// if iterated favour trait is not compatible with religion traits it should be skipped
			if sliceTools.Includes(t.OmitDominatedGenders, r.GenderDominance.DominatedSex.String()) {
				continue
			}
			// if iterated favour trait is not compatible with religion dominated gender it should be skipped
			if cross := sliceTools.GetCrossOfSlices(
				includedTraitSlugs,
				t.OmitTraitSlugs,
				func(left, right string) bool { return left == right },
			); len(cross) > 0 {
				continue
			}
			// if iterated favour trait is not compatible with religion stats keys it should be skipped
			if cross := sliceTools.GetCrossOfSlices(
				statsKeys,
				t.OmitStatsKeys,
				func(left, right string) bool { return left == right },
			); len(cross) > 0 {
				continue
			}

			out = append(out, t)
		}
	}

	return out, nil
}

func ExtractFavoursFromFavourTraits(ts []*FavourTrait) []favour.Favour {
	if len(ts) == 0 {
		return []favour.Favour{}
	}

	out := make([]favour.Favour, 0, len(ts))
	for _, t := range ts {
		out = append(out, t.Favour)
	}

	return out
}

func FindFavourTrait(ts []*FavourTrait, f favour.Favour) *FavourTrait {
	for _, el := range ts {
		if el.Favour == f {
			return el
		}
	}

	return nil
}

func PurifyFavourTrait(in *FavourTrait) *PureTrait {
	if in == nil {
		return nil
	}

	return &PureTrait{
		Slug:        in.Favour.String(),
		Description: "",
	}
}
