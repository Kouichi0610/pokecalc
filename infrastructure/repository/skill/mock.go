package skill

import (
	"fmt"
	"pokecalc/domain/skill"
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
	m.d = append(m.d, &data{id: 1, name: "ひのこ", ty: types.Fire, category: skill.Special, power: 40, accuracy: 100})
	m.d = append(m.d, &data{id: 2, name: "みずでっぽう", ty: types.Water, category: skill.Special, power: 40, accuracy: 100})
	m.d = append(m.d, &data{id: 3, name: "えだづき", ty: types.Grass, category: skill.Physical, power: 40, accuracy: 100})
	m.d = append(m.d, &data{id: 4, name: "サイコショック", ty: types.Psychic, category: skill.PsycoShock, power: 80, accuracy: 100})
	m.d = append(m.d, &data{id: 5, name: "ボディプレス", ty: types.Fighting, category: skill.BodyPress, power: 80, accuracy: 100})
	m.d = append(m.d, &data{id: 6, name: "イカサマ", ty: types.Dark, category: skill.FoulPlay, power: 90, accuracy: 100})
	m.d = append(m.d, &data{id: 7, name: "フレアドライブ", ty: types.Fire, category: skill.Physical, power: 120, accuracy: 100})
}

func (m *Mock) Load(name string) (*skill.Skill, error) {
	for _, p := range m.d {
		if p.name == name {
			return p.Skill(), nil
		}
	}
	return nil, fmt.Errorf("%s not found.", name)
}
