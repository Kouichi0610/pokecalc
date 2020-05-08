package types

type fire struct {
}

func (t fire) effect(d effector) Effect {
	switch d.(type) {
	case *fire, *water, *rock, *dragon:
		return notvery
	case *grass, *ice, *bug, *steel:
		return super
	}
	return flat
}

func (t fire) string() string {
	return "ほのお"
}
