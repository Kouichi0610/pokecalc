package types

type dark struct {
}

func (t dark) effect(d effector) Effect {
	switch d.(type) {
	case *fighting, *dark, *fairy:
		return notvery
	case *psychic, *ghost:
		return super
	}
	return flat
}

func (t dark) string() string {
	return "あく"
}
