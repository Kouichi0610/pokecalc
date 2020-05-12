package species

import (
	"fmt"
	"pokecalc/domain/stats"
	"pokecalc/domain/types"
)

type Mock struct {
	d []*data
}

func NewMock() Repository {
	res := &Mock{}
	res.initialize()
	return res
}

func (m *Mock) initialize() {
	m.d = make([]*data, 0)
	m.d = append(m.d, &data{id: 1, name: "ミュウ", type1: types.Psychic, type2: types.None, ability1: 0, ability2: 0, ability3: 0, hp: 100, attack: 100, defense: 100, spAttack: 100, spDefense: 100, speed: 100})
	m.d = append(m.d, &data{id: 2, name: "ゴウカザル", type1: types.Fire, type2: types.Fighting, ability1: 0, ability2: 0, ability3: 0, hp: 75, attack: 104, defense: 71, spAttack: 104, spDefense: 71, speed: 108})
	m.d = append(m.d, &data{id: 3, name: "エンペルト", type1: types.Water, type2: types.Steel, ability1: 0, ability2: 0, ability3: 0, hp: 84, attack: 86, defense: 88, spAttack: 111, spDefense: 101, speed: 60})
	m.d = append(m.d, &data{id: 4, name: "ドダイトス", type1: types.Grass, type2: types.Ground, ability1: 0, ability2: 0, ability3: 0, hp: 95, attack: 109, defense: 105, spAttack: 75, spDefense: 85, speed: 56})
}

func (m *Mock) Load(name string) (*stats.Species, *types.Types, error) {
	for _, p := range m.d {
		if p.name == name {
			s := p.Species()
			t := p.Types()
			return s, t, nil
		}
	}
	return nil, nil, fmt.Errorf("%s not found.", name)
}
