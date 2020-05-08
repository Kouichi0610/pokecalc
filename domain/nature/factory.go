package nature

type createFunc func() *Nature

var factory map[Name]createFunc

func init() {
	factory = map[Name]createFunc{
		Bashful: bashful,
		Lonely:  lonely,
		Adamant: adamant,
		Naughty: naughty,
		Brave:   brave,
		Bold:    bold,
		Impish:  impish,
		Lax:     lax,
		Relaxed: relaxed,
		Modest:  modest,
		Mild:    mild,
		Rash:    rash,
		Quiet:   quiet,
		Calm:    calm,
		Gentle:  gentle,
		Careful: careful,
		Sassy:   sassy,
		Timid:   timid,
		Hasty:   hasty,
		Jolly:   jolly,
		Naive:   naive,
	}
}

func (n *Nature) Shadinja() *Nature {
	n.hp = shadinjaHP
	return n
}

func bashful() *Nature {
	return &Nature{
		hp: hp,
		at: flat,
		df: flat,
		sa: flat,
		sd: flat,
		sp: flat,
	}
}
func lonely() *Nature {
	res := bashful()
	res.at = upper
	res.df = lower
	return res
}
func adamant() *Nature {
	res := bashful()
	res.at = upper
	res.sa = lower
	return res
}
func naughty() *Nature {
	res := bashful()
	res.at = upper
	res.sd = lower
	return res
}
func brave() *Nature {
	res := bashful()
	res.at = upper
	res.sp = lower
	return res
}

func bold() *Nature {
	res := bashful()
	res.df = upper
	res.at = lower
	return res
}
func impish() *Nature {
	res := bashful()
	res.df = upper
	res.sa = lower
	return res
}
func lax() *Nature {
	res := bashful()
	res.df = upper
	res.sd = lower
	return res
}
func relaxed() *Nature {
	res := bashful()
	res.df = upper
	res.sp = lower
	return res
}

func modest() *Nature {
	res := bashful()
	res.sa = upper
	res.at = lower
	return res
}
func mild() *Nature {
	res := bashful()
	res.sa = upper
	res.df = lower
	return res
}
func rash() *Nature {
	res := bashful()
	res.sa = upper
	res.sd = lower
	return res
}
func quiet() *Nature {
	res := bashful()
	res.sa = upper
	res.sp = lower
	return res
}

func calm() *Nature {
	res := bashful()
	res.sd = upper
	res.at = lower
	return res
}
func gentle() *Nature {
	res := bashful()
	res.sd = upper
	res.df = lower
	return res
}
func careful() *Nature {
	res := bashful()
	res.sd = upper
	res.sa = lower
	return res
}
func sassy() *Nature {
	res := bashful()
	res.sd = upper
	res.sp = lower
	return res
}

func timid() *Nature {
	res := bashful()
	res.sp = upper
	res.at = lower
	return res
}
func hasty() *Nature {
	res := bashful()
	res.sp = upper
	res.df = lower
	return res
}
func jolly() *Nature {
	res := bashful()
	res.sp = upper
	res.sa = lower
	return res
}
func naive() *Nature {
	res := bashful()
	res.sp = upper
	res.sd = lower
	return res
}
