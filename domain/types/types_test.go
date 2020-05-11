package types

import (
	"testing"
)

// effector for test
type dummy struct {
}

func (t dummy) effect(d effector) Effect {
	return 1.0
}

func (t dummy) string() string {
	return ""
}

func Test_New(t *testing.T) {
	ty, e := New(Fire, Water)
	if e != nil {
		t.Error()
	}
	if ty.String() != "ほのお/みず" {
		t.Error()
	}
	_, err := New()
	if err == nil {
		t.Error()
	}
}

func Test_NewError(t *testing.T) {
	_, e := New(None)
	if e == nil {
		t.Error()
	}
}

func Test_NewWithNone(t *testing.T) {
	ty, _ := New(Grass, None)
	if len(ty.t) != 1 {
		t.Error()
	}
}

func Test_New_範囲外(t *testing.T) {
	_, e := New(20)
	if e == nil {
		t.Error()
	}
}

func Test_TypesEffect(t *testing.T) {
	a, _ := New(Fire)
	d, _ := New(Grass, Steel)

	actual := a.Effect(d)
	if actual != 4.0 {
		t.Error()
	}
}

func Test_PartialMatch(t *testing.T) {
	a, _ := New(Fairy, Dark)
	s, _ := New(Dark)
	e, _ := New(Poison)

	if a.PartialMatch(s) == false {
		t.Error()
	}
	if a.PartialMatch(e) != false {
		t.Error()
	}
}

func Test_toInt_登録外(t *testing.T) {
	_, err := toInt(new(dummy))
	if err == nil {
		t.Error()
	}
}

func Test_fromIntおよびtoInt(t *testing.T) {
	names := map[Type]string{
		Normal:   "ノーマル",
		Fire:     "ほのお",
		Water:    "みず",
		Electric: "でんき",
		Grass:    "くさ",
		Ice:      "こおり",
		Fighting: "かくとう",
		Poison:   "どく",
		Ground:   "じめん",
		Flying:   "ひこう",
		Psychic:  "エスパー",
		Bug:      "むし",
		Rock:     "いわ",
		Ghost:    "ゴースト",
		Dragon:   "ドラゴン",
		Dark:     "あく",
		Steel:    "はがね",
		Fairy:    "フェアリー",
	}

	for k, expect := range names {
		actual, _ := fromInt(k)
		if actual.string() != expect {
			t.Errorf("fromInt failed actual %s expect %s", actual.string(), expect)
		}

		j, _ := toInt(actual)
		if j != k {
			t.Errorf("toInt failed %T %d", actual, j)
		}

	}
}

func Test_fromInt_範囲外(t *testing.T) {
	_, e := fromInt(20)
	if e == nil {
		t.Error()
	}
}

