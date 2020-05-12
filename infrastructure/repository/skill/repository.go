package skill

import (
	"pokecalc/domain/skill"
)

type Repository interface {
	Load(name string) (*skill.Skill, error)
}
