package religion

import (
	"fmt"

	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	we "github.com/olehmushka/golang-toolkit/wrapped_error"
	actionAcceptance "github.com/olehmushka/world-generator-engine/action_acceptance"
)

type StatsConfig struct {
	BaseCoef      float64 `json:"base_coef"`
	LowBaseCoef   float64 `json:"low_base_coef"`
	HighBaseCoef  float64 `json:"high_base_coef"`
	MaxStatsValue float64 `json:"max_stats_value"`
}

func NewStatsConfig() (StatsConfig, error) {
	lowBaseCoef, err := randomTools.RandFloat64InRange(0.45, 0.75)
	if err != nil {
		return StatsConfig{}, err
	}
	baseCoef, err := randomTools.RandFloat64InRange(0.95, 1.05)
	if err != nil {
		return StatsConfig{}, err
	}
	highBaseCoef, err := randomTools.RandFloat64InRange(1, 1.25)
	if err != nil {
		return StatsConfig{}, err
	}
	maxMetadataValue, err := randomTools.RandFloat64InRange(8, 10)
	if err != nil {
		return StatsConfig{}, err
	}

	return StatsConfig{
		BaseCoef:      baseCoef,
		LowBaseCoef:   lowBaseCoef,
		HighBaseCoef:  highBaseCoef,
		MaxStatsValue: maxMetadataValue,
	}, nil
}

type Stats struct {
	Restricted float64 `json:"restricted"`
	Open       float64 `json:"open"`

	Celestial float64 `json:"celestial"`
	Chthonic  float64 `json:"chthonic"`

	SexualEmancipated float64 `json:"sexual_emancipated"`
	SexualStrictness  float64 `json:"sexual_strictness"`

	Aggressive float64 `json:"aggressive"`
	Pacifistic float64 `json:"pacifistic"`

	Hedonistic float64 `json:"hedonistic"`
	Ascetic    float64 `json:"ascetic"`

	Lawful  float64 `json:"lawful"`
	Anarchy float64 `json:"anarchy"`

	Pragmatic  float64 `json:"pragmatic"`
	Altruistic float64 `json:"altruistic"`

	Naturalistic float64 `json:"naturalistic"`
	Urbanistic   float64 `json:"urbanistic"`

	Philosophic float64 `json:"philosophic"`
	Primitive   float64 `json:"primitive"`

	Authoritaristic float64 `json:"authoritaristic"`
	Liberal         float64 `json:"liberal"`

	Individualistic float64 `json:"individualistic"`
	Collectivistic  float64 `json:"collectivistic"`

	Spiritualustic float64 `json:"spiritualustic"`
	Materialistic  float64 `json:"materialistic"`
}

func (s *Stats) IsZero() bool {
	return s == nil
}

func (s *Stats) GetActualKeys() []string {
	out := make([]string, 0, 22)
	if s.IsRestricted() {
		out = append(out, s.GetRestrictedKey())
	}
	if s.IsOpen() {
		out = append(out, s.GetOpenKey())
	}
	if s.IsCelestial() {
		out = append(out, s.GetCelestialKey())
	}
	if s.IsChthonic() {
		out = append(out, s.GetChthonicKey())
	}
	if s.IsSexualEmancipated() {
		out = append(out, s.GetSexualEmancipatedKey())
	}
	if s.IsSexualStrictness() {
		out = append(out, s.GetSexualStrictnessKey())
	}
	if s.IsAggressive() {
		out = append(out, s.GetAggressiveKey())
	}
	if s.IsPacifistic() {
		out = append(out, s.GetPacifisticKey())
	}
	if s.IsHedonistic() {
		out = append(out, s.GetHedonisticKey())
	}
	if s.IsAscetic() {
		out = append(out, s.GetAsceticKey())
	}
	if s.IsLawful() {
		out = append(out, s.GetLawfulKey())
	}
	if s.IsAnarchy() {
		out = append(out, s.GetAnarchyKey())
	}
	if s.IsPragmatic() {
		out = append(out, s.GetPragmaticKey())
	}
	if s.IsAltruistic() {
		out = append(out, s.GetAltruisticKey())
	}
	if s.IsNaturalistic() {
		out = append(out, s.GetNaturalisticKey())
	}
	if s.IsUrbanistic() {
		out = append(out, s.GetUrbanisticKey())
	}
	if s.IsPhilosophic() {
		out = append(out, s.GetPhilosophicKey())
	}
	if s.IsPrimitive() {
		out = append(out, s.GetPrimitiveKey())
	}
	if s.IsAuthoritaristic() {
		out = append(out, s.GetAuthoritaristicKey())
	}
	if s.IsLiberal() {
		out = append(out, s.GetLiberalKey())
	}
	if s.IsIndividualistic() {
		out = append(out, s.GetIndividualisticKey())
	}
	if s.IsCollectivistic() {
		out = append(out, s.GetCollectivisticKey())
	}
	if s.IsSpiritualustic() {
		out = append(out, s.GetSpiritualusticKey())
	}
	if s.IsMaterialistic() {
		out = append(out, s.GetMaterialisticKey())
	}

	return out
}

