package types

type rock struct {
}

func (t rock) effect(d effector) Effect {
	switch d.(type) {
	case *fighting, *ground, *steel:
		return notvery
	case *fire, *ice, *flying, *bug:
		return super
	}
	return flat
}

func (t rock) string() string {
	return "いわ"
}
