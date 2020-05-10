// 能力値
package stats

import (
	"pokecalc/domain/nature"
)

type Stats struct {
	hp, at, df, sa, sd, sp uint
}

func NewStats(l Level, s *Species, i *Individual, b *BasePoints, n *nature.Nature) *Stats {
	return &Stats{
		hp: n.HP(uint(l), s.hp, i.hp, b.hp),
		at: n.Attack(uint(l), s.at, i.at, b.at),
		df: n.Defense(uint(l), s.df, i.df, b.df),
		sa: n.SpAttack(uint(l), s.sa, i.sa, b.sa),
		sd: n.SpDefense(uint(l), s.sd, i.sd, b.sd),
		sp: n.Speed(uint(l), s.sp, i.sp, b.sp),
	}
}

// 値を直接指定。基本テスト用
func NewStatsValue(hp, at, df, sa, sd, sp uint) *Stats {
	return &Stats{
		hp: hp,
		at: at,
		df: df,
		sa: sa,
		sd: sd,
		sp: sp,
	}
}

func (s *Stats) HP() uint {
	return s.hp
}

func (s *Stats) Attack() uint {
	return s.at
}

func (s *Stats) Defense() uint {
	return s.df
}

func (s *Stats) SpAttack() uint {
	return s.sa
}

func (s *Stats) SpDefense() uint {
	return s.sd
}

func (s *Stats) Speed() uint {
	return s.sp
}
