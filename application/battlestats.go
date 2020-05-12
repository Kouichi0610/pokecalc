package battle

import (
	"fmt"
	"pokecalc/domain/nature"
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
	SP "pokecalc/infrastructure/repository/species"
)

type BattleStats struct {
	Name   string
	Nature nature.Name
	Base   *stats.BasePoints
	spr    SP.Repository
}

func NewBattleStats(spr SP.Repository) *BattleStats {
	b, _ := stats.NewBasePoints(0, 0, 0, 0, 0, 0)
	return &BattleStats{
		Name:   "",
		Nature: nature.Bashful,
		Base:   b,
		spr:    spr,
	}
}

func (b *BattleStats) stats() (*stats.Stats, *types.Types, error) {
	sp, ty, err := b.spr.Load(b.Name)
	if err != nil {
		return nil, nil, fmt.Errorf("species %s not found.", b.Name)
	}
	in := stats.NewIndividual(31, 31, 31, 31, 31, 31)
	lv := defaultLevel
	nt := nature.New(b.Nature)

	return stats.NewStats(lv, sp, in, b.Base, nt), ty, nil
}
