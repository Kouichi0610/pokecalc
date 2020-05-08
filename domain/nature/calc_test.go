package nature

import (
	"testing"
)

// テストコード内で定義された定数、関数、メソッドはテスト外では使用できない

func Test_Flat(t *testing.T) {
	const l = 100
	const s = 95
	f := flat

	v := f(l, s, 31, 0)
	if v != 226 {
		t.Error()
	}
	v = f(l, s, 0, 0)
	if v != 195 {
		t.Error()
	}
	v = f(l, s, 31, 252)
	if v != 289 {
		t.Error()
	}
}

func Test_Upper(t *testing.T) {
	const l = 100
	const s = 95
	f := upper

	v := f(l, s, 31, 0)
	if v != 248 {
		t.Error()
	}
	v = f(l, s, 0, 0)
	if v != 214 {
		t.Error()
	}
	v = f(l, s, 31, 252)
	if v != 317 {
		t.Error()
	}
}

func Test_Lower(t *testing.T) {
	const l = 100
	const s = 95
	f := lower

	v := f(l, s, 31, 0)
	if v != 203 {
		t.Error()
	}
	v = f(l, s, 0, 0)
	if v != 175 {
		t.Error()
	}
	v = f(l, s, 31, 252)
	if v != 260 {
		t.Error()
	}
}

func Test_HP(t *testing.T) {
	const l = 100
	const s = 108
	f := hp

	v := f(l, s, 31, 0)
	if v != 357 {
		t.Error()
	}
	v = f(l, s, 30, 0)
	if v != 356 {
		t.Error()
	}
	v = f(l, s, 0, 0)
	if v != 326 {
		t.Error()
	}
	v = f(l, s, 31, 252)
	if v != 420 {
		t.Error()
	}

}
func Test_ShadinjaHP(t *testing.T) {
	const l = 100
	const s = 108
	f := shadinjaHP

	v := f(l, s, 31, 0)
	if v != 1 {
		t.Error()
	}
	v = f(l, s, 30, 0)
	if v != 1 {
		t.Error()
	}
	v = f(l, s, 0, 0)
	if v != 1 {
		t.Error()
	}
	v = f(l, s, 31, 252)
	if v != 1 {
		t.Error()
	}
}
