package damage

import (
	"pokecalc/domain/skill"
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
)

// ダメージ計算式
type Calculator struct {
	AttackersLevel stats.Level
	AttackerType   *types.Types
	AttackerStats  *stats.Stats
	Skill          *skill.Skill
	DefenderType   *types.Types
	DefenderStats  *stats.Stats
}

func New() *Calculator {
	return &Calculator{
		AttackersLevel: stats.NewLevel(1),
		AttackerType:   nil,
		AttackerStats:  nil,
		Skill:          nil,
		DefenderType:   nil,
		DefenderStats:  nil,
	}
}

func (c *Calculator) Enable() bool {
	if c.AttackerType == nil {
		return false
	}
	if c.AttackerStats == nil {
		return false
	}
	if c.DefenderType == nil {
		return false
	}
	if c.DefenderStats == nil {
		return false
	}
	if c.Skill == nil {
		return false
	}
	return true
}

// ダメージ計算
func (c *Calculator) Calculate() []uint {
	l := uint(c.AttackersLevel)
	s := c.Skill.Power()

	a, d := c.Skill.PickStats(c.AttackerStats, c.DefenderStats)
	dmg := calc(l, s, a, d)
	dmg = correct(dmg, c.Skill.Types(), c.AttackerType, c.DefenderType)

	res := make([]uint, 0)

	var m uint
	for m = 85; m <= 100; m += 1 {
		d = dmg * m / 100
		res = append(res, d)
	}
	return res
}

/*
	タイプ一致＆タイプ相性補正
	dmg ダメージ
	s わざタイプ
	a 攻撃側タイプ
	d 防御側タイプ
*/
func correct(dmg uint, s, a, d *types.Types) uint {
	// わざ->防御タイプ相性
	se := float32(s.Effect(d))
	// わざと攻撃側のタイプ一致
	var te float32 = 1.0
	if a.PartialMatch(s) {
		te = 1.5
	}
	return uint(float32(dmg) * se * te)
}

/*
	ダメージ計算(乱数抜き)
	l 攻撃側レベル
	s わざ威力
	a 攻撃側能力値
	d 防御側能力値
*/
func calc(l, s, a, d uint) uint {
	tmp := l*2/5 + 2
	tmp = tmp * s * a / d
	return tmp/50 + 2
}
