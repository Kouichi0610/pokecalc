package types

type water struct {
}

func (t water) effect(d effector) Effect {
	switch d.(type) {
	case *water, *grass, *dragon:
		return notvery
	case *fire, *ground, *rock:
		return super
	}
	return flat
}

func (t water) string() string {
	return "みず"
}
