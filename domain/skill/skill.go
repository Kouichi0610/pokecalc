// わざ
package skill

import (
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
)

type Category uint

const (
	Physical Category = iota
	Special
	PsycoShock
	BodyPress
	FoulPlay
)

type Skill struct {
	t *types.Types // 攻撃タイプ
	p uint         // 威力(power)
	a uint         // 命中(accuracy)
	c statsPicker  // 分類
}

func New(t *types.Types, p uint, a uint, c Category) *Skill {
	return &Skill{
		t: t,
		p: p,
		a: a,
		c: toPicker(c),
	}
}

//func special(a *stats.Stats, d *stats.Stats) (aval, dval uint) {

func (s *Skill) Types() *types.Types {
	return s.t
}

func (s *Skill) Power() uint {
	return s.p
}

func (s *Skill) Accuracy() uint {
	return s.a
}

func (s *Skill) PickStats(a, d *stats.Stats) (aval, dval uint) {
	return s.c(a, d)
}

func toPicker(c Category) statsPicker {
	switch c {
	case Physical:
		return physical
	case Special:
		return special
	case PsycoShock:
		return psycoShock
	case BodyPress:
		return bodyPress
	case FoulPlay:
		return foulPlay
	}
	return nil
}
