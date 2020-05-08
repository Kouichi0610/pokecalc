// 能力値の性格補正
package nature

type Name uint

const (
	Bashful Name = iota
	Lonely
	Adamant
	Naughty
	Brave
	Bold
	Impish
	Lax
	Relaxed
	Modest
	Mild
	Rash
	Quiet
	Calm
	Gentle
	Careful
	Sassy
	Timid
	Hasty
	Jolly
	Naive
)

type Nature struct {
	hp, at, df, sa, sd, sp statsCalc
}

func New(n Name) *Nature {
	return factory[n]()
}

func (n *Nature) HP(l, s, i, b uint) uint {
	return n.hp(l, s, i, b)
}
func (n *Nature) Attack(l, s, i, b uint) uint {
	return n.at(l, s, i, b)
}
func (n *Nature) Defense(l, s, i, b uint) uint {
	return n.df(l, s, i, b)
}
func (n *Nature) SpAttack(l, s, i, b uint) uint {
	return n.sa(l, s, i, b)
}
func (n *Nature) SpDefense(l, s, i, b uint) uint {
	return n.sd(l, s, i, b)
}
func (n *Nature) Speed(l, s, i, b uint) uint {
	return n.sp(l, s, i, b)
}
