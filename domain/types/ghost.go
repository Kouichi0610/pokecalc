package types

type ghost struct {
}

func (t ghost) effect(d effector) Effect {
	switch d.(type) {
	case *dark:
		return notvery
	case *psychic, *ghost:
		return super
	case *normal:
		return noeffect
	}
	return flat
}

func (t ghost) string() string {
	return "ゴースト"
}
