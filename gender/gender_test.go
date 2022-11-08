package gender

import (
	"testing"

	sliceTools "github.com/olehmushka/golang-toolkit/slice_tools"
)

func TestGetRandomSex(t *testing.T) {
	result, err := GetRandomSex()
	if err != nil {
		t.Fatalf("unexpected err (err=%+v)", err)
	}
	if !sliceTools.Includes([]string{
		MaleSex.String(),
		FemaleSex.String(),
	}, result.String()) {
		t.Errorf("result should be picked from available sexes")
	}
}
