package types

type ice struct {
}

func (t ice) effect(d effector) Effect {
	switch d.(type) {
	case *fire, *water, *ice, *steel:
		return notvery
	case *grass, *ground, *flying, *dragon:
		return super
	}
	return flat
}

func (t ice) string() string {
	return "こおり"
}
