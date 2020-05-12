package battle

import (
	"pokecalc/domain/damage"
	"pokecalc/domain/nature"
	"pokecalc/domain/skill"
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
	SK "pokecalc/infrastructure/repository/skill"
	SP "pokecalc/infrastructure/repository/species"
	"testing"
)

/*
 */
func Test_BattleService(t *testing.T) {
	spr := SP.NewMock()
	skr := SK.NewMock()

}
