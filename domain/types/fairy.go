package types

type fairy struct {
}

func (t fairy) effect(d effector) Effect {
	switch d.(type) {
	case *fire, *poison, *steel:
		return notvery
	case *fighting, *dragon, *dark:
		return super
	}
	return flat
}

func (t fairy) string() string {
	return "フェアリー"
}
