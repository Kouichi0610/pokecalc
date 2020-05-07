package types

type Type uint

const flat = 1.0
const notvery = 0.5
const super = 2.0
const noeffect = 0.0

const (
	None Type = iota
	Normal
	Fire
	Water
	Electric
	Grass
	Ice
	Fighting
	Poison
	Ground
	Flying
	Psychic
	Bug
	Rock
	Ghost
	Dragon
	Dark
	Steel
	Fairy
)
