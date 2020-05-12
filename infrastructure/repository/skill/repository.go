package skill

import (
	"pokecalc/domain/skill"
)

type Repository interface {
	Load(id uint) (*skill.Skill, error)
}
