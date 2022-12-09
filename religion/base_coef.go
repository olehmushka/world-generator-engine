package religion

type BaseCoefType string

const (
	LowBaseCoef     = "low_base_coef"
	RegularBaseCoef = "regular_base_coef"
	HighBaseCoef    = "high_base_coef"
)

func GetBaseCoef(cfg StatsConfig, t BaseCoefType) float64 {
	switch t {
	case LowBaseCoef:
		return cfg.LowBaseCoef
	case RegularBaseCoef:
		return cfg.BaseCoef
	case HighBaseCoef:
		return cfg.HighBaseCoef
	default:
		return 0
	}
}
