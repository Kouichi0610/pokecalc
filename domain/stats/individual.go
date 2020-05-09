/*
個体値
生まれついての資質
それぞれ0〜31までの値をとる
*/
package stats

type Individual struct {
	hp, at, df, sa, sd, sp uint
}

func NewIndividual(hp, at, df, sa, sd, sp uint) *Individual {
	clamp := func(v uint) uint {
		if v > 31 {
			return 31
		}
		return v
	}
	return &Individual{
		hp: clamp(hp),
		at: clamp(at),
		df: clamp(df),
		sa: clamp(sa),
		sd: clamp(sd),
		sp: clamp(sp),
	}
}
