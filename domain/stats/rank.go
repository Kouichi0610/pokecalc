/*
ランク補正
HP以外の能力値に補正をかける
-6から+6まで
*/
package stats

type rank int

type Ranks struct {
	at, df, sa, sd, sp rank
}

func NewRanks(at, df, sa, sd, sp int) *Ranks {
	return &Ranks{
		at: newRank(at),
		df: newRank(df),
		sa: newRank(sa),
		sd: newRank(sd),
		sp: newRank(sp),
	}
}

// ランク補正をかけたStatsを新たに生成する
func (r *Ranks) RankedStats(s *Stats) *Stats {
	return &Stats{
		hp: s.hp,
		at: r.at.rankedStats(s.at),
		df: r.df.rankedStats(s.df),
		sa: r.sa.rankedStats(s.sa),
		sd: r.sd.rankedStats(s.sd),
		sp: r.sp.rankedStats(s.sp),
	}
}

func newRank(r int) rank {
	if r > 6 {
		r = 6
	}
	if r < -6 {
		r = -6
	}
	return rank(r)
}
func (r rank) rankedStats(s uint) uint {
	if r < 0 {
		return r.minusStats(s)
	}
	return r.plusStats(s)
}
func (r rank) minusStats(s uint) uint {
	m := int(-r) + 2
	rs := int(s) * 2 / m
	return uint(rs)
}
func (r rank) plusStats(s uint) uint {
	m := int(r) + 2
	rs := int(s) * m / 2
	return uint(rs)
}
