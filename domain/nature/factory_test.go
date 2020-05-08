package nature

import (
	"testing"
)

func Test_Natures(t *testing.T) {
	const h = 404
	const f = 299
	const u = 328
	const l = 269
	expects := map[Name][6]uint{
		Bashful: [6]uint{h, f, f, f, f, f},
		Lonely:  [6]uint{h, u, l, f, f, f},
		Adamant: [6]uint{h, u, f, l, f, f},
		Naughty: [6]uint{h, u, f, f, l, f},
		Brave:   [6]uint{h, u, f, f, f, l},
		Bold:    [6]uint{h, l, u, f, f, f},
		Impish:  [6]uint{h, f, u, l, f, f},
		Lax:     [6]uint{h, f, u, f, l, f},
		Relaxed: [6]uint{h, f, u, f, f, l},
		Modest:  [6]uint{h, l, f, u, f, f},
		Mild:    [6]uint{h, f, l, u, f, f},
		Rash:    [6]uint{h, f, f, u, l, f},
		Quiet:   [6]uint{h, f, f, u, f, l},
		Calm:    [6]uint{h, l, f, f, u, f},
		Gentle:  [6]uint{h, f, l, f, u, f},
		Careful: [6]uint{h, f, f, l, u, f},
		Sassy:   [6]uint{h, f, f, f, u, l},
		Timid:   [6]uint{h, l, f, f, f, u},
		Hasty:   [6]uint{h, f, l, f, f, u},
		Jolly:   [6]uint{h, f, f, l, f, u},
		Naive:   [6]uint{h, f, f, f, l, u},
	}
	for key, expect := range expects {
		n := New(key)
		a := actual(n)
		if a != expect {
			t.Errorf("%d %v %v", key, expect, a)
		}
	}
	{ // ヌケニンはHPが1になること。性格補正は反映されていること。
		n := New(Brave).Shadinja()
		a := actual(n)
		expect := [6]uint{1, u, f, f, f, l}
		if a != expect {
			t.Error()
		}
	}
}

func actual(n *Nature) [6]uint {
	const l = 100
	const s = 100
	const i = 31
	const b = 252

	return [6]uint{
		n.HP(l, s, i, b),
		n.Attack(l, s, i, b),
		n.Defense(l, s, i, b),
		n.SpAttack(l, s, i, b),
		n.SpDefense(l, s, i, b),
		n.Speed(l, s, i, b),
	}
}
