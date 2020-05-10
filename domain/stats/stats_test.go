package stats

import (
	"pokecalc/domain/nature"
	"testing"
)

func Test_StatsValue(t *testing.T) {
	s := NewStatsValue(1, 2, 3, 4, 5, 6)

	if s.HP() != 1 {
		t.Error()
	}
	if s.Attack() != 2 {
		t.Error()
	}
	if s.Defense() != 3 {
		t.Error()
	}
	if s.SpAttack() != 4 {
		t.Error()
	}
	if s.SpDefense() != 5 {
		t.Error()
	}
	if s.Speed() != 6 {
		t.Error()
	}
}

func Test_Stats(t *testing.T) {
	l := NewLevel(50)
	s := NewSpecies(100, 80, 50, 40, 27, 130)
	b, _ := NewBasePoints(4, 252, 0, 0, 0, 252)
	i := NewIndividual(31, 31, 31, 31, 31, 31)
	n := nature.New(nature.Adamant)

	r := NewStats(l, s, i, b, n)
	if r.HP() != 176 {
		t.Error()
	}
	if r.Attack() != 145 {
		t.Error()
	}
	if r.Defense() != 70 {
		t.Error()
	}
	if r.SpAttack() != 54 {
		t.Error()
	}
	if r.SpDefense() != 47 {
		t.Error()
	}
	if r.Speed() != 182 {
		t.Error()
	}
}

func Test_Level(t *testing.T) {
	l := NewLevel(0)
	if l != 1 {
		t.Error()
	}
	l = NewLevel(101)
	if l != 100 {
		t.Error()
	}
}

func Test_Invididual(t *testing.T) {
	i := NewIndividual(32, 26, 4, 12, 18, 0)
	if i.hp != 31 {
		t.Error()
	}
	if i.at != 26 {
		t.Error()
	}
	if i.df != 4 {
		t.Error()
	}
	if i.sa != 12 {
		t.Error()
	}
	if i.sd != 18 {
		t.Error()
	}
	if i.sp != 0 {
		t.Error()
	}
}

func Test_BasePoints_正常(t *testing.T) {
	b, err := NewBasePoints(6, 100, 80, 60, 30, 120)
	if err != nil {
		t.Error()
	}
	if b.hp != 6 {
		t.Error()
	}
	if b.at != 100 {
		t.Error()
	}
	if b.df != 80 {
		t.Error()
	}
	if b.sa != 60 {
		t.Error()
	}
	if b.sd != 30 {
		t.Error()
	}
	if b.sp != 120 {
		t.Error()
	}
}

func Test_BasePoints_Clamp(t *testing.T) {
	b, err := NewBasePoints(511, 0, 0, 0, 0, 0)
	if err != nil {
		t.Error()
	}
	if b.hp != 252 {
		t.Error()
	}
}

func Test_BasePoints_異常(t *testing.T) {
	_, err := NewBasePoints(100, 100, 100, 100, 100, 11)
	if err == nil {
		t.Error()
	}
}

func Test_Species(t *testing.T) {
	s := NewSpecies(80, 120, 75, 60, 85, 130)
	if s.hp != 80 {
		t.Error()
	}
	if s.at != 120 {
		t.Error()
	}
	if s.df != 75 {
		t.Error()
	}
	if s.sa != 60 {
		t.Error()
	}
	if s.sd != 85 {
		t.Error()
	}
	if s.sp != 130 {
		t.Error()
	}
}

func Test_Rank(t *testing.T) {
	const s = 100
	expects := map[rank]uint{
		newRank(-6): 25,
		newRank(-5): 28,
		newRank(-4): 33,
		newRank(-3): 40,
		newRank(-2): 50,
		newRank(-1): 66,
		newRank(0):  100,
		newRank(1):  150,
		newRank(2):  200,
		newRank(3):  250,
		newRank(4):  300,
		newRank(5):  350,
		newRank(6):  400,
	}
	for r, e := range expects {
		a := r.rankedStats(s)
		if a != e {
			t.Errorf("failed Rank%d Expects:%d Actual:%d", r, e, a)
		}
	}

	if newRank(-7).rankedStats(s) != 25 {
		t.Error()
	}
	if newRank(7).rankedStats(s) != 400 {
		t.Error()
	}
}

func Test_Ranks(t *testing.T) {
	s := &Stats{hp: 100, at: 100, df: 100, sa: 100, sd: 100, sp: 100}
	r := NewRanks(0, 1, -1, 2, -2)
	e := &Stats{hp: 100, at: 100, df: 150, sa: 66, sd: 200, sp: 50}
	a := r.RankedStats(s)

	if *e != *a {
		t.Errorf("failed Expects(%v) Actual(%v)", a, e)
	}

}
