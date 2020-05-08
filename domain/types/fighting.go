package types

type fighting struct {
}

func (t fighting) effect(d effector) Effect {
	switch d.(type) {
	case *poison, *flying, *psychic, *bug, *fairy:
		return notvery
	case *normal, *ice, *rock, *dark, *steel:
		return super
	case *ghost:
		return noeffect
	}
	return flat
}

func (t fighting) string() string {
	return "かくとう"
}