func Test_むしタイプ相性(t *testing.T) {
	a := new(bug)
	expects := testExpects{
		fire:     notvery,
		fighting: notvery,
		poison:   notvery,
		flying:   notvery,
		ghost:    notvery,
		fairy:    notvery,
		steel:    notvery,
		grass:    super,
		psychic:  super,
		dark:     super,
		bug:      flat,
		dragon:   flat,
		water:    flat,
		ground:   flat,
		rock:     flat,
		electric: flat,
		ice:      flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_あくタイプ相性(t *testing.T) {
	a := new(dark)
	expects := testExpects{
		fighting: notvery,
		fairy:    notvery,
		dark:     notvery,
		ghost:    super,
		psychic:  super,
		fire:     flat,
		grass:    flat,
		poison:   flat,
		flying:   flat,
		bug:      flat,
		dragon:   flat,
		steel:    flat,
		water:    flat,
		ground:   flat,
		rock:     flat,
		electric: flat,
		ice:      flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_ドラゴンタイプ相性(t *testing.T) {
	a := new(dragon)
	expects := testExpects{
		steel:    notvery,
		dragon:   super,
		fairy:    noeffect,
		dark:     flat,
		fire:     flat,
		grass:    flat,
		poison:   flat,
		flying:   flat,
		bug:      flat,
		water:    flat,
		ground:   flat,
		rock:     flat,
		electric: flat,
		ice:      flat,
		ghost:    flat,
		psychic:  flat,
		fighting: flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_でんきタイプ相性(t *testing.T) {
	a := new(electric)
	expects := testExpects{
		electric: notvery,
		grass:    notvery,
		dragon:   notvery,
		water:    super,
		flying:   super,
		ground:   noeffect,
		dark:     flat,
		fire:     flat,
		poison:   flat,
		bug:      flat,
		steel:    flat,
		rock:     flat,
		ice:      flat,
		ghost:    flat,
		fairy:    flat,
		psychic:  flat,
		fighting: flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_フェアリータイプ相性(t *testing.T) {
	a := new(fairy)
	expects := testExpects{
		fire:     notvery,
		poison:   notvery,
		steel:    notvery,
		fighting: super,
		dragon:   super,
		dark:     super,
		grass:    flat,
		flying:   flat,
		bug:      flat,
		water:    flat,
		ground:   flat,
		rock:     flat,
		electric: flat,
		ice:      flat,
		ghost:    flat,
		fairy:    flat,
		psychic:  flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_かくとうタイプ相性(t *testing.T) {
	a := new(fighting)
	expects := testExpects{
		poison:   notvery,
		flying:   notvery,
		psychic:  notvery,
		bug:      notvery,
		fairy:    notvery,
		normal:   super,
		steel:    super,
		rock:     super,
		dark:     super,
		ice:      super,
		ghost:    noeffect,
		fire:     flat,
		grass:    flat,
		dragon:   flat,
		water:    flat,
		ground:   flat,
		electric: flat,
		fighting: flat,
	}
	testEffects(a, &expects, t)
}

func Test_ほのおタイプ相性(t *testing.T) {
	a := new(fire)
	expects := testExpects{
		fire:     notvery,
		water:    notvery,
		dragon:   notvery,
		rock:     notvery,
		grass:    super,
		steel:    super,
		ice:      super,
		bug:      super,
		dark:     flat,
		poison:   flat,
		flying:   flat,
		ground:   flat,
		electric: flat,
		ghost:    flat,
		fairy:    flat,
		psychic:  flat,
		fighting: flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_ひこうタイプ相性(t *testing.T) {
	a := new(flying)
	expects := testExpects{
		electric: notvery,
		rock:     notvery,
		steel:    notvery,
		grass:    super,
		fighting: super,
		bug:      super,
		dark:     flat,
		fire:     flat,
		poison:   flat,
		flying:   flat,
		dragon:   flat,
		water:    flat,
		ground:   flat,
		ice:      flat,
		ghost:    flat,
		fairy:    flat,
		psychic:  flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_ゴーストタイプ相性(t *testing.T) {
	a := new(ghost)
	expects := testExpects{
		dark:     notvery,
		psychic:  super,
		ghost:    super,
		normal:   noeffect,
		fire:     flat,
		grass:    flat,
		poison:   flat,
		flying:   flat,
		bug:      flat,
		dragon:   flat,
		steel:    flat,
		water:    flat,
		ground:   flat,
		rock:     flat,
		electric: flat,
		ice:      flat,
		fairy:    flat,
		fighting: flat,
	}
	testEffects(a, &expects, t)
}

func Test_くさタイプ相性(t *testing.T) {
	a := new(grass)
	expects := testExpects{
		fire:     notvery,
		grass:    notvery,
		poison:   notvery,
		flying:   notvery,
		bug:      notvery,
		dragon:   notvery,
		steel:    notvery,
		water:    super,
		ground:   super,
		rock:     super,
		electric: flat,
		ice:      flat,
		ghost:    flat,
		fairy:    flat,
		psychic:  flat,
		fighting: flat,
		dark:     flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_じめんタイプ相性(t *testing.T) {
	a := new(ground)
	expects := testExpects{
		grass:    notvery,
		bug:      notvery,
		fire:     super,
		electric: super,
		poison:   super,
		rock:     super,
		steel:    super,
		flying:   noeffect,
		water:    flat,
		ice:      flat,
		ground:   flat,
		dragon:   flat,
		ghost:    flat,
		fairy:    flat,
		psychic:  flat,
		fighting: flat,
		dark:     flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_こおりタイプ相性(t *testing.T) {
	a := new(ice)
	expects := testExpects{
		water:    notvery,
		ice:      notvery,
		fire:     notvery,
		steel:    notvery,
		grass:    super,
		ground:   super,
		flying:   super,
		dragon:   super,
		rock:     flat,
		ghost:    flat,
		poison:   flat,
		fairy:    flat,
		psychic:  flat,
		fighting: flat,
		dark:     flat,
		bug:      flat,
		electric: flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_ノーマルタイプ相性(t *testing.T) {
	a := new(normal)
	expects := testExpects{
		rock:     notvery,
		steel:    notvery,
		ghost:    noeffect,
		poison:   flat,
		ground:   flat,
		grass:    flat,
		fairy:    flat,
		psychic:  flat,
		fighting: flat,
		dark:     flat,
		fire:     flat,
		ice:      flat,
		flying:   flat,
		bug:      flat,
		water:    flat,
		electric: flat,
		dragon:   flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_どくタイプ相性(t *testing.T) {
	a := new(poison)
	expects := testExpects{
		poison:   notvery,
		ground:   notvery,
		rock:     notvery,
		ghost:    notvery,
		grass:    super,
		fairy:    super,
		steel:    noeffect,
		psychic:  flat,
		fighting: flat,
		dark:     flat,
		fire:     flat,
		ice:      flat,
		flying:   flat,
		bug:      flat,
		water:    flat,
		electric: flat,
		dragon:   flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_エスパータイプ相性(t *testing.T) {
	a := new(psychic)
	expects := testExpects{
		psychic:  notvery,
		steel:    notvery,
		fighting: super,
		poison:   super,
		dark:     noeffect,
		ground:   flat,
		fire:     flat,
		ice:      flat,
		flying:   flat,
		bug:      flat,
		water:    flat,
		electric: flat,
		rock:     flat,
		fairy:    flat,
		grass:    flat,
		dragon:   flat,
		ghost:    flat,
		normal:   flat,
	}
	testEffects(a, &expects, t)
}

func Test_いわタイプ相性(t *testing.T) {
	a := new(rock)
	expects := testExpects{
		fighting: notvery,
		ground:   notvery,
		steel:    notvery,
		fire:     super,
		ice:      super,
		flying:   super,
		bug:      super,
		water:    flat,
		electric: flat,
		rock:     flat,
		fairy:    flat,
		grass:    flat,
		dragon:   flat,
		dark:     flat,
		ghost:    flat,
		normal:   flat,
		poison:   flat,
		psychic:  flat,
	}
	testEffects(a, &expects, t)
}

func Test_はがねタイプ相性(t *testing.T) {
	a := new(steel)
	expects := testExpects{
		fire:     notvery,
		water:    notvery,
		electric: notvery,
		steel:    notvery,
		ice:      super,
		rock:     super,
		fairy:    super,
		grass:    flat,
		dragon:   flat,
		ground:   flat,
		bug:      flat,
		dark:     flat,
		fighting: flat,
		flying:   flat,
		ghost:    flat,
		normal:   flat,
		poison:   flat,
		psychic:  flat,
	}
	testEffects(a, &expects, t)
}

func Test_みずタイプ相性(t *testing.T) {
	a := new(water)
	expects := testExpects{
		water:    notvery,
		grass:    notvery,
		dragon:   notvery,
		fire:     super,
		ground:   super,
		rock:     super,
		bug:      flat,
		dark:     flat,
		electric: flat,
		fairy:    flat,
		fighting: flat,
		flying:   flat,
		ghost:    flat,
		ice:      flat,
		normal:   flat,
		poison:   flat,
		psychic:  flat,
		steel:    flat,
	}
	testEffects(a, &expects, t)
}

func (e *testExpects) toMap() map[effector]Effect {
	res := map[effector]Effect{
		new(bug):      e.bug,
		new(dark):     e.dark,
		new(dragon):   e.dragon,
		new(electric): e.electric,
		new(fairy):    e.fairy,
		new(fighting): e.fighting,
		new(fire):     e.fire,
		new(flying):   e.flying,
		new(ghost):    e.ghost,
		new(grass):    e.grass,
		new(ground):   e.ground,
		new(ice):      e.ice,
		new(normal):   e.normal,
		new(poison):   e.poison,
		new(psychic):  e.psychic,
		new(rock):     e.rock,
		new(steel):    e.steel,
		new(water):    e.water,
	}
	return res
}

func testEffects(a effector, e *testExpects, t *testing.T) {
	expects := e.toMap()
	for d, expect := range expects {
		actual := a.effect(d)
		if actual != expect {
			t.Errorf("failed (%T) expect (%f) actual (%f)\n", d, expect, actual)
		}
	}
}

// 期待値一覧
type testExpects struct {
	bug, dark, dragon, electric, fairy, fighting, fire, flying, ghost, grass, ground, ice, normal, poison, psychic, rock, steel, water Effect
}
