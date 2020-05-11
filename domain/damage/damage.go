package damage

import (
	"pokecalc/domain/skill"
	"pokecalc/domain/stats"
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

// タイプ一致＆タイプ相性補正
func (c *Calculator) correct(d uint) uint {
	// TODO:Effectを関数にしておきたい
	// タイプ相性
	effect := float32(c.Attacker.Types.Effect(c.Defender.Types))
	d = uint(float32(d) * effect)

	// タイプ一致
	if c.Skill.Types().PartialMatch(c.Attacker.Types) {
		d = d * 15 / 10
	}
	return d
}

// ダメージ計算
func (c *Calculator) Calculate() []uint {
	l := uint(c.AttackersLevel)
	s := c.Skill.Power()

	a, d := c.Skill.PickStats(c.Attacker.Stats, c.Defender.Stats)
	res := calcArray(l, s, a, d)

	for i := 0; i < len(res); i++ {
		res[i] = c.correct(res[i])
	}

	return res
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
