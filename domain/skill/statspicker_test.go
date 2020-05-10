package skill

import (
	"pokecalc/domain/stats"
	"testing"
)

func testPick(p statsPicker) (aval, dval uint) {
	at := stats.NewStatsValue(1, 2, 3, 4, 5, 6)
	df := stats.NewStatsValue(7, 8, 9, 10, 11, 12)
	return p(at, df)
}

func Test_StatsPicker_Physical(t *testing.T) {
	a, d := testPick(toPicker(Physical))
	if a != 2 {
		t.Error()
	}
	if d != 9 {
		t.Error()
	}
}
func Test_StatsPicker_Special(t *testing.T) {
	a, d := testPick(toPicker(Special))
	if a != 4 {
		t.Error()
	}
	if d != 11 {
		t.Error()
	}
}
func Test_StatsPicker_PsychoShock(t *testing.T) {
	a, d := testPick(toPicker(PsycoShock))
	if a != 4 {
		t.Error()
	}
	if d != 9 {
		t.Error()
	}
}
func Test_StatsPicker_BodyPress(t *testing.T) {
	a, d := testPick(toPicker(BodyPress))
	if a != 3 {
		t.Error()
	}
	if d != 9 {
		t.Error()
	}
}
func Test_StatsPicker_FoulPlay(t *testing.T) {
	a, d := testPick(toPicker(FoulPlay))
	if a != 8 {
		t.Error()
	}
	if d != 9 {
		t.Error()
	}
}

func Test_StatsPicker_Invalid(t *testing.T) {
	c := toPicker(5)
	if c != nil {
		t.Error()
	}
}
