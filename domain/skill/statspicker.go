// 攻撃側と防御側からダメージ計算に必要な能力値を取得する
package skill

import (
	"pokecalc/domain/stats"
)

type statsPicker func(a *stats.Stats, d *stats.Stats) (aval, dval uint)

// 物理
func physical(a *stats.Stats, d *stats.Stats) (aval, dval uint) {
	return a.Attack(), d.Defense()
}

// 特殊
func special(a *stats.Stats, d *stats.Stats) (aval, dval uint) {
	return a.SpAttack(), d.SpDefense()
}

// サイコショック(特殊攻撃、防御側は防御)
func psycoShock(a *stats.Stats, d *stats.Stats) (aval, dval uint) {
	return a.SpAttack(), d.Defense()
}

// ボディプレス(物理攻撃、攻撃側の防御)
func bodyPress(a *stats.Stats, d *stats.Stats) (aval, dval uint) {
	return a.Defense(), d.Defense()
}

// イカサマ(物理攻撃、防御側の攻撃)
func foulPlay(a *stats.Stats, d *stats.Stats) (aval, dval uint) {
	return d.Attack(), d.Defense()
}
