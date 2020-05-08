// タイプ取得
package types

import (
	"fmt"
	"reflect"
)

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

type Types struct {
	t []effector
}

// Type配列からTypesを生成 最低一つ以上指定すること
func New(args ...Type) (*Types, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("no type.")
	}
	res := new(Types)
	res.t = make([]effector, 0)

	for _, n := range args {
		ty, err := fromInt(n)
		if err != nil {
			return nil, err
		}
		res.t = append(res.t, ty)
	}
	return res, nil
}

func (at *Types) Effect(df *Types) Effect {
	var res Effect = flat
	for _, a := range at.t {
		for _, d := range df.t {
			ef := a.effect(d)
			res *= ef
		}
	}
	return res
}

func (t *Types) PartialMatch(o *Types) bool {
	for _, a := range t.t {
		for _, b := range o.t {
			if reflect.TypeOf(a) == reflect.TypeOf(b) {
				return true
			}
		}
	}
	return false
}

func (t *Types) String() string {
	var res string
	last := len(t.t) - 1
	for i, a := range t.t {
		res += a.string()
		if i == last {
			break
		}
		res += "/"
	}
	return res
}
