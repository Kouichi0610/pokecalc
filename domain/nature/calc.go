/*
	能力値計算式
	l level レベル
	s species 種族値
	i individual 個体値
	b basePoint 基礎ポイント
*/
package nature

type statsCalc func(l, s, i, b uint) uint

func shadinjaHP(l, s, i, b uint) uint {
	return 1
}

func hp(l, s, i, b uint) uint {
	x := s*2 + i + b/4
	y := x*l/100.0 + l + 10
	return uint(y)
}

func flat(l, s, i, b uint) uint {
	x := s*2 + i + b/4
	y := x*l/100.0 + 5
	return uint(y)
}

func upper(l, s, i, b uint) uint {
	a := float32(flat(l, s, i, b)) * 1.1
	return uint(a)
}

func lower(l, s, i, b uint) uint {
	a := float32(flat(l, s, i, b)) * 0.9
	return uint(a)
}
