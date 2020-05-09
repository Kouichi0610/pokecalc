/*
種族値
ポケモンごとの能力
*/
package stats

type Species struct {
	hp, at, df, sa, sd, sp uint
}

func NewSpecies(hp, at, df, sa, sd, sp uint) *Species {
	return &Species{
		hp: hp,
		at: at,
		df: df,
		sa: sa,
		sd: sd,
		sp: sp,
	}
}
