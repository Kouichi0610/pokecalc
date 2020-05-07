package types

type poison struct {
}

func (t poison) effect(d effector) Effect {
	switch d.(type) {
	case *poison, *ground, *rock, *ghost:
		return notvery
	case *grass, *fairy:
		return super
	case *steel:
		return noeffect
	}
	return flat
}

func (t poison) string() string {
	return "どく"
}
