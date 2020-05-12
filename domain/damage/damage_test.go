package damage

import (
	"pokecalc/domain/skill"
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
	"reflect"
	"testing"
)

func Test_New(t *testing.T) {
	c := New()
	if c.Enable() {
		t.Error()
	}
	st, _ := types.New(types.Fire)
	c.Skill = skill.New(st, 100, 80, skill.Physical)
	if c.Enable() {
		t.Error()
	}
	at, _ := types.New(types.Fire)
	as := stats.NewStatsValue(100, 100, 100, 100, 100, 100)
	c.Attacker = &Condition{Types: at, Stats: as}
	if c.Enable() {
		t.Error()
	}
	dt, _ := types.New(types.Fire)
	ds := stats.NewStatsValue(100, 100, 100, 100, 100, 100)
	c.Defender = &Condition{Types: dt, Stats: ds}
	if !c.Enable() {
		t.Error()
	}
}

func Test_Calculate(t *testing.T) {
	c := New()

	st, _ := types.New(types.Electric)
	c.Skill = skill.New(st, 120, 80, skill.Physical)

	at, _ := types.New(types.Electric)
	as := stats.NewStatsValue(100, 180, 80, 35, 80, 200)
	c.Attacker = &Condition{Types: at, Stats: as}

	dt, _ := types.New(types.Water, types.Flying)
	ds := stats.NewStatsValue(300, 20, 40, 40, 250, 30)
	c.Defender = &Condition{Types: dt, Stats: ds}

	actual := c.Calculate()
	expect := []uint{117, 118, 120, 121, 122, 124, 125, 126, 128, 129, 131, 132, 133, 135, 136, 138}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Actual %v", actual)
		t.Errorf("Expect %v", expect)
	}
}

func Test_Correct(t *testing.T) {
	const dmg uint = 100
	s, _ := types.New(types.Fire)
	a, _ := types.New(types.Fire, types.Dragon)

	d, _ := types.New(types.Water)

	actual := correct(dmg, s, a, d)
	var expect uint = 75
	if actual != expect {
		t.Errorf("failed Expect %d Actual %d", actual, expect)
	}
}

func Test_Calc(t *testing.T) {
	const l uint = 50
	const s uint = 100
	const a uint = 182
	const d uint = 55

	actual := calc(l, s, a, d)
	expect := uint(147)

	if actual != expect {
		t.Errorf("Failed Damage Expect %d Actual %d", expect, actual)
	}
}
