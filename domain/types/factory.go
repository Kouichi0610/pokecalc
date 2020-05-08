// effector生成に関わる
package types

import (
	"fmt"
	"reflect"
)

var effectors map[Type]effector

func fromInt(id Type) (effector, error) {
	res, ok := effectors[id]
	if !ok {
		return nil, fmt.Errorf("%d not supported.", id)
	}
	return res, nil
}

func toInt(d effector) (Type, error) {
	for key, value := range effectors {
		if reflect.TypeOf(value) == reflect.TypeOf(d) {
			return key, nil
		}
	}
	return 0, fmt.Errorf("%T not supported.", d)
}

func init() {
	effectors = map[Type]effector{
		Normal:   new(normal),
		Fire:     new(fire),
		Water:    new(water),
		Electric: new(electric),
		Grass:    new(grass),
		Ice:      new(ice),
		Fighting: new(fighting),
		Poison:   new(poison),
		Ground:   new(ground),
		Flying:   new(flying),
		Psychic:  new(psychic),
		Bug:      new(bug),
		Rock:     new(rock),
		Ghost:    new(ghost),
		Dragon:   new(dragon),
		Dark:     new(dark),
		Steel:    new(steel),
		Fairy:    new(fairy),
	}
}
