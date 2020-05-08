package types

type normal struct {
}

func (t normal) effect(d effector) Effect {
	switch d.(type) {
	case *rock, *steel:
		return notvery
	case *ghost:
		return noeffect
	}
	return flat
}

func (t normal) string() string {
	return "ノーマル"
}
