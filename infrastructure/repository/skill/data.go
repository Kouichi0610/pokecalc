package skill

import (
	"pokecalc/domain/skill"
	"pokecalc/domain/types"
)

type data struct {
	id       uint
	name     string
	category skill.Category
	ty       types.Type
	power    uint
	accuracy uint
}

func (d *data) Skill() *skill.Skill {
	ty, _ := types.New(d.ty)
	return skill.New(ty, d.power, d.accuracy, d.category)
}
