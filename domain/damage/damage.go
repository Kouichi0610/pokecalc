package damage

import (
	"pokecalc/domain/skill"
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
)

// ダメージ計算式
type Calculator struct {
	AttackersLevel stats.Level
	Attacker       *Condition
	Skill          *skill.Skill
	Defender       *Condition
}

func New() *Calculator {
	return &Calculator{
		AttackersLevel: stats.NewLevel(1),
		Attacker:       nil,
		Skill:          nil,
		Defender:       nil,
	}
}

func (c *Calculator) Enable() bool {
	if c.Attacker == nil {
		return false
	}
	if c.Defender == nil {
		return false
	}
	if c.Skill == nil {
		return false
	}
	return true
}

// ダメージ計算
// TODO:計算式については要検証。特に計算順
func (c *Calculator) Calculate() []uint {
	l := uint(c.AttackersLevel)
	s := c.Skill.Power()

	a, d := c.Skill.PickStats(c.Attacker.Stats, c.Defender.Stats)
	dmg := calc(l, s, a, d)
	dmg = correct(dmg, c.Skill.Types(), c.Attacker.Types, c.Defender.Types)

	res := make([]uint, 0)

	for m := 0.85; m <= 1.0; m += 0.01 {
		d = uint(float64(dmg) * m)
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
