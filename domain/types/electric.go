package types

type electric struct {
}

func (t electric) effect(d effector) Effect {
	switch d.(type) {
	case *electric, *grass, *dragon:
		return notvery
	case *water, *flying:
		return super
	case *ground:
		return noeffect
	}
	return flat
}

func (t electric) string() string {
	return "でんき"
}
