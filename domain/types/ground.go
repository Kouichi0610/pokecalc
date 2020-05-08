package types

type ground struct {
}

func (t ground) effect(d effector) Effect {
	switch d.(type) {
	case *grass, *bug:
		return notvery
	case *fire, *electric, *poison, *rock, *steel:
		return super
	case *flying:
		return noeffect
	}
	return flat
}

func (t ground) string() string {
	return "じめん"
}
