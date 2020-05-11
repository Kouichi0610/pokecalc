package species

import (
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
)

type data struct {
	id        uint
	name      string
	type1     types.Type
	type2     types.Type
	ability1  uint
	ability2  uint
	ability3  uint
	hp        uint
	attack    uint
	defense   uint
	spAttack  uint
	spDefense uint
	speed     uint
}

func (d *data) Species() *stats.Species {
	return stats.NewSpecies(d.hp, d.attack, d.defense, d.spAttack, d.spDefense, d.speed)
}
func (d *data) Types() *types.Types {
	res, _ := types.New(d.type1, d.type2)
	return res
}