func (s *Stats) GetRestrictedKey() string {
	return "restricted_stats_key"
}

func (s *Stats) IsRestricted() bool {
	if s.IsZero() {
		return false
	}

	return s.Restricted >= 2
}

func (s *Stats) GetOpenKey() string {
	return "open_stats_key"
}

func (s *Stats) IsOpen() bool {
	if s.IsZero() {
		return false
	}

	return s.Open >= 2
}

func (s *Stats) GetCelestialKey() string {
	return "celestial_stats_key"
}

func (s *Stats) IsCelestial() bool {
	if s.IsZero() {
		return false
	}

	return s.Celestial >= 2
}

func (s *Stats) GetChthonicKey() string {
	return "chthonic_stats_key"
}

func (s *Stats) IsChthonic() bool {
	if s.IsZero() {
		return false
	}

	return s.Chthonic >= 2
}

func (s *Stats) GetSexualEmancipatedKey() string {
	return "sexual_emancipated_stats_key"
}

func (s *Stats) IsSexualEmancipated() bool {
	if s.IsZero() {
		return false
	}

	return s.SexualEmancipated >= 2
}

func (s *Stats) GetSexualStrictnessKey() string {
	return "sexual_strictness_stats_key"
}

func (s *Stats) IsSexualStrictness() bool {
	if s.IsZero() {
		return false
	}

	return s.SexualStrictness >= 2
}

func (s *Stats) GetAggressiveKey() string {
	return "aggressive_stats_key"
}

func (s *Stats) IsAggressive() bool {
	if s.IsZero() {
		return false
	}

	return s.Aggressive >= 2
}

func (s *Stats) GetPacifisticKey() string {
	return "pacifistic_stats_key"
}

func (s *Stats) IsPacifistic() bool {
	if s.IsZero() {
		return false
	}

	return s.Pacifistic >= 2
}

func (s *Stats) GetHedonisticKey() string {
	return "hedonistic_stats_key"
}

func (s *Stats) IsHedonistic() bool {
	if s.IsZero() {
		return false
	}

	return s.Hedonistic >= 2
}

func (s *Stats) GetAsceticKey() string {
	return "ascetic_stats_key"
}

func (s *Stats) IsAscetic() bool {
	if s.IsZero() {
		return false
	}

	return s.Ascetic >= 2
}

func (s *Stats) GetLawfulKey() string {
	return "lawful_stats_key"
}

func (s *Stats) IsLawful() bool {
	if s.IsZero() {
		return false
	}

	return s.Lawful >= 2
}

func (s *Stats) GetAnarchyKey() string {
	return "anarchy_stats_key"
}

func (s *Stats) IsAnarchy() bool {
	if s.IsZero() {
		return false
	}

	return s.Anarchy >= 2
}

func (s *Stats) GetPragmaticKey() string {
	return "pragmatic_stats_key"
}

func (s *Stats) IsPragmatic() bool {
	if s.IsZero() {
		return false
	}

	return s.Pragmatic >= 2
}

func (s *Stats) GetAltruisticKey() string {
	return "altruistic_stats_key"
}

func (s *Stats) IsAltruistic() bool {
	if s.IsZero() {
		return false
	}

	return s.Altruistic >= 2
}

func (s *Stats) GetNaturalisticKey() string {
	return "naturalistic_stats_key"
}

func (s *Stats) IsNaturalistic() bool {
	if s.IsZero() {
		return false
	}

	return s.Naturalistic >= 2
}

func (s *Stats) GetUrbanisticKey() string {
	return "urbanistic_stats_key"
}

func (s *Stats) IsUrbanistic() bool {
	if s.IsZero() {
		return false
	}

	return s.Urbanistic >= 2
}

