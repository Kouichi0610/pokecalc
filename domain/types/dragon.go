package types

type dragon struct {
}

func (t dragon) effect(d effector) Effect {
	switch d.(type) {
	case *steel:
		return notvery
	case *dragon:
		return super
	case *fairy:
		return noeffect
	}
	return flat
}

func (t dragon) string() string {
	return "ドラゴン"
}
