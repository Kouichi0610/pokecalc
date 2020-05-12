package battle

import (
	"fmt"
	"pokecalc/domain/damage"
	"pokecalc/domain/stats"
	SK "pokecalc/infrastructure/repository/skill"
)

var defaultLevel stats.Level = stats.NewLevel(50)

type BattleService struct {
	Attacker *BattleStats
	Defender *BattleStats
	Skill    string
	skr      SK.Repository
}

func NewBattleService(skills SK.Repository) *BattleService {
	res := &BattleService{
		Attacker: nil,
		Defender: nil,
		Skill:    "",
		skr:      skills,
	}
	return res
}

func (b *BattleService) Calculator() (*damage.Calculator, error) {
	skill, err := b.skr.Load(b.Skill)
	if err != nil {
		return nil, fmt.Errorf("no skill %s", b.Skill)
	}

	res := damage.New()
	res.AttackersLevel = defaultLevel
	res.Skill = skill

	as, at, err := b.Attacker.stats()
	if err != nil {
		return nil, fmt.Errorf("no attacker %s", b.Attacker.Name)
	}
	res.AttackerStats = as
	res.AttackerType = at

	ds, dt, err := b.Defender.stats()
	if err != nil {
		return nil, fmt.Errorf("no defender %s", b.Defender.Name)
	}
	res.DefenderStats = ds
	res.DefenderType = dt

	return res, fmt.Errorf("no condition.")
}