func (s *Stats) GetPhilosophicKey() string {
	return "philosophic_stats_key"
}

func (s *Stats) IsPhilosophic() bool {
	if s.IsZero() {
		return false
	}

	return s.Philosophic >= 2
}

func (s *Stats) GetPrimitiveKey() string {
	return "primitive_stats_key"
}

func (s *Stats) IsPrimitive() bool {
	if s.IsZero() {
		return false
	}

	return s.Primitive >= 2
}

func (s *Stats) GetAuthoritaristicKey() string {
	return "authoritaristic_stats_key"
}

func (s *Stats) IsAuthoritaristic() bool {
	if s.IsZero() {
		return false
	}

	return s.Authoritaristic >= 2
}

func (s *Stats) GetLiberalKey() string {
	return "liberal_stats_key"
}

func (s *Stats) IsLiberal() bool {
	if s.IsZero() {
		return false
	}

	return s.Liberal >= 2
}

func (s *Stats) GetIndividualisticKey() string {
	return "individualistic_stats_key"
}

func (s *Stats) IsIndividualistic() bool {
	if s.IsZero() {
		return false
	}

	return s.Individualistic >= 2
}

func (s *Stats) GetCollectivisticKey() string {
	return "collectivistic_stats_key"
}

func (s *Stats) IsCollectivistic() bool {
	if s.IsZero() {
		return false
	}

	return s.Collectivistic >= 2
}

func (s *Stats) IsSpiritualustic() bool {
	if s.IsZero() {
		return false
	}

	return s.Spiritualustic >= 2
}

func (s *Stats) GetSpiritualusticKey() string {
	return "spiritualustic_stats_key"
}

func (s *Stats) IsMaterialistic() bool {
	if s.IsZero() {
		return false
	}

	return s.Materialistic >= 2
}

func (s *Stats) GetMaterialisticKey() string {
	return "materialistic_stats_key"
}

func mergeReligionStats(cfg StatsConfig, left, right *Stats, coef float64) (*Stats, error) {
	if right.Restricted > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Restricted+right.Restricted*coef)
		if err != nil {
			return nil, err
		}
		left.Restricted = p
	}

	if right.Open > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Open+right.Open*coef)
		if err != nil {
			return nil, err
		}
		left.Open = p
	}

	if right.Celestial > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Celestial+right.Celestial*coef)
		if err != nil {
			return nil, err
		}
		left.Celestial = p
	}

	if right.Chthonic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Chthonic+right.Chthonic*coef)
		if err != nil {
			return nil, err
		}
		left.Chthonic = p
	}

	if right.SexualEmancipated > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.SexualEmancipated+right.SexualEmancipated*coef)
		if err != nil {
			return nil, err
		}
		left.SexualEmancipated = p
	}

	if right.SexualStrictness > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.SexualStrictness+right.SexualStrictness*coef)
		if err != nil {
			return nil, err
		}
		left.SexualStrictness = p
	}

	if right.Aggressive > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Aggressive+right.Aggressive*coef)
		if err != nil {
			return nil, err
		}
		left.Aggressive = p
	}

	if right.Pacifistic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Pacifistic+right.Pacifistic*coef)
		if err != nil {
			return nil, err
		}
		left.Pacifistic = p
	}

	if right.Hedonistic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Hedonistic+right.Hedonistic*coef)
		if err != nil {
			return nil, err
		}
		left.Hedonistic = p
	}

	if right.Ascetic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Ascetic+right.Ascetic*coef)
		if err != nil {
			return nil, err
		}
		left.Ascetic = p
	}

	if right.Lawful > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Lawful+right.Lawful*coef)
		if err != nil {
			return nil, err
		}
		left.Lawful = p
	}

	if right.Anarchy > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Anarchy+right.Anarchy*coef)
		if err != nil {
			return nil, err
		}
		left.Anarchy = p
	}

	if right.Pragmatic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Pragmatic+right.Pragmatic*coef)
		if err != nil {
			return nil, err
		}
		left.Pragmatic = p
	}

	if right.Altruistic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Altruistic+right.Altruistic*coef)
		if err != nil {
			return nil, err
		}
		left.Altruistic = p
	}

	if right.Naturalistic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Naturalistic+right.Naturalistic*coef)
		if err != nil {
			return nil, err
		}
		left.Naturalistic = p
	}

	if right.Urbanistic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Urbanistic+right.Urbanistic*coef)
		if err != nil {
			return nil, err
		}
		left.Urbanistic = p
	}

	if right.Philosophic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Philosophic+right.Philosophic*coef)
		if err != nil {
			return nil, err
		}
		left.Philosophic = p
	}

	if right.Primitive > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Primitive+right.Primitive*coef)
		if err != nil {
			return nil, err
		}
		left.Primitive = p
	}

	if right.Authoritaristic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Authoritaristic+right.Authoritaristic*coef)
		if err != nil {
			return nil, err
		}
		left.Authoritaristic = p
	}

	if right.Liberal > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Liberal+right.Liberal*coef)
		if err != nil {
			return nil, err
		}
		left.Liberal = p
	}

	if right.Individualistic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Individualistic+right.Individualistic*coef)
		if err != nil {
			return nil, err
		}
		left.Individualistic = p
	}

	if right.Collectivistic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Collectivistic+right.Collectivistic*coef)
		if err != nil {
			return nil, err
		}
		left.Collectivistic = p
	}

	if right.Spiritualustic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Spiritualustic+right.Spiritualustic*coef)
		if err != nil {
			return nil, err
		}
		left.Spiritualustic = p
	}

	if right.Materialistic > 0 {
		p, err := prepareReligionStatsValueBeforeUpdate(cfg, left.Materialistic+right.Materialistic*coef)
		if err != nil {
			return nil, err
		}
		left.Materialistic = p
	}

	return left, nil
}

