package types

type steel struct {
}

func (t steel) effect(d effector) Effect {
	switch d.(type) {
	case *fire, *water, *electric, *steel:
		return notvery
	case *ice, *rock, *fairy:
		return super
	}
	return flat
}

func (t steel) string() string {
	return "はがね"
}
