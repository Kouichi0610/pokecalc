package battle

import (
	"pokecalc/domain/nature"
	"pokecalc/domain/stats"
	SK "pokecalc/infrastructure/repository/skill"
	SP "pokecalc/infrastructure/repository/species"
	"reflect"
	"testing"
)

/*
 */
func Test_BattleService(t *testing.T) {
	spr := SP.NewMock()
	skr := SK.NewMock()

	var err error

	at := NewBattleStats(spr)
	at.Name = "ゴウカザル"
	at.Nature = nature.Adamant
	at.Base, err = stats.NewBasePoints(6, 252, 0, 0, 0, 252)

	df := NewBattleStats(spr)
	df.Name = "ドダイトス"
	df.Nature = nature.Careful
	df.Base, err = stats.NewBasePoints(252, 0, 252, 0, 0, 0)

	sv := NewBattleService(skr)
	sv.Attacker = at
	sv.Defender = df

	sv.Skill = "ひのこ"

	c, err := sv.Calculator()
	if err == nil {
		t.Error()
	}
	if c == nil {
		t.Error()
	}

	actual := c.Calculate()
	expect := []uint{45, 46, 46, 47, 48, 48, 49, 49, 50, 50, 51, 51, 52, 52, 53, 54}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Actual %v", actual)
		t.Errorf("Expect %v", expect)
	}
}