func MergeReligionStats(cfg StatsConfig, left, right *Stats) (*Stats, error) {
	return mergeReligionStats(cfg, left, right, 1)
}

func MergeReligionStatsByAcceptance(cfg StatsConfig, left, right *Stats, a actionAcceptance.Acceptance) (*Stats, error) {
	var coef float64
	switch a {
	case actionAcceptance.Accepted:
		c, err := randomTools.RandFloat64InRange(0.3, 0.5)
		if err != nil {
			return nil, we.NewInternalServerError(err, fmt.Sprintf("can not merge religion stats by acceptance (acceptance=%s)", a))
		}
		coef = c
	case actionAcceptance.Shunned:
		c, err := randomTools.RandFloat64InRange(0.05, 0.25)
		if err != nil {
			return nil, we.NewInternalServerError(err, fmt.Sprintf("can not merge religion stats by acceptance (acceptance=%s)", a))
		}
		coef = -c
	case actionAcceptance.Criminal:
		c, err := randomTools.RandFloat64InRange(0.3, 0.5)
		if err != nil {
			return nil, we.NewInternalServerError(err, fmt.Sprintf("can not merge religion stats by acceptance (acceptance=%s)", a))
		}
		coef = -c
	}
	return mergeReligionStats(cfg, left, right, coef)
}

func prepareReligionStatsValueBeforeUpdate(cfg StatsConfig, v float64) (float64, error) {
	var out float64
	if v > 0 {
		out = v
	}

	switch rel := out / cfg.MaxStatsValue; {
	case rel < 0.25:
		p, err := randomTools.RandFloat64InRange(1, 1.01)
		if err != nil {
			return 0, we.NewInternalServerError(err, fmt.Sprintf("can not prepare religion stats value before update (rel=%.2f, out=%.2f, max_stats_value=%.2f)", rel, out, cfg.MaxStatsValue))
		}

		return out * p, nil
	case rel < 0.5:
		p, err := randomTools.RandFloat64InRange(1, 1.001)
		if err != nil {
			return 0, we.NewInternalServerError(err, fmt.Sprintf("can not prepare religion stats value before update (rel=%.2f, out=%.2f, max_stats_value=%.2f)", rel, out, cfg.MaxStatsValue))
		}

		return out * p, nil
	case rel < 0.75:
		return out, nil
	case rel < 0.9:
		p, err := randomTools.RandFloat64InRange(0.98, 0.9999)
		if err != nil {
			return 0, we.NewInternalServerError(err, fmt.Sprintf("can not prepare religion stats value before update (rel=%.2f, out=%.2f, max_stats_value=%.2f)", rel, out, cfg.MaxStatsValue))
		}

		return out * p, nil
	case rel < 1:
		p, err := randomTools.RandFloat64InRange(0.98, 0.99)
		if err != nil {
			return 0, we.NewInternalServerError(err, fmt.Sprintf("can not prepare religion stats value before update (rel=%.2f, out=%.2f, max_stats_value=%.2f)", rel, out, cfg.MaxStatsValue))
		}

		return out * p, nil
	default:
		return cfg.MaxStatsValue, nil
	}
}

