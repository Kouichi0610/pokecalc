package battle

import (
	"fmt"
	"pokecalc/domain/damage"
	"pokecalc/domain/nature"
	"pokecalc/domain/skill"
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
	SK "pokecalc/infrastructure/repository/skill"
	SP "pokecalc/infrastructure/repository/species"
)

type BattleService struct {
	lv       stats.Level
	attacker uint
	an       nature.Name
	ab       *stats.BasePoints
	defender uint
	dn       nature.Name
	db       *stats.BasePoints
	skill    uint
	skr      SK.Repository
	spr      SP.Repository
}

func New(species SP.Repository, skills SK.Repository) *BattleService {
	res := &BattleService{
		lv:       stats.NewLevel(50),
		attacker: 0,
		defender: 0,
		skill:    0,
		an:       nature.Bashful,
		dn:       nature.Bashful,
		skr:      skills,
		spr:      species,
	}
	return res
}

func (b *BattleService) Calculator() (*damage.Calculator, error) {
	skill, err := b.skr.Load(b.skill)
	if err != nil {
		return nil, fmt.Errorf("no skill %d", b.skill)
	}

	res := damage.New()
	res.AttackersLevel = b.lv

	return nil, fmt.Errorf("no condition.")
}

// 種族IDから能力値を返す
func (b *BattleService) stats(id uint, bs *stats.BasePoints, na nature.Name) (*stats.Stats, *types.Types, error) {
	sp, ty, err := b.spr.Load(id)
	if err != nil {
		return nil, nil, fmt.Errorf("species %d not found.", id)
	}
	in := stats.NewIndividual(31, 31, 31, 31, 31, 31)
	lv := b.lv
	nt := nature.New(na)

	return stats.NewStats(lv, sp, in, bs, nt), ty, nil
}

/*
必要なもの
(レベルは50で、個体値はさいこう固定で)
・攻撃側ID
　基礎ポイント
　性格
・防御側ID
　基礎ポイント
　性格
・わざ


*/
