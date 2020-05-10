package skill

import (
	"pokecalc/domain/types"
	"testing"
)

func Test_T(t *testing.T) {
	ty, _ := types.New(types.Dragon)

	s := New(ty, 130, 80, BodyPress)

	if s.Power() != 130 {
		t.Error()
	}
	if s.Accuracy() != 80 {
		t.Error()
	}

	if s.Types().String() != ty.String() {
		t.Error()
	}
}
