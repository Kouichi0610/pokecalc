/**
基礎ポイント
HP、攻撃、防御、特攻、特防、すばやさ
任意に振ることのできるポイント
１箇所252まで
合計510まで振ることができる
*/
package stats

import (
	"fmt"
)

type BasePoints struct {
	hp, at, df, sa, sd, sp uint
}

func NewBasePoints(hp, at, df, sa, sd, sp uint) (*BasePoints, error) {
	clamp := func(v uint) uint {
		if v > 252 {
			return 252
		}
		return v
	}
	r := &BasePoints{
		hp: clamp(hp),
		at: clamp(at),
		df: clamp(df),
		sa: clamp(sa),
		sd: clamp(sd),
		sp: clamp(sp),
	}
	t := r.hp + r.at + r.df + r.sa + r.sd + r.sp
	if t > 510 {
		return nil, fmt.Errorf("total over.")
	}
	return r, nil
}
