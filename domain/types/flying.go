package types

type flying struct {
}

func (t flying) effect(d effector) Effect {
	switch d.(type) {
	case *electric, *rock, *steel:
		return notvery
	case *grass, *fighting, *bug:
		return super
	}
	return flat
}

func (t flying) string() string {
	return "ひこう"
}
