package types

type Effect float32

type effector interface {
	effect(d effector) Effect
	string() string
}
