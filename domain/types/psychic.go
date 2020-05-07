package types

type psychic struct {
}

func (t psychic) effect(d effector) Effect {
	switch d.(type) {
	case *psychic, *steel:
		return notvery
	case *fighting, *poison:
		return super
	case *dark:
		return noeffect
	}
	return flat
}

func (t psychic) string() string {
	return "エスパー"
}
