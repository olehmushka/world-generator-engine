package worldgeneratorengine

import (
	"encoding/json"
	"testing"

	"github.com/olehmushka/world-generator-engine/religion"
	"github.com/stretchr/testify/require"
)

func TestNewReligion(t *testing.T) {
	types := make([]*religion.Trait, 0, 10)
	for chunk := range religion.LoadAllTypeTraits() {
		require.NoError(t, chunk.Err)
		types = append(types, chunk.Value...)
	}

	highGoals := make([]*religion.Trait, 0, 10)
	for chunk := range religion.LoadAllHighGoals() {
		require.NoError(t, chunk.Err)
		highGoals = append(highGoals, chunk.Value...)
	}

	socialTraits := make([]*religion.Trait, 0, 10)
	for chunk := range religion.LoadAllSocialTraits() {
		require.NoError(t, chunk.Err)
		socialTraits = append(socialTraits, chunk.Value...)
	}

	marriageKinds := make([]*religion.Trait, 0, 10)
	for chunk := range religion.LoadAllMarriageKinds() {
		require.NoError(t, chunk.Err)
		marriageKinds = append(marriageKinds, chunk.Value...)
	}

	bastardies := make([]*religion.Trait, 0, 10)
	for chunk := range religion.LoadAllBastardies() {
		require.NoError(t, chunk.Err)
		bastardies = append(bastardies, chunk.Value...)
	}

	consanguinities := make([]*religion.Trait, 0, 10)
	for chunk := range religion.LoadAllConsanguinities() {
		require.NoError(t, chunk.Err)
		consanguinities = append(consanguinities, chunk.Value...)
	}

	divorceTraditions := make([]*religion.PermissionTrait, 0, 10)
	for chunk := range religion.LoadAllDivorceOpts() {
		require.NoError(t, chunk.Err)
		divorceTraditions = append(divorceTraditions, chunk.Value...)
	}

	afterlifeParticipances := make([]religion.AfterlifeParticipance, 0, 10)
	for chunk := range religion.LoadAllAfterlifeParticipances() {
		require.NoError(t, chunk.Err)
		afterlifeParticipances = append(afterlifeParticipances, chunk.Value...)
	}

	afterlifeParticipants := make([]religion.AfterlifeParticipant, 0, 10)
	for chunk := range religion.LoadAllAfterlifeParticipants() {
		require.NoError(t, chunk.Err)
		afterlifeParticipants = append(afterlifeParticipants, chunk.Value...)
	}

	afterlifeExistOpts := make([]religion.AfterlifeExist, 0, 10)
	for chunk := range religion.LoadAllAfterlifeExists() {
		require.NoError(t, chunk.Err)
		afterlifeExistOpts = append(afterlifeExistOpts, chunk.Value...)
	}

	deityFavours := make([]*religion.FavourTrait, 0, 10)
	for chunk := range religion.LoadAllDeityFavours() {
		require.NoError(t, chunk.Err)
		deityFavours = append(deityFavours, chunk.Value...)
	}

	deityNatureTraits := make([]*religion.Trait, 0, 10)
	for chunk := range religion.LoadAllDeityNatureTraits() {
		require.NoError(t, chunk.Err)
		deityNatureTraits = append(deityNatureTraits, chunk.Value...)
	}

	humanNatureTraits := make([]*religion.Trait, 0, 10)
	for chunk := range religion.LoadAllHumanNatureTraits() {
		require.NoError(t, chunk.Err)
		humanNatureTraits = append(humanNatureTraits, chunk.Value...)
	}

	somlTraits := make([]*religion.Trait, 0, 10)
	for chunk := range religion.LoadAllSourcesOfMoralLaw() {
		require.NoError(t, chunk.Err)
		somlTraits = append(somlTraits, chunk.Value...)
	}

	r, err := religion.New(religion.CreateReligionOpts{
		Opts: religion.CreateReligionTraitsOpts{
			Slug:                    "quaralism",
			MinHighGoalsNum:         1,
			MaxHighGoalsNum:         3,
			MinDeityNatureTraitsNum: 0,
			MaxDeityNatureTraitsNum: 3,
			MinHumanNatureTraitsNum: 0,
			MaxHumanNatureTraitsNum: 2,
			MinSocialTraitsNum:      1,
			MaxSocialTraitsNum:      5,
		},
	}, religion.Data{
		Types:                   types,
		HighGoals:               highGoals,
		SocialTraits:            socialTraits,
		MarriageKinds:           marriageKinds,
		BastardyTraditions:      bastardies,
		ConsanguinityTraditions: consanguinities,
		DivorceTraditions:       divorceTraditions,
		AfterlifeParticipances:  afterlifeParticipances,
		AfterlifeParticipants:   afterlifeParticipants,
		AfterlifeExistOpts:      afterlifeExistOpts,
		DeityFavourTraits:       deityFavours,
		DeityNatureTraits:       deityNatureTraits,
		HumanNatureTraits:       humanNatureTraits,
		SourceOfMoralLawTraits:  somlTraits,
	})
	require.NoError(t, err)
	b, err := json.MarshalIndent(religion.PurifyReligion(r), "", "   ")
	require.NoError(t, err)
	t.Logf(string(b))
}
