package damage

import (
	"reflect"
	"testing"
)

func Test_Calc(t *testing.T) {
	const l uint = 50
	const s uint = 100
	const a uint = 182
	const d uint = 55

	actual := calc(l, s, a, d)
	expect := uint(147)

	if actual != expect {
		t.Errorf("Failed Damage Expect %d Actual %d", expect, actual)
	}
}

func Test_CalcArray(t *testing.T) {
	const l uint = 50
	const s uint = 100
	const a uint = 182
	const d uint = 55

	actual := calcArray(l, s, a, d)
	expect := []uint{124, 126, 127, 129, 130, 132, 133, 135, 136, 138, 139, 141, 142, 144, 145, 147}
	if len(actual) != 16 {
		t.Errorf("array length %d", len(actual))
	}
	if !reflect.DeepEqual(actual, expect) {
		t.Errorf("Actual %v", actual)
		t.Errorf("Expect %v", expect)
	}
}
