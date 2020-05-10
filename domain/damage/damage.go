package damage

import (
	"pokecalc/domain/skill"
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
)

// ダメージ計算式
// TODO:types & stats
type Calculator struct {
	AttackersLevel stats.Level
	Attacker       *stats.Stats
	AttackersType  *types.Types
	Skill          *skill.Skill
	Defender       *stats.Stats
	DefendersType  *types.Types
}

func New() *Calculator {
	return &Calculator{
		AttackersLevel: stats.NewLevel(1),
	}
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

/*
	取りうるダメージ値を全て取得
	(0.85〜1.0)
*/
func calcArray(l, s, a, d uint) []uint {
	res := make([]uint, 0)

	for m := 0.85; m <= 1.0; m += 0.01 {
		d := calc(l, s, a, d)
		d = uint(float64(d) * m)
		res = append(res, d)
	}
	return res
}
