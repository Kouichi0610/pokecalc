package types

import (
	"testing"
)

func Test_試作(t *testing.T) {
	n := new(normal)
	d := new(normal)

	e := n.effect(d)

	if e != flat {
		t.Error()
	}

}
