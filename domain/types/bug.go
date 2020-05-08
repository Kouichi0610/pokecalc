package types

type bug struct {
}

func (t bug) effect(d effector) Effect {
	switch d.(type) {
	case *fire, *fighting, *poison, *flying, *ghost, *steel, *fairy:
		return notvery
	case *grass, *psychic, *dark:
		return super
	}
	return flat
}

func (t bug) string() string {
	return "むし"
}
