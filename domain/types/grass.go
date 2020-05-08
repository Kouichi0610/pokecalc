package types

type grass struct {
}

func (t grass) effect(d effector) Effect {
	switch d.(type) {
	case *fire, *grass, *poison, *flying, *bug, *dragon, *steel:
		return notvery
	case *water, *ground, *rock:
		return super
	}
	return flat
}

func (t grass) string() string {
	return "くさ"
}