type CalcProbOpts struct {
	Log   bool
	Label string
}

func calcRSProb(coef float64, isMatchSame, isMatchContrary bool) (float64, error) {
	sameCoef, err := randomTools.RandFloat64InRange(0.4, 0.8)
	if err != nil {
		return 0, err
	}
	contraryCoef, err := randomTools.RandFloat64InRange(0.3, 0.7)
	if err != nil {
		return 0, err
	}
	newCoef, err := randomTools.RandFloat64InRange(0.5, 1)
	if err != nil {
		return 0, err
	}

	if isMatchSame && isMatchContrary {
		return coef * randomTools.PrepareProbability(sameCoef-contraryCoef), nil
	}
	if isMatchSame && !isMatchContrary {
		return coef * sameCoef, nil
	}
	if !isMatchSame && !isMatchContrary {
		return coef * sameCoef * newCoef, nil
	}

	return 0, nil
}

func CalcProbFromReligionStats(baseCoef float64, rStats, inStats *Stats, opts CalcProbOpts) (bool, error) {
	baseCoef = PrepareCoef(baseCoef)
	var (
		primaryProbability float64
		ideasCount         int
	)
	randCoef, err := randomTools.RandFloat64InRange(0.9, 1.1)
	if err != nil {
		return false, err
	}

	if inStats.Restricted > 0 {
		p, err := calcRSProb(randCoef, rStats.IsRestricted(), rStats.IsOpen())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Open > 0 {
		p, err := calcRSProb(randCoef, rStats.IsOpen(), rStats.IsRestricted())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Celestial > 0 {
		p, err := calcRSProb(randCoef, rStats.IsCelestial(), rStats.IsChthonic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Chthonic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsChthonic(), rStats.IsCelestial())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.SexualEmancipated > 0 {
		p, err := calcRSProb(randCoef, rStats.IsSexualEmancipated(), rStats.IsSexualStrictness())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.SexualStrictness > 0 {
		p, err := calcRSProb(randCoef, rStats.IsSexualStrictness(), rStats.IsSexualEmancipated())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Aggressive > 0 {
		p, err := calcRSProb(randCoef, rStats.IsAggressive(), rStats.IsPacifistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Pacifistic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsPacifistic(), rStats.IsAggressive())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Hedonistic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsHedonistic(), rStats.IsAscetic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Ascetic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsAscetic(), rStats.IsHedonistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Lawful > 0 {
		p, err := calcRSProb(randCoef, rStats.IsLawful(), rStats.IsAnarchy())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Anarchy > 0 {
		p, err := calcRSProb(randCoef, rStats.IsAnarchy(), rStats.IsLawful())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Pragmatic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsPragmatic(), rStats.IsAltruistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Altruistic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsAltruistic(), rStats.IsPragmatic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Naturalistic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsNaturalistic(), rStats.IsUrbanistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Urbanistic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsUrbanistic(), rStats.IsNaturalistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Philosophic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsPhilosophic(), rStats.IsPrimitive())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Primitive > 0 {
		p, err := calcRSProb(randCoef, rStats.IsPrimitive(), rStats.IsPhilosophic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Authoritaristic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsAuthoritaristic(), rStats.IsLiberal())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Liberal > 0 {
		p, err := calcRSProb(randCoef, rStats.IsLiberal(), rStats.IsAuthoritaristic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Individualistic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsIndividualistic(), rStats.IsCollectivistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Collectivistic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsCollectivistic(), rStats.IsIndividualistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	if inStats.Spiritualustic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsSpiritualustic(), rStats.IsMaterialistic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}
	if inStats.Materialistic > 0 {
		p, err := calcRSProb(randCoef, rStats.IsMaterialistic(), rStats.IsSpiritualustic())
		if err != nil {
			return false, err
		}
		primaryProbability += p
		ideasCount++
	}

	probability := baseCoef * primaryProbability
	if ideasCount > 0 {
		probability = probability / float64(ideasCount)
	}

	if opts.Log {
		fmt.Printf("\n>>>>>>>>>>\nLabel: %s\nprobability: %f\n<<<<<<<<<<<<\n", opts.Label, probability)
	}

	return randomTools.GetRandomBool(randomTools.PrepareProbability(probability))
}

func calcRSAcceptanceProb(coef float64, isMatchSame, isMatchContrary bool) (float64, float64, float64, error) {
	sameCoef, err := randomTools.RandFloat64InRange(0.4, 0.8)
	if err != nil {
		return 0, 0, 0, err
	}
	contraryCoef, err := randomTools.RandFloat64InRange(0.3, 0.7)
	if err != nil {
		return 0, 0, 0, err
	}
	newCoef, err := randomTools.RandFloat64InRange(0.5, 1)
	if err != nil {
		return 0, 0, 0, err
	}

	newProbability := coef * sameCoef * newCoef
	if isMatchSame && isMatchContrary {
		return coef * sameCoef, coef * randomTools.PrepareProbability(sameCoef-contraryCoef), coef * contraryCoef, nil
	}
	if isMatchSame && !isMatchContrary {
		return coef * sameCoef, newProbability, 0, nil
	}
	if !isMatchSame && isMatchContrary {
		return 0, newProbability, coef * contraryCoef, nil
	}
	if !isMatchSame && !isMatchContrary {
		return newProbability, newProbability, newProbability, nil
	}

	return 0, 0, 0, nil
}

func CalcAcceptanceFromReligionStats(acceptedBaseCoef, stunnedBaseCoef, criminalBaseCoef float64, rStats, inStats *Stats, opts CalcProbOpts) (actionAcceptance.Acceptance, error) {
	acceptedBaseCoef = PrepareCoef(acceptedBaseCoef)
	stunnedBaseCoef = PrepareCoef(stunnedBaseCoef)
	criminalBaseCoef = PrepareCoef(criminalBaseCoef)
	var (
		accepted, shunned, criminal float64
		ideasCount                  int
	)
	randCoef, err := randomTools.RandFloat64InRange(0.9, 1.1)
	if err != nil {
		return "", err
	}

	if inStats.Restricted > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsRestricted(), rStats.IsOpen())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Open > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsOpen(), rStats.IsRestricted())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Celestial > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsCelestial(), rStats.IsChthonic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Chthonic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsChthonic(), rStats.IsCelestial())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.SexualEmancipated > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsSexualEmancipated(), rStats.IsSexualStrictness())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.SexualStrictness > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsSexualStrictness(), rStats.IsSexualEmancipated())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Aggressive > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsAggressive(), rStats.IsPacifistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Pacifistic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsPacifistic(), rStats.IsAggressive())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Hedonistic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsHedonistic(), rStats.IsAscetic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Ascetic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsAscetic(), rStats.IsHedonistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Lawful > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsLawful(), rStats.IsAnarchy())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Anarchy > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsAnarchy(), rStats.IsLawful())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Pragmatic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsPragmatic(), rStats.IsAltruistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Altruistic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsAltruistic(), rStats.IsPragmatic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Naturalistic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsNaturalistic(), rStats.IsUrbanistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Urbanistic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsUrbanistic(), rStats.IsNaturalistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Philosophic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsPhilosophic(), rStats.IsPrimitive())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Primitive > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsPrimitive(), rStats.IsPhilosophic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Authoritaristic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsAuthoritaristic(), rStats.IsLiberal())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Liberal > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsLiberal(), rStats.IsAuthoritaristic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Individualistic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsIndividualistic(), rStats.IsCollectivistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Collectivistic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsCollectivistic(), rStats.IsIndividualistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	if inStats.Spiritualustic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsSpiritualustic(), rStats.IsMaterialistic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}
	if inStats.Materialistic > 0 {
		acc, shun, crim, err := calcRSAcceptanceProb(randCoef, rStats.IsMaterialistic(), rStats.IsSpiritualustic())
		if err != nil {
			return "", err
		}
		accepted += acc
		shunned += shun
		criminal += crim
		ideasCount++
	}

	accepted = accepted / float64(ideasCount)
	shunned = shunned / float64(ideasCount)
	criminal = criminal / float64(ideasCount)

	accepted = randomTools.PrepareProbability(acceptedBaseCoef * accepted)
	shunned = randomTools.PrepareProbability(stunnedBaseCoef * shunned)
	criminal = randomTools.PrepareProbability(criminalBaseCoef * criminal)
	if opts.Log {
		fmt.Printf("\n>>>>>>>>>>\nLabel: %s\naccepted: %f, shunned: %f, criminal: %f\n<<<<<<<<<<<<\n", opts.Label, accepted, shunned, criminal)
	}

	return actionAcceptance.GetAcceptanceByProb(accepted, shunned, criminal)
}
