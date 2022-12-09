package permission

import (
	"fmt"

	mapTools "github.com/olehmushka/golang-toolkit/map_tools"
	randomTools "github.com/olehmushka/golang-toolkit/random_tools"
	"github.com/olehmushka/golang-toolkit/wrapped_error"
)

type Permission string

func (p Permission) String() string {
	return string(p)
}

func (a Permission) IsAlwaysAllowed() bool {
	return a == AlwaysAllowed
}

func (a Permission) IsMustBeApproved() bool {
	return a == MustBeApproved
}

func (a Permission) IsDisallowed() bool {
	return a == Disallowed
}

const (
	AlwaysAllowed  Permission = "always_allowed"
	MustBeApproved Permission = "must_be_approved"
	Disallowed     Permission = "disallowed"
)

func GetPermissionByProb(alwaysAllowed, mustBeApproved, disallowed float64) (Permission, error) {
	out, err := mapTools.PickOneByProb(map[string]float64{
		string(AlwaysAllowed):  randomTools.PrepareProbability(alwaysAllowed),
		string(MustBeApproved): randomTools.PrepareProbability(mustBeApproved),
		string(Disallowed):     randomTools.PrepareProbability(disallowed),
	})
	if err != nil {
		return "", wrapped_error.NewInternalServerError(err, fmt.Sprintf("can not generate permission (always_allowed=%f, must_be_approved=%f, disallowed=%f)", alwaysAllowed, mustBeApproved, disallowed))
	}

	return Permission(out), nil
}

func IsValid(v string) bool {
	for _, valid := range []string{
		AlwaysAllowed.String(),
		MustBeApproved.String(),
		Disallowed.String(),
	} {
		if v == valid {
			return true
		}
	}

	return false
}
