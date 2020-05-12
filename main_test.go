package main

import (
	"testing"
)

type Hoge struct {
	Name string
}

type Fuga struct {
	hoge *Hoge
}

func Test_Structure(t *testing.T) {
	h := &Hoge{"hogehoge"}
	f := &Fuga{h}

	if f.hoge.Name != "hogehoge" {
		t.Error()
	}

	h.Name = "fugafuga"

	if f.hoge.Name != "fugafuga" {
		t.Error()
	}
}
